package msfrpc

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/vmihailenco/msgpack/v4"
	"msfrpc/models"
	"net/http"
)

type Msfrpc struct {
	buf   *bytes.Buffer
	host  string
	port  int
	user  string
	pass  string
	ssl   bool
	token string
}

func New(host string, port int, user string, pass string, ssl bool) (*Msfrpc, error) {
	msf := &Msfrpc{
		buf:  new(bytes.Buffer),
		host: host,
		port: port,
		user: user,
		pass: pass,
		ssl:  ssl,
	}
	if err := msf.Login(); err != nil {
		return nil, err
	}

	return msf, nil
}

func (msf *Msfrpc) send(req interface{}, res interface{}) error {
	var scheme string
	if msf.ssl {
		scheme = "https"
	} else {
		scheme = "http"
	}
	if err := msgpack.NewEncoder(msf.buf).Encode(req); err != nil {
		return err
	}
	dest := fmt.Sprintf("%s://%s:%d/api", scheme, msf.host, msf.port)

	tr := &http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	request, err := http.NewRequest("POST", dest, msf.buf)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "binary/message-pack")
	request.Header.Set("Accept", "binary/message-pack")
	request.Header.Set("Accept-Charset", "UTF-8")

	r, err := client.Do(request)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	switch r.StatusCode {
	case 200:
		if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil {
			return err
		}
		msf.buf.Reset()
		return nil
	case 500:
		var errRes models.ErrorRes
		if err := msgpack.NewDecoder(r.Body).Decode(&errRes); err != nil {
			return err
		}
		return errors.New(fmt.Sprintf("%t %s %s\n", errRes.Error, errRes.ErrorClass, errRes.ErrorMessage))
	case 401:
		return errors.New(fmt.Sprintf("The authentication credentials supplied were not valid"))
	case 403:
		return errors.New(fmt.Sprintf("The authentication credentials supplied were not granted access to the resource"))
	case 404:
		return errors.New(fmt.Sprintf("The request was sent to an invalid URI"))
	default:
		return errors.New(r.Status)
	}
}

func (msf *Msfrpc) Login() error {
	ctx := &models.AuthLoginReq{
		Method:   models.AuthLogin,
		Username: msf.user,
		Password: msf.pass,
	}
	var res models.AuthLoginRes
	if err := msf.send(ctx, &res); err != nil {
		return err
	}
	msf.token = res.Token
	return nil
}

func (msf *Msfrpc) Logout() error {
	ctx := &models.AuthLogoutReq{
		Method:          models.AuthLogout,
		Token:           msf.token,
		AuthLogoutToken: msf.token,
	}
	var res models.AuthLogoutRes
	if err := msf.send(ctx, &res); err != nil {
		return err
	}
	msf.token = ""
	return nil
}
