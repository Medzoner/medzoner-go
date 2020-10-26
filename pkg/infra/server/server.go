package server

import "net/http"

type Server struct {
	HTTPServer *http.Server
}
