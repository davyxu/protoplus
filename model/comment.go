package model

// 注释
type Comment struct {
	// 头注释
	Leading string `json:",omitempty"`

	// 尾注释
	Trailing string `json:",omitempty"`
}
