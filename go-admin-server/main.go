package main

import (
	"flag"
	"jackblog/conf"
	"jackblog/global"
	"jackblog/initconf"
	"jackblog/utils"
)

var (
	env = flag.String("env","dev","本地开发环境") // ./main --env=dev
)
func main()  {
	flag.Parse()
	var c conf.Config
	c.PathName = *env
	c.LoadYaml()
	global.DB = initconf.Db()
	initconf.Logger()
	if err := utils.InitTrans("zh"); err != nil {
		return
	}
	initconf.InitRouter()
}
