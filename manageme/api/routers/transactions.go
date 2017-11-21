package routers

/**
 *
 * @author Sika Kay
 * @date 20/07/17
 *
 */

import (
	"github.com/bitwiseTek/spottnow-dev/spottnow/api/common"
	"github.com/bitwiseTek/spottnow-dev/spottnow/api/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//SetTransactionRoutes sets route for transactions
func SetTransactionRoutes(router *mux.Router) *mux.Router {
	txRoute := mux.NewRouter()
	txRoute.HandleFunc("/api/v1/transactions/users/{id}", controllers.GetAllById).Methods("GET")
	txRoute.HandleFunc("/api/v1/transactions/users/{id}/transaction/{id}", controllers.GetTransactionByUserId).Methods("GET")
	txRoute.HandleFunc("/api/v1/admin/transactions/", controllers.GetAllTransactions).Methods("GET")
	txRoute.HandleFunc("/api/v1/admin/transactions/{id}", controllers.GetTransactionById).Methods("GET")
	txRoute.HandleFunc("/api/v1/admin/transactions/process/{id}", controllers.ProcessTransaction).Methods("PUT")
	txRoute.HandleFunc("/api/v1/admin/transactions/users/{id}", controllers.GetAllTransactionsByUserId).Methods("GET")
	txRoute.HandleFunc("/api/v1/admin/transactions/users/{id}/transaction/{id}", controllers.GetTransactionByUserId).Methods("GET")
	router.PathPrefix("/api/v1/transactions/users/{id}").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(txRoute),
	))
	router.PathPrefix("/api/v1/admin/transactions/users/{id}").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(txRoute),
	))
	return router
}
