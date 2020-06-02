package msfrpc

import (
	"bytes"
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
	r, err := http.Post(dest, "binary/message-pack", msf.buf)
	if err != nil {
		return err
	}

	if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil {
		return err
	}

	msf.buf.Reset()

	defer r.Body.Close()

	return nil
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
