package data

import (
	"github.com/gomodule/redigo/redis"
	"plan-to-remind/internal/conf"
)

type RedisClient struct {
	conn redis.Conn
}

func NewRedisClient(data *conf.Data) *RedisClient {
	conn, err := redis.Dial("tcp", data.Redis.Addr)
	if err != nil {
		panic(err)
	}
	return &RedisClient{conn: conn}
}

func (cli *RedisClient) SetEx(key string, value interface{}, dur uint32) error {
	_, err := cli.conn.Do("SETEX", key, dur, value)
	return err
}

func (cli *RedisClient) SetNx(key string, value interface{}, dur uint32) (bool, error) {
	res, err := redis.Int64(cli.conn.Do("SETNX", key, value))
	if err != nil {
		return false, err
	}
	if res == 0 {
		return false, nil
	}
	// set expire
	_, err = cli.conn.Do("EXPIRE", key, dur)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cli *RedisClient) Get(key string) (string, error) {
	return redis.String(cli.conn.Do("GET", key))
}

func (cli *RedisClient) Close() {
	_ = cli.conn.Close()
}
