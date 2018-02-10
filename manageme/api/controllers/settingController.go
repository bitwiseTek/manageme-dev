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

//AddDepartment for /orgs/settings/departments/add api
func AddDepartment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes SettingResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid dept data",
			500,
		)
		return
	}
	dept := &dataRes.Data
	dept.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("departments")
	repo := &data.SettingRepository{C: col}

	repo.AddDeptByOrgID(dept)
	j, err := json.Marshal(SettingResource{Data: *dept})
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

//AddDesignation for /orgs/settings/designations/add api
func AddDesignation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes SettingResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid desgn data",
			500,
		)
		return
	}
	desgn := &dataRes.Data
	desgn.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("designations")
	repo := &data.SettingRepository{C: col}

	repo.AddDesgnByOrgID(desgn)
	j, err := json.Marshal(SettingResource{Data: *desgn})
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

//AddSalaryMode for /orgs/settings/salary/modes/add api
func AddSalaryMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes SettingResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid salary mode data",
			500,
		)
		return
	}
	salMode := &dataRes.Data
	salMode.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarymodes")
	repo := &data.SettingRepository{C: col}

	repo.AddSalaryModeByOrgID(desgn)
	j, err := json.Marshal(SettingResource{Data: *salMode})
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

//AddBranch for /orgs/settings/branches/add api
func AddBranch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes SettingResource

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
	branch := &dataRes.Data
	branch.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("branches")
	repo := &data.SettingRepository{C: col}

	repo.AddBranchByOrgID(branch)
	j, err := json.Marshal(SettingResource{Data: *branch})
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

//AddLeaveType for /orgs/settings/leave/types/add api
func AddLeaveType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes SettingResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid leave type data",
			500,
		)
		return
	}
	leaveType := &dataRes.Data
	leaveType.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leavetypes")
	repo := &data.SettingRepository{C: col}

	repo.AddLeaveTypeByOrgID(leaveType)
	j, err := json.Marshal(SettingResource{Data: *leaveType})
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

//AddExpenseClaimType for /orgs/settings/claim/types/add api
func AddClaimType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes SettingResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid claim type data",
			500,
		)
		return
	}
	claimType := &dataRes.Data
	claimType.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("claimtypes")
	repo := &data.SettingRepository{C: col}

	repo.AddDesgnByOrgID(desgn)
	j, err := json.Marshal(SettingResource{Data: *claimType})
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

//AddProjectType for /orgs/settings/project/types/add api
func AddDesignation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes SettingResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid project type data",
			500,
		)
		return
	}
	projectType := &dataRes.Data
	projectType.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projectType")
	repo := &data.SettingRepository{C: col}

	repo.AddProjectTypeByOrgID(projectType)
	j, err := json.Marshal(SettingResource{Data: *projectType})
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

//GetDeptsByOrg for /orgs/settings/departments api
func GetDeptsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("departments")
	repo := &data.SettingRepository{C: col}
	depts, err := repo.GetDeptsByOrgID(orgid)
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

	j, err := json.Marshal(DepartmentsResource{Data: depts})
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

//GetSalaryModesByOrg for /orgs/settings/salary/modes api
func GetSalaryModesByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarymodes")
	repo := &data.SettingRepository{C: col}
	salaryModes, err := repo.GetModesByOrgID(orgid)
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

	j, err := json.Marshal(SalaryModesResource{Data: salaryModes})
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

//GetDesignationsByOrg for /orgs/settings/designations api
func GetDesignationsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("designations")
	repo := &data.SettingRepository{C: col}
	designations, err := repo.GetDesignationsByOrgID(orgid)
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

	j, err := json.Marshal(DesignationsResource{Data: designations})
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

//GetLeaveTypesByOrg for /orgs/settings/leave/types api
func GetLeaveTypesByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leavetypes")
	repo := &data.SettingRepository{C: col}
	leaveTypes, err := repo.GetLeaveTypesByOrgID(orgid)
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

	j, err := json.Marshal(LeaveTypesResource{Data: leaveTypes})
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

//GetBranchesByOrg for /orgs/settings/branches api
func GetBranchesByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("branches")
	repo := &data.SettingRepository{C: col}
	branches, err := repo.GetBranchesByOrgID(orgid)
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

	j, err := json.Marshal(BranchesResource{Data: branches})
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

//GetClaimTypesByOrg for /orgs/settings/claim/types api
func GetClaimTypesByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("claimtypes")
	repo := &data.SettingRepository{C: col}
	claimTypes, err := repo.GetClaimTypesByOrgID(orgid)
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

	j, err := json.Marshal(ClaimTypesResource{Data: claimTypes})
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

