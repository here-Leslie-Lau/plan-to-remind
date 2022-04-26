package limiter

type UserRequestLimiter interface {
	Limit(UserId uint64) error
}

type redisRequestLimiter struct {

}
