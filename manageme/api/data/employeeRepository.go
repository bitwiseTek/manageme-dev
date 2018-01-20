package data

/**
 *
 * @author Sika Kay
 * @date 17/01/18
 *
 */
import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
)

//EmpRepository Struct for Mongo Persistence
type EmployeeRepository struct {
	C *mgo.Collection
}

//AddEmp persists Employee associated with BranchID, DepartmentID, OrgUserID
func (r *EmployeeRepository) AddEmpByOrgUserID(orgUserID string, branchID string, deptID string) (emp models.Employee, err error) {
	objID := bson.NewObjectId()
	emp.ID = objID
	emp.OrgUserID = bson.ObjectIdHex(orgUserID)
	emp.BranchID = bson.ObjectIdHex(branchID)
	emp.DepartmentID = bson.ObjectIdHex(deptID)
	emp.Status = "Active"

	err = r.C.Insert(&emp)
	return
}

//AddEmp persists Employee associated with BranchID, DepartmentID, OrgUserID, EmpID
func (r *EmployeeRepository) AddChildEmpByOrgUserID(orgUserID string, branchID string, deptID string, empID string) (emp models.Employee, err error) {
	objID := bson.NewObjectId()
	emp.ID = objID
	emp.OrgUserID = bson.ObjectIdHex(orgUserID)
	emp.BranchID = bson.ObjectIdHex(branchID)
	emp.DepartmentID = bson.ObjectIdHex(deptID)
	emp.ReportsTo = bson.ObjectIdHex(empID)
	emp.Status = "Active"

	err = r.C.Insert(&emp)
	return
}

//GetEmps gets all employees
func (r *EmployeeRepository) GetEmployees() []models.Employee {
	var emps []models.Employee
	iter := r.C.Find(nil).Iter()
	result := models.Employee{}
	for iter.Next(&result) {
		emps = append(emps, result)
	}
	return emps
}

//GetEmpByID gets employee associated with an ID
func (r *EmployeeRepository) GetEmpByID(id string) (emp models.Employee, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&emp)
	return
}

//GetEmpsByOrgID gets emps associated with an OrgID
func (r *EmployeeRepository) GetEmpsByOrgUserID(orgID string) []models.Employee {
	var emps []models.Employee
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Employee{}
	for iter.Next(&result) {
		emps = append(emps, result)
	}
	return emps
}

//GetEmpsByDeptID gets emps associated with a DeptID
func (r *EmployeeRepository) GetEmpsByDeptID(deptID string) []models.Employee {
	var emps []models.Employee
	deptid := bson.ObjectIdHex(deptID)
	iter := r.C.Find(bson.M{"deptid": deptid}).Iter()
	result := models.Employee{}
	for iter.Next(&result) {
		emps = append(emps, result)
	}
	return emps
}

//GetEmpsByBranchID gets emps associated with a BranchID
func (r *EmployeeRepository) GetEmpsByBranchID(branchID string) []models.Employee {
	var emps []models.Employee
	branchid := bson.ObjectIdHex(branchID)
	iter := r.C.Find(bson.M{"branchid": branchid}).Iter()
	result := models.Employee{}
	for iter.Next(&result) {
		emps = append(emps, result)
	}
	return emps
}

//GetEmpsByMgrID gets emps associated with a MgrID
func (r *EmployeeRepository) GetEmpsByMgrID(mgrID string) []models.Employee {
	var emps []models.Employee
	mgrid := bson.ObjectIdHex(mgrID)
	iter := r.C.Find(bson.M{"mgrid": mgrid}).Iter()
	result := models.Employee{}
	for iter.Next(&result) {
		emps = append(emps, result)
	}
	return emps
}

//AddBiodata persists Biodata associated with EmpID
func (r *EmployeeRepository) AddBiodataByEmpID(empID string) (bio models.Biodata, err error) {
	objID := bson.NewObjectId()
	bio.ID = objID
	bio.EmployeeID = bson.ObjectIdHex(empID)
	bio.DisabilityType = "None"

	err = r.C.Insert(&bio)
	return
}