//GetProjectTypesByOrg for /orgs/settings/project/types api
func GetProjectTypesByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projecttypes")
	repo := &data.SettingRepository{C: col}
	projectTypes, err := repo.GetProjectTypesByOrgID(orgid)
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

	j, err := json.Marshal(ProjectTypesResource{Data: projectTypes})
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

//GetDept for /orgs/settings/departments/{id} api
func GetDepartment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("departments")
	repo := &data.SettingRepository{C: col}
	dept, err := repo.GetDeptByID(id)
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
	j, err := json.Marshal(dept)
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

//GetDesignation for /orgs/settings/designations/{id} api
func GetDesignation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("designations")
	repo := &data.SettingRepository{C: col}
	desgn, err := repo.GetDesgnByID(id)
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
	j, err := json.Marshal(desgn)
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

//GetBranch for /orgs/settings/branches/{id} api
func GetDepartment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("branches")
	repo := &data.SettingRepository{C: col}
	branch, err := repo.GetBranchByID(id)
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
	j, err := json.Marshal(branch)
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

//GetSalaryMode for /orgs/settings/salary/modes/{id} api
func GetSalaryMode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("salarymodes")
	repo := &data.SettingRepository{C: col}
	salaryMode, err := repo.GetSalaryModeByID(id)
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
	j, err := json.Marshal(salaryMode)
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

//GetLeaveType for /orgs/settings/leave/types/{id} api
func GetLeaveType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("leavetypes")
	repo := &data.SettingRepository{C: col}
	leaveType, err := repo.GetLeaveTypeByID(id)
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
	j, err := json.Marshal(leaveType)
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

//GetClaimType for /orgs/settings/claim/types/{id} api
func GetClaimType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("claimtypes")
	repo := &data.SettingRepository{C: col}
	claimType, err := repo.GetClaimTypeByID(id)
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
	j, err := json.Marshal(claimType)
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

//GetProjectType for /orgs/settings/project/types/{id} api
func GetProjectType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("projecttypes")
	repo := &data.SettingRepository{C: col}
	projectType, err := repo.GetProjectTypeByID(id)
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
	j, err := json.Marshal(projectType)
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

//EditDepatment for /orgs/settings/departments/edit/{id} api
func EditDepartmentByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource DepartmentResource
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
	dept := &dataResource.Data
	dept.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("departments")
	dept := &data.SettingRepository{C: col}
	
	err := repo.EditDepartmentByID(dept); err != nil {
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

//EditDesignation for /orgs/settings/designations/edit/{id} api
func EditDesignationByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource DesignationResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid desgn data",
			500,
		)
		return
	}
	desgn := &dataResource.Data
	desgn.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("designations")
	desgn := &data.SettingRepository{C: col}
	
	err := repo.EditDesignationByID(desgn); err != nil {
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

//EditBranch for /orgs/settings/branches/edit/{id} api
func EditBranchByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource BranchResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid branch data",
			500,
		)
		return
	}
	branch := &dataResource.Data
	branch.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("branches")
	branch := &data.SettingRepository{C: col}
	
	err := repo.EditBranchByID(branch); err != nil {
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

//EditSalaryMode for /orgs/settings/salary/modes/edit/{id} api
func EditSalaryModeByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource SalaryModeResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid salary mode data",
			500,
		)
		return
	}
	salaryMode := &dataResource.Data
	salaryMode.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarymodes")
	salaryMode := &data.SettingRepository{C: col}
	
	err := repo.EditSalaryModeByID(salaryMode); err != nil {
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

//EditLeaveType for /orgs/settings/leave/types/edit/{id} api
func EditLeaveTypeByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource LeaveTypeResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid leave type data",
			500,
		)
		return
	}
	leaveType := &dataResource.Data
	leaveType.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leavetypes")
	leaveType := &data.SettingRepository{C: col}
	
	err := repo.EditLeaveTypeByID(leaveType); err != nil {
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

//EditClaimType for /orgs/settings/claim/types/edit/{id} api
func EditClaimTypeByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource ClaimTypeResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid claim type data",
			500,
		)
		return
	}
	claimType := &dataResource.Data
	claimType.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("claimtypes")
	claimType := &data.SettingRepository{C: col}
	
	err := repo.EditClaimTypeByID(claimType); err != nil {
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

//EditProjectType for /orgs/settings/project/types/edit/{id} api
func EditProjectTypeByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource ProjectTypeResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid project type data",
			500,
		)
		return
	}
	projectType := &dataResource.Data
	projectType.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projecttypes")
	projectType := &data.SettingRepository{C: col}
	
	err := repo.EditProjectTypeByID(projectType); err != nil {
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

//Delete Dept - TBA

//Delete Designation - TBA

//Delete SalaryMode - TBA

//Delete ClaimType - TBA

//Delete ProjectType - TBA

//Delete LeaveType - TBA

//Delete Branch - TBA