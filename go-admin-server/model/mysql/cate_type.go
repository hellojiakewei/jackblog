package mysql

import "github.com/jinzhu/gorm"
// 文章类型
type Cate_type struct {
	gorm.Model
	Name string `gorm:"not null;type:varchar(100)"`
}
