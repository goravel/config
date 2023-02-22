package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/gookit/color"

	contracthttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	httplimit "github.com/goravel/framework/http/limit"
)

func Throttle(name string) contracthttp.Middleware {
	return func(ctx contracthttp.Context) {
		if limiter := facades.RateLimiter.Limiter(name); limiter != nil {
			if limits := limiter(ctx); len(limits) > 0 {
				for _, limit := range limits {
					if instance, ok := limit.(*httplimit.Limit); ok {

						// if no key is set, use the path and ip address as the default key
						if len(instance.Key) == 0 {
							hash := md5.Sum([]byte(ctx.Request().Path()))
							instance.Key = "throttle:" + hex.EncodeToString(hash[:]) + ":" + ctx.Request().Ip()
						} else {
							hash := md5.Sum([]byte(instance.Key))
							instance.Key = "throttle:" + hex.EncodeToString(hash[:])
						}

						// check if the timer exists in the cache
						if facades.Cache.Has(instance.Key + ":timer") {
							// if the timer exists, check if the number of attempts is greater than the max attempts
							value := facades.Cache.GetInt(instance.Key, 0)
							if value >= instance.MaxAttempts {
								if instance.ResponseCallback != nil {
									instance.ResponseCallback(ctx)
								} else {
									ctx.Request().AbortWithStatus(http.StatusTooManyRequests)
								}
							} else {
								err := facades.Cache.Put(instance.Key, value+1, time.Duration(instance.DecayMinutes)*time.Minute)
								if err != nil {
									color.Redln("[Throttle] Error: ", err.Error())
									ctx.Request().Next()
								}
							}
						} else {

							// if the timer does not exist, create it and set the number of attempts to 1
							err := facades.Cache.Put(instance.Key+":timer", time.Now().Unix(), time.Duration(instance.DecayMinutes)*time.Minute)
							if err != nil {
								color.Redln("[Throttle] Error: ", err.Error())
								ctx.Request().Next()
							}
							err = facades.Cache.Put(instance.Key, 1, time.Duration(instance.DecayMinutes)*time.Minute)
							if err != nil {
								color.Redln("[Throttle] Error: ", err.Error())
								ctx.Request().Next()
							}
						}
					}
				}
			}
		}

		ctx.Request().Next()
	}
}
