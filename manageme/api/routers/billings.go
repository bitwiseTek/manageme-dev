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

//SetBillingRoutes sets route for billings
func SetBillingRoutes(router *mux.Router) *mux.Router {
	billingRouter := mux.NewRouter()
	billingRouter.HandleFunc("/api/v1/orgs/billings/add", controllers.AddBilling).Methods("POST")
	billingRouter.HandleFunc("/api/v1/admin/orgs/billings/", controllers.GetBillings).Methods("GET")
	billingRouter.HandleFunc("/api/v1/orgs/billings/org/{orgId}", controllers.GetBillingsByOrg).Methods("GET")
	billingRouter.HandleFunc("/api/v1/orgs/billings/{id}", controllers.GetBilling).Methods("GET")
	billingRouter.HandleFunc("/api/v1/admin/orgs/billings/edit/{id}", controllers.EditBillingByAdmin).Methods("PUT")
	billingRouter.HandleFunc("/api/v1/orgs/billings/transactions/add", controllers.AddTransaction).Methods("POST")
	billingRouter.HandleFunc("/api/v1/admin/orgs/billings/transactions", controllers.GetTransactions).Methods("GET")
	billingRouter.HandleFunc("/api/v1/orgs/billings/transactions/billing/{billingId}", controllers.GetTransactionsByBilling).Methods("GET")
	billingRouter.HandleFunc("/api/v1/orgs/billings/transactions/{id}", controllers.GetTransaction).Methods("GET")
	billingRouter.HandleFunc("/api/v1/admin/orgs/billings/transactions/edit{id}", controllers.EditTransactionByAdmin).Methods("GET")
	router.PathPrefix("/api/v1/admin/org/users/").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(billingRouter),
	))
	return router
}
