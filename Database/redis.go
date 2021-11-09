package Database

import (
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
)

//redis链接
func RedisConnect() ( redis.Conn, error) {
	c, err := redis.Dial("tcp","localhost:6379")
	if err != nil {
		log.Error("redis:redis connection failed!")
		return nil, err
	}
	log.Print("redis:redis connection success!")
	return c, nil
}
