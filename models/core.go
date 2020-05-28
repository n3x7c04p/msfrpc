package models

type CoreAddModulePathReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	Path     string
}


type CoreModuleLengthRes struct {
	LenExploits 	uint64	`msgpack:"exploits"`
	LenAuxiliarys 	uint64	`msgpack:"auxiliary"`
	LenPosts 		uint64	`msgpack:"posts"`
	LenEncoders 	uint64	`msgpack:"encoders"`
	LenNops 		uint64	`msgpack:"nops"`
	LenPayloads 	uint64	`msgpack:"payloads"`
}

type CoreModuleStatsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreReloadModulesReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreSaveReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreSaveRes struct {
	Result 		string `msgpack:"result"`
}

type CoreSetGReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	OptName  string
	OptValue string
}

type CoreSetGRes struct {
	Result 		string `msgpack:"result"`
}

type CoreUnsetGReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	OptName  string
}

type CoreUnsetGRes struct {
	Result 		string `msgpack:"result"`
}

type CoreThreadListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreThreadListRes struct {
	Status 		string 	`msgpack:"status"`
	Critical 	bool 	`msgpack:"critical"`
	Name 		string 	`msgpack:"name"`
	Started 	string 	`msgpack:"started"`
}

type CoreThreadKillReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreThreadKillRes struct {
	Result 		string `msgpack:"result"`
}

type CoreVersionReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreVersionRes struct {
	Version 		string `msgpack:"version"`
	RubyVersion 	string `msgpack:"ruby"`
	ApiVersion 		string `msgpack:"api"`
}

type CoreStopReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type CoreStopRes struct {
	Result 		string `msgpack:"result"`
}

