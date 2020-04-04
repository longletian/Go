package config

//配置文件类,读取配置文件的信息

type  Config struct {
	MySqlConfig MySqlConfig `json:"mysql_config"`
	RedisConfig RedisConfig `json:"redis_config"`
	MongoConfig MongoConfig `json:"mongo_config"`
	VerifyCodeConfig VerifyCodeConfig `json:"verify_code_config"`
	LogConfig LogConfig  `json:"log_config"`
}

//日志配置文件
type LogConfig struct {
	//日志文件的位置
  	Path  string
	//在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxSize   int
	//保留旧文件的最大个数
	MaxBackups int
	//保留旧文件的最大天数
	MaxAge   int
	//是否压缩/归档旧文件
	Compress bool
}


//数据库配置
type  MySqlConfig struct {
	Username string
	Password string
	Databasename string
	Charsettype  string
}

//redis配置
type RedisConfig struct {

}
//mongodb配置
type  MongoConfig struct {
	
}
//验证码配置
type VerifyCodeConfig struct {

}

//初始化获取配置文件
func  GetConfig() *Config {
    config :=Config{
		MySqlConfig:      MySqlConfig{},
		RedisConfig:      RedisConfig{},
		MongoConfig:      MongoConfig{},
		VerifyCodeConfig: VerifyCodeConfig{},
		LogConfig:        LogConfig{
			Path:"",
		},
	}
    return  &config
}