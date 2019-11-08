package server

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/bnkamalesh/webgo"
	"github.com/bnkamalesh/webgo/middleware"
)

type HomeSyncServer struct {
	hardDriveOperations HardDriveOperations
}

func (server HomeSyncServer) Start() {
	server.hardDriveOperations.RootPath = "/home/roko/sharedTestRemote"

	cfg := &webgo.Config{
		Port:         "8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	router := webgo.NewRouter(cfg, server.getRoutes())
	router.Use(middleware.AccessLog)
	router.Start()
}

func (server HomeSyncServer) getRoutes() []*webgo.Route {
	return []*webgo.Route{
		&webgo.Route{
			Name:     "status",                                             // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodGet,                                       // request type
			Pattern:  "/status",                                            // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), server.status}, // route handler
		},
		&webgo.Route{
			Name:     "api/tree",                                               // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodGet,                                           // request type
			Pattern:  "/api/tree",                                              // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), server.folderTree}, // route handler
		},
		&webgo.Route{
			Name:     "api/upload",                                         // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodPost,                                      // request type
			Pattern:  "/api/upload",                                        // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), server.upload}, // route handler
		},
	}
}
func (server HomeSyncServer) upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		panic(err)
	}

	relativePath := r.FormValue("relativePath")
	file, header, err := r.FormFile("data")
	if err != nil {
		panic(err)
	}

	server.hardDriveOperations.Create(relativePath, header.Filename, file)

	defer file.Close()
	defer r.Body.Close()

	webgo.R200(w, nil)
}

func (server HomeSyncServer) folderTree(w http.ResponseWriter, r *http.Request) {
	service := new(HardDriveOperations)
	service.RootPath = "/home/roko/sharedTestRemote"

	webgo.R200(w, service.Tree())
}

func (server HomeSyncServer) status(w http.ResponseWriter, r *http.Request) {
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
