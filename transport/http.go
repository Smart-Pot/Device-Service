package transport

import (
	"deviceservice/endpoints"
	"net/http"

	pkghttp "github.com/Smart-Pot/pkg/common/http"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func MakeHTTPHandlers(e endpoints.Endpoints, l log.Logger) http.Handler {
	r := mux.NewRouter().PathPrefix("/device").Subrouter()
	return pkghttp.EnableCORS(r)
}


