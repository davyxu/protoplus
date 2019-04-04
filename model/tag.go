package model

import "strconv"

// 标记
type Tag struct {
	Key   string
	Value string
}

// 标记集合
type TagSet struct {
	Tags []Tag `json:",omitempty"`
}

func (self *TagSet) TagValueByKey(key string) (string, bool) {
	for _, tag := range self.Tags {
		if tag.Key == key {
			return tag.Value, true
		}
	}

	return "", false
}

func (self *TagSet) SetTagValue(key, value string) {

	for index, tag := range self.Tags {
		if tag.Key == key {
			self.Tags[index].Value = value
			break
		}
	}
}

func (self *TagSet) TagExists(key string) bool {
	for _, tag := range self.Tags {
		if tag.Key == key {
			return true
		}
	}
	return false
}

func (self *TagSet) TagValueInt(key string) int {
	if v, ok := self.TagValueByKey(key); ok {
		if r, err := strconv.Atoi(v); err == nil {
			return r
		}
	}

	return 0
}

func (self *TagSet) TagValueString(key string) string {
	v, _ := self.TagValueByKey(key)
	return v
}

func (self *TagSet) TagValueBool(key string) bool {
	if v, ok := self.TagValueByKey(key); ok {
		if r, err := strconv.ParseBool(v); err == nil {
			return r
		}
	}

	return false
}

func (self *TagSet) AddTag(tag Tag) {
	self.Tags = append(self.Tags, tag)
}
