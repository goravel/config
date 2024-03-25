package middleware

import (
	"math/rand"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/session"
	"github.com/goravel/framework/support/carbon"
)

func StartSession() http.Middleware {
	return func(ctx http.Context) {
		req := ctx.Request()

		// Check if session exists
		if req.HasSession() {
			req.Next()
			return
		}

		// Retrieve session driver
		d, err := session.Facade.Driver()
		if err != nil {
			return
		}

		// Build session
		s := session.Facade.BuildSession(d)
		s.SetID(req.Cookie(s.GetName()))

		// Start session
		s.Start()
		req.SetSession(s)

		// Perform garbage collection based on lottery
		lottery := session.ConfigFacade.Get("session.lottery").([]int)
		if len(lottery) == 2 {
			if rand.Intn(lottery[1])+1 <= lottery[0] {
				err := d.Gc(300)
				if err != nil {
					return
				}
			}
		}

		// Continue processing request
		req.Next()

		// Retrieve updated session
		s = req.Session()

		// Set session cookie in response
		config := session.ConfigFacade
		ctx.Response().Cookie(http.Cookie{
			Name:     s.GetName(),
			Value:    s.GetID(),
			Expires:  carbon.Now().AddMinutes(config.GetInt("session.lifetime")).StdTime(),
			Path:     config.GetString("session.path"),
			Domain:   config.GetString("session.domain"),
			Secure:   config.GetBool("session.secure"),
			HttpOnly: config.GetBool("session.http_only"),
			SameSite: config.GetString("session.same_site"),
		})

		// Save session
		err = s.Save()
		if err != nil {
			return
		}
	}
}
