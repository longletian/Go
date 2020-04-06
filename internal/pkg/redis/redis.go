package redis

import (
	"demoProject/config"
	"github.com/go-redis/redis"
)

func InitRedis() (*redis.Client, error){
	var  reconnect *redis.Client
	config,err:=config.GetConfig()
	if  err ==nil{
		if config.RedisConfig.SentryContect {
			//哨兵模式
			reconnect = redis.NewFailoverClient(&redis.FailoverOptions{
				MasterName:  config.RedisConfig.SentryConfig.MasterName,
				SentinelAddrs: config.RedisConfig.SentryConfig.SentinelAddrs,
			})
		} else {
			reconnect = redis.NewClient(&redis.Options{
				Addr:   config.RedisConfig.Address,
				Password: config.RedisConfig.Password,
				DB: config.RedisConfig.Database,
			})
		}
		//ping一下是否连接
		_, err := reconnect.Ping().Result()
		if err == nil {
			return reconnect,nil
		}
	}
	return reconnect,err
}

func SetData(key string, data ...interface{}) bool {
    flag :=false
	reconnect,err:= InitRedis()
	if err == nil{
	    err = reconnect.Set(key,data,0).Err()
		if err ==nil {
			flag= true
		}
	}
	return  flag
}

func GetData(key string) string{
	reconnect,err:= InitRedis()
	if err == nil{
		val, err := reconnect.Get(key).Result()
		if err == nil {
			return  val
		}
	}
	return ""
}

func  DeleteData(key string)  bool {
	flag :=false
	reconnect,err:= InitRedis()
	if err == nil{
		err = reconnect.Del(key).Err()
		if err ==nil{
			flag=true
		}
	}
	return flag
}





