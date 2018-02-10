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

//AddBilling for /orgs/billings/add api
func AddBilling(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes BillingResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid role data",
			500,
		)
		return
	}
	bill := &dataRes.Data
	bill.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("billings")
	repo := &data.BillingRepository{C: col}

	repo.AddBillingByOrgID(bill)
	j, err := json.Marshal(BillingResource{Data: *bill})
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

//GetBills for /orgs/billings api
func GetBillings(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("billings")
	repo := &data.BillingRepository{C: col}
	bills := repo.GetBillings()

	j, err := json.Marshal(BillingsResource{Data: bills})
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

//GetBillingsByOrg for /orgs/billings/org/{orgId} api
func GetBillingsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("billings")
	repo := &data.BillingRepository{C: col}
	bills, err := repo.GetBillingsByOrgID(orgid)
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

	j, err := json.Marshal(BillingsResource{Data: bills})
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

//GetBilling for /orgs/billings/{id} api
func GetBilling(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("billings")
	repo := &data.BillingRepository{C: col}
	bill, err := repo.GetBillingByID(id)
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
	j, err := json.Marshal(bill)
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

//EditBillingByAdmin for /admin/orgs/billings/edit/{id} api
func EditBillingByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource BillingResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid bill data",
			500,
		)
		return
	}
	bill := &dataResource.Data
	bill.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("billings")
	repo := &data.BillingRepository{C: col}
	
	err := repo.EditBillingByID(bill); err != nil {
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

//AddTransaction for /orgs/billings/transactions/add api
func AddTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	billid = bson.ObjectIdHex(vars["billingid"])
	var dataRes TransactionResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid tx data",
			500,
		)
		return
	}
	tx := &dataRes.Data
	tx.BillingID = billid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("transactions")
	repo := &data.BillingRepository{C: col}

	repo.AddTransactionByBillingID(tx)
	j, err := json.Marshal(TransactionResource{Data: *tx})
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

//GetTransactions for /orgs/billings/transactions api
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("transactions")
	repo := &data.BillingRepository{C: col}
	txs := repo.GetTransactions()

	j, err := json.Marshal(TransactionsResource{Data: txs})
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

//GetTransactionsByBilling for /org/billings/transactions/billing/{billingId} api
func GetTransactionsByBilling(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	billid = bson.ObjectIdHex(vars["billingid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("transactions")
	repo := &data.BillingRepository{C: col}
	txs, err := repo.GetTransactionsByBillingID(billid)
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

	j, err := json.Marshal(TransactionsResource{Data: txs})
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

//GetTransaction for /orgs/billings/transactions/{id} api
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("transactions")
	repo := &data.BillingRepository{C: col}
	tx, err := repo.GetTransactionByID(id)
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
	j, err := json.Marshal(tx)
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

//EditTransactionByAdmin for /admin/orgs/billings/transactions/edit/{id} api
func EditTransactionByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource TransactionResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid bill data",
			500,
		)
		return
	}
	tx := &dataResource.Data
	tx.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("transactions")
	repo := &data.BillingRepository{C: col}
	
	err := repo.EditTransactionByID(tx); err != nil {
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
