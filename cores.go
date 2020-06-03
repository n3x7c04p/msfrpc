package msfrpc

import "github.com/n3x7c04p/msfrpc/models"

//cores
func (msf *Msfrpc) CoreAddModulePath(Path string) (models.CoreModuleLengthRes, error) {
	req := &models.CoreAddModulePathReq{Method: models.CoreAddModulePath, Token: msf.token, Path: Path}
	var res models.CoreModuleLengthRes
	if err := msf.send(req, &res); err != nil {
		return models.CoreModuleLengthRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) CoreModuleStats() (models.CoreModuleLengthRes, error) {
	req := &models.CoreModuleStatsReq{Method: models.CoreModuleStats, Token: msf.token}
	var res models.CoreModuleLengthRes
	if err := msf.send(req, &res); err != nil {
		return models.CoreModuleLengthRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) CoreReloadModules() (models.CoreModuleLengthRes, error) {
	req := &models.CoreReloadModulesReq{Method: models.CoreReloadModules, Token: msf.token}
	var res models.CoreModuleLengthRes
	if err := msf.send(req, &res); err != nil {
		return models.CoreModuleLengthRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) CoreSave() (bool, error) {
	req := &models.CoreSaveReq{Method: models.CoreSave, Token: msf.token}
	var res models.CoreSaveRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) CoreSetG(OptionName string, OptionValue string) (bool, error) {
	req := &models.CoreSetGReq{Method: models.CoreSetG, Token: msf.token, OptName: OptionName, OptValue: OptionValue}
	var res models.CoreSetGRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) CoreUnsetG(OptionName string) (bool, error) {
	req := &models.CoreUnsetGReq{Method: models.CoreUnsetG, Token: msf.token, OptName: OptionName}
	var res models.CoreUnsetGRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) CoreThreadList() (map[int]models.CoreThreadListRes, error) {
	req := &models.CoreThreadListReq{Method: models.CoreThreadList, Token: msf.token}
	res := make(map[int]models.CoreThreadListRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (msf *Msfrpc) CoreThreadKill(ThreadID string) (bool, error) {
	req := &models.CoreThreadKillReq{Method: models.CoreThreadKill, Token: msf.token, ThreadID: ThreadID}
	var res models.CoreThreadKillRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}

func (msf *Msfrpc) CoreVersion() (models.CoreVersionRes, error) {
	req := &models.CoreVersionReq{Method: models.CoreVersion, Token: msf.token}
	var res models.CoreVersionRes
	if err := msf.send(req, &res); err != nil {
		return models.CoreVersionRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) CoreStop() (bool, error) {
	req := &models.CoreStopReq{Method: models.CoreStop, Token: msf.token}
	var res models.CoreStopRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}
