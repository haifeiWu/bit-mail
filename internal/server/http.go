package server

import (
	mail "bit-mail/api/mail/v1"
	user "bit-mail/api/user/v1"
	"bit-mail/internal/conf"
	"bit-mail/internal/service"
	
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userService *service.UserService, mailService *service.MailMessageService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	
	opts = append(
		opts,
		// http.ErrorEncoder(ErrorEncoder),
		http.Filter(
			// 跨域处理
			handlers.CORS(
				handlers.AllowedHeaders([]string{"Content-Type", "x-token"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
				handlers.AllowedOrigins([]string{"*"}))),
	)
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
	user.RegisterUserHTTPServer(srv, userService)
	mail.RegisterMailMessageHTTPServer(srv, mailService)
	return srv
}
