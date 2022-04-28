package limiter

import "fmt"

type LimitExpireTyp uint32

const (
	LimitExpireThreeSecond LimitExpireTyp = 3
	LimitExpireFiveSecond  LimitExpireTyp = 5
)

var limitExpireTimes = []LimitExpireTyp{LimitExpireThreeSecond, LimitExpireFiveSecond}

func GetCompletePlanKey(planId uint64) string {
	key := fmt.Sprintf("user::request::complete::plan::%d", planId)
	return key
}
