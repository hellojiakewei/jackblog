package mysql

import "github.com/jinzhu/gorm"

// 文章内容
type Cate_content struct {
	gorm.Model
	Content string  `gorm:"not null;type:text"`
}