package data

/**
 *
 * @author Sika Kay
 * @date 03/01/18
 *
 */
import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"time"

	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
)

//OrgRepository Struct for Mongo Persistence
type OrgRepository struct {
	C *mgo.Collection
}

//AddOrg persists Org associated with a UserID
func (r *OrgRepository) AddOrgByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.ID = objID
	org.UserID = bson.ObjectIdHex(userID)
	org.DefaultCurrency = "NGN"
	org.Status = "Active"
	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()

	err = r.C.Insert(&org)
	return
}

//AddOrgContact persists Org/Contact associated with a UserID
func (r *OrgRepository) AddOrgContactByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.CompanyContacts.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgBilling persists Org/Billing associated with a UserID
func (r *OrgRepository) AddOrgBillingByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Billings.ID = objID
	org.Billings.LastBilled = time.Now()

	err = r.C.Insert(&org)
	return
}

//AddOrgBillingTransaction persists Org/Billing/Transaction associated with a UserID
func (r *OrgRepository) AddOrgBillingTransactionByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Billings.Transactions.ID = objID
	org.Billings.Transactions.CreatedAt = time.Now()
	org.Billings.Transactions.Status = "Pending"

	err = r.C.Insert(&org)
	return
}

//AddOrgUsers persists Org/Users associated with a UserID
func (r *OrgRepository) AddOrgUserByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.ID = objID
	hpass, err := bcrypt.GenerateFromPassword([]byte(org.Users.TempPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	org.Users.HashPassword = hpass
	org.Users.TempPassword = ""
	org.Users.Status = "Active"
	org.Users.CreatedAt = time.Now()
	org.Users.UpdatedAt = time.Now()

	err = r.C.Insert(&org)
	return
}

//AddOrgUsersRoles persists Org/Users/Roles associated with a UserID
func (r *OrgRepository) AddOrgUsersRolesByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Roles.ID = objID
	org.Users.Roles.Status = "Active"
	org.Users.Roles.CreatedAt = time.Now()
	org.Users.Roles.UpdatedAt = time.Now()

	err = r.C.Insert(&org)
	return
}