//AddPersonalID persists ID associated with BioID
func (r *EmployeeRepository) AddPIDByBioID(bioID string) (pid models.PersonalIdentification, err error) {
	objID := bson.NewObjectId()
	pid.ID = objID
	pid.BioID = bson.ObjectIdHex(bioID)

	err = r.C.Insert(&pid)
	return
}

//AddHealthDet persists HealthDets associated with BIoID
func (r *EmployeeRepository) AddHealthDetByBioID(bioID string) (health models.HealthDetail, err error) {
	objID := bson.NewObjectId()
	health.ID = objID
	health.BioID = bson.ObjectIdHex(bioID)

	err = r.C.Insert(&health)
	return
}

//AddEdu persists edu associated with BioID
func (r *EmployeeRepository) AddEduByBioID(bioID string) (edu models.Education, err error) {
	objID := bson.NewObjectId()
	edu.ID = objID
	edu.BioID = bson.ObjectIdHex(bioID)

	err = r.C.Insert(&edu)
	return
}

//AddWork persists work associated with BioID
func (r *EmployeeRepository) AddWorkByBioID(bioID string) (work models.WorkExperience, err error) {
	objID := bson.NewObjectId()
	work.ID = objID
	work.BioID = bson.ObjectIdHex(bioID)

	err = r.C.Insert(&work)
	return
}

//GetBioByID gets biodata associated with an ID
func (r *EmployeeRepository) GetBioByID(id string) (bio models.Biodata, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&bio)
	return
}

//GetBiosByEmpID gets bios associated with an EmpID
func (r *EmployeeRepository) GetBiosByEmpID(empID string) []models.Biodata {
	var bios []models.Biodata
	empid := bson.ObjectIdHex(empID)
	iter := r.C.Find(bson.M{"empid": empid}).Iter()
	result := models.Biodata{}
	for iter.Next(&result) {
		bios = append(bios, result)
	}
	return bios
}

//GetPIDByID gets personal ID associated with an ID
func (r *EmployeeRepository) GetPIDByID(id string) (pid models.PersonalIdentification, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&pid)
	return
}

//GetPIDByBioID gets pids associated with a BioID
func (r *EmployeeRepository) GetPIDssByBioID(bioID string) []models.PersonalIdentification {
	var pids []models.PersonalIdentification
	bioid := bson.ObjectIdHex(bioID)
	iter := r.C.Find(bson.M{"bioid": bioid}).Iter()
	result := models.PersonalIdentification{}
	for iter.Next(&result) {
		pids = append(pids, result)
	}
	return pids
}

//GetEduByID gets edu associated with an ID
func (r *EmployeeRepository) GetEduByID(id string) (edu models.Education, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&edu)
	return
}

//GetEdusByBioID gets edus associated with a BioID
func (r *EmployeeRepository) GetEdusByBioID(bioID string) []models.Education {
	var edus []models.Education
	bioid := bson.ObjectIdHex(bioID)
	iter := r.C.Find(bson.M{"bioid": bioid}).Iter()
	result := models.Education{}
	for iter.Next(&result) {
		edus = append(edus, result)
	}
	return edus
}

//GetWorkByID gets work exoperience associated with an ID
func (r *EmployeeRepository) GetWorkByID(id string) (work models.WorkExperience, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&work)
	return
}

//GetWorksByMgrID gets works associated with a BioID
func (r *EmployeeRepository) GetWorksByBioID(bioID string) []models.WorkExperience {
	var works []models.WorkExperience
	bioid := bson.ObjectIdHex(bioID)
	iter := r.C.Find(bson.M{"bioid": bioid}).Iter()
	result := models.WorkExperience{}
	for iter.Next(&result) {
		works = append(works, result)
	}
	return works
}

//AddExpenseClaimByEmpID persists expense claim associated with EmpID, MgrID
func (r *EmployeeRepository) AddExpenseClaimByEmpID(orgID string, empID string, mgrID string, typedID string, projectID string, taskID string, accountID string) (claim models.ExpenseClaim, err error) {
	objID := bson.NewObjectId()
	claim.ID = objID
	claim.OrgID = bson.ObjectIdHex(orgID)
	claim.ExpApplier = bson.ObjectIdHex(empID)
	claim.ExpApprover = bson.ObjectIdHex(mgrID)
	claim.ExpenseType = bson.ObjectIdHex(typeID)
	claim.ProjectID = bson.ObjectIdHex(projectID)
	claim.TaskID = bson.ObjectIdHex(taskID)
	claim.PayableAccount = bson.ObjectIdHex(accountID)
	claim.PostingDate = time.Now()
	claim.Status = "Applied"

	err = r.C.Insert(&claim)
	return
}

