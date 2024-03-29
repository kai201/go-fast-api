package routers

import (
	"net/http"

	"github.com/fast/internal/config"
	"github.com/fast/pkg/ginx/metrics"
	"github.com/fast/pkg/ginx/middleware"
	"github.com/fast/pkg/ginx/validator"
	"github.com/fast/pkg/jwt"
	"github.com/fast/pkg/logger"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var routerFns []func(r *gin.RouterGroup)

// NewRouter create a new router
func NewRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// request id middleware
	r.Use(middleware.RequestID())

	// logger middleware
	r.Use(middleware.Logging(
		middleware.WithLog(logger.Get()),
		middleware.WithRequestIDFromContext(),
		middleware.WithIgnoreRoutes("/metrics"), // ignore path
	))

	pprof.Register(r)
	// init jwt middleware
	jwt.Init(
	//jwt.WithExpire(time.Hour*24),
	//jwt.WithSigningKey("123456"),
	//jwt.WithSigningMethod(jwt.HS384),
	)

	// metrics middleware
	if config.Instance.Server.EnableMetrics {
		r.Use(metrics.Metrics(r,
			metrics.WithMetricsPath("/metrics"),                // default is /metrics
			metrics.WithIgnoreStatusCodes(http.StatusNotFound), // ignore 404 status codes
		))
	}
	// trace middleware
	if config.Instance.Server.EnableTrace {
		r.Use(middleware.Tracing(config.Instance.Server.Name))
	}

	// validator
	binding.Validator = validator.Init()

	rg := &r.RouterGroup

	for _, fn := range routerFns {
		fn(rg)
	}
	return r
}
