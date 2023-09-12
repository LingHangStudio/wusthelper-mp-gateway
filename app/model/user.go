package model

type UserBasic struct {
	Uid         int64  `json:"uid,omitempty"`          // uid
	WxOid       string `json:"wx_oid,omitempty"`       // 微信openid
	QqOid       string `json:"qq_oid,omitempty"`       // qq openid
	Sid         string `json:"sid,omitempty"`          // 学号
	Type        int8   `json:"type,omitempty"`         // 用户类型，0：本科生，1：研究生
	OfficialPwd string `json:"official_pwd,omitempty"` // 教务处密码
	LibPwd      string `json:"lib_pwd,omitempty"`      // 图书馆密码
	PhysicsPwd  string `json:"physics_pwd,omitempty"`  // 物理实验系统密码
	CreateTime  string `json:"create_time,omitempty"`
	UpdateTime  string `json:"update_time,omitempty"`
	Deleted     int8   `json:"deleted,omitempty"`
}

func (UserBasic) TableName() string {
	return "user_basic"
}

type QQUserProfile struct {
	Oid        string `json:"oid,omitempty"` // openid
	Nickname   string `json:"nickname,omitempty"`
	Gender     int32  `json:"gender,omitempty"`
	City       string `json:"city,omitempty"`
	Province   string `json:"province,omitempty"`
	Country    string `json:"country,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
	Deleted    int8   `json:"deleted,omitempty"`
}

func (QQUserProfile) TableName() string {
	return "qq_profile"
}

type WxUserProfile struct {
	Oid        string `json:"oid,omitempty"` // openid
	Nickname   string `json:"nickname,omitempty"`
	Gender     int32  `json:"gender,omitempty"`
	City       string `json:"city,omitempty"`
	Province   string `json:"province,omitempty"`
	Country    string `json:"country,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
	Deleted    int8   `json:"deleted,omitempty"`
}

func (WxUserProfile) TableName() string {
	return "wx_profile"
}
