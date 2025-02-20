package model

type ServerErr struct {
	location []string
	err      error
}

func (e *ServerErr) Error() string {
	return e.err.Error()
}

type ConnectionErr struct {
	location string
	err      error
}

func (e *ConnectionErr) Error() string {
	return e.err.Error()
}
