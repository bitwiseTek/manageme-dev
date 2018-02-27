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
func (r *EmployeeRepository) AddEmpByOrgUserID(orgUserID string, orgID string, branchID string, deptID string) (emp models.Employee, err error) {
	objID := bson.NewObjectId()
	emp.ID = objID
	emp.OrgID = bson.ObjectIdHex(orgID)
	emp.OrgUserID = bson.ObjectIdHex(orgUserID)
	emp.BranchID = bson.ObjectIdHex(branchID)
	emp.DepartmentID = bson.ObjectIdHex(deptID)
	emp.Status = "Active"

	err = r.C.Insert(&emp)
	return
}

//AddEmp persists Employee associated with BranchID, DepartmentID, OrgUserID, EmpID
func (r *EmployeeRepository) AddChildEmpByOrgUserID(orgUserID string, orgID string, branchID string, deptID string, empID string) (emp models.Employee, err error) {
	objID := bson.NewObjectId()
	emp.ID = objID
	emp.OrgID = bson.ObjectIdHex(orgID)
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

//GetEmpByUserID gets employee associated with a UserID
func (r *EmployeeRepository) GetEmpByUserID(orgUserID string) (emp models.Employee, err error) {
	err = r.C.Find(bson.ObjectIdHex(orgUserID)).One(&emp)
	return
}

//GetEmpsByOrgID gets emps associated with an OrgID
func (r *EmployeeRepository) GetEmpsByOrgID(orgID string) []models.Employee {
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
	iter := r.C.Find(bson.M{"empid": mgrid}).Iter()
	result := models.Employee{}
	for iter.Next(&result) {
		emps = append(emps, result)
	}
	return emps
}

//EditEmpByID edits emp associated with an ID
func (r *EmployeeRepository) EditEmpByID(emp *models.Employee) error {
	err := r.C.Update(bson.M{"_id": emp.ID},
		bson.M{"$set": bson.M{
			"joiningdate":               emp.JoiningDate,
			"employmenttype":            emp.EmploymentType,
			"scheduledconfirmationdate": emp.ScheduledConfirmationDate,
			"finalconfirmationdate":     emp.FinalConfirmationDate,
			"contractenddate":           emp.ContractEndDate,
			"retirementdate":            emp.RetirementDate,
			"status":                    emp.Status,
			"bankname":                  emp.BankName,
			"bankaccount":               emp.BankAccount,
			"firstname":                 emp.FirstName,
			"lastname":                  emp.LastName,
			"middlename":                emp.MiddleName,
			"addresspermanent":          emp.AddressPermanent,
			"addresscurrent":            emp.AddressCurrent,
			"emailpersonal":             emp.EmailPersonal,
			"emailcompany":              emp.EmailCompany,
			"image":                     emp.Image,
			"accommodationtype":         emp.AccommodationType,
			"primaryphone":              emp.PrimaryPhone,
			"secondaryphone":            emp.SecondaryPhone,
			"emfullname":                emp.EmFullName,
			"emaddress":                 emp.EmAddress,
			"ememail":                   emp.EmEmail,
			"emrelationshiptype":        emp.EmRelationshipType,
			"emaccommodationtype":       emp.EmAccommodationType,
			"emphone":                   emp.EmPhone,
			"updatedat":                 time.Now(),
		}})
	return err
}

//DeleteEmpById deletes emp out of the system by Id
func (r *EmployeeRepository) DeleteByEmpById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//EditEmpByUserID edits emp associated with a UserID
func (r *EmployeeRepository) EditEmpByUserID(emp *models.Employee) error {
	err := r.C.Update(bson.M{"userid": emp.OrgUserID},
		bson.M{"$set": bson.M{
			"joiningdate":               emp.JoiningDate,
			"employmenttype":            emp.EmploymentType,
			"scheduledconfirmationdate": emp.ScheduledConfirmationDate,
			"finalconfirmationdate":     emp.FinalConfirmationDate,
			"contractenddate":           emp.ContractEndDate,
			"retirementdate":            emp.RetirementDate,
			"status":                    emp.Status,
			"bankname":                  emp.BankName,
			"bankaccount":               emp.BankAccount,
			"firstname":                 emp.FirstName,
			"lastname":                  emp.LastName,
			"middlename":                emp.MiddleName,
			"addresspermanent":          emp.AddressPermanent,
			"addresscurrent":            emp.AddressCurrent,
			"emailpersonal":             emp.EmailPersonal,
			"emailcompany":              emp.EmailCompany,
			"image":                     emp.Image,
			"accommodationtype":         emp.AccommodationType,
			"primaryphone":              emp.PrimaryPhone,
			"secondaryphone":            emp.SecondaryPhone,
			"emfullname":                emp.EmFullName,
			"emaddress":                 emp.EmAddress,
			"ememail":                   emp.EmEmail,
			"emrelationshiptype":        emp.EmRelationshipType,
			"emaccommodationtype":       emp.EmAccommodationType,
			"emphone":                   emp.EmPhone,
			"updatedat":                 time.Now(),
		}})
	return err
}

//DeleteEmpByUserId deletes emp out of the system by a UserId
func (r *EmployeeRepository) DeleteByEmpByUserId(id string) error {
	err := r.C.Remove(bson.M{"orguserid": bson.ObjectIdHex(id)})
	return err
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

//GetBioByEmpID gets biodata associated with an EmpID
func (r *EmployeeRepository) GetBioByEmpID(empID string) (bio models.Biodata, err error) {
	err = r.C.Find(bson.ObjectIdHex(empID)).One(&bio)
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

//EditBioByEmpID edits bio associated with an EmpID
func (r *EmployeeRepository) EditBiodataByEmpID(bio *models.Biodata) error {
	err := r.C.Update(bson.M{"empid": bio.EmployeeID},
		bson.M{"$set": bson.M{
			"dateofbirth":    bio.DateOfBirth,
			"sex":            bio.Sex,
			"bloodgroup":     bio.BloodGroup,
			"maritalstatus":  bio.MaritalStatus,
			"disabilitytype": bio.DisabilityType,
			"nationality":    bio.Nationality,
			"stateoforigin":  bio.StateOfOrigin,
		}})
	return err
}

//DeleteBioByEmpId deletes biodata out of the system by EmpId
func (r *EmployeeRepository) DeleteBiodataByEmpId(id string) error {
	err := r.C.Remove(bson.M{"empid": bson.ObjectIdHex(id)})
	return err
}

//EditBioByID edits bio associated with an ID
func (r *EmployeeRepository) EditBiodataByID(bio *models.Biodata) error {
	err := r.C.Update(bson.M{"_id": bio.ID},
		bson.M{"$set": bson.M{
			"dateofbirth":    bio.DateOfBirth,
			"sex":            bio.Sex,
			"bloodgroup":     bio.BloodGroup,
			"maritalstatus":  bio.MaritalStatus,
			"disabilitytype": bio.DisabilityType,
			"nationality":    bio.Nationality,
			"stateoforigin":  bio.StateOfOrigin,
		}})
	return err
}

//DeleteBioById deletes biodata out of the system by Id
func (r *EmployeeRepository) DeleteBiodataById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetPIDByID gets personal ID associated with an ID
func (r *EmployeeRepository) GetPIDByID(id string) (pid models.PersonalIdentification, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&pid)
	return
}

//GetPIDByBioID gets personal ID associated with a BioID
func (r *EmployeeRepository) GetPIDByBioID(bioID string) (pid models.PersonalIdentification, err error) {
	err = r.C.Find(bson.ObjectIdHex(bioID)).One(&pid)
	return
}

//GetPIDsByBioID gets pids associated with a BioID
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

//EditPIDByBioID edits PID associated with a BioID
func (r *EmployeeRepository) EditPIDByBioID(pid *models.PersonalIdentification) error {
	err := r.C.Update(bson.M{"bioid": pid.BioID},
		bson.M{"$set": bson.M{
			"attachments": pid.Attachments,
			"idtype":      pid.IDType,
			"idno":        pid.IDNo,
			"validtill":   pid.ValidTill,
			"issueplace":  pid.IssuePlace,
			"issuedate":   pid.IssueDate,
		}})
	return err
}

//EditPIDByID edits PID associated with an ID
func (r *EmployeeRepository) EditPIDByID(pid *models.PersonalIdentification) error {
	err := r.C.Update(bson.M{"_id": pid.ID},
		bson.M{"$set": bson.M{
			"attachments": pid.Attachments,
			"idtype":      pid.IDType,
			"idno":        pid.IDNo,
			"validtill":   pid.ValidTill,
			"issueplace":  pid.IssuePlace,
			"issuedate":   pid.IssueDate,
		}})
	return err
}

//GetEduByID gets edu associated with an ID
func (r *EmployeeRepository) GetEduByID(id string) (edu models.Education, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&edu)
	return
}

//GetEduByBioID gets edu associated with a BioID
func (r *EmployeeRepository) GetEduByBioID(bioID string) (edu models.Education, err error) {
	err = r.C.Find(bson.ObjectIdHex(bioID)).One(&edu)
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

//EditEduByBioID edits Edu associated with a BioID
func (r *EmployeeRepository) EditEduByBioID(edu *models.Education) error {
	err := r.C.Update(bson.M{"bioid": edu.BioID},
		bson.M{"$set": bson.M{
			"attachments":    edu.Attachments,
			"qualification":  edu.Qualification,
			"cgpa":           edu.CGPA,
			"graduatedon":    edu.GraduatedOn,
			"schoolattended": edu.SchoolAttended,
			"honortype":      edu.HonorType,
		}})
	return err
}

//EditEduByID edits Edu associated with an ID
func (r *EmployeeRepository) EditEduByID(edu *models.Education) error {
	err := r.C.Update(bson.M{"_id": edu.ID},
		bson.M{"$set": bson.M{
			"attachments":    edu.Attachments,
			"qualification":  edu.Qualification,
			"cgpa":           edu.CGPA,
			"graduatedon":    edu.GraduatedOn,
			"schoolattended": edu.SchoolAttended,
			"honortype":      edu.HonorType,
		}})
	return err
}

//GetWorkByID gets work experience associated with an ID
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

//GetWorkByBioID gets work experience associated with a BioID
func (r *EmployeeRepository) GetWorkByBioID(bioID string) (work models.WorkExperience, err error) {
	err = r.C.Find(bson.ObjectIdHex(bioID)).One(&work)
	return
}

//EditWorkByBioID edits work associated with a BioID
func (r *EmployeeRepository) EditWorkByBioID(work *models.WorkExperience) error {
	err := r.C.Update(bson.M{"bioid": work.BioID},
		bson.M{"$set": bson.M{
			"attachments":    work.Attachments,
			"companyworked":  work.CompanyWorked,
			"workhistoryext": work.WorkHistoryExt,
			"designation":    work.Designation,
			"resignedon":     work.ResignedOn,
			"address":        work.Address,
			"salary":         work.Salary,
		}})
	return err
}

//EditWorkByID edits work associated with an ID
func (r *EmployeeRepository) EditWorkByID(work *models.WorkExperience) error {
	err := r.C.Update(bson.M{"_id": work.ID},
		bson.M{"$set": bson.M{
			"attachments":    work.Attachments,
			"companyworked":  work.CompanyWorked,
			"workhistoryext": work.WorkHistoryExt,
			"designation":    work.Designation,
			"resignedon":     work.ResignedOn,
			"address":        work.Address,
			"salary":         work.Salary,
		}})
	return err
}

//GetHealthByID gets health detail associated with an ID
func (r *EmployeeRepository) GetHealthDetailByID(id string) (health models.HealthDetail, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&health)
	return
}

//GetWHealthByBioID gets health detail associated with a BioID
func (r *EmployeeRepository) GetHealthDetailsByBioID(bioID string) []models.HealthDetail {
	var healths []models.HealthDetail
	bioid := bson.ObjectIdHex(bioID)
	iter := r.C.Find(bson.M{"bioid": bioid}).Iter()
	result := models.HealthDetail{}
	for iter.Next(&result) {
		healths = append(healths, result)
	}
	return healths
}

//GetHealthByBioID gets health detail associated with a BioID
func (r *EmployeeRepository) GetHealthDetailByBioID(bioID string) (health models.HealthDetail, err error) {
	err = r.C.Find(bson.ObjectIdHex(bioID)).One(&health)
	return
}

//EditHealthByBioID edits health detail associated with a BioID
func (r *EmployeeRepository) EditHealthByBioID(health *models.HealthDetail) error {
	err := r.C.Update(bson.M{"bioid": health.BioID},
		bson.M{"$set": bson.M{
			"attachments":    health.Attachments,
			"height":         health.Height,
			"weight":         health.Weight,
			"eyecolor":       health.EyeColor,
			"knownallergies": health.KnownAllergies,
			"healthconcerns": health.HealthConcerns,
		}})
	return err
}

//EditHealthByID edits health detail associated with an ID
func (r *EmployeeRepository) EditHealthByID(health *models.HealthDetail) error {
	err := r.C.Update(bson.M{"_id": health.ID},
		bson.M{"$set": bson.M{
			"attachments":    health.Attachments,
			"height":         health.Height,
			"weight":         health.Weight,
			"eyecolor":       health.EyeColor,
			"knownallergies": health.KnownAllergies,
			"healthconcerns": health.HealthConcerns,
		}})
	return err
}

//AddExpenseClaimByEmpID persists expense claim associated with EmpID, MgrID
func (r *EmployeeRepository) AddExpenseClaimByEmpID(orgID string, empID string, mgrID string, typeID string, projectID string, taskID string, accountID string) (claim models.ExpenseClaim, err error) {
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
func (r *EmployeeRepository) AddLeaveAllByEmpID(orgID string, empID string, mgrID string, typeID string) (leaveAll models.LeaveAllocation, err error) {
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
func (r *EmployeeRepository) AddLeaveAppByEmpID(orgID string, empID string, mgrID string, typeID string) (leaveApp models.LeaveApplication, err error) {
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

//EditClaimsByEmpID edits claim associated with an EmpID
func (r *EmployeeRepository) EditClaimByEmpID(claim *models.ExpenseClaim) error {
	err := r.C.Update(bson.M{"empid": claim.ExpApplier},
		bson.M{"$set": bson.M{
			"description": claim.Description,
			"claimamount": claim.ClaimAmount,
			"remarks":     claim.Remarks,
			"expensedate": claim.ExpenseDate,
			"updatedat":   time.Now(),
		}})
	return err
}

//EditClaimsByMgrID edits claim associated with a MgrID
func (r *EmployeeRepository) EditClaimByMgrID(claim *models.ExpenseClaim) error {
	err := r.C.Update(bson.M{"mgrid": claim.ExpApprover},
		bson.M{"$set": bson.M{
			"ispaid":                claim.Description,
			"approvalstatus":        claim.ClaimAmount,
			"totalsanctionedamount": claim.Remarks,
			"paymentmode":           claim.PaymentMode,
			"status":                claim.Status,
			"sanctionedamount":      claim.SanctionedAmount,
			"updatedat":             time.Now(),
		}})
	return err
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

//EditLeaveAllocationByID edits leave allocation associated with an ID
func (r *EmployeeRepository) EditLeaveAllocationByID(leave *models.LeaveAllocation) error {
	err := r.C.Update(bson.M{"_id": leave.ID},
		bson.M{"$set": bson.M{
			"description": leave.Description,
			"fromdate":    leave.FromDate,
			"todate":      leave.ToDate,
			"updatedat":   time.Now(),
		}})
	return err
}

//DeleteLeaveAllocationById deletes leave allocation out of the system by Id
func (r *EmployeeRepository) DeleteLeaveAllocationById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetLeaveApplicationByID gets leave application associated with an ID
func (r *EmployeeRepository) GetLeaveApplicationByID(id string) (leaveApp models.LeaveApplication, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&leaveApp)
	return
}

//GetLeaveApplicationsByMgrID gets leave applications associated with a MgrID
func (r *EmployeeRepository) GetLeaveApplicationsByMgrID(mgrID string) []models.LeaveApplication {
	var apps []models.LeaveApplication
	mgrid := bson.ObjectIdHex(mgrID)
	iter := r.C.Find(bson.M{"mgrid": mgrid}).Iter()
	result := models.LeaveApplication{}
	for iter.Next(&result) {
		apps = append(apps, result)
	}
	return apps
}

//GetLeaveApplicationsByOrgID gets leave applications associated with a OrgID
func (r *EmployeeRepository) GetLeaveApplicationsByOrgID(orgID string) []models.LeaveApplication {
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

//EditLeaveApplicationByID edits leave application associated with an ID
func (r *EmployeeRepository) EditLeaveApplicationByID(leave *models.LeaveApplication) error {
	err := r.C.Update(bson.M{"_id": leave.ID},
		bson.M{"$set": bson.M{
			"description": leave.Description,
			"fromdate":    leave.FromDate,
			"todate":      leave.ToDate,
			"updatedat":   time.Now(),
		}})
	return err
}

//DeleteLeaveApplicationById deletes leave application out of the system by Id
func (r *EmployeeRepository) DeleteLeaveApplicationById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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

//EditLeaveBlockListByID edits leave block list associated with an ID
func (r *EmployeeRepository) EditLeaveBlockListByID(blk *models.LeaveBlockList) error {
	err := r.C.Update(bson.M{"_id": blk.ID},
		bson.M{"$set": bson.M{
			"name":      blk.Name,
			"blockdate": blk.BlockDate,
			"reason":    blk.Reason,
			"applyall":  blk.ApplyAll,
			"updatedat": time.Now(),
		}})
	return err
}

//DeleteLeaveBlockListById deletes leave blocklist out of the system by Id
func (r *EmployeeRepository) DeleteLeaveBlockListById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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

//EditHolidayListByID edits holiday list associated with an ID
func (r *EmployeeRepository) EditHolidayListByID(hol *models.HolidayList) error {
	err := r.C.Update(bson.M{"_id": hol.ID},
		bson.M{"$set": bson.M{
			"name":        hol.Name,
			"fromdate":    hol.FromDate,
			"todate":      hol.ToDate,
			"weeklyoff":   hol.WeeklyOff,
			"description": hol.Description,
			"holidaydate": hol.HolidayDate,
		}})
	return err
}

//DeleteLHolidayListById deletes holiday list out of the system by Id
func (r *EmployeeRepository) DeleteHolidayListById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Appraisal{}
	for iter.Next(&result) {
		appraisals = append(appraisals, result)
	}
	return appraisals
}

//EditAppraisalByID edits appraisal associated with an ID
func (r *EmployeeRepository) EditAppraisalByID(app *models.Appraisal) error {
	err := r.C.Update(bson.M{"_id": app.ID},
		bson.M{"$set": bson.M{
			"stardate":     app.StartDate,
			"enddate":      app.EndDate,
			"remarks":      app.Remarks,
			"title":        app.Title,
			"description":  app.Description,
			"kra":          app.KRA,
			"perweightage": app.PerWeightage,
			"score":        app.Score,
			"status":       app.Status,
		}})
	return err
}

//DeleteAppraisalById deletes appraisal out of the system by Id
func (r *EmployeeRepository) DeleteAppraisalById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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

//EditExitByID edits exit associated with an ID
func (r *EmployeeRepository) EditExitByID(ex *models.Exit) error {
	err := r.C.Update(bson.M{"_id": ex.ID},
		bson.M{"$set": bson.M{
			"leavingreason":     ex.LeavingReason,
			"relievingdate":     ex.RelievingDate,
			"leaveencashed":     ex.LeaveEncashed,
			"encashmentdate":    ex.EncashmentDate,
			"interviewdate":     ex.InterviewDate,
			"resignationreason": ex.ResignationReason,
			"newworkplace":      ex.NewWorkPlace,
			"feedback":          ex.Feedback,
		}})
	return err
}

//DeleteExitById deletes exit out of the system by Id
func (r *EmployeeRepository) DeleteExitById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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
func (r *EmployeeRepository) AddActivityCostByOrgID(orgID string, typeID string, empID string) (cost models.ActivityCost, err error) {
	objID := bson.NewObjectId()
	cost.ID = objID
	cost.OrgID = bson.ObjectIdHex(orgID)
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

//EditSalaryComponentByID edits salary component associated with an ID
func (r *EmployeeRepository) EditSalaryComponentByID(sal *models.SalaryComponent) error {
	err := r.C.Update(bson.M{"_id": sal.ID},
		bson.M{"$set": bson.M{
			"name":        sal.Name,
			"type":        sal.Type,
			"description": sal.Description,
		}})
	return err
}

//GetSalaryEmployeeByID gets salary employee associated with an ID
func (r *EmployeeRepository) GetSalaryEmployeeByID(id string) (sal models.SalaryEmployee, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&sal)
	return
}

//GetSalaryEmployeesByOrgID gets salary employees associated with an OrgID
func (r *EmployeeRepository) GetSalaryEmployeesByOrgID(orgID string) []models.SalaryEmployee {
	var salEmployees []models.SalaryEmployee
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.SalaryEmployee{}
	for iter.Next(&result) {
		employees := result.Employees
		for _, e := range employees {
			//e = append(e, result)
		}
	}
	return salEmployees
}

//EditSalaryEmployeeByID edits salary employee associated with an ID
func (r *EmployeeRepository) EditSalaryEmployeeByID(sal *models.SalaryEmployee) error {
	err := r.C.Update(bson.M{"_id": sal.ID},
		bson.M{"$set": bson.M{
			"fromdate": sal.FromDate,
			"todate":   sal.ToDate,
			"base":     sal.Base,
			"variable": sal.Variable,
		}})
	return err
}

//DeleteSalaryEmployeeById deletes salary employee out of the system by Id
func (r *EmployeeRepository) DeleteSalaryEmployeeById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetSalaryStructureByID gets salary structure associated with an ID
func (r *EmployeeRepository) GetSalaryStructureByID(id string) (sal models.SalaryStructure, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&sal)
	return
}

//GetSalaryStructuresByOrgID gets salary structures associated with an EmpID
func (r *EmployeeRepository) GetSalaryStructuresByOrgID(orgID string) []models.SalaryStructure {
	var structures []models.SalaryStructure
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.SalaryStructure{}
	for iter.Next(&result) {
		structures = append(structures, result)
	}
	return structures
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

//EditSalaryStructureByID edits salary structure associated with an ID
func (r *EmployeeRepository) EditSalaryStructureByID(sal *models.SalaryStructure) error {
	err := r.C.Update(bson.M{"_id": sal.ID},
		bson.M{"$set": bson.M{
			"payrollfrequency":     sal.PayrollFrequency,
			"isactive":             sal.IsActive,
			"isdefault":            sal.IsDefault,
			"hourrate":             sal.HourRate,
			"paymentmode":          sal.PaymentMode,
			"isamountformulabased": sal.IsAmountFormulaBased,
			"isamountlwpbased":     sal.IsAmountLwpBased,
			"defaultamount":        sal.DefaultAmount,
			"amount":               sal.Amount,
			"updatedat":            time.Now(),
		}})
	return err
}

//DeleeSalaryStructurelById deletes salary structure out of the system by Id
func (r *EmployeeRepository) DeleteSalaryStructureById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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

//EditActivityTypeByID edits activity type associated with an ID
func (r *EmployeeRepository) EditActivityTypeByID(actType *models.ActivityType) error {
	err := r.C.Update(bson.M{"_id": actType.ID},
		bson.M{"$set": bson.M{
			"name":        actType.Name,
			"billingrate": actType.BillingRate,
			"costingrate": actType.CostingRate,
			"disabled":    actType.Disabled,
			"updatedat":   time.Now(),
		}})
	return err
}

//DeleteActivityTypeById deletes activity type out of the system by Id
func (r *EmployeeRepository) DeleteActivityTypeById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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

//EditWorkingHourByID edits working hour associated with an ID
func (r *EmployeeRepository) EditWorkingHourByID(work *models.WorkingHour) error {
	err := r.C.Update(bson.M{"_id": work.ID},
		bson.M{"$set": bson.M{
			"starttime": work.StartTime,
			"endtime":   work.EndTime,
			"enabled":   work.Enabled,
		}})
	return err
}

//DeleteWorkingHourById deletes working hour out of the system by Id
func (r *EmployeeRepository) DeleteWorkingHourById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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

//EditOperationByID edits operation associated with an ID
func (r *EmployeeRepository) EditOperationByID(op *models.Operation) error {
	err := r.C.Update(bson.M{"_id": op.ID},
		bson.M{"$set": bson.M{
			"description": op.Description,
		}})
	return err
}

//DeleteOperationById deletes operation out of the system by Id
func (r *EmployeeRepository) DeleteOperationById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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

//EditWorkstationByID edits workstation associated with an ID
func (r *EmployeeRepository) EditWorkstationByID(work *models.Workstation) error {
	err := r.C.Update(bson.M{"_id": work.ID},
		bson.M{"$set": bson.M{
			"name":                work.Name,
			"description":         work.Description,
			"hourrate":            work.HourRate,
			"hourrateelectricity": work.HourRateElectricity,
			"hourraterent":        work.HourRateRent,
			"hourratelabor":       work.HourRateLabor,
			"hourrateconsumable":  work.HourRateConsumable,
		}})
	return err
}

//DeleteWorkstationById deletes workstation out of the system by Id
func (r *EmployeeRepository) DeleteWorkstationById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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

//EditTimesheetByID edits timesheet associated with an ID
func (r *EmployeeRepository) EditTimesheetByID(timeSheet *models.Timesheet) error {
	err := r.C.Update(bson.M{"_id": timeSheet.ID},
		bson.M{"$set": bson.M{
			"stardate":      timeSheet.StartDate,
			"enddate":       timeSheet.EndDate,
			"fromtime":      timeSheet.FromTime,
			"totime":        timeSheet.ToTime,
			"completedqty":  timeSheet.CompletedQty,
			"hours":         timeSheet.Hours,
			"billable":      timeSheet.Billable,
			"billinghours":  timeSheet.BillingHours,
			"billingamount": timeSheet.BillingAmount,
			"costingamount": timeSheet.CostingAmount,
			"note":          timeSheet.Note,
			"updatedat":     time.Now(),
		}})
	return err
}

//DeleteTimesheetById deletes timesheet out of the system by Id
func (r *EmployeeRepository) DeleteTimesheetById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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

//EditSalarySlipByID edits salary slip associated with an ID
func (r *EmployeeRepository) EditSalarySlipByID(sal *models.SalarySlip) error {
	err := r.C.Update(bson.M{"_id": sal.ID},
		bson.M{"$set": bson.M{
			"stardate":              sal.StartDate,
			"enddate":               sal.EndDate,
			"issalarysliptimesheet": sal.IsSalarySlipTimesheet,
			"workinghours":          sal.WorkingHours,
			"paymentdays":           sal.PaymentDays,
			"designation":           sal.Designation,
			"grosspay":              sal.GrossPay,
			"interestamount":        sal.InterestAmount,
			"updatedat":             time.Now(),
		}})
	return err
}

//DeleteSalarySlipById deletes salary slip out of the system by Id
func (r *EmployeeRepository) DeleteSalarySlipById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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

//EditActivityCostByID edits activity cost associated with an ID
func (r *EmployeeRepository) EditActivityCostByID(cost *models.ActivityCost) error {
	err := r.C.Update(bson.M{"_id": cost.ID},
		bson.M{"$set": bson.M{
			"activitytype": cost.ActivityType,
			"costingrate":  cost.CostingRate,
			"billingrate":  cost.BillingRate,
			"status":       cost.Status,
			"updatedat":    time.Now(),
		}})
	return err
}

//DeleteActivityCostById deletes activity cost out of the system by Id
func (r *EmployeeRepository) DeleteActivityCostById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
