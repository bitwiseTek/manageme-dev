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

//SetOrgUserRoutes sets route for org users
func SetOrgUserRoutes(router *mux.Router) *mux.Router {
	orgUserRouter := mux.NewRouter()
	orgUserRouter.HandleFunc("/api/v1/org/users/add", controllers.AddOrgUser).Methods("POST")
	orgUserRouter.HandleFunc("/api/v1/org/users/signin", controllers.OrgUserSignIn).Methods("POST")
	orgUserRouter.HandleFunc("/api/v1/admin/org/users/", controllers.GetOrgUsers).Methods("GET")
	orgUserRouter.HandleFunc("/api/v1/admin/org/users/org/{orgId}", controllers.GetOrgUsersByOrg).Methods("GET")
	orgUserRouter.HandleFunc("/api/v1/org/users/edit/{id}", controllers.EditOrgUser).Methods("PUT")
	orgUserRouter.HandleFunc("/api/v1/org/users/{id}", controllers.GetOrgUser).Methods("GET")
	orgUserRouter.HandleFunc("/api/v1/admin/org/users/edit/{id}", controllers.EditOrgUserByAdmin).Methods("PUT")
	router.PathPrefix("/api/v1/admin/org/users/").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(orgUserRouter),
	))
	return router
}
