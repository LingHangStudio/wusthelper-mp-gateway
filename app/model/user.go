package model

import "time"

const (
	UndergradUser = 0
	GraduateUser  = 1
)

type UserBasic struct {
	Uid         uint64    `json:"uid,omitempty"`          // uid
	Oid         string    `json:"oid,omitempty"`          // 用户唯一标识，openid
	Sid         string    `json:"sid,omitempty"`          // 学号
	Unionid     string    `json:"unionid,omitempty"`      // 用户开放平台的唯一标识符
	Platform    uint8     `json:"platform,omitempty"`     // 用户平台，0：微信，1：QQ
	Type        int8      `json:"type,omitempty"`         // 用户类型，0：本科生，1：研究生
	OfficialPwd string    `json:"official_pwd,omitempty"` // 教务处密码
	LibPwd      string    `json:"lib_pwd,omitempty"`      // 图书馆密码
	PhysicsPwd  string    `json:"physics_pwd,omitempty"`  // 物理实验系统密码
	CreateTime  time.Time `json:"create_time,omitempty"`
	UpdateTime  time.Time `json:"update_time,omitempty"`
	Deleted     int8      `json:"deleted,omitempty"`
}

func (UserBasic) TableName() string {
	return "user_basic"
}

type QQUserProfile struct {
	Oid        string    `json:"oid,omitempty"` // openid
	Nickname   string    `json:"nickname,omitempty"`
	Gender     int32     `json:"gender,omitempty"`
	City       string    `json:"city,omitempty"`
	Province   string    `json:"province,omitempty"`
	Country    string    `json:"country,omitempty"`
	Avatar     string    `json:"avatar,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
	Deleted    int8      `json:"deleted,omitempty"`
}

func (QQUserProfile) TableName() string {
	return "qq_profile"
}

type WxUserProfile struct {
	Oid        string    `json:"oid,omitempty"` // openid
	Nickname   string    `json:"nickname,omitempty"`
	Gender     int32     `json:"gender,omitempty"`
	City       string    `json:"city,omitempty"`
	Province   string    `json:"province,omitempty"`
	Country    string    `json:"country,omitempty"`
	Avatar     string    `json:"avatar,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	UpdateTime time.Time `json:"update_time,omitempty"`
	Deleted    int8      `json:"deleted,omitempty"`
}

func (WxUserProfile) TableName() string {
	return "wx_profile"
}
