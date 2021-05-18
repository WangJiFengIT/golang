package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

func initClinet() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         "192.168.10.46:6379",
		Password:     "21cnjycom",
		PoolSize:     10000,
		DB:           15,
		ReadTimeout:  time.Millisecond * time.Duration(100),
		WriteTimeout: time.Millisecond * time.Duration(100),
		IdleTimeout:  time.Second * time.Duration(60),
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	} else {
		fmt.Println("redis connection success")
	}
	return nil
}

func get(key string) (string, bool) {
	r, err := rdb.Get(key).Result()
	if err != nil {
		return "get failed", false
	}
	return r, true
}

func set(key string, val string, expTime int32) {
	err := rdb.Set(key, val, time.Duration(expTime)*time.Second)
	if err != nil {
		return
	}
}

func remove(key string) {
	err := rdb.Del(key)
	if err != nil {
		return
	}
}

//func main() {
//	var name = "GQROUP"
//	initClinet()
//	set(name, "6379", 100)
//	s, b := get(name)
//	fmt.Println(s, b)
//	//remove(name)
//}
