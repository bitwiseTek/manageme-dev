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

//SetRoleRoutes sets route for roles
func SetRoleRoutes(router *mux.Router) *mux.Router {
	roleRouter := mux.NewRouter()
	roleRouter.HandleFunc("/api/v1/orgs/roles/add", controllers.AddRole).Methods("POST")
	roleRouter.HandleFunc("/api/v1/admin/orgs/roles/", controllers.GetRoles).Methods("GET")
	roleRouter.HandleFunc("/api/v1/orgs/roles/org/{orgId}", controllers.GetRolesByOrg).Methods("GET")
	roleRouter.HandleFunc("/api/v1/orgs/roles/{id}", controllers.GetRole).Methods("GET")
	roleRouter.HandleFunc("/api/v1/orgs/roles/edit/{id}", controllers.EditRoleByAdmin).Methods("PUT")
	roleRouter.HandleFunc("/api/v1/orgs/roles/permissions/add", controllers.AddPermission).Methods("POST")
	roleRouter.HandleFunc("/api/v1/orgs/roles/permissions/", controllers.GetPermissions).Methods("GET")
	roleRouter.HandleFunc("/api/v1/orgs/roles/permissions/org/{orgId}", controllers.GetPermissionsByOrg).Methods("GET")
	roleRouter.HandleFunc("/api/v1/orgs/roles/permissions/{id}", controllers.GetPermission).Methods("GET")
	roleRouter.HandleFunc("/api/v1/orgs/roles/permissions/edit{id}", controllers.EditPermissionByAdmin).Methods("GET")
	router.PathPrefix("/api/v1/admin/org/roles/").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(roleRouter),
	))
	return router
}
