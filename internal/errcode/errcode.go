package errcode

import (
	"fmt"
	"net/http"
)

var (
	UnknownError     = NewError(10000, "伺服器錯誤")
	InvalidParams    = NewError(10001, "輸入參數錯誤")
	NotFound         = NewError(10002, "頁面不存在")
	DuplicateRecords = NewError(10003, "已存在相同的記錄")
	RecordNotExists  = NewError(10004, "記錄不存在")
)

var ErrorList = map[int]string{}

type Error struct {
	code int
	msg  string
}

func NewError(code int, msg string) *Error {
	if _, ok := ErrorList[code]; ok {
		panic(fmt.Sprintf("錯誤 %d 已經存在", code))
	}
	ErrorList[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Message() string {
	return e.msg
}

func (e *Error) StatusCode() int {
	switch e.code {
	case UnknownError.code:
		return http.StatusInternalServerError
	case NotFound.code:
		return http.StatusNotFound
	case InvalidParams.code:
		fallthrough
	case DuplicateRecords.code:
		fallthrough
	case RecordNotExists.code:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("錯誤: %d, 錯誤訊息: %s", e.Code(), e.Message())
}