//AddOrgUsersRolesPermissions persists Org/Users/Roles/Permissions associated with a UserID
func (r *OrgRepository) AddOrgUsersRolesPermissionsByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Roles.Permissions.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployees persists Org/Users/Employees associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.ID = objID
	org.Users.Employees.Status = "Active"

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesContact persists Org/Users/Employees/Contact associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesContactByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.EmployeeContact.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesContactEmergency persists Org/Users/Employees/Contact/Emergency associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesContactEmergencyByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.EmployeeContact.EmergencyContact.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesBiodata persists Org/Users/Employees/Biodata associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesBiodataByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Biodata.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesBiodataPersonalID persists Org/Users/Employees/Biodata/PersonalID associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesBiodataPersoanlIDByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Biodata.PersonalIdentification.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesBiodataHealth persists Org/Users/Employees/Biodata/Health associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesBiodataHealthByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Biodata.HealthDetails.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesBiodataEducation persists Org/Users/Employees/Biodata/Education associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesBiodataEducationByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Biodata.Education.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesBiodataWork persists Org/Users/Employees/Biodata/Work associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesBiodataWorkByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Biodata.WorkExperience.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesExpenseClaim persists Org/Users/Employees/ExpenseClaim associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesExpenseClaimByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.ExpenseClaim.ID = objID
	org.Users.Employees.ExpenseClaim.ApprovalStatus = "Pending"
	org.Users.Employees.ExpenseClaim.Status = "Applied"
	org.Users.Employees.ExpenseClaim.PostingDate = time.Now()

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesExpenseClaimDetail persists Org/Users/Employees/ExpenseClaim/Detail associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesExpenseClaimDetailByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.ExpenseClaim.ExpenseDetail.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesLeaveAllocation persists Org/Users/Employees/LeaveAllocation associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesLeaveAllocationByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.LeaveAllocation.ID = objID
	org.Users.Employees.LeaveAllocation.AllocatedOn = time.Now()

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesLeaveApplication persists Org/Users/Employees/LeaveApplication associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesLeaveApplicationByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.LeaveApplication.ID = objID
	org.Users.Employees.LeaveApplication.AppliedAt = time.Now()
	org.Users.Employees.LeaveApplication.Status = "Pending"

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesLeaveBlockList persists Org/Users/Employees/LeaveBlockList associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesLeaveBlockListByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.LeaveBlockList.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesLeaveBlockListDates persists Org/Users/Employees/LeaveBlockList/Dates associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesLeaveBlockListDatesByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.LeaveBlockList.BlockedDates.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesLeaveBlockListUsers persists Org/Users/Employees/LeaveBlockList/Users associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesLeaveBlockListUsersByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.LeaveBlockList.AllowedUsers.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesHolidayList persists Org/Users/Employees/HolidayList associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesHolidayListByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.HolidayList.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesHolidayListHolidays persists Org/Users/Employees/HolidayList/Holidays associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesHolidayListHolidaysByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.HolidayList.Holidays.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesAppraisal persists Org/Users/Employees/Appraisal associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesAppraisalByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Appraisal.ID = objID
	org.Users.Employees.Appraisal.AppraisedAt = time.Now()
	org.Users.Employees.Appraisal.Status = "Initiated"

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesAppraisalTemplate persists Org/Users/Employees/Appraisal/Template associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesAppraisalTemplateByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Appraisal.AppraisalTemplate.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesAppraisalGoal persists Org/Users/Employees/Appraisal/Goal associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesAppraisalGoalByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Appraisal.AppraisalGoal.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesExit persists Org/Users/Employees/Exit associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesExitByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Exit.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalaryStructure persists Org/Users/Employees/SalaryStructure associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalaryStructureByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalaryStructure.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalaryStructureEmployees persists Org/Users/Employees/SalaryStructure/Employees associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalaryStructureEmployeesByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalaryStructure.SalaryEmployees.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalaryStructureDetails persists Org/Users/Employees/SalaryStructure/Details associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalaryStructureDetailsByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalaryStructure.SalaryDetail.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalaryStructureDetailComponents persists Org/Users/Employees/SalaryStructure/Details/Components associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalaryStructureDetailComponentsByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalaryStructure.SalaryDetail.SalaryComponent.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalarySlip persists Org/Users/Employees/SalarySlip associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalarySlipByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalarySlip.ID = objID
	org.Users.Employees.SalarySlip.PostingDate = time.Now()

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalarySlipByTimesheet persists Org/Users/Employees/SalarySlip/Timesheet associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalarySlipByTimesheetByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalarySlip.Timesheet.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalarySlipTimesheetDetail persists Org/Users/Employees/SalarySlip/Timesheet/Detail associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalarySlipTimesheetDetailByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalarySlip.Timesheet.TimesheetDetail.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalarySlipTimesheetDetailActivityType persists Org/Users/Employees/SalarySlip/Timesheet/Detail/ActivityType associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalarySlipTimesheetDetailActivityTypeByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalarySlip.Timesheet.TimesheetDetail.ActivityType.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalarySlipTimesheetDetailWorkStation persists Org/Users/Employees/SalarySlip/Timesheet/Detail/WorkStation associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalarySlipTimesheetDetailWorkStationByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalarySlip.Timesheet.TimesheetDetail.Workstation.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalarySlipTimesheetDetailWorkStationHours persists Org/Users/Employees/SalarySlip/Timesheet/Detail/WorkStation/Hours associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalarySlipTimesheetDetailWorkStationHoursByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalarySlip.Timesheet.TimesheetDetail.Workstation.WorkingHours.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalarySlipTimesheetDetailWorkStationOperation persists Org/Users/Employees/SalarySlip/Timesheet/Detail/WorkStation/Operation associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalarySlipTimesheetDetailWorkStationOperationByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalarySlip.Timesheet.TimesheetDetail.Workstation.Operation.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesSalarySlipTimesheet persists Org/Users/Employees/SalarySlipTimesheet associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesSalarySlipTimesheetByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.SalarySlipTimesheet.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesActivityCost persists Org/Users/Employees/ActivityCost associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesActivityCostByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.ActivityCost.ID = objID
	org.Users.Employees.ActivityCost.Status = "Active"

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesProject persists Org/Users/Employees/Project associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesProjectByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Project.ID = objID
	org.Users.Employees.Project.Status = "Not Started"

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesProjectTasks persists Org/Users/Employees/Project/Tasks associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesProjectTasksByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Project.ProjectTasks.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesProjectTask persists Org/Users/Employees/Project/Tasks/Task associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesProjectTaskByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Project.ProjectTasks.Task.ID = objID
	org.Users.Employees.Project.ProjectTasks.Task.Status = "Not Started"

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesProjectTaskDependent persists Org/Users/Employees/Project/Tasks/Task/Dependent associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesProjectTaskDependentByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.Project.ProjectTasks.Task.TaskDependent.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesJouranlEntry persists Org/Users/Employees/JournalEntry associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesJournalEntryByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.JournalEntry.ID = objID
	org.Users.Employees.JournalEntry.PostingDate = time.Now()

	err = r.C.Insert(&org)
	return
}

