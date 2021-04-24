
package setup

import (
	"crypto/tls"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/hojabri/suss/pkg/api/middleware"
	"github.com/hojabri/suss/pkg/entities"
	"github.com/hojabri/suss/pkg/susslogger"
	"github.com/opentracing/opentracing-go"
	"github.com/valyala/fasthttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Service interface {
	NewUserSessionEvent(c *fiber.Ctx) *entities.Response
}



type Routes struct {
	*fiber.App
	EventRoutes fiber.Router
}

func initializeRoutes() *Routes {
	engine := html.New("./static", ".html")
	engine.Reload(true)
	engine.Debug(true)
	engine.Layout("embed")
	engine.Delims("{{", "}}")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	routes := &Routes{App: app}
	
	routes.EventRoutes = app.Group("/v1/event", logger.New(), middleware.AuthReq())
	
	return routes
}

type Config struct {
	Address      string
	InsecureHTTP bool
	TLSCertFile  string
	TLSKeyFile   string
	Tracer       opentracing.Tracer
}

// Server defines the Server service.
type Server struct {
	Routes    *Routes
	service   Service
	config    *Config
	Server    *fasthttp.Server
	Title     string
	Version   string
	TLSConfig *tls.Config
}

func NewServer(svc Service, config *Config) *Server {
	server := &Server{
		Routes:  initializeRoutes(),
		service: svc,
		config:  config,
		Title:   "SUSS",
		Version: "1.0.0",
	}
	server.Server = &fasthttp.Server{
		Handler:      server.Routes.App.Handler(),
		Name:         "SUSS",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	
	return server
}

// ConfigureRoutes starts the internal configureRoutes methode.
func (s *Server) ConfigureRoutes() {
	s.configureRoutes()
}

func (s *Server) configureRoutes() {
	s.Routes.Get("/",do(index))
    s.Routes.Get("/ping",do(ping))

	s.Routes.EventRoutes.Post("/", do(s.service.NewUserSessionEvent))

}


func do(handler func(c *fiber.Ctx) *entities.Response) fiber.Handler {
	return func(c *fiber.Ctx) error {
		
		resp := handler(c)
		err :=c.Status(resp.Code).JSON(resp.Body)
		if err!=nil {
			err = c.Status(http.StatusInternalServerError).JSON(entities.Response{
				Code:     http.StatusInternalServerError,
				Body:     err.Error(),
				Title:    "InternalServerError",
				Message:  err.Error(),
				Instance: "InternalServerError",
			})
			if err != nil {
				return err
			}
		}
		
		return nil
	}
}


func ping(c *fiber.Ctx) *entities.Response {
	return &entities.Response{
		Code:     http.StatusOK,
		Body:     "Pong",
		Title:    "OK",
		Message:  "It is working fine!",
		Instance: "Health",
	}
}
func index(c *fiber.Ctx) *entities.Response {
	return &entities.Response{
		Code:     http.StatusOK,
		Body:     "SUSS index. Call POST to /v1/event (see documents for more info)",
		Title:    "OK",
		Message:  "",
		Instance: "Index",
	}
}
func (s *Server) Run() error {
	
	// configure service routes
	s.configureRoutes()
	
	susslogger.Log().Infof("Serving '%s - %s' ", s.Title, s.Version)
	

	if s.config.InsecureHTTP {
		susslogger.Log().Info("Insecure HTTP")
		s.Routes.App.Listen(s.config.Address)
		return nil
	}
	susslogger.Log().Info("Secure HTTP")
	ln, err := tls.Listen("tcp", s.config.Address, s.TLSConfig)
	if err != nil {
		panic(err)
	}
	
	
	susslogger.Log().Fatal(s.Routes.App.Listener(ln))
	
	return nil
}

func (s *Server) Shutdown() error {
	return s.Routes.App.Shutdown()
}

// RunWithSigHandler runs the Server server with SIGTERM handling automatically
// enabled. The server will listen for a SIGTERM signal and gracefully shutdown
// the web server.
func (s *Server) RunWithSigHandler(shutdown ...func() error) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		<-sigCh
		s.Shutdown()
	}()
	
	err := s.Run()
	if err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}
	
	for _, fn := range shutdown {
		err := fn()
		if err != nil {
			return err
		}
	}
	
	return nil
}

