package mysql

import "github.com/jinzhu/gorm"

// 文章栏目
type Cate_catalog struct {
	gorm.Model
	Name string `gorm:"not null;type:varchar(100)"`
}
