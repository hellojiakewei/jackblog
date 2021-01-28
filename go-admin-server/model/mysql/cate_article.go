package mysql

import "github.com/jinzhu/gorm"

//文章

type Cate_article struct {
	gorm.Model
	Title string `gorm:"not null;type:varchar(64);comment:'文章标题'"`
	Author int `gorm:"not null;type:tinyint;comment:'作者'"`
	Like int `gorm:"not null;type:tinyint;comment:'文章点赞'"`
	Down int `gorm:"not null;type:tinyint;comment:'文章踩'"`
	ContentId int `gorm:"not null;type:int"`
	Cate_content Cate_content `gorm:"foreignKey:ContentId"`
	//CataLogId []Cate_catalog `gorm:"many2many:cate_catalog"`
}