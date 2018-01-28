package data

/**
 *
 * @author Sika Kay
 * @date 18/01/18
 *
 */
import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
)

//SettingRepository Struct for Mongo Persistence
type SettingRepository struct {
	C *mgo.Collection
}

//AddDept persists Dept associated with OrgID
func (r *SettingRepository) AddDeptByOrgID(orgID string) (dept models.Department, err error) {
	objID := bson.NewObjectId()
	dept.ID = objID
	dept.OrgID = bson.ObjectIdHex(orgID)
	dept.Status = "Active"
	dept.CreatedAt = time.Now()
	dept.UpdatedAt = time.Now()

	err = r.C.Insert(&dept)
	return
}

//AddDesignation persists Designation associated with OrgID
func (r *SettingRepository) AddDesgnByOrgID(orgID string) (des models.Designation, err error) {
	objID := bson.NewObjectId()
	des.ID = objID
	des.OrgID = bson.ObjectIdHex(orgID)
	des.Status = "Active"
	des.CreatedAt = time.Now()
	des.UpdatedAt = time.Now()

	err = r.C.Insert(&des)
	return
}

//AddSalaryMode persists salary mode associated with OrgID
func (r *SettingRepository) AddSalaryModeByOrgID(orgID string) (sal models.SalaryMode, err error) {
	objID := bson.NewObjectId()
	sal.ID = objID
	sal.OrgID = bson.ObjectIdHex(orgID)
	sal.Status = "Active"
	sal.CreatedAt = time.Now()
	sal.UpdatedAt = time.Now()

	err = r.C.Insert(&sal)
	return
}

//AddBranch persists branch associated with OrgID
func (r *SettingRepository) AddBranchByOrgID(orgID string) (branch models.Branch, err error) {
	objID := bson.NewObjectId()
	branch.ID = objID
	branch.OrgID = bson.ObjectIdHex(orgID)
	branch.Status = "Active"
	branch.CreatedAt = time.Now()
	branch.UpdatedAt = time.Now()

	err = r.C.Insert(&branch)
	return
}

//AddLeaveType persists leave type associated with OrgID
func (r *SettingRepository) AddLeaveTypeByOrgID(orgID string) (leaveType models.LeaveType, err error) {
	objID := bson.NewObjectId()
	leaveType.ID = objID
	leaveType.OrgID = bson.ObjectIdHex(orgID)
	leaveType.Status = "Active"
	leaveType.CreatedAt = time.Now()
	leaveType.UpdatedAt = time.Now()

	err = r.C.Insert(&leaveType)
	return
}

//AddExpenseClaimType persists claim type associated with OrgID
func (r *SettingRepository) AddClaimTypeByOrgID(orgID string) (claimType models.ExpenseClaimType, err error) {
	objID := bson.NewObjectId()
	claimType.ID = objID
	claimType.OrgID = bson.ObjectIdHex(orgID)
	claimType.Status = "Active"
	claimType.CreatedAt = time.Now()
	claimType.UpdatedAt = time.Now()

	err = r.C.Insert(&claimType)
	return
}

//AddProjectType persists project type associated with OrgID
func (r *SettingRepository) AddProjectTypeByOrgID(orgID string) (projectType models.ProjectType, err error) {
	objID := bson.NewObjectId()
	projectType.ID = objID
	projectType.OrgID = bson.ObjectIdHex(orgID)
	projectType.Status = "Active"
	projectType.CreatedAt = time.Now()
	projectType.UpdatedAt = time.Now()

	err = r.C.Insert(&projectType)
	return
}

//GetDeptsByOrgID gets depts associated with an OrgID
func (r *SettingRepository) GetDeptsByOrgID(orgID string) []models.Department {
	var depts []models.Department
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Deparmtent{}
	for iter.Next(&result) {
		depts = append(depts, result)
	}
	return depts
}

//EditDeptByID edits dept associated with an ID
func (r *SettingRepository) EditDepartmentByID(dept *models.Deparmtent) error {
	err := r.C.Update(bson.M{"_id": dept.ID},
		bson.M{"$set": bson.M{
			"name":        dept.Name,
			"description": dept.Description,
			"updatedat":   time.Now(),
			"status":      dept.Status,
		}})
	return err
}

//DeleteDeptById deletes dept out of the system by Id
func (r *SettingRepository) DeleteDeptById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetBranchesByOrgID gets branches associated with a OrgID
func (r *SettingRepository) GetBranchesByOrgID(orgID string) []models.Branch {
	var branches []models.Branch
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Branch{}
	for iter.Next(&result) {
		branches = append(branches, result)
	}
	return branches
}

//EditBranchByID edits  branch associated with an ID
func (r *SettingRepository) EditBranchByID(branch *models.Branch) error {
	err := r.C.Update(bson.M{"_id": branch.ID},
		bson.M{"$set": bson.M{
			"name":      branch.Name,
			"updatedat": time.Now(),
			"status":    dept.Status,
		}})
	return err
}

