package middleware

import (
	"fmt"
	"strconv"
	"time"

	httpcontract "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/http"
	httplimit "github.com/goravel/framework/http/limit"
)

const (
	// HeaderRateLimitLimit, HeaderRateLimitRemaining, and HeaderRateLimitReset
	// are the recommended return header values from IETF on rate limiting. Reset
	// is in UTC time.
	HeaderRateLimitLimit     = "X-RateLimit-Limit"
	HeaderRateLimitRemaining = "X-RateLimit-Remaining"
	HeaderRateLimitReset     = "X-RateLimit-Reset"

	// HeaderRetryAfter is the header used to indicate when a client should retry
	// requests (when the rate limit expires), in UTC time.
	HeaderRetryAfter = "Retry-After"
)

func Throttle(name string) httpcontract.Middleware {
	return func(ctx httpcontract.Context) {
		if limiter := http.RateLimiterFacade.Limiter(name); limiter != nil {
			if limits := limiter(ctx); len(limits) > 0 {
				for _, limit := range limits {
					if instance, exist := limit.(*httplimit.Limit); exist {
						tokens, remaining, reset, ok, err := instance.Store.Take(ctx, key(ctx, instance))
						if err != nil {
							response(ctx, instance)
							return
						}

						resetTime := time.Unix(0, int64(reset)).UTC()
						retryAfter := resetTime.Sub(time.Now().UTC()).Seconds()
						ctx.Response().Header(HeaderRateLimitLimit, strconv.FormatUint(tokens, 10))
						ctx.Response().Header(HeaderRateLimitRemaining, strconv.FormatUint(remaining, 10))

						if !ok {
							ctx.Response().Header(HeaderRateLimitReset, strconv.Itoa(int(resetTime.Unix())))
							ctx.Response().Header(HeaderRetryAfter, strconv.Itoa(int(retryAfter)))
							response(ctx, instance)
							break
						}
					}
				}
			}
		}

		ctx.Request().Next()
	}
}

func key(ctx httpcontract.Context, limit *httplimit.Limit) string {
	// if no key is set, use the path and ip address as the default key
	if len(limit.Key) == 0 {
		limit.Key = fmt.Sprintf("%s:%s", ctx.Request().Ip(), ctx.Request().Path())
	}

	return limit.Key
}

func response(ctx httpcontract.Context, limit *httplimit.Limit) {
	if limit.ResponseCallback != nil {
		limit.ResponseCallback(ctx)
	} else {
		ctx.Request().AbortWithStatus(httpcontract.StatusTooManyRequests)
	}
}
