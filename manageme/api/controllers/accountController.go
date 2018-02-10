package controllers

/**
 *
 * @author Sika Kay
 * @date 05/02/18
 *
 */
 
import (
	"encoding/json"
	"net/http"

	httpcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
	"github.com/bitwiseTek/manageme-dev/manageme/api/data"
	"github.com/bitwiseTek/manageme-dev/manageme/api/common"
)

//AddAccount for /orgs/accounts/add api
func AddAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes AccountResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid account data",
			500,
		)
		return
	}
	acct := &dataRes.Data
	acct.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}

	repo.AddAcctByOrgID(acct)
	j, err := json.Marshal(AccountResource{Data: *acct})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//AddChildAccount for /orgs/accounts/child/add api
func AddChildAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	acctid = bson.ObjectIdHex(vars["acctid"])
	var dataRes AccountResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid account data",
			500,
		)
		return
	}
	acct := &dataRes.Data
	acct.OrgID = orgid
	acct.ParentAccount = acctid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}

	repo.AddChildAcctByOrgID(acct)
	j, err := json.Marshal(AccountResource{Data: *acct})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//GetAccounts for /orgs/accounts api
func GetAccounts(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	accts := repo.GetAccounts()

	j, err := json.Marshal(AccountsResource{Data: bills})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetAccountsByOrg for /orgs/accounts/org/{orgId} api
func GetAccountsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	accounts, err := repo.GetAccountsByOrgID(orgid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(AccountsResource{Data: accounts})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetAccountsByAccount for /orgs/accounts/account/{accountId} api
func GetAccountsByAcct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	acctid = bson.ObjectIdHex(vars["accountid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	accounts, err := repo.GetAccountsByAcctID(acctid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(AccountsResource{Data: accounts})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetAccount for /orgs/accounts/{id} api
func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	account, err := repo.GetAcctByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occured",
				500,
			)
		}
		return
	}
	j, err := json.Marshal(account)
	if err != nil {
		common.DisplayAppError(
			w, 
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.Header().Set("Content-Application", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//EditAccountByEmp for /orgs/accounts/employees/edit/{id} api
func EditAccountByMgr(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource AccountResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid account data",
			500,
		)
		return
	}
	acct := &dataResource.Data
	acct.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	
	err := repo.EditAccountByID(acct); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//EditChildAccountByMgr for /orgs/accounts/employees/child/edit/{id} api
func EditChildAccountByMgr(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	accountid = bson.ObjectIdHex(vars["accountid"])
	var dataResource AccountResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid account data",
			500,
		)
		return
	}
	acct := &dataResource.Data
	acct.ParentAccount = accountid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("accounts")
	repo := &data.AccountRepository{C: col}
	
	err := repo.EditChildAccountByID(acct); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//Delete Account - TBA

//AddJournalAccount for /orgs/accounts/journal/accounts/add api
func AddJournalAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountid = bson.ObjectIdHex(vars["accountid"])
	projectid = bson.ObjectIdHex(vars["projectid"])
	var dataRes JournalAccountResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid journal account data",
			500,
		)
		return
	}
	journalAcct := &dataRes.Data
	journalAcct.AccountID = accountid
	journalAcct.ProjectID = projectid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("journalaccounts")
	repo := &data.AccountRepository{C: col}

	repo.AddJournalAcctByAccountID(journalAcct)
	j, err := json.Marshal(JournalAccountResource{Data: *journalAcct})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//AddJournalEntry for /orgs/accounts/journal/entries/add api
func AddJournalEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountid = bson.ObjectIdHex(vars["accountid"])
	var dataRes JournalEntryResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid journal entry data",
			500,
		)
		return
	}
	journalEntry := &dataRes.Data
	journalEntry.JournalAccountID = accountid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("journalentries")
	repo := &data.AccountRepository{C: col}

	repo.AddJournalEntryByAcctID(journalEntry)
	j, err := json.Marshal(JournalAccountResource{Data: *journalEntry})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//GetJournalAccountsByAcct for /orgs/accounts/journal/acccounts/account/{accountId} api
func GetJournalAccountsByAcct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountid = bson.ObjectIdHex(vars["accountid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("journalaccounts")
	repo := &data.AccountRepository{C: col}
	journalAccts, err := repo.GetJournalAccountsByAcctID(accountid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(JournalAccountsResource{Data: journalAccts})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetJournalEntriesByAccount for /orgs/accounts/journal/entries/account/{accountId} api
func GetJournalEntriesByAcct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountid = bson.ObjectIdHex(vars["accountid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("journalentries")
	repo := &data.AccountRepository{C: col}
	journalEntries, err := repo.GetJournalEntriesByAcctID(accountid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(JournalEntriesResource{Data: journalEntries})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetJournalAccountsByOrg for /orgs/accounts/journal/acccounts/org/{orgId} api
func GetJournalAccountsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("journalaccounts")
	repo := &data.AccountRepository{C: col}
	journalAccts, err := repo.GetJournalAccountsByAcctID(orgid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(JournalAccountsResource{Data: journalAccts})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetJournalEntriesByOrg for /orgs/accounts/journal/entries/org/{orgId} api
func GetJournalEntriesByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("journalentries")
	repo := &data.AccountRepository{C: col}
	journalEntries, err := repo.GetJournalEntriesByAcctID(orgid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(JournalEntriesResource{Data: journalEntries})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetJournalAccount for /orgs/accounts/journal/accounts/{id} api
func GetJournalAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("journalaccounts")
	repo := &data.AccountRepository{C: col}
	journalAcct, err := repo.GetJournalAcctByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occured",
				500,
			)
		}
		return
	}
	j, err := json.Marshal(journalAcct)
	if err != nil {
		common.DisplayAppError(
			w, 
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.Header().Set("Content-Application", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetJournalEntry for /orgs/accounts/journal/entries/{id} api
func GetJournalEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("journalentries")
	repo := &data.AccountRepository{C: col}
	journalEntry, err := repo.GetEntryByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occured",
				500,
			)
		}
		return
	}
	j, err := json.Marshal(journalEntry)
	if err != nil {
		common.DisplayAppError(
			w, 
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.Header().Set("Content-Application", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//EditJournalAccountByEmp for /orgs/accounts/journal/accounts/edit/{id} api
func EditJournalAccountByEmp(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource JournalAccountResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid journal account data",
			500,
		)
		return
	}
	journalAcct := &dataResource.Data
	journalAcct.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("journalaccounts")
	repo := &data.AccountRepository{C: col}
	
	err := repo.EditJournalAccountByID(journalAcct); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//EditJournalEntryByEmp for /orgs/accounts/journal/entries/edit/{id} api
func EditJournalEntryByEmp(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource JournalEntryResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid journal entry data",
			500,
		)
		return
	}
	journalEntry := &dataResource.Data
	journalEntry.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("journalentries")
	repo := &data.AccountRepository{C: col}
	
	err := repo.EditJournalEntryByID(journalEntry); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//Delete Journal Account - TBA

//Delete Journal Entry - TBA