//AddLeaveAllocationByEmpID persists leave allocation associated with EmpID, MgrID
func (r *EmployeeRepository) AddLeaveAllByEmpID(orgID string, empID string, mgrID string, typedID string) (leaveAll models.LeaveAllocation, err error) {
	objID := bson.NewObjectId()
	leaveAll.ID = objID
	leaveAll.OrgID = bson.ObjectIdHex(orgID)
	leaveAll.LeaveReceiver = bson.ObjectIdHex(empID)
	leaveAll.LeaveAllocator = bson.ObjectIdHex(mgrID)
	leaveAll.LeaveType = bson.ObjectIdHex(typeID)
	leaveAll.AllocatedOn = time.Now()

	err = r.C.Insert(&leaveAll)
	return
}

//AddLeaveApplicationByEmpID persists leave application associated with EmpID, MgrID
func (r *EmployeeRepository) AddLeaveAppByEmpID(orgID string, empID string, mgrID string, typedID string) (leaveApp models.LeaveApplication, err error) {
	objID := bson.NewObjectId()
	leaveApp.ID = objID
	leaveApp.OrgID = bson.ObjectIdHex(orgID)
	leaveApp.LeaveApplier = bson.ObjectIdHex(empID)
	leaveApp.LeaveApprover = bson.ObjectIdHex(mgrID)
	leaveApp.LeaveType = bson.ObjectIdHex(typeID)
	leaveApp.AppliedAt = time.Now()
	leaveApp.UpdatedAt = time.Now()
	leaveApp.Status = "Applied"

	err = r.C.Insert(&leaveApp)
	return
}

//AddLeaveBlockListByOrgID persists leave block list associated with OrgID
func (r *EmployeeRepository) AddLeaveBlockListByOrgID(orgID string) (leaveBlk models.LeaveBlockList, err error) {
	objID := bson.NewObjectId()
	leaveBlk.ID = objID
	leaveBlk.OrgID = bson.ObjectIdHex(orgID)

	err = r.C.Insert(&leaveBlk)
	return
}

//AddHolidayListByOrgID persists holiday list associated with OrgID
func (r *EmployeeRepository) AddHolidayListByOrgID(orgID string) (hol models.HolidayList, err error) {
	objID := bson.NewObjectId()
	hol.ID = objID
	hol.OrgID = bson.ObjectIdHex(orgID)

	err = r.C.Insert(&hol)
	return
}

//AddAppraisalByEmpID persists appraisal associated with EmpID, MgrID
func (r *EmployeeRepository) AddAppraisalByEmpID(orgID string, empID string, mgrID string) (appr models.Appraisal, err error) {
	objID := bson.NewObjectId()
	appr.ID = objID
	appr.OrgID = bson.ObjectIdHex(orgID)
	appr.ForEmployee = bson.ObjectIdHex(empID)
	appr.ApGenerator = bson.ObjectIdHex(mgrID)
	appr.AppraisedAt = time.Now()
	appr.Status = "Started"

	err = r.C.Insert(&appr)
	return
}

//AddExitByEmpID persists exit associated with EmpID, MgrID
func (r *EmployeeRepository) AddExitByEmpID(orgID string, empID string, mgrID string) (exit models.Exit, err error) {
	objID := bson.NewObjectId()
	exit.ID = objID
	exit.OrgID = bson.ObjectIdHex(orgID)
	exit.ForEmployee = bson.ObjectIdHex(empID)
	exit.OverseenBy = bson.ObjectIdHex(mgrID)
	exit.ExitedAt = time.Now()

	err = r.C.Insert(&exit)
	return
}

//GetClaimByID gets expense claim associated with an ID
func (r *EmployeeRepository) GetClaimByID(id string) (claim models.ExpenseClaim, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&claim)
	return
}

