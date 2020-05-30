package models

type JobListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type JobListRes struct {
	ID     int    `msgpack:",omitempty"`
	Module string `msgpack:",omitempty"`
}

type JobInfoReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	JobID    string
}

type JobInfoRes struct {
	JobID     int                    `msgpack:"jid"`
	Name      string                 `msgpack:"name"`
	StartTime int                    `msgpack:"start_time"`
	UriPath   string                 `msgpack:"uripath"`
	DataStore map[string]interface{} `msgpack:"datastore"`
}

type JobStopReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	JobID    string
}

type JobStopRes struct {
	Result string `msgpack:"result"`
}
