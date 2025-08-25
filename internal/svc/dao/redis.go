package dao

import (
	"context"
	"github.com/greyireland/log"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	R *redis.Client
}

func NewRedis(url string) *Redis {
	//url := "redis://user:password@localhost:6379/0?protocol=3"
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	r := redis.NewClient(opts)
	return &Redis{R: r}
}
func (d *Redis) Push(key, value string) {
	ret := d.R.RPush(context.Background(), key, value)
	if ret.Err() != nil {
		log.Warn("LPush error", "err", ret.Err())
	}
	//log.Info("LPush", "key", key, "value", value)
}
func (d *Redis) Peek(key string) string {
	if d.R.LLen(context.Background(), key).Val() == 0 {
		return ""
	}
	ret := d.R.LIndex(context.Background(), key, 0)
	if ret.Err() != nil {
		log.Warn("HSet error", "err", ret.Err())
	}
	return ret.Val()
}
func (d *Redis) Pop(key string) string {
	ret := d.R.LPop(context.Background(), key)
	if ret.Err() != nil {
		log.Warn("HSet error", "err", ret.Err())
	}
	//log.Info("LPop", "key", key, "value", ret.Val())
	return ret.Val()
}
func (d *Redis) Get(key string) string {
	ret := d.R.Get(context.Background(), key)
	if ret.Err() != nil {
		return ""
	}
	return ret.Val()
}
func (d *Redis) Set(key, val string) error {
	ret := d.R.Set(context.Background(), key, val, 0)
	if ret.Err() != nil {
		log.Warn("Set error", "err", ret.Err())
		return ret.Err()
	}
	return nil
}
func (d *Redis) SetExpire(key, val string, dur time.Duration) error {
	ret := d.R.Set(context.Background(), key, val, dur)
	if ret.Err() != nil {
		log.Warn("SetExpire error", "err", ret.Err())
		return ret.Err()
	}
	return nil
}

// Del delete key
func (d *Redis) Del(key string) error {
	ret := d.R.Del(context.Background(), key)
	if ret.Err() != nil {
		log.Warn("Del error", "err", ret.Err())
		return ret.Err()
	}
	return nil
}

// HGet hget key field
func (d *Redis) HGet(key, field string) string {
	ret := d.R.HGet(context.Background(), key, field)
	if ret.Err() != nil {
		//log.Warn("HGet error", "err", ret.Err())
		return ""
	}
	return ret.Val()
}

func (d *Redis) HSet(key string, values []interface{}) error {
	ret := d.R.HSet(context.Background(), key, values...)
	if ret.Err() != nil {
		log.Warn("HSet error", "err", ret.Err())
		return ret.Err()
	}
	return nil
}
func (d *Redis) HGetAll(key string) map[string]string {
	ret := d.R.HGetAll(context.Background(), key)
	if ret.Err() != nil {
		log.Warn("HGetAll error", "err", ret.Err())
		return map[string]string{}
	}
	return ret.Val()
}
func (d *Redis) HDel(key string, field string) error {
	ret := d.R.HDel(context.Background(), key, field)
	if ret.Err() != nil {
		log.Warn("HDel error", "err", ret.Err())
		return ret.Err()
	}
	return nil
}