//GetExpenseClaimsByEmpID gets expense claims associated with an EmpID
func (r *EmployeeRepository) GetClaimsByEmpID(empID string) []models.ExpenseClaim {
	var claims []models.ExpenseClaim
	empid := bson.ObjectIdHex(empID)
	iter := r.C.Find(bson.M{"empid": empid}).Iter()
	result := models.ExpenseClaim{}
	for iter.Next(&result) {
		claims = append(claims, result)
	}
	return claims
}

//GetExpenseClaimsByMgrID gets expense claims associated with a MgrID
func (r *EmployeeRepository) GetClaimsByMgrID(mgrID string) []models.ExpenseClaim {
	var claims []models.ExpenseClaim
	mgrid := bson.ObjectIdHex(mgrID)
	iter := r.C.Find(bson.M{"mgrid": mgrid}).Iter()
	result := models.ExpenseClaim{}
	for iter.Next(&result) {
		claims = append(claims, result)
	}
	return claims
}

//GetExpenseClaimsByOrgID gets expense claims associated with a OrgID
func (r *EmployeeRepository) GetClaimsByOrgID(orgID string) []models.ExpenseClaim {
	var claims []models.ExpenseClaim
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.ExpenseClaim{}
	for iter.Next(&result) {
		claims = append(claims, result)
	}
	return claims
}

//GetAllcoationByID gets leave allocation associated with an ID
func (r *EmployeeRepository) GetLeaveAllocationByID(id string) (leaveAll models.LeaveAllocation, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&leaveAll)
	return
}

//GetLeaveAllocationsByMgrID gets leave allocations associated with a MgrID
func (r *EmployeeRepository) GetAllocationsByEmpID(mgrID string) []models.LeaveAllocation {
	var allocations []models.LeaveAllocation
	mgrid := bson.ObjectIdHex(mgrID)
	iter := r.C.Find(bson.M{"mgrid": mgrid}).Iter()
	result := models.LeaveAllocation{}
	for iter.Next(&result) {
		allocations = append(allocations, result)
	}
	return allocations
}

//GetLeaveAllocationsByOrgID gets leave allocations associated with a OrgID
func (r *EmployeeRepository) GetAllocationsByOrgID(orgID string) []models.LeaveAllocation {
	var allocations []models.LeaveAllocation
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.LeaveAllocation{}
	for iter.Next(&result) {
		allocations = append(allocations, result)
	}
	return allocations
}

//GetLeaveApplicationByID gets leave application associated with an ID
func (r *EmployeeRepository) GetLeaveApplicationByID(id string) (leaveApp models.LeaveApplication, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&leaveApp)
	return
}

//GetLeaveApplicationsByMgrID gets leave allocations associated with a MgrID
func (r *EmployeeRepository) GetApplicationsByMgrID(mgrID string) []models.LeaveApplication {
	var apps []models.LeaveApplication
	mgrid := bson.ObjectIdHex(mgrID)
	iter := r.C.Find(bson.M{"mgrid": mgrid}).Iter()
	result := models.LeaveApplication{}
	for iter.Next(&result) {
		apps = append(apps, result)
	}
	return apps
}

//GetLeaveApplicationsByOrgID gets leave allocations associated with a OrgID
func (r *EmployeeRepository) GetApplicationsByOrgID(orgID string) []models.LeaveApplication {
	var apps []models.LeaveApplication
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.LeaveApplication{}
	for iter.Next(&result) {
		apps = append(apps, result)
	}
	return apps
}

//GetLeaveApplicationsByEmpID gets leave applications associated with a EmpID
func (r *EmployeeRepository) GetApplicationsByEmpID(empID string) []models.LeaveApplication {
	var apps []models.LeaveApplication
	empid := bson.ObjectIdHex(empID)
	iter := r.C.Find(bson.M{"empid": empid}).Iter()
	result := models.LeaveApplication{}
	for iter.Next(&result) {
		apps = append(apps, result)
	}
	return apps
}

//GetBlockListByID gets block list associated with an ID
func (r *EmployeeRepository) GetLeaveBlockListByID(id string) (blk models.LeaveBlockList, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&blk)
	return
}

