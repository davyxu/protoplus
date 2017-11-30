package model

type Comment struct {
	Leading  string `json:",omitempty"`
	Trailing string `json:",omitempty"`
}
