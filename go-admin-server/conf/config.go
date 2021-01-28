package conf

import "github.com/spf13/viper"

var (
	WebPort string
	DbPath string
	DbConfig string
	DbName string
	BbUserName string
	DbPassWord string
)

type Config struct {
	PathName string
	V *viper.Viper
}

func (c *Config) LoadYaml(){
	c.V = viper.New()
	c.V.AddConfigPath("conf")
	c.V.SetConfigName(c.PathName)
	c.V.SetConfigType("yaml")
	err := c.V.ReadInConfig()
	if err !=nil {
		panic("文件读取失败")
	}
	port(c.V)
	mysql(c.V)
}
func port(v *viper.Viper)  {
	WebPort = v.GetString("system.port")
}

func mysql(v *viper.Viper)  {
	DbPath = v.GetString("mysql.path")
	DbConfig = v.GetString("mysql.config")
	DbName = v.GetString("mysql.dbName")
	BbUserName = v.GetString("mysql.username")
	DbPassWord = v.GetString("mysql.password")
}
