package models

type ConsoleCreateReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type ConsoleCreateRes struct {
	ID 			uint64	`msgpack:"id"`
	Prompt 		string	`msgpack:"prompt"`
	Busy 		bool	`msgpack:"busy"`
}

type ConsoleDestroyReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleID string
}

type ConsoleDestroyRes struct {
	Result 		string	`msgpack:"result"`
}

type ConsoleListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

//ConsoleListRes ???

type ConsoleWriteReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleID string
	Data      string
}

type ConsoleWriteRes struct {
	Wrote 		uint64	`msgpack:"wrote"`
}

type ConsoleReadReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleID string
}

type ConsoleReadRes struct {
	Data 		string	`msgpack:"data"`
	Prompt 		string	`msgpack:"prompt"`
	Busy 		bool	`msgpack:"busy"`
}

type ConsoleSessionDetachReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleID string
}

type ConsoleSessionDetachRes struct {
	Result 		string	`msgpack:"result"`
}

type ConsoleSessionKillReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleID string
}

type ConsoleSessionKillRes struct {
	Result 		string	`msgpack:"result"`
}

type ConsoleTabsReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	ConsoleID string
	InputLine string
}

type ConsoleTabsRes struct {
	Tabs 		[]string	`msgpack:"tabs"`
}




