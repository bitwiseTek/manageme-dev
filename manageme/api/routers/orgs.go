package routers

/**
 *
 * @author Sika Kay
 * @date 11/02/18
 *
 */

import (
	"github.com/bitwiseTek/manageme-dev/manageme/api/common"
	"github.com/bitwiseTek/manageme-dev/manageme/api/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//SetOrgRoutes sets route for orgs
func SetOrgRoutes(router *mux.Router) *mux.Router {
	orgRouter := mux.NewRouter()
	orgRouter.HandleFunc("/api/v1/orgs/add", controllers.AddOrg).Methods("POST")
	orgRouter.HandleFunc("/api/v1/admin/orgs/", controllers.GetOrgs).Methods("GET")
	orgRouter.HandleFunc("/api/v1/orgs/users/{userId}", controllers.GetOrgByUser).Methods("GET")
	orgRouter.HandleFunc("/api/v1/admin/orgs/edit/{id}", controllers.EditOrg).Methods("PUT")
	orgRouter.HandleFunc("/api/v1/orgs/{id}", controllers.GetOrg).Methods("GET")
	orgRouter.HandleFunc("/api/v1/orgs/edit/{id}", controllers.EditOrg).Methods("PUT")
	router.PathPrefix("/api/v1/admin/orgs/").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(orgRouter),
	))
	return router
}
