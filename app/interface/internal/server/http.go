package server

import (
	v1 "github.com/Casper-Mars/cloud-restaurant/api/interface/v1"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/conf"
	"github.com/Casper-Mars/cloud-restaurant/app/interface/internal/service"
	"github.com/Casper-Mars/cloud-restaurant/pkg/jwt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, logger log.Logger,
	auth *service.AuthService,
	health *service.HealthService,
	user *service.UserService,
	food *service.FoodService,
	comment *service.CommentService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
			jwt.Server(ac.AccessSecret),
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
