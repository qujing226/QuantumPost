package redis_slide_window

import (
	"context"
	_ "embed"
	"github.com/qujing226/quantum_post/pkg/ginx/ratelimit"
	"github.com/redis/go-redis/v9"
	"time"
)

//go:embed slide_window.lua
var luaSlideWindow string

type RedisSlidingWindowLimiter struct {
	cmd redis.Cmdable

	interval time.Duration
	rate     int
}

func (r *RedisSlidingWindowLimiter) Limit(ctx context.Context, key string) (bool, error) {
	return r.cmd.Eval(ctx, luaSlideWindow,
		[]string{key},
		r.interval.Milliseconds(), r.rate,
		time.Now().UnixMilli()).Bool()
}

func NewRedisSlidingWindowLimiter(cmd redis.Cmdable, interval time.Duration, rate int) ratelimit.Limiter {
	return &RedisSlidingWindowLimiter{
		cmd:      cmd,
		interval: interval,
		rate:     rate,
	}

}
