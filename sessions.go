package msfrpc

import "github.com/n3x7c04p/msfrpc/models"

//sessions
func (msf *Msfrpc) SessionList() (map[int]models.SessionListRes, error) {
	req := &models.SessionListReq{Method: models.SessionList, Token: msf.token}
	res := make(map[int]models.SessionListRes)
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
	req := &models.SessionStopReq{Method: models.SessionStop, Token: msf.token, SessionID: SessionID}
	var res models.SessionStopRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) SessionCompatibleModules(SessionID string) ([]string, error) {
	req := &models.SessionCompatibleModulesReq{Method: models.SessionCompatibleModules, Token: msf.token, SessionID: SessionID}
	var res models.SessionCompatibleModulesRes
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res.Modules, nil
}

func (msf *Msfrpc) SessionShellRead(SessionID string) (models.SessionShellReadRes, error) {
	req := &models.SessionShellReadReq{Method: models.SessionShellRead, Token: msf.token, SessionID: SessionID}
	var res models.SessionShellReadRes
	if err := msf.send(req, &res); err != nil {
		return models.SessionShellReadRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) SessionShellWrite(SessionID string, Data string) (int, error) {
	req := &models.SessionShellWriteReq{Method: models.SessionShellRead, Token: msf.token, SessionID: SessionID, Data: Data}
	var res models.SessionShellWriteRes
	if err := msf.send(req, &res); err != nil {
		return 0, err
	}
	return res.WriteCount, nil
}

func (msf *Msfrpc) SessionShellUpgrade(SessionID string, ConnHost string, ConnPort int) (bool, error) {
	req := &models.SessionShellUpgradeReq{Method: models.SessionShellUpgrade, Token: msf.token, SessionID: SessionID, ConnHost: ConnHost, ConnPort: ConnPort}
	var res models.SessionShellUpgradeRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) SessionMeterpreterRead(SessionID string) (string, error) {
	req := &models.SessionMeterpreterReadReq{Method: models.SessionMeterpreterRead, Token: msf.token, SessionID: SessionID}
	var res models.SessionMeterpreterReadRes
	if err := msf.send(req, &res); err != nil {
		return "", err
	}
	return res.Data, nil
}

func (msf *Msfrpc) SessionMeterpreterWrite(SessionID string, Data string) (bool, error) {
	req := &models.SessionMeterpreterWriteReq{Method: models.SessionMeterpreterWrite, Token: msf.token, SessionID: SessionID, Data: Data}
	var res models.SessionMeterpreterWriteRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) SessionMeterpreterTabs(SessionID string, InputLine string) ([]string, error) {
	req := &models.SessionMeterpreterTabsReq{Method: models.SessionMeterpreterWrite, Token: msf.token, SessionID: SessionID, InputLine: InputLine}
	var res models.SessionMeterpreterTabsRes
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res.Tabs, nil
}

func (msf *Msfrpc) SessionMeterpreterRunSingle(SessionID string, Data string) (bool, error) {
	req := &models.SessionMeterpreterRunSingleReq{Method: models.SessionMeterpreterRunSingle, Token: msf.token, SessionID: SessionID, Data: Data}
	var res models.SessionMeterpreterRunSingleRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) SessionMeterpreterScript(SessionID string, ScriptName string) (bool, error) {
	req := &models.SessionMeterpreterScriptReq{Method: models.SessionMeterpreterScript, Token: msf.token, SessionID: SessionID, ScriptName: ScriptName}
	var res models.SessionMeterpreterScriptRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) SessionMeterpreterSessionDetach(SessionID string) (bool, error) {
	req := &models.SessionMeterpreterSessionDetachReq{Method: models.SessionMeterpreterSessionDetach, Token: msf.token, SessionID: SessionID}
	var res models.SessionMeterpreterSessionDetachRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) SessionMeterpreterSessionKill(SessionID string) (bool, error) {
	req := &models.SessionMeterpreterSessionKillReq{Method: models.SessionMeterpreterSessionKill, Token: msf.token, SessionID: SessionID}
	var res models.SessionMeterpreterSessionKillRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) SessionRingClear(SessionID string) (bool, error) {
	req := &models.SessionRingClearReq{Method: models.SessionRingClear, Token: msf.token, SessionID: SessionID}
	var res models.SessionRingClearRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) SessionRingLast(SessionID string) (int, error) {
	req := &models.SessionRingLastReq{Method: models.SessionRingLast, Token: msf.token, SessionID: SessionID}
	var res models.SessionRingLastRes
	if err := msf.send(req, &res); err != nil {
		return 0, err
	}
	return res.Seq, nil
}

func (msf *Msfrpc) SessionRingPut(SessionID string, Data string) (int, error) {
	req := &models.SessionRingPutReq{Method: models.SessionRingPut, Token: msf.token, SessionID: SessionID, Data: Data}
	var res models.SessionRingPutRes
	if err := msf.send(req, &res); err != nil {
		return 0, err
	}
	return res.WriteCount, nil
}
