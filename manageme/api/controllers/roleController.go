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

//AddRole for /roles/add api
func AddRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes RoleResource

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
	role := &dataRes.Data
	role.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("roles")
	repo := &data.RoleRepository{C: col}

	repo.AddRole(role)
	j, err := json.Marshal(RoleResource{Data: *role})
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

//GetRoles for /roles api
func GetRoles(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("roles")
	repo := &data.RoleRepository{C: col}
	roles := repo.GetRoles()

	j, err := json.Marshal(RolesResource{Data: roles})
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

//GetRolesByOrg for /org/roles api
func GetRolesByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("roles")
	repo := &data.RoleRepository{C: col}
	roles, err := repo.GetRolesByOrgID(orgid)
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

	j, err := json.Marshal(RolesResource{Data: roles})
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

//GetRole for /roles/{id} api
func GetRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("roles")
	repo := &data.RoleRepository{C: col}
	role, err := repo.GetRoleByID(id)
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
	j, err := json.Marshal(role)
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

//EditRole for /roles/edit/{id} api
func EditRoleByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource RoleResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid role data",
			500,
		)
		return
	}
	role := &dataResource.Data
	role.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("roles")
	repo := &data.RoleRepository{C: col}
	
	err := repo.EditRoleByID(role); err != nil {
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

//Delete Role - TBA

//AddPermission for /permissions/add api
func AddPermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes PermissionResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid permission data",
			500,
		)
		return
	}
	perm := &dataRes.Data
	perm.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("permissions")
	repo := &data.RoleRepository{C: col}

	repo.AddPermission(perm)
	j, err := json.Marshal(PermissionResource{Data: *perm})
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

//GetPermissions for /permissions api
func GetPermissions(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("permissions")
	repo := &data.RoleRepository{C: col}
	perms := repo.GetPermissions()

	j, err := json.Marshal(PermissionsResource{Data: perms})
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

//GetPermissionsByOrg for /org/permissions api
func GetPermissionsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("permissions")
	repo := &data.RoleRepository{C: col}
	perms, err := repo.GetPermissionsByOrgID(orgid)
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

	j, err := json.Marshal(PermissionsResource{Data: perms})
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

//GetPermission for /permissions/{id} api
func GetPermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("permissions")
	repo := &data.RoleRepository{C: col}
	perm, err := repo.GetPermByID(id)
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
	j, err := json.Marshal(perm)
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

//EditPermission for /permissions/edit/{id} api
func EditPermissionByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource PermissionResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid permission data",
			500,
		)
		return
	}
	perm := &dataResource.Data
	perm.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("permissions")
	repo := &data.RoleRepository{C: col}
	
	err := repo.EditPermissionByID(perm); err != nil {
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

//Delete Permission - TBA


