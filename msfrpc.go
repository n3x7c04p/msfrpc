package msfrpc

import (
	"bytes"
	"fmt"
	"github.com/vmihailenco/msgpack/v4"
	models "msfrpc/models"
	"net/http"
)

type Msfrpc struct {
	host  string
	port  uint64
	user  string
	pass  string
	ssl   bool 
	token string
}

func New(host string, port uint64, user string, pass string, ssl bool) (*Msfrpc, error) {
	msf := &Msfrpc{
		host: host,
		port: port,
		user: user,
		pass: pass,
		ssl: ssl,
	}

	if err := msf.Login(); err != nil {
		return nil, err
	}

	return msf, nil
}

func (msf *Msfrpc) send(req interface{}, res interface{}) error {
	buf := new(bytes.Buffer)
	var scheme string
	if msf.ssl { scheme = "https" }else{ scheme = "http" }
	if err := msgpack.NewEncoder(buf).Encode(req); err != nil {
		return err
	}
	dest := fmt.Sprintf("%s://%s:%d/api", scheme, msf.host,msf.port)
	r, err := http.Post(dest, "binary/message-pack", buf)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil {
		return err
	}

	return nil
}

func (msf *Msfrpc) Login() error {
	ctx := &models.AuthLoginReq{
		Method:   AuthLogin,
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
		Method:          "auth.logout",
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

//consoles
func (msf *Msfrpc) ConsoleCreate() (models.ConsoleCreateRes, error) {
	req := &models.ConsoleCreateReq{Method: ConsoleCreate, Token: msf.token}
	var res models.ConsoleCreateRes
	if err := msf.send(req, &res); err != nil {
		return models.ConsoleCreateRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) ConsoleDestroy(ConsoleID string) (bool, error) {
	req := &models.ConsoleDestroyReq{Method: ConsoleDestroy, Token: msf.token, ConsoleID: ConsoleID}
	var res models.ConsoleDestroyRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

func (msf *Msfrpc) ConsoleList() (map[uint64]models.ConsoleCreateRes, error) {
	req := &models.ConsoleListReq{Method: ConsoleList, Token: msf.token}
	res := make(map[uint64]models.ConsoleCreateRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	for id,console := range res {
		id = console.ID
		res[id] = console
	}
	return res, nil
}

func (msf *Msfrpc) ConsoleWrite(ConsoleID string, Data string) (uint64, error) {
	req := &models.ConsoleWriteReq{Method: ConsoleWrite, Token: msf.token, ConsoleID: ConsoleID, Data: Data}
	var res models.ConsoleWriteRes
	if err := msf.send(req, &res); err != nil {
		return 0, err
	}
	return res.Wrote, nil
}

func (msf *Msfrpc) ConsoleRead(ConsoleID string) (models.ConsoleReadRes, error) {
	req := &models.ConsoleReadReq{Method: ConsoleRead, Token: msf.token, ConsoleID: ConsoleID}
	var res models.ConsoleReadRes
	if err := msf.send(req, &res); err != nil {
		return models.ConsoleReadRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) ConsoleSessionDetach(ConsoleID string) (bool, error) {
	req := &models.ConsoleSessionDetachReq{Method: ConsoleSessionDetach, Token: msf.token, ConsoleID: ConsoleID}
	var res models.ConsoleSessionDetachRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

func (msf *Msfrpc) ConsoleSessionKill(ConsoleID string) (bool, error) {
	req := &models.ConsoleSessionKillReq{Method: ConsoleSessionKill, Token: msf.token, ConsoleID: ConsoleID}
	var res models.ConsoleSessionKillRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

func (msf *Msfrpc) ConsoleTabs(ConsoleID string,InputLine string) ([]string, error) {
	req := &models.ConsoleTabsReq{Method: ConsoleTabs, Token: msf.token, ConsoleID: ConsoleID, InputLine: InputLine}
	var res models.ConsoleTabsRes
	if err := msf.send(req, &res); err != nil {
		return []string{}, err
	}
	return res.Tabs, nil
}

//modules
func (msf *Msfrpc) ModuleList(ModuleType string) ([]string, error) {
	var req interface{}
	var res map[string][]string
	switch ModuleType {
	case "exploits":
		req = &models.ModuleExploitsReq{Method: ModuleExploits, Token: msf.token}
	case "payloads":
		req = &models.ModulePayloadsReq{Method: ModulePayloads, Token: msf.token}
	case "auxiliary":
		req = &models.ModuleAuxiliaryReq{Method: ModuleAuxiliary, Token: msf.token}
	case "encoders":
		req = &models.ModuleEncodersReq{Method: ModuleEncoders, Token: msf.token}
	case "post":
		req = &models.ModulePostReq{Method: ModulePost, Token: msf.token}
	case "nops":
		req = &models.ModuleNopsReq{Method: ModuleNops, Token: msf.token}
	default:
		req = &models.ModuleExploitsReq{Method: ModuleExploits, Token: msf.token}
	}

	if err := msf.send(req, &res); err != nil {
		return nil, err
	}

	return res["modules"], nil
}

func (msf *Msfrpc) ModuleInfo(ModuleType string,ModuleName string) (models.ModuleInfoRes, error) {
	req := &models.ModuleInfoReq{Method: ModuleInfo, Token: msf.token, ModType: ModuleType, ModName: ModuleName}
	var res models.ModuleInfoRes
	if err := msf.send(req, &res); err != nil {
		return models.ModuleInfoRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) ModuleOptions(ModuleType string,ModuleName string) (map[string]models.ModuleOptionsRes, error) {
	req := &models.ModuleOptionsReq{Method: ModuleOptions, Token: msf.token, ModType: ModuleType, ModName: ModuleName}
	res := make(map[string]models.ModuleOptionsRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (msf *Msfrpc) ModuleCompatiablePayloads(ModuleName string) ([]string, error) {
	req := &models.ModuleCompatiblePayloadsReq{Method: ModuleCompatiblePayloads, Token: msf.token, ModName: ModuleName}
	var res models.ModuleCompatiblePayloadsRes
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res.Payloads, nil
}

func (msf *Msfrpc) ModuleTargetCompatiablePayloads(ModuleName string, Target uint64) ([]string, error) {
	req := &models.ModuleTargetCompatiblePayloadsReq{Method: ModuleTargetCompatiblePayloads, Token: msf.token, ModName: ModuleName, Target: Target}
	var res models.ModuleTargetCompatiblePayloadsRes
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res.Payloads, nil
}

/*func (msf *Msfrpc) ModuleCompatiableSessions(ModuleName string) ([]uint64, error) {
	req := &models.ModuleCompatibleSessionsReq{Method: ModuleTargetCompatiblePayloads, Token: msf.token, ModName: ModuleName}
	res := make(map[string]models.ModuleCompatiableSessionRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res["sessions"].Sessions, nil
}*/

func (msf *Msfrpc) ModuleExecute(ModuleType string, ModuleName string) (models.ModuleExecuteRes, error) {
	req := &models.ModuleExecuteReq{Method: ModuleExecute, Token: msf.token, ModType: ModuleType, ModName: ModuleName}
	var res models.ModuleExecuteRes
	if err := msf.send(req, &res); err != nil {
		return models.ModuleExecuteRes{}, err
	}
	return res, nil
}


//jobs
func (msf *Msfrpc) JobList() (map[uint64]models.JobListRes,error) {
	req := &models.JobListReq{Method: JobList, Token: msf.token}
	res := make(map[uint64]models.JobListRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	for id, module := range res {
		res[id] = module
	}
	return res, nil
}

func (msf *Msfrpc) JobInfo(JobID string) (models.JobInfoRes,error) {
	req := &models.JobInfoReq{Method: JobInfo, Token: msf.token, JobID: JobID}
	var res models.JobInfoRes
	if err := msf.send(req, &res); err != nil {
		return models.JobInfoRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) JobStop(JobID string) (bool,error) {
	req := &models.JobStopReq{Method: JobStop, Token: msf.token, JobID: JobID}
	var res models.JobStopRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

//sessions
func (msf *Msfrpc) SessionList() (map[uint64]models.SessionListRes, error) {
	req := &models.SessionListReq{Method: SessionList, Token: msf.token}
	res := make(map[uint64]models.SessionListRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}

	for id, session := range res {
		session.ID = id
		res[id] = session
	}
	return res, nil
}

func (msf *Msfrpc) SessionStop(SessionID string) (bool, error) {
	req := &models.SessionStopReq{Method: SessionList, Token: msf.token, SessionID: SessionID}
	var res models.SessionStopRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

func (msf *Msfrpc) SessionShellRead(SessionID string) (models.SessionShellReadRes, error) {
	req := &models.SessionShellReadReq{Method: SessionShellRead, Token: msf.token, SessionID: SessionID}
	var res models.SessionShellReadRes
	if err := msf.send(req, &res); err != nil {
		return models.SessionShellReadRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) SessionShellWrite(SessionID string, Data string) (uint64, error) {
	req := &models.SessionShellWriteReq{Method: SessionShellRead, Token: msf.token, SessionID: SessionID, Data: Data}
	var res models.SessionShellWriteRes
	if err := msf.send(req, &res); err != nil {
		return 0, err
	}
	return res.WriteCount, nil
}

func (msf *Msfrpc) SessionShellUpgrade(SessionID string, ConnHost string, ConnPort uint64) (bool, error) {
	req := &models.SessionShellUpgradeReq{Method: SessionShellUpgrade, Token: msf.token, SessionID: SessionID, ConnHost: ConnHost, ConnPort: ConnPort}
	var res models.SessionShellUpgradeRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

func (msf *Msfrpc) SessionMeterpreterRead(SessionID string) (string, error) {
	req := &models.SessionMeterpreterReadReq{Method: SessionMeterpreterRead, Token: msf.token, SessionID: SessionID}
	var res models.SessionMeterpreterReadRes
	if err := msf.send(req, &res); err != nil {
		return "", err
	}
	return res.Data, nil
}

func (msf *Msfrpc) SessionMeterpreterWrite(SessionID string, Data string) (bool, error) {
	req := &models.SessionMeterpreterWriteReq{Method: SessionMeterpreterRead, Token: msf.token, SessionID: SessionID, Data: Data}
	var res models.SessionMeterpreterWriteRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

/*func (msf *Msfrpc) SessionMeterpreterRunSingle(SessionID string, Data string) (bool, error) {
	req := &models.SessionMeterpreterRunSingleReq{Method: SessionMeterpreterRunSingle, Token: msf.token, SessionID: SessionID, Data: Data}
	var res models.SessionMeterpreterRunSingleRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}*/

func (msf *Msfrpc) SessionMeterpreterScript(SessionID string, ScriptName string) (bool, error) {
	req := &models.SessionMeterpreterScriptReq{Method: SessionMeterpreterScript, Token: msf.token, SessionID: SessionID, ScriptName: ScriptName}
	var res models.SessionMeterpreterScriptRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

func (msf *Msfrpc) SessionMeterpreterSessionDetach(SessionID string) (bool, error) {
	req := &models.SessionMeterpreterSessionDetachReq{Method: SessionMeterpreterSessionDetach, Token: msf.token, SessionID: SessionID}
	var res models.SessionMeterpreterSessionDetachRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

func (msf *Msfrpc) SessionMeterpreterSessionKill(SessionID string) (bool, error) {
	req := &models.SessionMeterpreterSessionKillReq{Method: SessionMeterpreterSessionKill, Token: msf.token, SessionID: SessionID}
	var res models.SessionMeterpreterSessionKillRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

func (msf *Msfrpc) SessionRingClear(SessionID string) (bool, error) {
	req := &models.SessionRingClearReq{Method: SessionRingClear, Token: msf.token, SessionID: SessionID}
	var res models.SessionRingClearRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result != "success", nil
}

func (msf *Msfrpc) SessionRingPut(SessionID string, Data string) (uint64, error) {
	req := &models.SessionRingPutReq{Method: SessionRingPut, Token: msf.token, SessionID: SessionID, Data: Data}
	var res models.SessionRingPutRes
	if err := msf.send(req, &res); err != nil {
		return 0, err
	}
	return res.WriteCount, nil
}





