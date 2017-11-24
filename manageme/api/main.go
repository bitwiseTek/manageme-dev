package main

/**
 *
 * @author Sika Kay
 * @date 22/11/17
 *
 */
import (
	"log"
	"net/http"

	common "github.com/bitwiseTek/manageme-dev/manageme/api/common"
	"github.com/bitwiseTek/manageme-dev/manageme/api/routers"
	"github.com/codegangsta/negroni"
)

func main() {
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening on :80")
	server.ListenAndServe()
}