//AddOrgUserEmployeesJouranlEntryAccount persists Org/Users/Employees/JournalEntry/Account associated with a UserID
func (r *OrgRepository) AddOrgUserEmployeesJournalEntryAccountByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Users.Employees.JournalEntry.JournalAccount.ID = objID

	err = r.C.Insert(&org)
	return
}

//AddOrgDepartments persists Org/Departments associated with a UserID
func (r *OrgRepository) AddOrgDepartmentsByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Departments.ID = objID
	org.Departments.CreatedAt = time.Now()
	org.Departments.UpdatedAt = time.Now()
	org.Departments.Status = "Active"

	err = r.C.Insert(&org)
	return
}

//AddOrgDesignations persists Org/Designations associated with a UserID
func (r *OrgRepository) AddOrgDesignationsByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Designations.ID = objID
	org.Designations.CreatedAt = time.Now()
	org.Designations.UpdatedAt = time.Now()
	org.Designations.Status = "Active"

	err = r.C.Insert(&org)
	return
}

//AddOrgSalaryModes persists Org/SalaryModes associated with a UserID
func (r *OrgRepository) AddOrgSalaryModesByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.SalaryModes.ID = objID
	org.SalaryModes.CreatedAt = time.Now()
	org.SalaryModes.UpdatedAt = time.Now()
	org.SalaryModes.Status = "Active"

	err = r.C.Insert(&org)
	return
}

//AddOrgBranches persists Org/Branches associated with a UserID
func (r *OrgRepository) AddOrgBranchesByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Branches.ID = objID
	org.Branches.CreatedAt = time.Now()
	org.Branches.UpdatedAt = time.Now()
	org.Branches.Status = "Active"

	err = r.C.Insert(&org)
	return
}

//AddOrgLeaveType persists Org/LeaveType associated with a UserID
func (r *OrgRepository) AddOrgLeaveTypeByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.LeaveType.ID = objID
	org.LeaveType.CreatedAt = time.Now()
	org.LeaveType.UpdatedAt = time.Now()
	org.LeaveType.Status = "Active"

	err = r.C.Insert(&org)
	return
}

//AddOrgExpenseClaimType persists Org/ExpenseClaimType associated with a UserID
func (r *OrgRepository) AddOrgExpenseClaimTypeByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.ExpenseClaimType.ID = objID
	org.ExpenseClaimType.CreatedAt = time.Now()
	org.ExpenseClaimType.UpdatedAt = time.Now()
	org.ExpenseClaimType.Status = "Active"

	err = r.C.Insert(&org)
	return
}

