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

//AddOrgUser for /orgs/users/add api
func AddOrgUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	roleid = bson.ObjectIdHex(vars["roleid"])
	var dataResource OrgUserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid user data",
			500,
		)
		return
	}
	orgUser := &dataResource.Data
	orgUser.OrgID = orgid
	orgUser.RoleID = roleid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("orgusers")
	repo := &data.OrgUserRepository{C: col}

	repo.AddUserByOrgID(orgUser)
	orgUser.HashPassword = nil
	j, err := json.Marshal(UserResource{Data: *orgUser})
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

//OrgUserSignIn for /orgs/users/signin api
func OrgUserSignIn(w http.ResponseWriter, r *http.Request) {
	var dataRes OrgUserSignInResource
	var token string
	
	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"Invalid login token", 
			500,
		)
		return
	}
	orgUserSignInModel := &dataRes.Data
	orgSignInUser := models.OrgUser{
		Username:	orgUserSignInModel.Username, 
		Password:	orgUserSignInModel.Password,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("orgusers")
	repo := &data.UserRepository{C:	col}

	orgUser, err := repo.SignIn(orgSignInUser)
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"Invalid login credentials", 
			401,
		)
		return
	}

	token, err := common.GenerateJWT(orgUser.Username, "member")
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"Error generating access token", 
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	authOrgUser := AuthOrgUserModel{
		User:	orgUser, 
		Token:	token, 
	}
	j, err := json.Marshal(AuthOrgUserResource{Data: autOrgUser})
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
	w.Write(j)
}

//GetOrgUsers for /orgs/users api
func GetOrgUsers(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("orgusers")
	repo := &data.OrgUserRepository{C: col}
	orgUsers := repo.GetOrgUsers()

	j, err := json.Marshal(OrgUsersResource{Data: orgUsers})
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

//GetUsersByOrg for /orgs/users/org/{orgID} api
func GetUsersByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = vars["orgid"]
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("orgusers")
	repo := &data.OrgUserRepository{C: col}
	orgUsers, err := repo.GetUsersByOrgID(orgid)
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

	j, err := json.Marshal(OrgUsersResource{Data: orgUsers})
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

//GetOrgUser for /orgs/users/{id} api
func GetOrgUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("orgusers")
	repo := &data.OrgUserRepository{C: col}
	orgUser, err := repo.GetOrgUserByID(id)
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
	j, err := json.Marshal(orgUser)
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

//EditOrgUserByAdmin for /orgs/users/edit/{id} api
func EditOrgUserByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource OrgUserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid user data",
			500,
		)
		return
	}
	orgUser := &dataResource.Data
	orgUser.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("orgusers")
	repo := &data.OrgUserRepository{C: col}
	
	err := repo.EditOrgUserByID(orgUser); err != nil {
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

//EditOrgUser for /orgs/users/edit/{userID} api
func EditOrgUser(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource OrgUserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid user data",
			500,
		)
		return
	}
	orgUser := &dataResource.Data
	orgUser.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("orgusers")
	repo := &data.OrgUserRepository{C: col}
	
	err := repo.EditOrgUserByOrgUserID(orgUser); err != nil {
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

//DeleteOrgUser - TBA