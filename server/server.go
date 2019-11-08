package server

import (
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/bnkamalesh/webgo"
	"github.com/bnkamalesh/webgo/middleware"
)

type HomeSyncServer struct {
}

func (server HomeSyncServer) Start() {
	cfg := &webgo.Config{
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
			Name:     "status",                                      // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodGet,                                // request type
			Pattern:  "/status",                                     // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), status}, // route handler
		},
		&webgo.Route{
			Name:     "api/tree",                                        // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodGet,                                    // request type
			Pattern:  "/api/tree",                                       // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), folderTree}, // route handler
		},
		&webgo.Route{
			Name:     "api/upload",                                  // A label for the API/URI, this is not used anywhere.
			Method:   http.MethodPost,                               // request type
			Pattern:  "/api/upload",                                 // Pattern for the route
			Handlers: []http.HandlerFunc{middleware.Cors(), upload}, // route handler
		},
	}
}
func upload(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	// err := r.ParseMultipartForm(0)
	// if err != nil {
	// 	panic(err)
	// }

	// out, err := os.Create("filename.ext")
	// if err != nil {
	// 	// panic?
	// }
	// io.Copy(out, r.Body)

	// defer out.Close()
	// defer r.Body.Close()
	// _, header, _ := r.FormFile("data")
	webgo.R200(w, body)
}
func folderTree(w http.ResponseWriter, r *http.Request) {
	service := new(HardDriveOperations)
	service.RootPath = "/home/roko/sharedTestRemote"
	webgo.R200(w, service.Tree())
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
