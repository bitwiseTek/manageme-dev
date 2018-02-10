package controllers

/**
 *
 * @author Sika Kay
 * @date 07/02/18
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

//AddEmployee for /orgs/employees/add api
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	orguserid = bson.ObjectIdHex(vars["orguserid"])
	branchid = bson.ObjectIdHex(vars["branchid"])
	deptid = bson.ObjectIdHex(vars["departmentid"])
	var dataRes EmployeeResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid employee data",
			500,
		)
		return
	}
	emp := &dataRes.Data
	emp.OrgUserID = orguserid
	emp.OrgID = orgid
	emp.BranchID = branchid
	emp.DepartmentID = deptid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("employees")
	repo := &data.EmployeeRepository{C: col}

	repo.AddEmpByOrgUserID(emp)
	j, err := json.Marshal(EmployeeResource{Data: *emp})
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

//AddChildEmployee for /orgs/employees/child/add api
func AddChildEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orguserid = bson.ObjectIdHex(vars["orguserid"])
	branchid = bson.ObjectIdHex(vars["branchid"])
	deptid = bson.ObjectIdHex(vars["departmentid"])
	empid = bson.ObjectIdHex(vars["empid"])
	var dataRes EmployeeResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid employee data",
			500,
		)
		return
	}
	emp := &dataRes.Data
	emp.OrgUserID = orguserid
	emp.BranchID = branchid
	emp.DepartmentID = deptid
	emp.ReportsTo = empid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("employees")
	repo := &data.EmployeeRepository{C: col}

	repo.AddChildEmpByOrgUserID(emp)
	j, err := json.Marshal(EmployeeResource{Data: *emp})
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

//GetEmployees for /orgs/employees api
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("employees")
	repo := &data.EmployeeRepository{C: col}
	emps := repo.GetEmployees()

	j, err := json.Marshal(EmployeesResource{Data: emps})
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

//GetEmployeesByOrg for /orgs/employees/org/{orgId} api
func GetEmployeesByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("employees")
	repo := &data.EmployeeRepository{C: col}
	emps, err := repo.GetEmpsByOrgID(orgid)
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

	j, err := json.Marshal(EmployeesResource{Data: emps})
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

//GetEmployeesByDept for /orgs/employees/department/{deptId} api
func GetEmployeesByDept(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deptid = bson.ObjectIdHex(vars["departmentid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("employees")
	repo := &data.EmployeeRepository{C: col}
	emps, err := repo.GetEmpsByDeptID(deptid)
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

	j, err := json.Marshal(EmployeesResource{Data: emps})
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

//GetEmployeesByBranch for /orgs/employees/branch/{branchId} api
func GetEmployeesByBranch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	branchid = bson.ObjectIdHex(vars["branchid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("employees")
	repo := &data.EmployeeRepository{C: col}
	emps, err := repo.GetEmpsByBranchID(branchid)
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

	j, err := json.Marshal(EmployeesResource{Data: emps})
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

//GetEmployeesByMgr for /orgs/employees/child/{empId} api
func GetEmployeesByMgr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("employees")
	repo := &data.EmployeeRepository{C: col}
	emps, err := repo.GetEmpsByMgrID(empid)
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

	j, err := json.Marshal(EmployeesResource{Data: emps})
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

//GetEmployee for /orgs/employees/{id} api
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("employees")
	repo := &data.ProjectRepository{C: col}
	emp, err := repo.GetEmpByID(id)
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
	j, err := json.Marshal(emp)
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

//GetEmployeeByUser for /orgs/employees/user/{userId} api
func GetEmployeeByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid = vars["orguserid"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("employees")
	repo := &data.ProjectRepository{C: col}
	emp, err := repo.GetEmpByUserID(userid)
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
	j, err := json.Marshal(emp)
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

//EditEmpByAdmin for /orgs/employees/edit/{id} api
func EditEmployeeByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource EmployeeResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid employee data",
			500,
		)
		return
	}
	emp := &dataResource.Data
	emp.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("employees")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditEmpByID(emp); err != nil {
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

//EditEmpByUser for /orgs/employees/user/edit/{id} api
func EditEmployeeByUser(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	userid = bson.ObjectIdHex(vars["orguserid"])
	var dataResource EmployeeResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid employee data",
			500,
		)
		return
	}
	emp := &dataResource.Data
	emp.OrgUserID = userid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("employees")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditEmpByID(emp); err != nil {
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

//Delete Employee - TBA

//AddBiodataByEmp for /orgs/employees/biodata/add
func AddBiodataByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	var dataRes BiodataResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid bio data",
			500,
		)
		return
	}
	bio := &dataRes.Data
	bio.EmployeeID = empid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("biodatas")
	repo := &data.EmployeeRepository{C: col}

	repo.AddBiodataByEmpID(bio)
	j, err := json.Marshal(BiodataResource{Data: *bio})
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

//AddPIDByBio for /orgs/employees/pid/add
func AddPIDByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	var dataRes PIDResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid pid data",
			500,
		)
		return
	}
	pid := &dataRes.Data
	pid.BioID = bioid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("pids")
	repo := &data.EmployeeRepository{C: col}

	repo.AddPIDByBioID(pid)
	j, err := json.Marshal(PIDResource{Data: *pid})
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

//AddHealthDetailByBio for /orgs/employees/health/detail/add
func AddHealthDetByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	var dataRes HealthDetailResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid pid data",
			500,
		)
		return
	}
	health := &dataRes.Data
	health.BioID = bioid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("pids")
	repo := &data.EmployeeRepository{C: col}

	repo.AddHealthDetByBioID(health)
	j, err := json.Marshal(PIDResource{Data: *pid})
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

//AddWorkExperienceByBio for /orgs/employees/work/experience/add
func AddWorkExpByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	var dataRes WorkExperienceResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid work data",
			500,
		)
		return
	}
	work := &dataRes.Data
	work.BioID = bioid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("workexperiences")
	repo := &data.EmployeeRepository{C: col}

	repo.AddWorkByBioID(work)
	j, err := json.Marshal(WorkExperienceResource{Data: *work})
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

//AddEduByBio for /orgs/employees/education/add
func AddEduByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	var dataRes EducationResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid pid data",
			500,
		)
		return
	}
	edu := &dataRes.Data
	edu.BioID = bioid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("educations")
	repo := &data.EmployeeRepository{C: col}

	repo.AddEduByBioID(edu)
	j, err := json.Marshal(EducationResource{Data: *edu})
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

//GetBiosByEmp for /orgs/employees/biodatas/employee/{empId} api
func GetBiosByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("biodatas")
	repo := &data.EmployeeRepository{C: col}
	bios, err := repo.GetBiosByEmpID(empid)
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

	j, err := json.Marshal(BiodatasResource{Data: bios})
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

//GetBio for /orgs/employees/biodatas/{id} api
func GetBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("biodatas")
	repo := &data.EmployeeRepository{C: col}
	bio, err := repo.GetBioByID(id)
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
	j, err := json.Marshal(bio)
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

//GetBioByEmp for /orgs/employees/biodatas/employee/{empId} api
func GetBioByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("biodatas")
	repo := &data.EmployeeRepository{C: col}
	bio, err := repo.GetBioByEmpID(empid)
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
	j, err := json.Marshal(bio)
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

//EditBioByEmp for /orgs/employees/biodatas/employee/edit/{id} api
func EditBioByEmp(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	var dataResource BiodataResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid task data",
			500,
		)
		return
	}
	bio := &dataResource.Data
	bio.EmployeeID = empid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("biodatas")
	repo := &data.AccountRepository{C: col}
	
	err := repo.EditBioByEmpID(bio); err != nil {
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

//EditBioByAdmin for /orgs/employees/biodatas/edit/{id} api
func EditBioByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource BiodataResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid bio data",
			500,
		)
		return
	}
	bio := &dataResource.Data
	bio.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("biodatas")
	repo := &data.AccountRepository{C: col}
	
	err := repo.EditBioByID(bio); err != nil {
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

//Delete Bio - TBA

//GetPIDsByBio for /orgs/employees/pids/biodata/{bioId} api
func GetPIDsByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("pids")
	repo := &data.EmployeeRepository{C: col}
	pids, err := repo.GetPIDsByBioID(bioid)
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

	j, err := json.Marshal(PIDsResource{Data: pids})
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

//GetPID for /orgs/employees/pids/{id} api
func GetPID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("pids")
	repo := &data.EmployeeRepository{C: col}
	pid, err := repo.GetPIDByID(id)
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
	j, err := json.Marshal(pid)
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

//GetPIDByBio for /orgs/employees/pids/biodata/{bioId} api
func GetPIDByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("biodatas")
	repo := &data.EmployeeRepository{C: col}
	pid, err := repo.GetPIDByBioID(bioid)
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
	j, err := json.Marshal(pid)
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

//EditPIDByBio for /orgs/employees/pids/biodata/edit/{bioId} api
func EditPIDByBio(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	var dataResource PIDResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid pid data",
			500,
		)
		return
	}
	pid := &dataResource.Data
	pid.BioID = bioid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("pids")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditPIDByBioID(pid); err != nil {
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

//EditPIDByAdmin for /orgs/employees/pids/edit/{bioId} api
func EditPIDByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource PIDResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid pid data",
			500,
		)
		return
	}
	pid := &dataResource.Data
	pid.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("pids")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditPIDByID(pid); err != nil {
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

//Delete PID - TBA

//GetHealthDetailsByBio for /orgs/employees/health/details/{bioId} api
func GetHealthDetsByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("healthdetails")
	repo := &data.EmployeeRepository{C: col}
	healthDets, err := repo.GetHealthDetailsByBioID(bioid)
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

	j, err := json.Marshal(HealthDetailsResource{Data: healthDets})
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

//GetHealthDetail for /orgs/employees/health/details/{id} api
func GetHealthDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("healthdetails")
	repo := &data.EmployeeRepository{C: col}
	healthDet, err := repo.GetHealthDetailByID(id)
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
	j, err := json.Marshal(healthDet)
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

//GetHealthDetailByBio for /orgs/employees/health/details/biodata/{bioId} api
func GetHealthDetailByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("healthdetails")
	repo := &data.EmployeeRepository{C: col}
	healthDet, err := repo.GetHealthDetailByBioID(bioid)
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
	j, err := json.Marshal(healthDet)
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

//EditHealthDetailByBio for /orgs/employees/health/details/edit/biodata/{bioId} api
func EditHealthDetailByBio(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	var dataResource HealthDetailResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid health data",
			500,
		)
		return
	}
	healthDet := &dataResource.Data
	healthDet.BioID = bioid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("healthdetails")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditHealthDetailByBioID(healthDet); err != nil {
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

//EditHealthDetailByAdmin for /orgs/employees/health/details/edit/{id} api
func EditHealthDetailByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource HealthDetailResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid health data",
			500,
		)
		return
	}
	healthDet := &dataResource.Data
	healthDet.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("healthdetails")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditHealthDetailByBioID(healthDet); err != nil {
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

//Delete Health Detail - TBA

//GetWorkExpereincesByBio for /orgs/employees/work/experiences/biodata/{bioId} api
func GetWorkExpsByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("workexperiences")
	repo := &data.EmployeeRepository{C: col}
	workExps, err := repo.GetWorksByBioID(bioid)
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

	j, err := json.Marshal(WorkExperiencesResource{Data: workExps})
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

//GetWorkExperience for /orgs/employees/work/experiences/{id} api
func GetWorkExperience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("workexperiences")
	repo := &data.EmployeeRepository{C: col}
	workExp, err := repo.GetWorkByID(id)
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
	j, err := json.Marshal(workExp)
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

//GetWorkExperienceByBio for /orgs/employees/work/experiences/biodata/{bioId} api
func GetWorkExperienceByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("workexperiences")
	repo := &data.EmployeeRepository{C: col}
	workExp, err := repo.GetWorkByBioID(bioid)
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
	j, err := json.Marshal(workExp)
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

//EditWorkExperienceByBio for /orgs/employees/work/experiences/edit/biodata/{bioId} api
func EditWorkExperienceByBio(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	var dataResource WorkExperienceResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid health data",
			500,
		)
		return
	}
	workExp := &dataResource.Data
	workExp.BioID = bioid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("workexperiences")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditWorkByBioID(bioid); err != nil {
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

//EditWorkExperienceByAdmin for /orgs/employees/work/experiences/edit/{id} api
func EditWorkExperienceByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource WorkExperienceResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid health data",
			500,
		)
		return
	}
	workExp := &dataResource.Data
	workExp.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("workexperiences")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditWorkByID(id); err != nil {
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

//Delete Work Experience - TBA

//GetEducationsByBio for /orgs/employees/educations/biodata/{bioId} api
func GetEdusByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("educations")
	repo := &data.EmployeeRepository{C: col}
	edus, err := repo.GetEdusByBioID(bioid)
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

	j, err := json.Marshal(EducationsResource{Data: edus})
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

//GetEducation for /orgs/employees/educations/{id} api
func GetEducation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("educations")
	repo := &data.EmployeeRepository{C: col}
	edu, err := repo.GetEduByID(id)
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
	j, err := json.Marshal(edu)
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

//GetEduByBio for /orgs/employees/educations/biodata/{bioId} api
func GetEduByBio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("educations")
	repo := &data.EmployeeRepository{C: col}
	edu, err := repo.GetEduByBioID(bioid)
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
	j, err := json.Marshal(edu)
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

//EditEduByBio for /orgs/employees/educations/edit/biodata/{bioId} api
func EditEduByBio(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	bioid = bson.ObjectIdHex(vars["bioid"])
	var dataResource EducationResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid health data",
			500,
		)
		return
	}
	edu := &dataResource.Data
	edu.BioID = bioid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("educations")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditEduByBioID(bioid); err != nil {
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

//EditEduByAdmin for /orgs/employees/educations/edit/{id} api
func EditEduByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource EducationResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid health data",
			500,
		)
		return
	}
	edu := &dataResource.Data
	edu.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("educations")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditEduByID(id); err != nil {
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

//Delete Education - TBA

//AddExpenseClaimByEmp for /orgs/employees/expense/claims/add api
func AddExpenseClaimByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	projectid = bson.ObjectIdHex(vars["projectid"])
	taskid = bson.ObjectIdHex(vars["taskid"])	
	typeid = bson.ObjectIdHex(vars["exptypeid"])
	accountid = bson.ObjectIdHex(vars["accountid"])
	var dataRes ExpenseClaimResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid expense claim data",
			500,
		)
		return
	}
	exp := &dataRes.Data
	exp.OrgID = orgid
	exp.ExpApplier = empid
	exp.ExpApprover = mgrid
	exp.PayableAccount = accountid
	exp.ProjectID = projectid
	exp.TaskID = taskid
	exp.ExpenseType = typeid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("expenseclaims")
	repo := &data.EmployeeRepository{C: col}

	repo.AddExpenseClaimByEmpID(exp)
	j, err := json.Marshal(ExpenseClaimResource{Data: *exp})
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

//AddLeaveAllocationByEmp for /orgs/employees/leave/allocations/add api
func AddLeaveAllocationByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	typeid = bson.ObjectIdHex(vars["leavetypeid"])
	var dataRes LeaveAllocationResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid leave allocation data",
			500,
		)
		return
	}
	leaveAll := &dataRes.Data
	leaveAll.OrgID = orgid
	leaveAll.LeaveAllocator = mgrid
	leaveAll.LeaveReceiver = empid
	leaveAll.LeaveType = typeid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveallocations")
	repo := &data.EmployeeRepository{C: col}

	repo.AddLeaveAllocationByEmpID(leaveAll)
	j, err := json.Marshal(LeaveAllocationResource{Data: *leaveAll})
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

//AddLeaveApplicationByEmp for /orgs/employees/leave/applications/add api
func AddLeaveApplicationByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	typeid = bson.ObjectIdHex(vars["leavetypeid"])
	var dataRes LeaveApplicationResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid leave aaplication data",
			500,
		)
		return
	}
	leaveApp := &dataRes.Data
	leaveApp.OrgID = orgid
	leaveApp.LeaveApprover = mgrid
	leaveApp.LeaveApplier = empid
	leaveApp.LeaveType = typeid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveapplications")
	repo := &data.EmployeeRepository{C: col}

	repo.AddLeaveApplicationByEmpID(leaveApp)
	j, err := json.Marshal(LeaveApplicationResource{Data: *leaveApp})
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

//AddLeaveBlockListByOrg for /orgs/employees/leave/blocklist/add api
func AddLeaveBlockListByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes LeaveBlockListResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid leave block list data",
			500,
		)
		return
	}
	leaveBlk := &dataRes.Data
	leaveBlk.OrgID = orgid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveblocklists")
	repo := &data.EmployeeRepository{C: col}

	repo.AddLeaveBlockListByOrgID(leaveBlk)
	j, err := json.Marshal(LeaveBlockListResource{Data: *leaveBlk})
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

//AddHolidayListByOrg for /orgs/employees/holiday/list/add api
func AddHolidayListByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes HolidayListResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid holiday list data",
			500,
		)
		return
	}
	hol := &dataRes.Data
	hol.OrgID = orgid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("holidaylists")
	repo := &data.EmployeeRepository{C: col}

	repo.AddHolidayListByOrgID(hol)
	j, err := json.Marshal(HolidayListResource{Data: *hol})
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

//AddAppraisalByEmp for /orgs/employees/appraisals/add api
func AddAppraisalByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	var dataRes AppraisalResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid appraisal data",
			500,
		)
		return
	}
	app := &dataRes.Data
	app.OrgID = orgid
	app.ForEmployee = empid
	app.ApGenerator = mgrid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("appraisals")
	repo := &data.EmployeeRepository{C: col}

	repo.AddAppraisalByEmpID(app)
	j, err := json.Marshal(AppraisalResource{Data: *app})
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

//AddExitByEmp for /orgs/employees/exits/add api
func AddExitByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	var dataRes ExitResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid exit data",
			500,
		)
		return
	}
	ex := &dataRes.Data
	ex.OrgID = orgid
	ex.ForEmployee = empid
	ex.OverseenBy = mgrid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("exits")
	repo := &data.EmployeeRepository{C: col}

	repo.AddExitByEmpID(app)
	j, err := json.Marshal(ExitResource{Data: *ex})
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

//GetExpenseClaimsByOrg for /orgs/employees/expense/claims/org/{orgId} api
func GetExpenseClaimsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("expenseclaims")
	repo := &data.EmployeeRepository{C: col}
	claims, err := repo.GetClaimsByOrgID(orgid)
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

	j, err := json.Marshal(ExpenseClaimsResource{Data: claims})
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

//GetExpenseClaimsByEmp for /orgs/employees/expense/claims/employee/{empId} api
func GetExpenseClaimsByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("expenseclaims")
	repo := &data.EmployeeRepository{C: col}
	claims, err := repo.GetClaimsByEmpID(empid)
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

	j, err := json.Marshal(ExpenseClaimsResource{Data: claims})
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

//GetExpenseClaimsByMgr for /orgs/employees/expense/claims/employees/{mgrId} api
func GetExpenseClaimsByMgr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("expenseclaims")
	repo := &data.EmployeeRepository{C: col}
	claims, err := repo.GetClaimsByOrgID(mgrid)
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

	j, err := json.Marshal(ExpenseClaimsResource{Data: claims})
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

//GetExpenseClaim for /orgs/employees/expense/claims/{id} api
func GetExpenseClaim(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("expenseclaims")
	repo := &data.EmployeeRepository{C: col}
	exp, err := repo.GetClaimByID(id)
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
	j, err := json.Marshal(exp)
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

//EditExpenseClaimByMgr for /orgs/employees/expense/claims/employees/edit/{mgrId} api
func EditExpenseClaimByMgr(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	var dataResource ExpenseClaimResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid expense claim data",
			500,
		)
		return
	}
	exp := &dataResource.Data
	exp.ExpApprover = mgrid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("expenseclaims")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditClaimByMgrID(exp); err != nil {
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

//EditExpenseClaimByEmp for /orgs/employees/expense/claims/employee/edit/{empId} api
func EditExpenseClaimByEmp(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	var dataResource ExpenseClaimResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid expense claim data",
			500,
		)
		return
	}
	exp := &dataResource.Data
	exp.ExpApplier = empid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("expenseclaims")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditClaimByEmpID(exp); err != nil {
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

//Delete Expense Claim - TBA

//GetLeaveAllocationsByOrg for /orgs/employees/leave/allocations/org/{orgId} api
func GetLeaveAllocationsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveallocations")
	repo := &data.EmployeeRepository{C: col}
	leaveAlls err := repo.GetLeaveAllocationsByOrgID(orgid)
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

	j, err := json.Marshal(LeaveAllocationsResource{Data: leaveAlls})
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

//GetLeaveAllocationsByMgr for /orgs/employees/leave/allocations/employees/{mgrId} api
func GetLeaveAllocationsByMgr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveallocations")
	repo := &data.EmployeeRepository{C: col}
	leaveAlls, err := repo.GetLeaveAllocationsByEmpID(mgrid)
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

	j, err := json.Marshal(LeaveAllocationsResource{Data: leaveAlls})
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

//GetLeaveAllocation for /orgs/employees/leave/allocations/{id} api
func GetLeaveAllocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("leaveallocations")
	repo := &data.EmployeeRepository{C: col}
	leaveAll, err := repo.GetLeaveAllocationByID(id)
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
	j, err := json.Marshal(leaveAll)
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

//EditLeaveAllocationByAdmin for /orgs/employees/leave/allocations/edit/{id} api
func EditLeaveAllocationByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource LeaveAllocationResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid leave allocation data",
			500,
		)
		return
	}
	leaveAll := &dataResource.Data
	leaveAll.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveallocations")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditLeaveAllocationByID(exp); err != nil {
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

//Delete Leave Allocation - TBA

//GetLeaveApplicationsByOrg for /orgs/employees/leave/applications/org/{orgId} api
func GetLeaveApplicationsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveapplications")
	repo := &data.EmployeeRepository{C: col}
	leaveApps err := repo.GetLeaveApplicationsByOrgID(orgid)
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

	j, err := json.Marshal(LeaveApplicationsResource{Data: leaveApps})
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

//GetLeaveApplicationsByMgr for /orgs/employees/leave/applications/employees/{mgrId} api
func GetLeaveApplicationsByMgr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveapplications")
	repo := &data.EmployeeRepository{C: col}
	leaveApps, err := repo.GetLeaveApplicationsByMgrID(mgrid)
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

	j, err := json.Marshal(LeaveApplicationsResource{Data: leaveApps})
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

//GetLeaveApplicationsByEmp for /orgs/employees/leave/applications/employee/{empId} api
func GetLeaveApplicationsByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveapplications")
	repo := &data.EmployeeRepository{C: col}
	leaveApps, err := repo.GetLeaveApplicationsByEmpID(empid)
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

	j, err := json.Marshal(LeaveApplicationsResource{Data: leaveApps})
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

//GetLeaveApplication for /orgs/employees/leave/applications/{id} api
func GetLeaveApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("leaveapplications")
	repo := &data.EmployeeRepository{C: col}
	leaveApp, err := repo.GetLeaveApplicationByID(id)
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
	j, err := json.Marshal(leaveApp)
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

//EditLeaveApplicationByAdmin for /orgs/employees/leave/applications/edit/{id} api
func EditLeaveApplicationByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource LeaveApplicationResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid leave application data",
			500,
		)
		return
	}
	leaveApp := &dataResource.Data
	leaveApp.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveapplications")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditLeaveApplicationByID(leaveApp); err != nil {
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

//Delete Leave Application - TBA

//GetLeaveBlockListsByOrg for /orgs/employees/leave/blocklists/org/{orgId} api
func GetLeaveBlockListsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveblocklists")
	repo := &data.EmployeeRepository{C: col}
	leaveBlk, err := repo.GetLeaveBlockListsByOrgID(orgid)
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

	j, err := json.Marshal(LeaveBlockListsResource{Data: leaveBlk})
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

//GetLeaveBlockList for /orgs/employees/leave/blocklists/{id} api
func GetLeaveBlockList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("leaveblocklists")
	repo := &data.EmployeeRepository{C: col}
	leaveBlk, err := repo.GetLeaveBlockListByID(id)
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
	j, err := json.Marshal(leaveBlk)
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

//EditLeaveBlockListByAdmin for /orgs/employees/leave/blocklists/edit/{id} api
func EditLeaveBlockListByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource LeaveBlockListResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid leave block list data",
			500,
		)
		return
	}
	leaveBlk := &dataResource.Data
	leaveBlk.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("leaveblocklists")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditLeaveBlockListByID(leaveBlk); err != nil {
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

//Delete Leave Block List - TBA

//GetHolidayListsByOrg for /orgs/employees/holiday/lists/org/{orgId} api
func GetHolidayListsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("holidaylists")
	repo := &data.EmployeeRepository{C: col}
	hols, err := repo.GetHolidayListsByOrgID(orgid)
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

	j, err := json.Marshal(HolidayListsResource{Data: hols})
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

//GetHolidayList for /orgs/employees/holiday/lists/{id} api
func GetHolidayList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("holidaylists")
	repo := &data.EmployeeRepository{C: col}
	hol, err := repo.GetHolidayListByID(id)
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
	j, err := json.Marshal(hol)
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

//EditHolidayListByAdmin for /orgs/employees/holiday/lists/edit/{id} api
func EditHolidayListByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource HolidayListResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid leave block list data",
			500,
		)
		return
	}
	hol := &dataResource.Data
	hol.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("holidaylists")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditHolidayListByID(hol); err != nil {
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

//Delete Holiday List - TBA

//GetAppraisalsByOrg for /orgs/employees/appraisals/org/{orgId} api
func GetAppraisalsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("appraisals")
	repo := &data.EmployeeRepository{C: col}
	apps, err := repo.GetAppraisalsByOrgID(orgid)
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

	j, err := json.Marshal(AppraisalsResource{Data: apps})
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

//GetAppraisalsByEmp for /orgs/employees/appraisals/employee/{empId} api
func GetAppraisalsByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("appraisals")
	repo := &data.EmployeeRepository{C: col}
	apps, err := repo.GetAppraisalsByEmpID(empid)
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

	j, err := json.Marshal(AppraisalsResource{Data: apps})
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

//GetAppraisal for /orgs/employees/appraisals/{id} api
func GetAppraisal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("appraisals")
	repo := &data.EmployeeRepository{C: col}
	app, err := repo.GetAppraisalByID(id)
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
	j, err := json.Marshal(app)
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

//EditAppraisalByAdmin for /orgs/employees/appraisals/edit/{id} api
func EditAppraisalByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource AppraisalResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid appraisal data",
			500,
		)
		return
	}
	app := &dataResource.Data
	app.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("appraisals")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditAppraisalByID(app); err != nil {
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

//Delete Appraisal - TBA

//GetExitsByOrg for /orgs/employees/exits/org/{orgId} api
func GetExitsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("exits")
	repo := &data.EmployeeRepository{C: col}
	exs, err := repo.GetExitsByOrgID(orgid)
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

	j, err := json.Marshal(ExitsResource{Data: exs})
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

//GetExitsByEmp for /orgs/employees/exits/employee/{empId} api
func GetExitsByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("exits")
	repo := &data.EmployeeRepository{C: col}
	exs, err := repo.GetExitsByEmpID(empid)
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

	j, err := json.Marshal(ExitsResource{Data: apps})
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

//GetExit for /orgs/employees/exits/{id} api
func GetExit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("exits")
	repo := &data.EmployeeRepository{C: col}
	ex, err := repo.GetExitByID(id)
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
	j, err := json.Marshal(ex)
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

//EditExitByAdmin for /orgs/employees/exits/edit/{id} api
func EditExitByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource ExitResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid exit data",
			500,
		)
		return
	}
	ex := &dataResource.Data
	ex.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("exits")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditExitByID(ex); err != nil {
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

//Delete Exit - TBA

//AddSalaryComponentByOrg for /orgs/employees/salary/components/add api
func AddSalaryComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes SalaryComponentResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid component data",
			500,
		)
		return
	}
	sal := &dataRes.Data
	sal.OrgID = orgid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarycomponents")
	repo := &data.EmployeeRepository{C: col}

	repo.AddSalaryComponentByOrgID(sal)
	j, err := json.Marshal(SalaryComponentResource{Data: *sal})
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

//AddSalaryEmployeesByOrg for /orgs/employees/salary/employees/add api
func AddSalaryEmployees(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes SalaryEmployeeResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid employee data",
			500,
		)
		return
	}
	sal := &dataRes.Data
	sal.OrgID = orgid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salaryemployees")
	repo := &data.EmployeeRepository{C: col}

	repo.AddSalaryEmployeesByOrgID(sal)
	j, err := json.Marshal(SalaryEmployeeResource{Data: *sal})
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

//AddSalaryStructureByOrg for /orgs/employees/salary/structures/add api
func AddSalaryStructure(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	accountid = bson.ObjectIdHex(vars["accountid"])
	salempid = bson.ObjectIdHex(vars["salempid"])
	salcompid = bson.ObjectIdHex(vars["salcompid"])
	var dataRes SalaryStructureResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid employee data",
			500,
		)
		return
	}
	sal := &dataRes.Data
	sal.OrgID = orgid
	sal.GeneratedBy = empid
	sal.SalaryEmployeeID = salempid
	sal.SalaryComponentID = salcompid
	sal.PayableAccount = accountid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarystructures")
	repo := &data.EmployeeRepository{C: col}

	repo.AddSalaryStructureByMgrID(sal)
	j, err := json.Marshal(SalaryStructureResource{Data: *sal})
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

//AddActivityTypeByOrg for /orgs/employees/activity/types/add api
func AddActivityType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes ActivityTypeResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid employee data",
			500,
		)
		return
	}
	actType := &dataRes.Data
	actType.OrgID = orgid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("activitytypes")
	repo := &data.EmployeeRepository{C: col}

	repo.AddActivityTypeByOrgID(actType)
	j, err := json.Marshal(ActivityTypeResource{Data: *actType})
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

//AddWorkingHourByOrg for /orgs/employees/working/hours/add api
func AddWorkingHour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes WorkingHourResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid employee data",
			500,
		)
		return
	}
	workHr := &dataRes.Data
	workHr.OrgID = orgid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("workinghours")
	repo := &data.EmployeeRepository{C: col}

	repo.AddWorkingHourByOrgID(workHr)
	j, err := json.Marshal(WorkingHourResource{Data: *workHr})
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

//AddOperationByOrg for /orgs/employees/operations/add api
func AddOperation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes OperationResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid operation data",
			500,
		)
		return
	}
	op := &dataRes.Data
	op.OrgID = orgid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("operations")
	repo := &data.EmployeeRepository{C: col}

	repo.AddOperationByOrgID(sal)
	j, err := json.Marshal(OperationResource{Data: *op})
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

//AddWorkstationByOrg for /orgs/employees/workstations/add api
func AddWorkstation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	opid = bson.ObjectIdHex(vars["operationid"])
	workhrid = bson.ObjectIdHex(vars["workhourid"])
	var dataRes WorkstationResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid workstation data",
			500,
		)
		return
	}
	workSt := &dataRes.Data
	workSt.OrgID = orgid
	workSt.OperationID = opid
	workSt.WorkingHourID = workhrid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("workstations")
	repo := &data.EmployeeRepository{C: col}

	repo.AddWorkstationByOrgID(workSt)
	j, err := json.Marshal(WorkstationResource{Data: *workSt})
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

//AddTimesheetByEmp for /orgs/employees/timesheets/add api
func AddTimesheet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	projectid = bson.ObjectIdHex(vars["projectid"])
	typeid = bson.ObjectIdHex(vars["activitytypeid"])
	stationid = bson.ObjectIdHex(vars["workstationid"])
	var dataRes TimesheetResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid timesheet data",
			500,
		)
		return
	}
	timeSh := &dataRes.Data
	timeSh.OrgID = orgid
	timeSh.ForEmployee = empid
	timeSh.ApprovedBy = mgrId
	timeSh.ActivityTypeID = typeid
	timeSh.WorkStationID = stationid
	timeSh.ProjectID = projectid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("timesheets")
	repo := &data.EmployeeRepository{C: col}

	repo.AddTimesheetByEmpID(timeSh)
	j, err := json.Marshal(TimesheetResource{Data: *timeSh})
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

//AddSalarySlipByEmp for /orgs/employees/salary/slips/add api
func AddSalarySlip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	deptid = bson.ObjectIdHex(vars["deptid"])
	branchid = bson.ObjectIdHex(vars["branchid"])
	timesheetid = bson.ObjectIdHex(vars["timesheetid"])
	var dataRes SalarySlipResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid salary slip data",
			500,
		)
		return
	}
	sal := &dataRes.Data
	sal.OrgID = orgid
	sal.ForEmployee = empid
	sal.ApprovedBy = mgrId
	sal.TimesheetID = timesheetid
	sal.DepartmentID = deptid
	sal.BranchID = branchid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salaryslips")
	repo := &data.EmployeeRepository{C: col}

	repo.AddSalarySlipByEmpID(sal)
	j, err := json.Marshal(SalarySlipResource{Data: *sal})
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

//AddActivityCostByOrg for /orgs/employees/activity/costs/add api
func AddActivityCost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	typeid = bson.ObjectIdHex(vars["activitytypeid"])
	var dataRes ActivityCostResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid activity cost data",
			500,
		)
		return
	}
	actCost := &dataRes.Data
	actCost.OrgID = orgid
	actCost.EmployeeID = empid
	actCost.ActivityType = typeid

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("activitycosts")
	repo := &data.EmployeeRepository{C: col}

	repo.AddActivityCostByOrgID(actCost)
	j, err := json.Marshal(ActivityCostResource{Data: *actCost})
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

//GetSalaryComponentByOrg for /orgs/employees/salary/components/org/{orgId} api
func GetSalaryComponentsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarycomponents")
	repo := &data.EmployeeRepository{C: col}
	salComps, err := repo.GetSalaryComponentsByOrgID(orgid)
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

	j, err := json.Marshal(SalaryComponentsResource{Data: salComps})
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

//GetSalaryComponent for /orgs/employees/salary/components/{id} api
func GetSalaryComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("salarycomponents")
	repo := &data.EmployeeRepository{C: col}
	salComp, err := repo.GetSalaryComponentByID(id)
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
	j, err := json.Marshal(salComp)
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

//EditSalaryComponentByAdmin for /orgs/employees/salary/components/edit/{id} api
func EditSalaryComponentByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource SalaryComponentResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid component data",
			500,
		)
		return
	}
	salComp := &dataResource.Data
	salComp.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarycomponents")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditSalaryComponentByID(leaveApp); err != nil {
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

//Delete Salary Component - TBA

//GetSalaryEmployeesByOrg for /orgs/employees/salary/employees/org/{orgId} api
func GetSalaryEmployeesByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salaryemployees")
	repo := &data.EmployeeRepository{C: col}
	salEmps, err := repo.GetSalaryEmployeesByOrgID(orgid)
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

	j, err := json.Marshal(SalaryEmployeesResource{Data: salEmps})
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

//GetSalaryEmployee for /orgs/employees/salary/employees/{id} api
func GetSalaryEmnployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("salaryemployees")
	repo := &data.EmployeeRepository{C: col}
	salEmp, err := repo.GetSalaryEmployeeByID(id)
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
	j, err := json.Marshal(salEmp)
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

//EditSalaryEmployeeByAdmin for /orgs/employees/salary/employees/edit/{id} api
func EditSalaryEmployeeByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource SalaryComponeEmployeeResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid component data",
			500,
		)
		return
	}
	salEmp := &dataResource.Data
	salEmp.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salaryemployees")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditSalaryEmployeeByID(salEmp); err != nil {
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

//Delete Salary Employee - TBA

//GetSalaryStructuresByOrg for /orgs/employees/salary/structures/org/{orgId} api
func GetSalaryStructuresByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarystructures")
	repo := &data.EmployeeRepository{C: col}
	sals, err := repo.GetSalaryStructuresByOrgID(orgid)
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

	j, err := json.Marshal(SalaryStructuresResource{Data: sals})
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

//GetSalaryStructuresByAcct for /orgs/employees/salary/structures/account/{acctId} api
func GetSalaryStructuresByAcct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	acctid = bson.ObjectIdHex(vars["acctid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarystructures")
	repo := &data.EmployeeRepository{C: col}
	sals, err := repo.GetSalaryStructuresByAcctID(acctid)
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

	j, err := json.Marshal(SalaryStructuresResource{Data: sals})
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

//GetSalaryStructuresByEmp for /orgs/employees/salary/structures/employees/{empId} api
func GetSalaryStructuresByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarystructures")
	repo := &data.EmployeeRepository{C: col}
	sals, err := repo.GetSalaryStructuresByEmpID(empid)
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

	j, err := json.Marshal(SalaryStructuresResource{Data: sals})
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

//GetSalaryStructure for /orgs/employees/salary/structures/{id} api
func GetSalaryStructure(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("salarystructures")
	repo := &data.EmployeeRepository{C: col}
	sal, err := repo.GetSalaryStructureByID(id)
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
	j, err := json.Marshal(sal)
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

//EditSalaryStructureByAdmin for /orgs/employees/salary/structures/edit/{id} api
func EditSalaryStructureByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource SalaryStructureResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid structure data",
			500,
		)
		return
	}
	sal := &dataResource.Data
	sal.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salarystructures")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditSalaryStructureByID(sal); err != nil {
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

//Delete Salary Structure - TBA

//GetActivityTypesByOrg for /orgs/employees/activity/types/org/{orgId} api
func GetActivitytTypesByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("activitytypes")
	repo := &data.EmployeeRepository{C: col}
	actTypes, err := repo.GetActivityTypesByOrgID(orgid)
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

	j, err := json.Marshal(ActivityTypsResource{Data: actTypes})
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

//GetActivityType for /orgs/employees/activity/types/{id} api
func GetSalaryActivityType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("activitytypes")
	repo := &data.EmployeeRepository{C: col}
	actType, err := repo.GetActivityTypeByID(id)
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
	j, err := json.Marshal(actType)
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

//EditActivityTypeByAdmin for /orgs/employees/activity/types/edit/{id} api
func EditActivityTypeByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource ActivityTypeResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid type data",
			500,
		)
		return
	}
	actType := &dataResource.Data
	actType.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("activitytypes")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditActivityTypeByID(sal); err != nil {
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

//Delete Activity Type - TBA

//GetOperationsByOrg for /orgs/employees/operations/org/{orgId} api
func GetOperationsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("operations")
	repo := &data.EmployeeRepository{C: col}
	ops, err := repo.GetOperationsByOrgID(orgid)
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

	j, err := json.Marshal(OperationsResource{Data: ops})
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

//GetOperation for /orgs/employees/operations/{id} api
func GetOperation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("operations")
	repo := &data.EmployeeRepository{C: col}
	op, err := repo.GetOperationByID(id)
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
	j, err := json.Marshal(op)
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

//EditOperationByAdmin for /orgs/employees/operations/edit/{id} api
func EditOperationByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource OperationResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid op data",
			500,
		)
		return
	}
	op := &dataResource.Data
	op.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("operations")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditOperationByID(sal); err != nil {
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

//Delete Operation - TBA

//GetWorkingHoursByOrg for /orgs/employees/working/hours/org/{orgId} api
func GetWorkingHoursByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("workinghours")
	repo := &data.EmployeeRepository{C: col}
	works, err := repo.GetWorkingHoursByOrgID(orgid)
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

	j, err := json.Marshal(WorkingHoursResource{Data: works})
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

//GetWorkingHour for /orgs/employees/working/hours/{id} api
func GetWorkingHour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("workinghours")
	repo := &data.EmployeeRepository{C: col}
	work, err := repo.GetWorkingHourByID(id)
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
	j, err := json.Marshal(work)
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

//EditWorkingHourByAdmin for /orgs/employees/working/hours/edit/{id} api
func EditWorkingHourByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource WorkingHourResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid hour data",
			500,
		)
		return
	}
	work := &dataResource.Data
	work.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("workinghours")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditWorkingHourByID(sal); err != nil {
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

//Delete Working Hour - TBA

//GetWorkstationsByOrg for /orgs/employees/workstations/org/{orgId} api
func GetWorkstationsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("workstations")
	repo := &data.EmployeeRepository{C: col}
	works, err := repo.GetWorkstationsByOrgID(orgid)
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

	j, err := json.Marshal(WorkstationsResource{Data: works})
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

//GetWorkstation for /orgs/employees/workstations/{id} api
func GetWorkingHour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("workstations")
	repo := &data.EmployeeRepository{C: col}
	work, err := repo.GetWorkstationByID(id)
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
	j, err := json.Marshal(work)
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

//EditWorkstationByAdmin for /orgs/employees/workstations/edit/{id} api
func EditWorkstationByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource WorkingHourResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid workstation data",
			500,
		)
		return
	}
	work := &dataResource.Data
	work.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("workstations")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditWorkstationByID(sal); err != nil {
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

//Delete Workstation - TBA

//GetTimesheetsByOrg for /orgs/employees/timesheets/org/{orgId} api
func GetTimesheetsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("timesheets")
	repo := &data.EmployeeRepository{C: col}
	times, err := repo.GetTimesheetsByOrgID(orgid)
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

	j, err := json.Marshal(TimesheetsResource{Data: times})
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

//GetTimesheetsByMgr for /orgs/employees/timesheets/employees/{mgrId} api
func GetTimesheetsByMgr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("timesheets")
	repo := &data.EmployeeRepository{C: col}
	times, err := repo.GetTimesheetsByMgrID(mgrid)
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

	j, err := json.Marshal(TimesheetsResource{Data: times})
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

//GetTimesheetsByEmp for /orgs/employees/timesheets/employees/{empId} api
func GetTimesheetsByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("timesheets")
	repo := &data.EmployeeRepository{C: col}
	times, err := repo.GetTimesheetsByEmpID(empid)
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

	j, err := json.Marshal(TimesheetsResource{Data: times})
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

//GetTimesheet for /orgs/employees/timesheets/{id} api
func GetTimesheet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("timesheets")
	repo := &data.EmployeeRepository{C: col}
	timeSh, err := repo.GetTimesheetByID(id)
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
	j, err := json.Marshal(timeSh)
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

//EditTimesheetByAdmin for /orgs/employees/timesheets/edit/{id} api
func EditTimesheetByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource TimesheetResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid time data",
			500,
		)
		return
	}
	timeSh := &dataResource.Data
	timeSh.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("timesheets")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditWorkingHourByID(sal); err != nil {
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

//Delete Timesheet - TBA

//GetSalarySlipsByOrg for /orgs/employees/salary/slips/org/{orgId} api
func GetSalarySlipsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salaryslips")
	repo := &data.EmployeeRepository{C: col}
	sals, err := repo.GetSalarySlipsByOrgID(orgid)
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

	j, err := json.Marshal(SalarySlipsResource{Data: sals})
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

//GetSalarySlipsByEmp for /orgs/employees/salary/slips/employee/{empId} api
func GetSalarySlipsByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salaryslips")
	repo := &data.EmployeeRepository{C: col}
	sals, err := repo.GetSalarySlipsByEmpID(empid)
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

	j, err := json.Marshal(SalarySlipsResource{Data: sals})
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

//GetSalarySlipsByMgr for /orgs/employees/salary/slips/employees/{mgrId} api
func GetSalarySlipsByMgr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mgrid = bson.ObjectIdHex(vars["mgrid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salaryslips")
	repo := &data.EmployeeRepository{C: col}
	sals, err := repo.GetSalarySlipsByMgrID(mgrid)
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

	j, err := json.Marshal(SalarySlipsResource{Data: sals})
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

//GetSalarySlip for /orgs/employees/salary/slips/{id} api
func GetSalarySlip(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("salaryslips")
	repo := &data.EmployeeRepository{C: col}
	sal, err := repo.GetSalarySlipByID(id)
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
	j, err := json.Marshal(sal)
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

//EditSalarySlipByAdmin for /orgs/employees/salary/slips/edit/{id} api
func EditSalarySlipByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource SalarySlipResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid slip data",
			500,
		)
		return
	}
	sal := &dataResource.Data
	sal.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("salaryslips")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditSalarySlipByID(sal); err != nil {
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

//Delete Salary Slip - TBA

//GetActivtyCostsByOrg for /orgs/employees/activity/costs/org/{orgId} api
func GetActivityCostsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("activitycosts")
	repo := &data.EmployeeRepository{C: col}
	acts, err := repo.GetActivtyCostsByOrgID(orgid)
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

	j, err := json.Marshal(ActivityCostsResource{Data: acts})
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

//GetActivtyCostsByEmp for /orgs/employees/activity/costs/employees/{empId} api
func GetActivityCostsByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("activitycosts")
	repo := &data.EmployeeRepository{C: col}
	acts, err := repo.GetActivtyCostsByEmpID(empid)
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

	j, err := json.Marshal(ActivityCostsResource{Data: acts})
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

//GetActivtyCostsByType for /orgs/employees/activity/costs/types/{typeId} api
func GetActivityCostsByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeid = bson.ObjectIdHex(vars["activitytypeid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("activitycosts")
	repo := &data.EmployeeRepository{C: col}
	acts, err := repo.GetActivtyCostsByTypeID(typeid)
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

	j, err := json.Marshal(ActivityCostsResource{Data: acts})
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

//GetActivityCost for /orgs/employees/activity/costs/{id} api
func GetActivityCost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("activitycosts")
	repo := &data.EmployeeRepository{C: col}
	act, err := repo.GetActivityCostByID(id)
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
	j, err := json.Marshal(act)
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

//EditActivityCostByAdmin for /orgs/employees/activity/costs/edit/{id} api
func EditActivityCostByAdmin(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource ActivityCostResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid cost data",
			500,
		)
		return
	}
	act := &dataResource.Data
	act.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("activitycosts")
	repo := &data.EmployeeRepository{C: col}
	
	err := repo.EditActivityCostByID(sal); err != nil {
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

//Delete Activity Cost - TBA