//GetLeaveBlockListsByOrgID gets leave block lists associated with a OrgID
func (r *EmployeeRepository) GetBlockListsByOrgID(orgID string) []models.LeaveBlockList {
	var blockLists []models.LeaveBlockList
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.LeaveBlockList{}
	for iter.Next(&result) {
		blockLists = append(blockLists, result)
	}
	return blockLists
}

//GetHolidayListByID gets holiday list associated with an ID
func (r *EmployeeRepository) GetHolidayListByID(id string) (hol models.HolidayList, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&hol)
	return
}

//GetHolidayListsByOrgID gets holiday lists associated with a OrgID
func (r *EmployeeRepository) GetHolidayListsByOrgID(orgID string) []models.HolidayList {
	var holLists []models.HolidayList
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.HolidayList{}
	for iter.Next(&result) {
		holLists = append(holLists, result)
	}
	return holLists
}

//GetAppraisalByID gets appraisal associated with an ID
func (r *EmployeeRepository) GetAppraisalByID(id string) (appraisal models.Appraisal, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&appraisal)
	return
}

//GetAppraisalsByEmpID gets appraisals associated with an EmpID
func (r *EmployeeRepository) GetAppraisalsByEmpID(empID string) []models.Appraisal {
	var appraisals []models.Appraisal
	empid := bson.ObjectIdHex(empID)
	iter := r.C.Find(bson.M{"empid": empid}).Iter()
	result := models.Appraisal{}
	for iter.Next(&result) {
		appraisals = append(appraisals, result)
	}
	return appraisals
}

//GetAppraisalsByOrgID gets appraisals associated with an OrgID
func (r *EmployeeRepository) GetAppraisalsByOrgID(orgID string) []models.Appraisal {
	var appraisals []models.Appraisal
	orgid := bson.ObjectIdHex(empID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Appraisal{}
	for iter.Next(&result) {
		appraisals = append(appraisals, result)
	}
	return appraisals
}

//GetExitByID gets exit associated with an ID
func (r *EmployeeRepository) GetExitByID(id string) (exit models.Exit, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&exit)
	return
}

//GetExitsByEmpID gets exits associated with an EmpID
func (r *EmployeeRepository) GetExitsByEmpID(empID string) []models.Exit {
	var exits []models.Exit
	empid := bson.ObjectIdHex(empID)
	iter := r.C.Find(bson.M{"empid": empid}).Iter()
	result := models.Exit{}
	for iter.Next(&result) {
		exits = append(exits, result)
	}
	return exits
}

//GetExitsByOrgID gets exits associated with an OrgID
func (r *EmployeeRepository) GetExitsByOrgID(orgID string) []models.Exit {
	var exits []models.Exit
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Exit{}
	for iter.Next(&result) {
		exits = append(exits, result)
	}
	return exits
}

//AddSalaryComponentByOrgID persists salary component associated with OrgID
func (r *EmployeeRepository) AddSalaryComponentByOrgID(orgID string) (sal models.SalaryComponent, err error) {
	objID := bson.NewObjectId()
	sal.ID = objID
	sal.OrgID = bson.ObjectIdHex(orgID)

	err = r.C.Insert(&sal)
	return
}

//AddSalaryEmployeeByOrgID persists salary employees associated with OrgID
func (r *EmployeeRepository) AddSalaryEmployeesByOrgID(orgID string) (sal models.SalaryEmployee, err error) {
	objID := bson.NewObjectId()
	sal.ID = objID
	sal.OrgID = bson.ObjectIdHex(orgID)

	err = r.C.Insert(&sal)
	return
}

//AddSalaryStructureByMgrID persists salary structure associated with MgrID
func (r *EmployeeRepository) AddSalaryStructureByMgrID(orgID string, mgrID string, acctID string, compID string, salEmpID string) (sal models.SalaryStructure, err error) {
	objID := bson.NewObjectId()
	sal.ID = objID
	sal.OrgID = bson.ObjectIdHex(orgID)
	sal.GeneratedBy = bson.ObjectIdHex(mgrID)
	sal.PayableAccount = bson.ObjectIdHex(acctID)
	sal.SalaryComponentID = bson.ObjectIdHex(compID)
	sal.SalaryEmployeeID = bson.ObjectIdHex(salEmpID)

	err = r.C.Insert(&sal)
	return
}

