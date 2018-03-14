package main

import (
	"net/http"
	"os"

	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/server"
	"github.com/gorilla/mux"
)

var redir = http.StatusTemporaryRedirect
var landingPage = "https://www.ons.gov.uk/help/localstatistics"
var apiResponse = "This service is no longer available. Please visit https://www.ons.gov.uk/help/localstatistics for more information."

func main() {
	log.Namespace = "dp-legacy-redirector"

	bindAddr := ":8080"
	if v := os.Getenv("BIND_ADDR"); len(v) > 0 {
		bindAddr = v
	}

	srv := server.New(bindAddr, getRouter())

	log.Debug("starting http server", log.Data{"bind_addr": bindAddr})
	if err := srv.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}

func getRouter() *mux.Router {
	router := mux.NewRouter()

	// NeSS website
	router.Host("neighbourhood.statistics.gov.uk").Path("/HTMLDocs/{uri:.*}").HandlerFunc(dataVisHandler)
	router.Host("{subdomain:[a-z]+}.neighbourhood.statistics.gov.uk").Path("/HTMLDocs/{uri:.*}").HandlerFunc(dataVisHandler)
	// NeSS API
	router.Host("neighbourhood.statistics.gov.uk").Path("/NDE2/{uri:.*}").HandlerFunc(apiHandler)
	router.Host("{subdomain:[a-z]+}.neighbourhood.statistics.gov.uk").Path("/NDE2/{uri:.*}").HandlerFunc(apiHandler)
	// WDA website
	router.Host("web.ons.gov.uk").Path("/ons/apiservice/web/{uri:.*}").HandlerFunc(defaultHandler)
	// WDA API
	router.Host("web.ons.gov.uk").Path("/ons/apiservice/{uri:.*}").HandlerFunc(apiHandler)
	router.Host("web.ons.gov.uk").Path("/ons/api/{uri:.*}").HandlerFunc(apiHandler)
	router.Host("data.ons.gov.uk").Path("/{uri:.*}").HandlerFunc(apiHandler)
	// Visual.ONS
	router.Host("visual.ons.gov.uk").Path("/{uri:.*").HandlerFunc(visualHandler)
	// Catch-all
	router.Path("/{uri:.*}").HandlerFunc(defaultHandler)

	return router
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	log.DebugR(req, "redirecting to landing page", log.Data{
		"host": req.Host,
		"path": req.URL.Path,
		"dest": landingPage,
	})
	w.Header().Set("Location", landingPage)
	w.WriteHeader(redir)
}

func dataVisHandler(w http.ResponseWriter, req *http.Request) {
	dest := "https://www.ons.gov.uk/visualisations/nesscontent/" + mux.Vars(req)["uri"]
	log.DebugR(req, "redirecting visualisation", log.Data{
		"host": req.Host,
		"path": req.URL.Path,
		"dest": dest,
	})
	w.Header().Set("Location", dest)
	w.WriteHeader(redir)
}

func apiHandler(w http.ResponseWriter, req *http.Request) {
	log.DebugR(req, "returning api help text", log.Data{
		"host": req.Host,
		"path": req.URL.Path,
	})
	w.WriteHeader(410)
	w.Write([]byte(apiResponse))
}

func visualHandler(w http.ResponseWriter, req *http.Request) {
	log.DebugR(req, "redirecting visual.ons.gov.uk", log.Data{
		"host": req.Host,
		"path": req.URL.Path,
	})
	w.WriteHeader(410)
	w.Write([]byte("FIXME"))
}
