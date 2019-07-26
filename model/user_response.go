package model

type UserRes struct {
	Result Result
	Token  string
}

type Result struct {
	Code MsgCode
	Msg  string
}

type MsgCode int32

const (
	MsgCode_MsgCode_FAIL = iota
	MsgCode_MsgCode_SUCCESS
)