//AddActivityTypeByOrgID persists activity type associated with OrgID
func (r *EmployeeRepository) AddActivityTypeByOrgID(orgID string) (actType models.ActivityType, err error) {
	objID := bson.NewObjectId()
	actType.ID = objID
	actType.OrgID = bson.ObjectIdHex(orgID)

	err = r.C.Insert(&actType)
	return
}

//AddWorkingHourByOrgID persists working hour associated with OrgID
func (r *EmployeeRepository) AddWorkingHourByOrgID(orgID string) (hour models.WorkingHour, err error) {
	objID := bson.NewObjectId()
	hour.ID = objID
	hour.OrgID = bson.ObjectIdHex(orgID)

	err = r.C.Insert(&hour)
	return
}

//AddOperationByOrgID persists operation associated with OrgID
func (r *EmployeeRepository) AddOperationByOrgID(orgID string) (op models.Operation, err error) {
	objID := bson.NewObjectId()
	op.ID = objID
	op.OrgID = bson.ObjectIdHex(orgID)

	err = r.C.Insert(&op)
	return
}

//AddWorkStationByOrgID persists workstation associated with OrgID
func (r *EmployeeRepository) AddWorkstationByOrgID(orgID string, opID string, hourID string) (station models.Workstation, err error) {
	objID := bson.NewObjectId()
	station.ID = objID
	station.OrgID = bson.ObjectIdHex(orgID)
	station.OperationID = bson.ObjectIdHex(opID)
	station.WorkingHourID = bson.ObjectIdHex(hourID)

	err = r.C.Insert(&station)
	return
}

//AddTimesheetByEmpID persists timesheet associated with EmpID
func (r *EmployeeRepository) AddTimesheetByEmpID(orgID string, empID string, mgrID string, projectID string, typeID string, stationID string) (time models.Timesheet, err error) {
	objID := bson.NewObjectId()
	time.ID = objID
	time.OrgID = bson.ObjectIdHex(orgID)
	time.ForEmployee = bson.ObjectIdHex(empID)
	time.ApprovedBy = bson.ObjectIdHex(mgrID)
	time.ProjectID = bson.ObjectIdHex(projectID)
	time.WorkStationID = bson.ObjectIdHex(stationID)
	time.ActivityTypeID = bson.ObjectIdHex(typeID)

	err = r.C.Insert(&time)
	return
}

//AddSalarySlipByEmpID persists salary slip associated with EmpID
func (r *EmployeeRepository) AddSalarySlipByEmpID(orgID string, empID string, mgrID string, deptID string, branchID string, sheetID string) (slip models.SalarySlip, err error) {
	objID := bson.NewObjectId()
	slip.ID = objID
	slip.OrgID = bson.ObjectIdHex(orgID)
	slip.ForEmployee = bson.ObjectIdHex(empID)
	slip.ApprovedBy = bson.ObjectIdHex(mgrID)
	slip.DepartmentID = bson.ObjectIdHex(deptID)
	slip.BranchID = bson.ObjectIdHex(branchID)
	slip.TimesheetID = bson.ObjectIdHex(sheetID)
	slip.PostingDate = time.Now()

	err = r.C.Insert(&slip)
	return
}

//AddActivityCostByEmpID persists activity cost associated with EmpID
func (r *EmployeeRepository) AddActivityCostByOrgID(typeID string, empID string) (cost models.ActivityCost, err error) {
	objID := bson.NewObjectId()
	cost.ID = objID
	cost.EmployeeID = bson.ObjectIdHex(empID)
	cost.ActivityType = bson.ObjectIdHex(typeID)

	err = r.C.Insert(&cost)
	return
}

//GetSalaryComponentByID gets salary component associated with an ID
func (r *EmployeeRepository) GetSalaryComponentByID(id string) (comp models.SalaryComponent, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&comp)
	return
}

//GetSalaryComponentsByOrgID gets salary components associated with an OrgID
func (r *EmployeeRepository) GetSalaryComponentsByOrgID(orgID string) []models.SalaryComponent {
	var components []models.SalaryComponent
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.SalaryComponent{}
	for iter.Next(&result) {
		components = append(components, result)
	}
	return components
}

