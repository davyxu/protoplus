package model

type ServiceCall struct {
	Comment
	TagSet

	IsRPC bool

	Name        string // 函数名
	RequestName string // 请求消息名
	RespondName string // 回应消息名
}
