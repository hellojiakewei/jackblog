package initconf

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"jackblog/conf"
	"jackblog/model/mysql"
)

func Db() *gorm.DB {
	fmt.Println(conf.BbUserName)
	fmt.Println(conf.DbPassWord)
	fmt.Println(conf.DbName)
	fmt.Println(conf.DbPath)
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", conf.BbUserName, conf.DbPassWord, conf.DbPath, conf.DbName, conf.DbConfig))
	if err != nil {
		panic("数据库连接失败")
	}
	db.SingularTable(true)
	db.AutoMigrate(&mysql.Cate_content{},&mysql.Cate_catalog{},&mysql.Cate_article{},&mysql.Cate_type{},&mysql.Cate_user{})
	return db
}
