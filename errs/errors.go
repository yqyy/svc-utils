package errs

import "fmt"

type commonErr struct {
	code int
	msg  string
}

func (e commonErr) Error() string {
	return fmt.Sprintf("code:%d,msg:%v", e.code, e.msg)
}

func New(code int, msg string) error {
	return commonErr{
		code: code,
		msg:  msg,
	}
}

func Is(err error) bool {
	_, ok := err.(commonErr)
	return ok
}

func GetCode(err error) int {
	if e, ok := err.(commonErr); ok {
		return e.code
	}

	return -1
}

func GetMsg(err error) string {
	if e, ok := err.(commonErr); ok {
		return e.msg
	}

	return err.Error()
}
