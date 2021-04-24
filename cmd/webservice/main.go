package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hojabri/suss/pkg/api/service"
	"github.com/hojabri/suss/pkg/api/setup"
	"github.com/hojabri/suss/pkg/config"
	"github.com/hojabri/suss/pkg/maxmind"
	"github.com/hojabri/suss/pkg/repository/sqlite"
	"github.com/hojabri/suss/pkg/susslogger"
	"golang.org/x/crypto/acme/autocert"
)

var err error

func main() {
	
	// Connect to db
	
	err = sqlite.Connect()
	if err != nil {
		susslogger.Log().Fatalf("Can not connect to sqllite db: %s", err.Error())
	}
	
	// Connect to GeoLite2 db
	err = maxmind.OpenDB()
	if err != nil {
		susslogger.Log().Fatalf("Can not connect to GeoLite2 db: %s", err.Error())
	}
	defer func() {
		err = maxmind.CloseDB()
		if err != nil {
			susslogger.Log().Fatalf("Can not close GeoLite2 db: %s", err.Error())
		}
	}()
	
	address := config.Config.GetString("WEBSERVICE.ADDRESS")
	port := config.Config.GetString("WEBSERVICE.PORT")
	domain := config.Config.GetString("WEBSERVICE.DOMAIN")
	enableAutoCert := config.Config.GetBool("WEBSERVICE.ENABLE_AUTOCERT")
	
	var setupConfig setup.Config
	
	susslogger.Log().Infof("PORT:%s", port)
	susslogger.Log().Infof("ENABLE_AUTOCERT:%v", enableAutoCert)
	susslogger.Log().Infof("DOMAIN:%s", domain)
	
	setupConfig.Address = fmt.Sprintf("%s:%s", address, port)
	setupConfig.InsecureHTTP = !enableAutoCert
	var scheme string
	if setupConfig.InsecureHTTP {
		scheme = "http"
	} else {
		scheme = "https"
	}
	svc := &service.SUSSService{}
	api := setup.NewServer(svc, &setupConfig)
	
	// Certificate manager
	m := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		// Replace with your domain
		HostPolicy: autocert.HostWhitelist(domain),
		// Folder to store the certificates
		Cache: autocert.DirCache("cert"),
	}
	
	// TLS Config
	cfg := &tls.Config{
		// Get Certificate from Let's Encrypt
		GetCertificate: m.GetCertificate,
		// By default NextProtos contains the "h2"
		// This has to be removed since Fasthttp does not support HTTP/2
		// Or it will cause a flood of PRI method logs
		// http://webconcepts.info/concepts/http-method/PRI
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}
	api.TLSConfig = cfg
	
	api.Routes.Static("/static", "./static")
	api.Routes.Static("/openapi", "./openapi")
	swaggerUrl := fmt.Sprintf("%s://%s:%s/openapi/suss-openapi.yml", scheme, address, port)

	// Place ReDoc file to render swagger specification document in the root GET of webservice
	api.Routes.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"url": swaggerUrl,
		})
	})
	
	// Run the webservice
	err = api.RunWithSigHandler()
	if err != nil {
		susslogger.Log().Fatal(err)
	}
	
}
