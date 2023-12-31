package model

import "time"

type Student struct {
	Sid        string    `json:"sid,omitempty"`     // 学号
	Name       string    `json:"name,omitempty"`    // 姓名
	College    string    `json:"college,omitempty"` // 学院
	Major      string    `json:"major,omitempty"`   // 专业
	Clazz      string    `json:"clazz,omitempty"`   // 班级
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
	Deleted    int8      `json:"deleted,omitempty"`
}

func (Student) TableName() string {
	return "student"
}