//GetSalaryEmployeeByID gets salary employee associated with an ID
func (r *EmployeeRepository) GetSalaryEmployeeByID(id string) (sal models.SalaryEmployee, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&sal)
	return
}

//GetSalaryEmployeesByOrgID gets salary employees associated with an OrgID
func (r *EmployeeRepository) GetSalaryEmployeesByOrgID(orgID string) []models.SalaryEmployee {
	var salEmployees []models.SalaryEmployees
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.SalaryEmployees{}
	for iter.Next(&result) {
		employees := result.Employees
		for _, e := range employees {
			salEmployees = append(e, result)
		}
	}
	return salEmployees
}

//GetSalaryStructureByID gets salary structure associated with an ID
func (r *EmployeeRepository) GetSalaryStructureByID(id string) (sal models.SalaryStructure, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&sal)
	return
}

//GetSalaryStructuresByOrgID gets salary structures associated with an EmpID
func (r *EmployeeRepository) GetSalaryStructuresByEmpID(mgrID string) []models.SalaryStructure {
	var structures []models.SalaryStructure
	mgrid := bson.ObjectIdHex(mgrID)
	iter := r.C.Find(bson.M{"mgrid": mgrid}).Iter()
	result := models.SalaryStructure{}
	for iter.Next(&result) {
		structures = append(structures, result)
	}
	return structures
}

//GetSalaryStructuresByAcctID gets salary structures associated with an AcctID
func (r *EmployeeRepository) GetSalaryStructuresByAcctID(acctID string) []models.SalaryStructure {
	var structures []models.SalaryStructure
	accountid := bson.ObjectIdHex(acctID)
	iter := r.C.Find(bson.M{"accountid": accountid}).Iter()
	result := models.SalaryStructure{}
	for iter.Next(&result) {
		structures = append(structures, result)
	}
	return structures
}

//GetActivityTypeByID gets activity type associated with an ID
func (r *EmployeeRepository) GetActivityTypeByID(id string) (actType models.ActivityType, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&actType)
	return
}

//GetActivityTypesByOrgID gets activity types associated with an OrgID
func (r *EmployeeRepository) GetActivityTypesByOrgID(orgID string) []models.ActivityType {
	var types []models.ActivityType
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.ActivityType{}
	for iter.Next(&result) {
		types = append(types, result)
	}
	return types
}

//GetWorkingHourByID gets working hour associated with an ID
func (r *EmployeeRepository) GetWorkingHourByID(id string) (hour models.WorkingHour, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&hour)
	return
}

//GetWorkingHoursByOrgID gets working hours associated with an OrgID
func (r *EmployeeRepository) GetWorkingHoursByOrgID(orgID string) []models.WorkingHour {
	var hours []models.WorkingHour
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.WorkingHour{}
	for iter.Next(&result) {
		hours = append(hours, result)
	}
	return hours
}

//GetOperationByID gets operation associated with an ID
func (r *EmployeeRepository) GetOperationByID(id string) (op models.Operation, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&op)
	return
}

//GetOperationByOrgID gets operations associated with an OrgID
func (r *EmployeeRepository) GetOperationsByOrgID(orgID string) []models.Operation {
	var ops []models.Operation
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Operation{}
	for iter.Next(&result) {
		ops = append(ops, result)
	}
	return ops
}

//GetWorkstationByID gets workstation associated with an ID
func (r *EmployeeRepository) GetWorkstationByID(id string) (work models.Workstation, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&work)
	return
}

//GetWorkstationsByOrgID gets workstation associated with an OrgID
func (r *EmployeeRepository) GetWorkstationsByOrgID(orgID string) []models.Workstation {
	var stations []models.Workstation
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Workstation{}
	for iter.Next(&result) {
		stations = append(stations, result)
	}
	return stations
}

//GetTimesheetByID gets timesheet associated with an ID
func (r *EmployeeRepository) GetTimesheetByID(id string) (time models.Timesheet, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&time)
	return
}

