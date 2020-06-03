package models

type ErrorRes struct {
	Error        bool   `msgpack:"error,omitempty"`
	ErrorClass   string `msgpack:"error_class,omitempty"`
	ErrorMessage string `msgpack:"error_message,omitempty"`
}
