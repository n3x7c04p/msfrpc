package msfrpc

import "msfrpc/models"

//jobs
func (msf *Msfrpc) JobList() (map[int]models.JobListRes, error) {
	req := &models.JobListReq{Method: models.JobList, Token: msf.token}
	res := make(map[int]models.JobListRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	for id, module := range res {
		res[id] = module
	}
	return res, nil
}

func (msf *Msfrpc) JobInfo(JobID string) (models.JobInfoRes, error) {
	req := &models.JobInfoReq{Method: models.JobInfo, Token: msf.token, JobID: JobID}
	var res models.JobInfoRes
	if err := msf.send(req, &res); err != nil {
		return models.JobInfoRes{}, err
	}
	return res, nil
}

func (msf *Msfrpc) JobStop(JobID string) (bool, error) {
	req := &models.JobStopReq{Method: models.JobStop, Token: msf.token, JobID: JobID}
	var res models.JobStopRes
	if err := msf.send(req, &res); err != nil {
		return false, err
	}
	return res.Result == "success", nil
}
