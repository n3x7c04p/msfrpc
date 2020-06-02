package models

type SessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type SessionListRes struct {
	ID          int    `msgpack:",omitempty"`
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"Session_host"`
	SessionPort int    `msgpack:"Session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid"`
}

type SessionStopReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
}

type SessionStopRes struct {
	Result string `msgpack:"result"`
}

type SessionCompatibleModulesReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
}

type SessionCompatibleModulesRes struct {
	Modules []string `msgpack:"modules"`
}

type SessionShellReadReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
}

type SessionShellReadRes struct {
	Seq  int    `msgpack:"seq"`
	Data string `msgpack:"data"`
}

type SessionShellWriteReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
	Data      string
}

type SessionShellWriteRes struct {
	WriteCount int `msgpack:"write_count"`
}

type SessionShellUpgradeReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
	ConnHost  string
	ConnPort  int
}

type SessionShellUpgradeRes struct {
	Result string `msgpack:"result"`
}

type SessionMeterpreterReadReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
}

type SessionMeterpreterReadRes struct {
	Data string `msgpack:"data"`
}

type SessionMeterpreterWriteReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
	Data      string
}

type SessionMeterpreterWriteRes struct {
	Result string `msgpack:"result"`
}

type SessionMeterpreterTabsReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
	InputLine string
}

type SessionMeterpreterTabsRes struct {
	Tabs []string `msgpack:"result"`
}

type SessionMeterpreterRunSingleReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
	Data      string
}

type SessionMeterpreterRunSingleRes struct {
	Result string `msgpack:"result"`
}

type SessionMeterpreterScriptReq struct {
	_msgpack   struct{} `msgpack:",asArray"`
	Method     string
	Token      string
	SessionID  string
	ScriptName string
}

type SessionMeterpreterScriptRes struct {
	Result string `msgpack:"result"`
}

type SessionMeterpreterSessionDetachReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
}

type SessionMeterpreterSessionDetachRes struct {
	Result string `msgpack:"result"`
}

type SessionMeterpreterSessionKillReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
}

type SessionMeterpreterSessionKillRes struct {
	Result string `msgpack:"result"`
}

type SessionRingClearReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
}

type SessionRingClearRes struct {
	Result string `msgpack:"result"`
}

type SessionRingLastReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
}

type SessionRingLastRes struct {
	Seq int `msgpack:"seq"`
}

type SessionRingPutReq struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Method    string
	Token     string
	SessionID string
	Data      string
}

type SessionRingPutRes struct {
	WriteCount int `msgpack:"write_count"`
}
