package models

type ModuleExploitsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type ModulePayloadsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type ModuleAuxiliaryReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type ModuleEncodersReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type ModulePostReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type ModuleNopsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type ModuleInfoReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	ModType  string
	ModName  string
}

type ModuleInfoRes struct {
	Type			string 		`msgpack:"type"`
	Name 			string 		`msgpack:"name"`
	FullName 		string 		`msgpack:"fullname"`
	Rank 			string 		`msgpack:"rank"`
	DisClosureDate	string 		`msgpack:"disclosuredate"`
	Description 	string 		`msgpack:"description"`
	License 		string 		`msgpack:"license"`
	Filepath 		string 		`msgpack:"filepath"`
	Arch			[]string	`msgpack:"arch"`
	Platforms		[]string 	`msgpack:"platform"`
	Authors 		[]string 	`msgpack:"authors"`
	Privileged		bool 		`msgpack:"privileged"`
	References 		[][]interface{} `msgpack:"references"`
	Targets 		map[int]string `msgpack:"targets"`
	DefaultTarget	uint64	`msgpack:"default_target"`
	Stance			string 	`msgpack:"stance"`
	Options 		map[string]interface{}	`msgpack:"options"`
}

type ModuleOptionsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	ModType  string
	ModName  string
}

type ModuleOptionsRes struct {
	Type 		string		`msgpack:"type"`
	Required    bool		`msgpack:"required"`
	Advanced 	bool		`msgpack:"advanced"`
	Evasion 	bool 		`msgpack:"evasion"`
	Desc 		string 		`msgpack:"desc"`
	Default 	interface{} `msgpack:"default,omitempty"`
	Enums		[]string	`msgpack:"enums,omitempty"`
}

type ModuleCompatiblePayloadsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	ModName  string
}

type ModuleCompatiblePayloadsRes struct {
	Payloads []string `msgpack:"payloads"`
}

type ModuleTargetCompatiblePayloadsReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	ModName  string
	Target   uint64
}

type ModuleTargetCompatiblePayloadsRes struct {
	Payloads []string `msgpack:"payloads"`
}

type ModuleExecuteReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	ModType  string
	ModName  string
	Option   map[string]interface{}
}

type ModuleExecuteRes struct {
	JobID	uint64	`msgpack:"job_id,omitempty"`
	UUID    string  `msgpack:"uuid"`
}
















