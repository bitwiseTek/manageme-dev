package controllers

/**
 *
 * @author Sika Kay
 * @date 22/11/17
 *
 */
import (
	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
)

type (
	//UserResource For Post/Put/Get - /api/v1/users/{id}
	UserResource struct {
		Data models.User `json:"data"`
	}

	//UsersResource For Post/Put/Get - /api/v1/users/
	UsersResource struct {
		Data []models.User `json:"data"`
	}

	//SignInResource For Post/Put/Get - /api/v1/users/signin
	SignInResource struct {
		Data SignInModel `json:"data"`
	}

	//AuthUserResource Response For Authorized Users - /api/v1/users/signin
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	//OrgUserResource For Post/Put/Get - /api/v1/org/users/{id}
	OrgUserResource struct {
		Data models.OrgUser `json:"data"`
	}

	//OrgUsersResource For Post/Put/Get - /api/v1/org/users/
	OrgUsersResource struct {
		Data []models.OrgUser `json:"data"`
	}

	//OrgUserSignInResource For Post/Put/Get - /api/v1/org/users/signin
	OrgUserSignInResource struct {
		Data OrgUserSignInModel `json:"data"`
	}

	//AuthOrgUserResource Response For Authorized Org Users - /api/v1/org/users/signin
	AuthOrgUserResource struct {
		Data OrgAuthUserModel `json:"data"`
	}

	//OrgResource For Post/Put/Get - /api/v1/orgs/{id}
	OrgResource struct {
		Data models.Org `json:"data"`
	}

	//OrgsResource For Post/Put/Get - /api/v1/orgs/
	OrgsResource struct {
		Data []models.Org `json:"data"`
	}

	//PermissionResource For Post/Put/Get - /api/v1/permissions/{id}
	PermissionResource struct {
		Data models.Permission `json:"data"`
	}

	//PermissionsResource For Post/Put/Get - /api/v1/permissions/
	PermissionsResource struct {
		Data []models.Permission `json:"data"`
	}

	//RoleResource For Post/Put/Get - /api/v1/roles/{id}
	RoleResource struct {
		Data models.Role `json:"data"`
	}

	//RolesResource For Post/Put/Get - /api/v1/roles/
	RolesResource struct {
		Data []models.Roles `json:"data"`
	}

	//EmployeeResource For Post/Put/Get - /api/v1/employees/{id}
	EmployeeResource struct {
		Data models.Employee `json:"data"`
	}

	//EmployeesResource For Post/Put/Get - /api/v1/employees/
	EmployeesResource struct {
		Data []models.Employee `json:"data"`
	}

	//AccountResource For Post/Put/Get - /api/v1/accounts/{id}
	AccountResource struct {
		Data models.Account `json:"data"`
	}

	//AccountsResource For Post/Put/Get - /api/v1/accounts/
	AccountsResource struct {
		Data []models.Account `json:"data"`
	}

	//ExpenseClaimResource For Post/Put/Get - /api/v1/claims/{id}
	ExpenseClaimResource struct {
		Data models.ExpenseClaim `json:"data"`
	}

	//ExpenseClaimsResource For Post/Put/Get - /api/v1/claims/
	ExpenseClaimsResource struct {
		Data []models.ExpenseClaim `json:"data"`
	}

	//LeaveAllocationResource For Post/Put/Get - /api/v1/leave/allocations/{id}
	LeaveAllocationResource struct {
		Data models.LeaveAllocation `json:"data"`
	}

	//LeaveAllocationsResource For Post/Put/Get - /api/v1/leave/allocations/
	LeaveAllocationsResource struct {
		Data []models.LeaveAllocation `json:"data"`
	}

	//LeaveApplicationResource For Post/Put/Get - /api/v1/leave/applications/{id}
	LeaveApplicationResource struct {
		Data models.LeaveApplication `json:"data"`
	}

	//LeaveApplicationsResource For Post/Put/Get - /api/v1/leave/applications/
	LeaveApplicationsResource struct {
		Data []models.LeaveApplication `json:"data"`
	}

	//LeaveBlockListResource For Post/Put/Get - /api/v1/leave/blocklists/{id}
	LeaveBlockListResource struct {
		Data models.LeaveBlockList `json:"data"`
	}

	//LeaveBlockListsResource For Post/Put/Get - /api/v1/leave/blocklists/
	LeaveBlockListsResource struct {
		Data []models.LeaveBlockList `json:"data"`
	}

	//HolidayListResource For Post/Put/Get - /api/v1/holidaylists/{id}
	HolidayListResource struct {
		Data models.HolidayList `json:"data"`
	}

	//HolidayListsResource For Post/Put/Get - /api/v1/holidaylists/
	HolidayListsResource struct {
		Data []models.HolidayList `json:"data"`
	}

	//AppraisalResource For Post/Put/Get - /api/v1/appraisals/{id}
	AppraisalResource struct {
		Data models.Appraisal `json:"data"`
	}

	//AppraisalsResource For Post/Put/Get - /api/v1/appraisals/
	AppraisalsResource struct {
		Data []models.Appraisal `json:"data"`
	}

	//ExitResource For Post/Put/Get - /api/v1/exits/{id}
	ExitResource struct {
		Data models.Exit `json:"data"`
	}

	//ExitsResource For Post/Put/Get - /api/v1/exits/
	ExitsResource struct {
		Data []models.Exit `json:"data"`
	}

	//BiodataResource For Post/Put/Get - /api/v1/biodatas/{id}
	BiodataResource struct {
		Data models.Biodata `json:"data"`
	}

	//BiodatasResource For Post/Put/Get - /api/v1/biodatas/
	BiodatasResource struct {
		Data []models.Biodata `json:"data"`
	}

	//PIDResource For Post/Put/Get - /api/v1/pids/{id}
	PIDResource struct {
		Data models.PersonalIdentification `json:"data"`
	}

	//PIDsResource For Post/Put/Get - /api/v1/pids/
	PIDsResource struct {
		Data []models.PersonalIdentification `json:"data"`
	}

	//HealthDetailResource For Post/Put/Get - /api/v1/health/details/{id}
	HealthDetailResource struct {
		Data models.HealthDetail `json:"data"`
	}

	//HealthDetailsResource For Post/Put/Get - /api/v1/health/details/
	HealthDetailsResource struct {
		Data []models.HealthDetail `json:"data"`
	}

	//EducationResource For Post/Put/Get - /api/v1/educations/{id}
	EducationResource struct {
		Data models.Education `json:"data"`
	}

	//EducationsResource For Post/Put/Get - /api/v1/educations/{
	EducationsResource struct {
		Data []models.Education `json:"data"`
	}

	//WorkExperienceResource For Post/Put/Get - /api/v1/work/experiences/{id}
	WorkExperienceResource struct {
		Data models.WorkExperience `json:"data"`
	}

	//WorkExperiencesResource For Post/Put/Get - /api/v1/work/experiences/
	WorkExperiencesResource struct {
		Data []models.WorkExperience `json:"data"`
	}

	//SalaryEmployeeResource For Post/Put/Get - /api/v1/salary/employees/{id}
	SalaryEmployeeResource struct {
		Data models.SalaryEmployee `json:"data"`
	}

	//SalaryEmployeesResource For Post/Put/Get - /api/v1/salary/employees/
	SalaryEmployeesResource struct {
		Data []models.SalaryEmployee `json:"data"`
	}

	//SalaryComponentResource For Post/Put/Get - /api/v1/salary/components/{id}
	SalaryComponentResource struct {
		Data models.SalaryComponent `json:"data"`
	}

	//SalaryComponentsResource For Post/Put/Get - /api/v1/salary/components/
	SalaryComponentsResource struct {
		Data []models.SalaryComponent `json:"data"`
	}

	//SalaryStructureResource For Post/Put/Get - /api/v1/salary/structures/{id}
	SalaryStructureResource struct {
		Data models.SalaryStructure `json:"data"`
	}

	//SalaryStructuresResource For Post/Put/Get - /api/v1/salary/structures/
	SalaryStructuresResource struct {
		Data []models.SalaryStructure `json:"data"`
	}

	//ActivityTypeResource For Post/Put/Get - /api/v1/activity/types/{id}
	ActivityTypeResource struct {
		Data models.ActivityType `json:"data"`
	}

	//ActivityTypesResource For Post/Put/Get - /api/v1/activity/types/
	ActivityTypesResource struct {
		Data []models.ActivityType `json:"data"`
	}

	//WorkingHourResource For Post/Put/Get - /api/v1/working/hours/{id}
	WorkingHourResource struct {
		Data models.WorkingHour `json:"data"`
	}

	//WorkingHoursResource For Post/Put/Get - /api/v1/working/hours/
	WorkingHoursResource struct {
		Data []models.WorkingHour `json:"data"`
	}

	//OperationResource For Post/Put/Get - /api/v1/operations/{id}
	OperationResource struct {
		Data models.Operation `json:"data"`
	}

	//OperationsResource For Post/Put/Get - /api/v1/operations/
	OperationsResource struct {
		Data []models.Operation `json:"data"`
	}

	//WorkstationResource For Post/Put/Get - /api/v1/workstations/{id}
	WorkstationResource struct {
		Data models.Workstation `json:"data"`
	}

	//WorkstationsResource For Post/Put/Get - /api/v1/workstations/
	WorkstationsResource struct {
		Data []models.Workstation `json:"data"`
	}

	//TimesheetResource For Post/Put/Get - /api/v1/timesheet/{id}
	TimesheetResource struct {
		Data models.Timesheet `json:"data"`
	}

	//TimesheetsResource For Post/Put/Get - /api/v1/timesheet/
	TimesheetsResource struct {
		Data []models.Timesheet `json:"data"`
	}

	//SalarySlipResource For Post/Put/Get - /api/v1/salary/slips/{id}
	SalarySlipResource struct {
		Data models.SalarySlip `json:"data"`
	}

	//SalarySlipsResource For Post/Put/Get - /api/v1/salary/slips/
	SalarySlipsResource struct {
		Data []models.SalarySlip `json:"data"`
	}

	//ActivityCostResource For Post/Put/Get - /api/v1/activity/costs/{id}
	ActivityCostResource struct {
		Data models.ActivityCost `json:"data"`
	}

	//ActivityCostsResource For Post/Put/Get - /api/v1/activity/costs/
	ActivityCostsResource struct {
		Data []models.ActivityCost `json:"data"`
	}

	//BillingResource For Post/Put/Get - /api/v1/billings/{id}
	BillingResource struct {
		Data models.Billing `json:"data"`
	}

	//BillingsResource For Post/Put/Get - /api/v1/billings/
	BillingsResource struct {
		Data []models.Billing `json:"data"`
	}

	//TransactionResource For Post/Put/Get - /api/v1/transactions/{id}
	TransactionResource struct {
		Data models.Transaction `json:"data"`
	}

	//TransactionsResource For Post/Put/Get - /api/v1/transactions/
	TransactionsResource struct {
		Data []models.Transaction `json:"data"`
	}

	//DepartmentResource For Post/Put/Get - /api/v1/departments/{id}
	DepartmentResource struct {
		Data models.Department `json:"data"`
	}

	//DepartmentsResource For Post/Put/Get - /api/v1/departments/
	DepartmentsResource struct {
		Data []models.Department `json:"data"`
	}

	//DesignationResource For Post/Put/Get - /api/v1/designations/{id}
	DesignationResource struct {
		Data models.Designation `json:"data"`
	}

	//DesignationsResource For Post/Put/Get - /api/v1/designations/
	DesignationsResource struct {
		Data []models.Designation `json:"data"`
	}

	//SalaryModeResource For Post/Put/Get - /api/v1/salary/modes/{id}
	SalaryModeResource struct {
		Data models.SalaryMode `json:"data"`
	}

	//SalaryModesResource For Post/Put/Get - /api/v1/salary/modes/
	SalaryModesResource struct {
		Data []models.SalaryMode `json:"data"`
	}

	//BranchResource For Post/Put/Get - /api/v1/branches/{id}
	BranchResource struct {
		Data models.Branch `json:"data"`
	}

	//BranchesResource For Post/Put/Get - /api/v1/branches/
	BranchesResource struct {
		Data []models.Branch `json:"data"`
	}

	//LeaveTypeResource For Post/Put/Get - /api/v1/leave/types/{id}
	LeaveTypeResource struct {
		Data models.LeaveType `json:"data"`
	}

	//LeaveTypesResource For Post/Put/Get - /api/v1/leave/types/
	LeaveTypesResource struct {
		Data []models.LeaveType `json:"data"`
	}

	//ExpenseClaimTypeResource For Post/Put/Get - /api/v1/claim/types/{id}
	ExpenseClaimTypeResource struct {
		Data models.ExpenseClaimType `json:"data"`
	}

	//ExpenseClaimTypesResource For Post/Put/Get - /api/v1/claim/types/
	ExpenseClaimTypesResource struct {
		Data []models.ExpenseClaimType `json:"data"`
	}

	//ProjectTypeResource For Post/Put/Get - /api/v1/project/types/{id}
	ProjectTypeResource struct {
		Data models.ProjectType `json:"data"`
	}

	//ProjectTypesResource For Post/Put/Get - /api/v1/project/types/
	ProjectTypesResource struct {
		Data []models.ProjectType `json:"data"`
	}

	//TaskResource For Post/Put/Get - /api/v1/tasks/{id}
	TaskResource struct {
		Data models.Task `json:"data"`
	}

	//TasksResource For Post/Put/Get - /api/v1/tasks/
	TasksResource struct {
		Data []models.Task `json:"data"`
	}

	//ProjectTaskResource For Post/Put/Get - /api/v1/project/tasks/{id}
	ProjectTaskResource struct {
		Data models.ProjectTask `json:"data"`
	}

	//ProjectTasksResource For Post/Put/Get - /api/v1/project/tasks/
	ProjectTasksResource struct {
		Data []models.ProjectTask `json:"data"`
	}

	//ProjectResource For Post/Put/Get - /api/v1/projects/{id}
	ProjectResource struct {
		Data models.Project `json:"data"`
	}

	//ProjectsResource For Post/Put/Get - /api/v1/projects/
	ProjectsResource struct {
		Data []models.Project `json:"data"`
	}

	//JournalAccountResource For Post/Put/Get - /api/v1/journal/accounts/{id}
	JournalAccountResource struct {
		Data models.JournalAccount `json:"data"`
	}

	//JournalAccountsResource For Post/Put/Get - /api/v1/journal/accounts/
	JournalAccountsResource struct {
		Data []models.JournalAccount `json:"data"`
	}

	//JournalEntryResource For Post/Put/Get - /api/v1/journal/entries/{id}
	JournalEntryResource struct {
		Data models.JournalEntry `json:"data"`
	}

	//JournalEntriesResource For Post/Put/Get - /api/v1/journal/entries/
	JournalEntriesResource struct {
		Data []models.JournalEntry `json:"data"`
	}

	//SignInModel For /users/signin
	SignInModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//AuthUserModel For Authorized User with Access Token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}

	//OrgUserSignInModel For /org/users/signin
	OrgUserSignInModel struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	//AuthOrgUserModel For Authorized Org User with Access Token
	AuthOrgUserModel struct {
		OrgUser models.OrgUser `json:"user"`
		Token 	string         `json:"token"`
	}
)
