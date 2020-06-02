package msfrpc

import "msfrpc/models"

//consoles
func (msf *Msfrpc) ConsoleCreate() (models.ConsoleCreateRes, error) {
	req := &models.ConsoleCreateReq{Method: models.ConsoleCreate, Token: msf.token}
	var res models.ConsoleCreateRes
	if err := msf.send(req, &res); err != nil {
		return models.ConsoleCreateRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) ConsoleDestroy(ConsoleID string) (bool, error) {
	req := &models.ConsoleDestroyReq{Method: models.ConsoleDestroy, Token: msf.token, ConsoleID: ConsoleID}
	var res models.ConsoleDestroyRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) ConsoleList() (map[int]models.ConsoleCreateRes, error) {
	req := &models.ConsoleListReq{Method: models.ConsoleList, Token: msf.token}
	res := make(map[int]models.ConsoleCreateRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	for id, console := range res {
		id = console.ID
		res[id] = console
	}
	return res, nil
}

func (msf *Msfrpc) ConsoleWrite(ConsoleID string, Data string) (int, error) {
	req := &models.ConsoleWriteReq{Method: models.ConsoleWrite, Token: msf.token, ConsoleID: ConsoleID, Data: Data}
	var res models.ConsoleWriteRes
	if err := msf.send(req, &res); err != nil {
		return 0, err
	}
	return res.Wrote, nil
}

func (msf *Msfrpc) ConsoleRead(ConsoleID string) (models.ConsoleReadRes, error) {
	req := &models.ConsoleReadReq{Method: models.ConsoleRead, Token: msf.token, ConsoleID: ConsoleID}
	var res models.ConsoleReadRes
	if err := msf.send(req, &res); err != nil {
		return models.ConsoleReadRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) ConsoleSessionDetach(ConsoleID string) (bool, error) {
	req := &models.ConsoleSessionDetachReq{Method: models.ConsoleSessionDetach, Token: msf.token, ConsoleID: ConsoleID}
	var res models.ConsoleSessionDetachRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) ConsoleSessionKill(ConsoleID string) (bool, error) {
	req := &models.ConsoleSessionKillReq{Method: models.ConsoleSessionKill, Token: msf.token, ConsoleID: ConsoleID}
	var res models.ConsoleSessionKillRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) ConsoleTabs(ConsoleID string, InputLine string) ([]string, error) {
	req := &models.ConsoleTabsReq{Method: models.ConsoleTabs, Token: msf.token, ConsoleID: ConsoleID, InputLine: InputLine}
	var res models.ConsoleTabsRes
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res.Tabs, nil
}
