package common

type DBError struct {
	Msg string
	Err error
}

type ServerError struct {
	Msg    string
	Status int
}

func NewDBError(msg string, err error) *DBError {
	return &DBError{Msg: msg, Err: err}
}

func NewServerError(msg string, status int) *ServerError {
	return &ServerError{Msg: msg, Status: status}
}
