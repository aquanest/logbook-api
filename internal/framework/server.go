package framework

import (
	"fmt"
	"logbook-api/internal/application"
	"logbook-api/internal/config"
	"logbook-api/internal/infrastructure"
	"logbook-api/pkg/http"
	impl "logbook-api/pkg/router"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server struct
type Server struct {
	Config     *config.Config
	Repository *infrastructure.Repository
	Usecase    *application.Usecase
}

// New returns Server struct
func New(c *config.Config, r *infrastructure.Repository, u *application.Usecase) Server {
	return Server{
		Config:     c,
		Repository: r,
		Usecase:    u,
	}
}

// Boot the server
func (s *Server) Boot(addr string, port int, debugMode bool) {
	c := NewController(s.Config, s.Repository, s.Usecase)
	routes := c.InvokeRouteProvider()
	e := echo.New()

	// Set debug mode
	e.Debug = debugMode

	// Customize the startup banner
	e.HideBanner = true
	s.printStartupBanner()

	// Customize the startup message
	e.HidePort = true
	s.printBootMessage(&addr, &port)

	// Use custom error handler
	e.HTTPErrorHandler = s.customHTTPErrorHandler

	// Register middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowOrigins: []string{http.FQDNPrivate, http.FQDNPublic},
		AllowMethods: []string{http.MethodGet, http.MethodHead},
	}))

	// Attach router
	routes.RegisterDefaultRoute(e)
	impl.RegisterHandlers(e, &routes)
	s.printRoutes(e.Routes())

	// Boot server
	if err := e.Start(addr + ":" + strconv.Itoa(port)); err != nil {
		s.printHaltMessage(&err)
	}
}

func (s *Server) customHTTPErrorHandler(err error, ctx echo.Context) {
	c := NewController(s.Config, s.Repository, s.Usecase)
	routes := c.InvokeRouteProvider()
	err = routes.InvokeUnexpectedError(err.(*echo.HTTPError), ctx)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