//AddOrgProjectType persists Org/ProjectType associated with a UserID
func (r *OrgRepository) AddOrgProjectTypeByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.ProjectType.ID = objID
	org.ProjectType.CreatedAt = time.Now()
	org.ProjectType.UpdatedAt = time.Now()
	org.ProjectType.Status = "Active"

	err = r.C.Insert(&org)
	return
}

//AddOrgAccount persists Org/Account associated with a UserID
func (r *OrgRepository) AddOrgAccountByUserID(userID string) (org models.Org, err error) {
	objID := bson.NewObjectId()
	org.Account.ID = objID
	org.Account.CreatedAt = time.Now()
	org.Account.UpdatedAt = time.Now()
	org.Account.Status = "Active"

	err = r.C.Insert(&org)
	return
}

//GetOrgs gets all orgs
func (r *OrgRepository) GetOrgs() []models.Org {
	var orgs []models.Org
	iter := r.C.Find(nil).Iter()
	result := models.Org{}
	for iter.Next(&result) {
		orgs = append(orgs, result)
	}
	return orgs
}

//GetOrgByID gets org associated with an ID
func (r *OrgRepository) GetOrgByID(id string) (orgs models.Org, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&orgs)
	return
}

//GetOrgByUserID gets org associated with a UserID
func (r *OrgRepository) GetOrgByUserID(userID string) (orgs models.Org, err error) {
	err = r.C.Find(bson.ObjectIdHex(userID)).One(&orgs)
	return
}

//GetOrgTransactiosBillingByUserID gets org/billings/transactions associated with a UserID
func (r *OrgRepository) GetOrgTransactionBillingsByUserID(userID string) []models.Org {
	var orgs []models.Org
	userid := bson.ObjectIdHex(userID)
	iter := r.C.Find(bson.M{"userid": userid}).Iter()
	result := models.Org{}
	for iter.Next(&result) {
		billings := result.Billings
		for iter.Next(&billings) {
			for _, t := range billings.Transactions {
				orgs = append(orgs, t)
			}
		}
	}
	return orgs
}

//GetOrgBillingsByUserID gets org/billings associated with a UserID
func (r *OrgRepository) GetOrgBillingsByUserID(userID string) []models.Org {
	var orgs []models.Org
	userid := bson.ObjectIdHex(userID)
	iter := r.C.Find(bson.M{"userid": userid}).Iter()
	result := models.Org{}
	for iter.Next(&result) {
		billings := result.Billings
		for _, b := range billings {
			orgs = append(orgs, b)
		}
	}
	return orgs
}

//GetOrgUsersByUserID gets org/users associated with a UserID
func (r *OrgRepository) GetOrgUsersByUserID(userID string) []models.Org {
	var orgs []models.Org
	userid := bson.ObjectIdHex(userID)
	iter := r.C.Find(bson.M{"userid": userid}).Iter()
	result := models.Org{}
	for iter.Next(&result) {
		musers := result.Users
		for _, m := range musers {
			orgs = append(orgs, m)
		}
	}
	return orgs
}

//EditOrgByID edits org associated with an ID
func (r *OrgRepository) EditOrgByID(org *models.Org) error {
	err := r.C.Update(bson.M{"_id": org.ID},
		bson.M{"$set": bson.M{
			"status":    org.Status,
			"updatedat": time.Now(),
		}})
	return err
}

//EditOrgByUserID edits org associated with an UserID
func (r *OrgRepository) EditOrgByUserID(org *models.Org) error {
	err := r.C.Update(bson.M{"userid": org.ID},
		bson.M{"$set": bson.M{
			"financialyearstartdate": org.FinancialYearStartDate,
			"financialyearenddate":   org.FinancialYearEndDate,
			"defaultcurrency":        org.DefaultCurrency,
			"updatedat":              time.Now(),
		}})
	return err
}
