package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode    int32 = 0
	ServiceErrCode       = iota + 10000
	ParamErrCode
	AuthorizationFailedErrCode

	UserAlreadyExistErrCode
	UserIsNotExistErrCode
)

const (
	SuccessMsg    = "Success"
	ServiceErrMsg = "Service is unable to start successfully"

	PasswordIsNotVerifiedMsg = "username or password not verified"
	UserIsNotExistErrMsg     = "User does not exist"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (err ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", err.ErrCode, err.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (err ErrNo) WithMessage(msg string) ErrNo {
	err.ErrMsg = msg
	return err
}

var (
	Success    = NewErrNo(SuccessCode, SuccessMsg)
	ServiceErr = NewErrNo(ServiceErrCode, ServiceErrMsg)

	UserIsNotExistErr    = NewErrNo(UserIsNotExistErrCode, UserIsNotExistErrMsg)
	PasswordIsNotVerfied = NewErrNo(AuthorizationFailedErrCode, PasswordIsNotVerifiedMsg)
)

func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
