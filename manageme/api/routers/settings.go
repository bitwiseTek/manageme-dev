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

//SetSettingRoutes sets route for accounts
func SetSettingRoutes(router *mux.Router) *mux.Router {
	settingRouter := mux.NewRouter()
	settingRouter.HandleFunc("/api/v1/orgs/settings/departments/add", controllers.AddDepartment).Methods("POST")
	settingRouter.HandleFunc("/api/v1/orgs/settings/designations/add", controllers.AddDesignation).Methods("POST")
	settingRouter.HandleFunc("/api/v1/orgs/settings/salary/modes/add", controllers.AddSalaryMode).Methods("POST")
	settingRouter.HandleFunc("/api/v1/orgs/settings/leave/types/add", controllers.AddLeaveType).Methods("POST")
	settingRouter.HandleFunc("/api/v1/orgs/settings/claim/types/add", controllers.AddClaimType).Methods("POST")
	settingRouter.HandleFunc("/api/v1/orgs/settings/branches/add", controllers.AddBranch).Methods("POST")
	settingRouter.HandleFunc("/api/v1/orgs/settings/project/types/add", controllers.AddProjectType).Methods("POST")
	settingRouter.HandleFunc("/api/v1/orgs/settings/departments/org/{orgId}", controllers.GetDepartmentsByOrg).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/designations/org/{orgId}", controllers.GetDesignationsByOrg).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/salary/modes/org/{orgId}", controllers.GetSalaryModesByOrg).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/leave/types/org/{orgId}", controllers.GetLeaveTypesByOrg).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/claim/types/org/{orgId}", controllers.GetClaimTypesByOrg).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/project/types/org/{orgId}", controllers.GetProjectTypesByOrg).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/branches/org/{orgId}", controllers.GetBranchesByOrg).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/departments/{id}", controllers.GetDepartment).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/designations/{id}", controllers.GetDesignation).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/branches/{id}", controllers.GetBranch).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/leave/types/{id}", controllers.GetLeaveType).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/claim/types/{id}", controllers.GetClaimType).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/project/types/{id}", controllers.GetProjectType).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/salary/modes/{id}", controllers.GetSalaryMode).Methods("GET")
	settingRouter.HandleFunc("/api/v1/orgs/settings/departments/edit/{id}", controllers.EditDepartmentByAdmin).Methods("PUT")
	settingRouter.HandleFunc("/api/v1/orgs/settings/designations/edit/{id}", controllers.EditDesignationByAdmin).Methods("PUT")
	settingRouter.HandleFunc("/api/v1/orgs/settings/branches/edit/{id}", controllers.EditBranchByAdmin).Methods("PUT")
	settingRouter.HandleFunc("/api/v1/orgs/settings/leave/types/edit/{id}", controllers.EditLeaveTypeByAdmin).Methods("PUT")
	settingRouter.HandleFunc("/api/v1/orgs/settings/claim/types/edit/{id}", controllers.EditClaimTypeByAdmin).Methods("PUT")
	settingRouter.HandleFunc("/api/v1/orgs/settings/project/types/edit/{id}", controllers.EditProjectTypeByAdmin).Methods("PUT")
	settingRouter.HandleFunc("/api/v1/orgs/settings/salary/modes/edit/{id}", controllers.EditSalaryModeByAdmin).Methods("PUT")
	router.PathPrefix("/api/v1/org/settings/departments").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(settingRouter),
	))
	return router
}
