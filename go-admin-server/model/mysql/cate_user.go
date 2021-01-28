package mysql

import (
	"github.com/jinzhu/gorm"
)
// 管理员信息表
type Cate_user struct {
	gorm.Model
	Username string `gorm:"not null;type:varchar(64);comment:'用户名'"`
	Password string `gorm:"not null;type:varchar(64);comment:'密码'"`
	Email string `gorm:"not null;type:varchar(64);comment:'邮箱'"`
	Nickname string `gorm:"not null;type:varchar(64);comment:'别名'"`
	Avatar string `gorm:"not null;type:varchar(255);comment:'头像'"`
}
