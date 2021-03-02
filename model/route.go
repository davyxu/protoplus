package model

// 路由规则
type RouteRule struct {
	MsgName string
	MsgID   int

	SvcName string
	Router  string `json:",omitempty"`
}

// 路由表，包含多条路由规则
type RouteTable struct {
	Rule []*RouteRule
}