//GetTimesheetsByEmpID gets timesheets associated with an EmpID
func (r *EmployeeRepository) GetTimesheetsByEmpID(empID string) []models.Timesheet {
	var sheets []models.Timesheet
	empid := bson.ObjectIdHex(empID)
	iter := r.C.Find(bson.M{"empid": empid}).Iter()
	result := models.Timesheet{}
	for iter.Next(&result) {
		sheets = append(sheets, result)
	}
	return sheets
}

//GetTimesheetsByOrgID gets timesheets associated with an OrgID
func (r *EmployeeRepository) GetTimesheetsByOrgID(orgID string) []models.Timesheet {
	var sheets []models.Timesheet
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Timesheet{}
	for iter.Next(&result) {
		sheets = append(sheets, result)
	}
	return sheets
}

//GetTimesheetsByMgrID gets timesheets associated with an MgrID
func (r *EmployeeRepository) GetTimesheetsByMgrID(mgrID string) []models.Timesheet {
	var sheets []models.Timesheet
	mgrid := bson.ObjectIdHex(mgrID)
	iter := r.C.Find(bson.M{"mgrid": mgrid}).Iter()
	result := models.Timesheet{}
	for iter.Next(&result) {
		sheets = append(sheets, result)
	}
	return sheets
}

//GetSalarySlipByID gets salary slip associated with an ID
func (r *EmployeeRepository) GetSalarySlipByID(id string) (slip models.SalarySlip, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&slip)
	return
}

//GetSalarySlipsByEmpID gets salary slips associated with an EmpID
func (r *EmployeeRepository) GetSalarySlipsByEmpID(empID string) []models.SalarySlip {
	var slips []models.SalarySlip
	empid := bson.ObjectIdHex(empID)
	iter := r.C.Find(bson.M{"empid": empid}).Iter()
	result := models.SalarySlip{}
	for iter.Next(&result) {
		slips = append(slips, result)
	}
	return slips
}

//GetSalarySlipsByOrgID gets salary slips associated with an OrgID
func (r *EmployeeRepository) GetSalarySlipsByOrgID(orgID string) []models.SalarySlip {
	var slips []models.SalarySlip
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.SalarySlip{}
	for iter.Next(&result) {
		slips = append(slips, result)
	}
	return slips
}

//GetSalarySlipsByMgrID gets salary slips associated with an MgrID
func (r *EmployeeRepository) GetSalarySlipsByMgrID(mgrID string) []models.SalarySlip {
	var slips []models.SalarySlip
	mgrid := bson.ObjectIdHex(mgrID)
	iter := r.C.Find(bson.M{"mgrid": mgrid}).Iter()
	result := models.SalarySlip{}
	for iter.Next(&result) {
		slips = append(slips, result)
	}
	return slips
}

//GetActivityCostByID gets activity cost associated with an ID
func (r *EmployeeRepository) GetActivityCostByID(id string) (cost models.ActivityCost, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&cost)
	return
}

//GetActivityCostsByEmpID gets activty costs associated with an EmpID
func (r *EmployeeRepository) GetActivityCostsByEmpID(empID string) []models.ActivityCost {
	var costs []models.ActivityCost
	empid := bson.ObjectIdHex(empID)
	iter := r.C.Find(bson.M{"empid": empid}).Iter()
	result := models.ActivityCost{}
	for iter.Next(&result) {
		costs = append(costs, result)
	}
	return costs
}

//GetActivityCostsByOrgID gets activty costs associated with an OrgID
func (r *EmployeeRepository) GetActivityCostsByOrgID(orgID string) []models.ActivityCost {
	var costs []models.ActivityCost
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.ActivityCost{}
	for iter.Next(&result) {
		costs = append(costs, result)
	}
	return costs
}

//GetActivityCostsByTypeID gets activty costs associated with a TypeID
func (r *EmployeeRepository) GetActivityCostsByTypeID(typeID string) []models.ActivityCost {
	var costs []models.ActivityCost
	typeid := bson.ObjectIdHex(typeID)
	iter := r.C.Find(bson.M{"typeid": typeid}).Iter()
	result := models.ActivityCost{}
	for iter.Next(&result) {
		costs = append(costs, result)
	}
	return costs
}
