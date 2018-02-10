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

//AddOrg for /orgs/add api
func AddOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid = bson.ObjectIdHex(vars["userid"])
	var dataRes OrgResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid org data",
			500,
		)
		return
	}
	org := &dataRes.Data
	org.UserID = userid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("orgs")
	repo := &data.OrgRepository{C: col}

	repo.AddOrgByUserID(org)
	j, err := json.Marshal(OrgResource{Data: *org})
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

//GetOrgs for /admin/orgs api
func GetOrgs(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("orgs")
	repo := &data.OrgRepository{C: col}
	orgs := repo.GetOrgs()

	j, err := json.Marshal(OrgsResource{Data: orgs})
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

//GetOrg for /orgs/{id} api
func GetOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("orgs")
	repo := &data.OrgRepository{C: col}
	org, err := repo.GetOrgByID(id)
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
	j, err := json.Marshal(org)
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

//GetOrgByUser for /orgs/users/{userId} api
func GetOrgByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid = vars["userid"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("orgs")
	repo := &data.OrgRepository{C: col}
	org, err := repo.GetOrgByUserID(userid)
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
	j, err := json.Marshal(org)
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

//EditOrgByAdmin for /admin/orgs/edit/{id} api
func EditOrgByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource OrgResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid org data",
			500,
		)
		return
	}
	org := &dataResource.Data
	org.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("orgs")
	repo := &data.OrgRepository{C: col}
	
	err := repo.EditOrgByID(org); err != nil {
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

//EditOrg for /orgs/edit/{id} api
func EditOrg(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource OrgResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid org data",
			500,
		)
		return
	}
	org := &dataResource.Data
	org.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("orgs")
	repo := &data.OrgRepository{C: col}
	
	err := repo.EditOrgByID(org); err != nil {
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

//Delete Org - TBA