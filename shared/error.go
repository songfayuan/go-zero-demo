package shared

import (
	"errors"
	"net/http"
)

const (
	defaultCode = "50001"
)

var (
	ErrNotFound               = errors.New("data not find")
	ErrorUserNotFound         = errors.New("用户不存在")
	ErrorNoRequiredParameters = errors.New("必要参数不能为空")
	ErrorUserOperation        = errors.New("用户正在操作中，请稍后重试")

	DbError = errors.New("db error")

	ErrorNotSelect        = errors.New("未选中记录")
	ErrorRecordInvalid    = errors.New("无效的记录")
	ErrorIdsRecordInvalid = errors.New("所选记录无效，请刷新后再操作")
	ErrorOperateTooFast   = errors.New("操作频繁，请稍后再试")
)

type (
	CodeError struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func ErrorHandler(err error) (int, interface{}) {
	return http.StatusConflict, CodeError{
		Code:    "50000",
		Message: err.Error(),
	}
}

func NewCodeError(code string, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Message
}
