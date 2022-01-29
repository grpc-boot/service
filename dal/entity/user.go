package entity

import "time"

type User struct {
	Id           int64  `borm:"id,primary"`
	NickName     string `borm:"nickname"`
	Password     string `borm:"passwd"`
	Phone        string `borm:"phone"`
	Email        string `borm:"email"`
	WxUuid       string `borm:"wx_uuid"`
	WxPubOpenId  string `borm:"wx_pub_openid"`
	WxSappOpenId string `borm:"wx_sapp_openid"`
	LastLoginAt  int64  `borm:"last_login_at"`
	IsOn         int8   `borm:"is_on"`
	CreatedAt    int64  `borm:"created_at"`
	UpdatedAt    int64  `borm:"updated_at"`
}

/*func (u *User) TableName() string {
	return `user`
}*/

func (u *User) BeforeCreate() {
	u.CreatedAt = time.Now().Unix()
	u.LastLoginAt = u.CreatedAt
}

func (u *User) BeforeSave() {
	u.UpdatedAt = time.Now().Unix()
}
