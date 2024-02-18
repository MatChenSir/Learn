package test

type JSONData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  Result `json:"result"`
}
type Result struct {
	AppID        int    `json:"appId"`
	AccessToken  string `json:"accessToken"`
	ExpireTime   int    `json:"expireTime"`
	RefreshToken string `json:"refreshToken"`
}

type User struct {
	ID            int64  `gorm:"primary_key;column:id" json:"id,omitempty"`
	Username      string `gorm:"column:username" json:"username"`
	Role          string `gorm:"column:role" json:"role"`
	Authorization string `gorm:"column:authorization" json:"authorization,omitempty"`
	Deleted       bool   `gorm:"column:deleted" json:"deleted,omitempty"`
	CreateTime    int64  `gorm:"column:create_time" json:"create_time,omitempty"`
	UpdateTime    int64  `gorm:"column:update_time" json:"update_time,omitempty"`
}
