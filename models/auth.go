package models

type AuthLoginReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

type AuthLoginRes struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error,omitempty"`
	ErrorClass   string `msgpack:"error_class,omitempty"`
	ErrorMessage string `msgpack:"error_message,omitempty"`
}

type AuthLogoutReq struct {
	_msgpack        struct{} `msgpack:",asArray"`
	Method          string
	Token           string
	AuthLogoutToken string
}

type AuthLogoutRes struct {
	Result string `msgpack:"result"`
}

type AuthTokenAddReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	NewToken string
}

type AuthTokenAddRes struct {
	Result 		string `msgpack:"result"`
}

type AuthTokenGenerateReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
	NewToken string
}

type AuthTokenGenerateRes struct {
	Result 		string `msgpack:"result"`
	Token 		string `msgpack:"token"`
}

type AuthTokenListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type AuthTokenListRes struct {
	Tokens 		[]string `msgpack:"tokens"`
}

type AuthTokenRemoveReq struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	TokenRemove string
}

type AuthTokenRemoveRes struct {
	Result 		string `msgpack:"result"`
}

