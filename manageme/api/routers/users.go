package routers

/**
 *
 * @author Sika Kay
 * @date 22/11/17
 *
 */

import (
	"github.com/bitwiseTek/manageme-dev/manageme/api/common"
	"github.com/bitwiseTek/manageme-dev/manageme/api/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//SetUserRoutes sets route for users
func SetUserRoutes(router *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/api/v1/users/signup", controllers.SignUp).Methods("POST")
	userRouter.HandleFunc("/api/v1/users/signin", controllers.SignIn).Methods("POST")
	userRouter.HandleFunc("/api/v1/admin/users/", controllers.GetUsers).Methods("GET")
	userRouter.HandleFunc("/api/v1/users/edit/{id}", controllers.Edit).Methods("PUT")
	userRouter.HandleFunc("/api/v1/users/{id}", controllers.GetById).Methods("GET")
	userRouter.HandleFunc("/api/v1/admin/users/edit/{id}", controllers.EditUserByAdmin).Methods("PUT")
	router.PathPrefix("/api/v1/admin/users/").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(userRouter),
	))
	return router
}
