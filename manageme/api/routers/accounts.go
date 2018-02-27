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

//SetAccountRoutes sets route for accounts
func SetAccountRoutes(router *mux.Router) *mux.Router {
	accountRouter := mux.NewRouter()
	accountRouter.HandleFunc("/api/v1/orgs/accounts/add", controllers.AddAccount).Methods("POST")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/child/add", controllers.AddChildAccount).Methods("POST")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/", controllers.GetAccounts).Methods("GET")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/org/{orgId}", controllers.GetAccountsByOrg).Methods("GET")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/account/{accountId}", controllers.GetAccountsByAcct).Methods("GET")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/{id}", controllers.GetAccount).Methods("GET")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/employees/edit/{id}", controllers.EditAccountByMgr).Methods("PUT")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/employees/child/edit/{id}", controllers.EditChildAccountByMgr).Methods("PUT")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/journal/accounts/add", controllers.AddJournalAccount).Methods("POST")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/journal/entries/add", controllers.AddJournalEntry).Methods("POST")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/journal/accounts/account/{accountId}", controllers.GetJournalAccountsByAcct).Methods("GET")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/journal/entries/account/{accountId}", controllers.GetJournalEntriesByAcct).Methods("GET")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/journal/accounts/org/{orgId}", controllers.GetJournalAccountsByOrg).Methods("GET")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/journal/entries/org/{orgId}", controllers.GetJournalEntriesByOrg).Methods("GET")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/journal/accounts/{id}", controllers.GetJournalAccount).Methods("GET")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/journal/entries/{id}", controllers.GetJournalEntry).Methods("GET")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/journal/accounts/edit{id}", controllers.EditJournalAccountByEmp).Methods("PUT")
	accountRouter.HandleFunc("/api/v1/orgs/accounts/journal/entries/edit{id}", controllers.EditJournalEntryByEmp).Methods("PUT")
	router.PathPrefix("/api/v1/admin/org/users/").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(accountRouter),
	))
	return router
}
