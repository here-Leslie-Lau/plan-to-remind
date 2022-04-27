package limiter

import (
	"errors"
	"plan-to-remind/internal/data"
)

type UserRequestLimiter interface {
	Limit(key string) error
}

type redisRequestLimiter struct {
	opts *option
	rdb  *data.RedisClient
}

type option struct {
	// limit expire time
	expire uint32
}

type Option func(*option)

func WithExpire(expire LimitExpireTyp) Option {
	return func(o *option) {
		o.expire = uint32(expire)
	}
}

func newRedisRequestLimiter(data *data.Data, opts ...Option) UserRequestLimiter {
	o := new(option)
	for _, opt := range opts {
		opt(o)
	}
	return &redisRequestLimiter{rdb: nil, opts: o}
}

func (limiter *redisRequestLimiter) Limit(key string) error {
	ok, err := limiter.rdb.SetNx(key, "request", limiter.opts.expire)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("frequent user requests")
	}
	return nil
}

type Pool struct {
	m map[LimitExpireTyp]UserRequestLimiter
}

var RequestPool *Pool

func InitPool(data *data.Data) {
	m := make(map[LimitExpireTyp]UserRequestLimiter)
	for _, expire := range limitExpireTimes {
		m[expire] = newRedisRequestLimiter(data, WithExpire(expire))
	}
	RequestPool = &Pool{m: m}
}

func (p *Pool) GetLimiter(expire LimitExpireTyp) UserRequestLimiter {
	return p.m[expire]
}
