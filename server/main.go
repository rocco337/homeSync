package main

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/bnkamalesh/webgo"
	"github.com/bnkamalesh/webgo/middleware"
)

func main() {
	cfg := &webgo.Config{
		Host:         "",
		Port:         "8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	router := webgo.NewRouter(cfg, getRoutes())
	router.Use(middleware.AccessLog)
	router.Start()
}

func getRoutes() []*webgo.Route {
	return []*webgo.Route{
		&webgo.Route{
			Name:     "api",                                         // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodGet,                                // request type
			Pattern:  "/status",                                     // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), status}, // route handler
		},
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	data := struct {
		Hostname, Ip, DateTime string
	}{
		Hostname: hostName,
		DateTime: time.Now().String(),
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		data.Ip = data.Ip + ", " + addr.String()
	}

	webgo.R200(w, data)
}