//DeleteBranchById deletes branch out of the system by Id
func (r *SettingRepository) DeleteBrancById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetDesignationsByOrgID gets designations associated with a OrgID
func (r *SettingRepository) GetDesignationsByOrgID(orgID string) []models.Designation {
	var des []models.Designation
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Designation{}
	for iter.Next(&result) {
		des = append(des, result)
	}
	return des
}

//EditDesignationByID edits designation associated with an ID
func (r *SettingRepository) EditDesignationByID(des *models.Designation) error {
	err := r.C.Update(bson.M{"_id": des.ID},
		bson.M{"$set": bson.M{
			"name":      des.Name,
			"updatedat": time.Now(),
			"status":    des.Status,
		}})
	return err
}

//DeleteDesignationById deletes designation out of the system by Id
func (r *SettingRepository) DeleteDesById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetSalaryModesByOrgID gets salary modes associated with a OrgID
func (r *SettingRepository) GetModesByOrgID(orgID string) []models.SalaryMode {
	var salModes []models.SalaryMode
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.SalaryMode{}
	for iter.Next(&result) {
		salModes = append(salModes, result)
	}
	return salModes
}

//EditSalaryModeByID edits salary mode associated with an ID
func (r *SettingRepository) EditSalaryModeByID(sal *models.SalaryMode) error {
	err := r.C.Update(bson.M{"_id": sal.ID},
		bson.M{"$set": bson.M{
			"name":      sal.Name,
			"updatedat": time.Now(),
			"status":    sal.Status,
		}})
	return err
}

//DeleteSalaryModeById deletes salary mode out of the system by Id
func (r *SettingRepository) DeleteSalaryModeById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetLeaveTypesByOrgID gets leave types associated with a OrgID
func (r *SettingRepository) GetLeaveTypesByOrgID(orgID string) []models.LeaveType {
	var leaveTypes []models.LeaveType
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.LeaveType{}
	for iter.Next(&result) {
		leaveTypes = append(leaveTypes, result)
	}
	return leaveTypes
}

//EditLeaveTypeByID edits leave type associated with an ID
func (r *SettingRepository) EditLeaveTypeByID(leave *models.LeaveType) error {
	err := r.C.Update(bson.M{"_id": leave.ID},
		bson.M{"$set": bson.M{
			"name":      leave.Name,
			"updatedat": time.Now(),
			"status":    leave.Status,
		}})
	return err
}

//DeleteLeaveTypeById deletes leave type out of the system by Id
func (r *SettingRepository) DeleteLeaveTypeById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetClaimTypesByOrgID gets leave types associated with a OrgID
func (r *SettingRepository) GetClaimTypesByOrgID(orgID string) []models.ExpenseClaimType {
	var claimTypes []models.ExpenseClaimType
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.ExpenseClaimType{}
	for iter.Next(&result) {
		claimTypes = append(claimTypes, result)
	}
	return claimTypes
}

//EditClaimTypeByID edits claim type associated with an ID
func (r *SettingRepository) EditClaimTypeByID(claimType *models.ExpenseClaimType) error {
	err := r.C.Update(bson.M{"_id": claimType.ID},
		bson.M{"$set": bson.M{
			"name":      claimType.Name,
			"updatedat": time.Now(),
			"status":    claimType.Status,
		}})
	return err
}

//DeleteClaimTypeById deletes claim type out of the system by Id
func (r *SettingRepository) DeleteClaimTypeById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetProjectTypesByOrgID gets project types associated with a OrgID
func (r *SettingRepository) GetProjectTypesByOrgID(orgID string) []models.ProjectType {
	var projectTypes []models.ProjectType
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.ProjectType{}
	for iter.Next(&result) {
		projectTypes = append(projectTypes, result)
	}
	return projectTypes
}

//EditProjectTypeByID edits project type associated with an ID
func (r *SettingRepository) EditProjectTypeByID(projectType *models.ProjectType) error {
	err := r.C.Update(bson.M{"_id": projectType.ID},
		bson.M{"$set": bson.M{
			"name":      projectType.Name,
			"updatedat": time.Now(),
			"status":    projectType.Status,
		}})
	return err
}

//DeleteProjectTypeById deletes project type out of the system by Id
func (r *SettingRepository) DeleteProjectTypeById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetDeptByID gets dept associated with an ID
func (r *SettingRepository) GetDeptByID(id string) (dept models.Department, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&dept)
	return
}

//GetBranchByID gets branch associated with an ID
func (r *SettingRepository) GetBranchByID(id string) (branch models.Branch, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&branch)
	return
}

//GetSalaryModeByID gets salary mode associated with an ID
func (r *SettingRepository) GetSalaryModeByID(id string) (sal models.SalaryMode, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&sal)
	return
}

//GetLeaveTypeByID gets leave type associated with an ID
func (r *SettingRepository) GetLeaveTypeByID(id string) (leaveType models.LeaveType, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&leaveType)
	return
}

//GetClaimTypeByID gets claim type associated with an ID
func (r *SettingRepository) GetClaimTypeByID(id string) (claimType models.ExpenseClaimType, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&claimType)
	return
}

//GetDesignationByID gets designation associated with an ID
func (r *SettingRepository) GetDesgnByID(id string) (des models.Designation, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&des)
	return
}
