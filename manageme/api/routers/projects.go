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

//SetProjectRoutes sets route for accounts
func SetProjectRoutes(router *mux.Router) *mux.Router {
	projectRouter := mux.NewRouter()
	projectRouter.HandleFunc("/api/v1/orgs/tasks/add", controllers.AddTask).Methods("POST")
	projectRouter.HandleFunc("/api/v1/orgs/tasks/child/add", controllers.AddChildTask).Methods("POST")
	projectRouter.HandleFunc("/api/v1/orgs/tasks/", controllers.GetTasks).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/tasks/org/{orgId}", controllers.GetTasksByOrg).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/tasks/task/{taskId}", controllers.GetTasksByTask).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/tasks/employees/{employeeId}", controllers.GetTasksByEmp).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/tasks/{id}", controllers.GetAccount).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/tasks/employees/edit/{id}", controllers.EditTaskByEmp).Methods("PUT")
	projectRouter.HandleFunc("/api/v1/orgs/projects/tasks/add", controllers.AddProjectTask).Methods("POST")
	projectRouter.HandleFunc("/api/v1/orgs/projects/tasks/", controllers.GetProjectTasks).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/projects/tasks/org/{orgId}", controllers.GetProjectTasksByOrg).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/projects/tasks/edit/{id}", controllers.EditProjectTaskByEmp).Methods("PUT")
	projectRouter.HandleFunc("/api/v1/orgs/projects/tasks/{id}", controllers.GetProjectTaskById).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/projects/add", controllers.AddProject).Methods("POST")
	projectRouter.HandleFunc("/api/v1/orgs/projects/org/{orgId}", controllers.GetProjectsByOrg).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/projects/employees/{mgrId}", controllers.GetProjectsByMgr).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/projects/types/{typeId}", controllers.GetProjectsByType).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/projects/{id}", controllers.GetProject).Methods("GET")
	projectRouter.HandleFunc("/api/v1/orgs/projects/employees/edit{id}", controllers.EditProjectByMgr).Methods("PUT")
	router.PathPrefix("/api/v1/org/tasks/").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(projectRouter),
	))
	return router
}
