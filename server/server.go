package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	mw "github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/steveoc64/godev/config"
	"github.com/steveoc64/godev/db"
	"github.com/steveoc64/godev/echocors"
	"github.com/steveoc64/godev/smt"
	runner "gopkg.in/mgutz/dat.v1/sqlx-runner"

	// "github.com/facebookgo/grace/gracehttp"
	"golang.org/x/net/websocket"
)

var e *echo.Echo
var DB *runner.DB

var Config config.ConfigType

func main() {

	Config = config.LoadConfig()
	cpus := smt.Init()
	fmt.Printf("Alice & the Cat running on %d CPU cores\n", cpus)

	// Make sure the SMS stuff is all working before we go too far
	smsbal, smserr := GetSMSBalance()
	if smserr != nil {
		log.Fatal("Cannot retrieve SMS account info", smserr.Error())
	}
	log.Println("... Remaining SMS Balance =", smsbal)

	// Start up the basic web server
	e = echo.New()
	e.SetDebug(true)
	e.Static("/", "public")

	// e.Index("public/index.html")
	// e.ServeDir("/", "public/")
	e.SetHTTPErrorHandler(func(err error, context echo.Context) {
		httpError, ok := err.(*echo.HTTPError)
		if ok {
			// errorCode := httpError.Code()
			errorCode := httpError.Code
			switch errorCode {
			case http.StatusNotFound:
				// TODO handle not found case
				// log.Println("Not Found", err.Error())
				// We are usually here due to an F5 refresh, in which case
				// the URL is not expected to be there
				context.Redirect(http.StatusMovedPermanently, "/")
			default:
				// TODO handle any other case
			}
		}
	})

	// e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())
	if Config.Debug {
		e.SetDebug(true)
	}
	echocors.Init(e, Config.Debug)

	// Do a database backup before we begin
	out, err := exec.Command("../scripts/backup.sh").Output()
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(out))
	}

	// Connect to the database
	DB = db.Init(Config.DataSourceName)

	// Add the all important Websocket handler
	Connections = new(ConnectionsList)
	registerRPC()

	e.Get("/ws", standard.WrapHandler(websocket.Handler(webSocket)))
	// e.Get("/ws", fasthttp.WrapHandler(websocket.Handler(webSocket)))

	e.SetDebug(true)
	// e.WebSocket("/ws", webSocket)

	// Start the web server
	if Config.Debug {
		log.Printf("... Starting Web Server on port %d", Config.WebPort)
	}
	std := standard.New(fmt.Sprintf(":%d", Config.WebPort))
	e.Run(std)

	// std.SetHandler(e)
	// gracehttp.Serve(std.Server)

}
