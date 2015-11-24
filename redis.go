package main

import "github.com/garyburd/redigo/redis"

type Redis struct {
	conn redis.Conn
}

func (r *Redis) SCARD(key string) (int, error) {
	return redis.Int(r.conn.Do("SCARD", key))
}

func (r *Redis) SMEMBERS(key string) ([]string, error) {
	return redis.Strings(r.conn.Do("SMEMBERS", key))
}

func (r *Redis) SRANDMEMBER(key string) (string, error) {
	return redis.String(r.conn.Do("SRANDMEMBER", key))
}

func (r *Redis) SADD(key, member string) (int, error) {
	return redis.Int(r.conn.Do("SADD", key, member))
}

func (r *Redis) SREM(key, member string) (int, error) {
	return redis.Int(r.conn.Do("SREM", key, member))
}

func (r *Redis) DEL(key string) (int, error) {
	return redis.Int(r.conn.Do("DEL", key))
}

func NewRedis(conn redis.Conn) *Redis {
	return &Redis{conn}
}
