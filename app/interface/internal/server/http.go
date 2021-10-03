package server

import (
	v1 "github.com/Casper-Mars/cloud-restaurant/api/interface/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/conf"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/service"
	authpkg "github.com/Casper-Mars/cloud-restaurant/pkg/middleware/auth"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"regexp"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, logger log.Logger,
	auth *service.AuthService,
	health *service.HealthService,
	user *service.UserService,
	food *service.FoodService,
	comment *service.CommentService,
) *http.Server {

	keyFunc := func(token *jwt2.Token) (interface{}, error) {
		return []byte(ac.AccessSecret), nil
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
			selector.Server(
				jwt.Server(keyFunc),
				authpkg.Server(),
			).
				Match(func(operation string) bool {
					// 白名单
					r, err := regexp.Compile("/interface.v1.Auth/.*")
					if err != nil {
						return false
					}
					return r.FindString(operation) != operation
				}).
				Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterAuthHTTPServer(srv, auth)
	v1.RegisterHealthHTTPServer(srv, health)
	v1.RegisterUserHTTPServer(srv, user)
	v1.RegisterCommentHTTPServer(srv, comment)
	v1.RegisterFoodHTTPServer(srv, food)
	return srv
}
