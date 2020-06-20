package msfrpc

import (
	"fmt"
	"github.com/n3x7c04p/msfrpc/models"
	"strings"
)

//modules
func (msf *Msfrpc) ModuleList(ModuleType string) ([]string, error) {
	var req interface{}
	var res map[string][]string
	switch ModuleType {
	case "exploits":
		req = &models.ModuleExploitsReq{Method: models.ModuleExploits, Token: msf.token}
	case "payloads":
		req = &models.ModulePayloadsReq{Method: models.ModulePayloads, Token: msf.token}
	case "auxiliary":
		req = &models.ModuleAuxiliaryReq{Method: models.ModuleAuxiliary, Token: msf.token}
	case "encoders":
		req = &models.ModuleEncodersReq{Method: models.ModuleEncoders, Token: msf.token}
	case "post":
		req = &models.ModulePostReq{Method: models.ModulePost, Token: msf.token}
	case "nops":
		req = &models.ModuleNopsReq{Method: models.ModuleNops, Token: msf.token}
	default:
		req = &models.ModuleExploitsReq{Method: models.ModuleExploits, Token: msf.token}
	}

	if err := msf.send(req, &res); err != nil {
		return nil, err
	}

	return res["modules"], nil
}

func (msf *Msfrpc) ModuleInfo(ModuleType string, ModuleName string) (models.ModuleInfoRes, error) {
	req := &models.ModuleInfoReq{Method: models.ModuleInfo, Token: msf.token, ModType: ModuleType, ModName: ModuleName}
	var res models.ModuleInfoRes
	if err := msf.send(req, &res); err != nil {
		return models.ModuleInfoRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) ModuleOptions(ModuleType string, ModuleName string) (map[string]models.ModuleOptionsRes, error) {
	req := &models.ModuleOptionsReq{Method: models.ModuleOptions, Token: msf.token, ModType: ModuleType, ModName: ModuleName}
	res := make(map[string]models.ModuleOptionsRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	for key, options := range res {
		if w, ok := options.Default.([]uint8); ok {
			var buf []string = nil
			for _, c := range w {
				buf = append(buf, fmt.Sprintf("%c", c))
			}
			str := strings.Join(buf, "")
			options.Default = str
		}
		res[key] = options
	}
	return res, nil
}

func (msf *Msfrpc) ModuleCompatiblePayloads(ModuleName string) ([]string, error) {
	req := &models.ModuleCompatiblePayloadsReq{Method: models.ModuleCompatiblePayloads, Token: msf.token, ModName: ModuleName}
	var res models.ModuleCompatiblePayloadsRes
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res.Payloads, nil
}

func (msf *Msfrpc) ModuleTargetCompatiblePayloads(ModuleName string, Target int) ([]string, error) {
	req := &models.ModuleTargetCompatiblePayloadsReq{Method: models.ModuleTargetCompatiblePayloads, Token: msf.token, ModName: ModuleName, Target: Target}
	var res models.ModuleTargetCompatiblePayloadsRes
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res.Payloads, nil
}

func (msf *Msfrpc) ModuleCompatibleSessions(ModuleName string) ([]int, error) {
	req := &models.ModuleCompatibleSessionsReq{Method: models.ModuleCompatibleSessions, Token: msf.token, ModName: ModuleName}
	var res models.ModuleCompatibleSessionsRes
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	return res.Sessions, nil
}

func (msf *Msfrpc) ModuleExecute(ModuleType string, ModuleName string, options map[string]interface{}) (models.ModuleExecuteRes, error) {
	req := &models.ModuleExecuteReq{Method: models.ModuleExecute, Token: msf.token, ModType: ModuleType, ModName: ModuleName, Options: options}
	var res models.ModuleExecuteRes
	if err := msf.send(req, &res); err != nil {
		return models.ModuleExecuteRes{}, err
	}
	return res, nil
}
