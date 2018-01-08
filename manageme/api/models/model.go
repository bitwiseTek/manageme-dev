package models

/**
 *
 * @author Sika Kay
 * @date 22/11/17
 *
 */

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	//User Struct for DB Model
	User struct {
		ID                bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Role              string        `json:"role"`
		Email             string        `json:"email"`
		Password          string        `json:"password,omitempty"`
		HashPassword      []byte        `json:"hashpassword,omitempty"`
		LastLogin         time.Time     `json:"lastlogin"`
		LastLoginLocation string        `json:"lastloginlocation"`
		LastLoginIP       string        `json:"lastloginip"`
		Status            string        `json:"status,omitempty"`
		CreatedAt         time.Time     `json:"createdat,omitempty"`
		UpdatedAt         time.Time     `json:"updatedat,omitempty"`
	}

	// Org/Contact for DB Model
	CompanyContact struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CompanyName string        `json:"companyname"`
		Address     string        `json:"address"`
		State       string        `json:"state"`
		Country     string        `json:"country"`
		Logo        string        `json:"logo"`
		PostalCode  string        `json:"postalcode"`
		Phone       string        `json:"phone"`
		Website     string        `json:"website"`
	}

	// Billing/Transactions for DB Model
	Transactions struct {
		ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
		RefCode       string        `json:"refcode"`
		Type          string        `json:"type"`
		Statement     string        `json:"statement"`
		PaymentRef    string        `json:"paymentref"`
		PaymentMethod string        `json:"paymentmethod"`
		ResponseCode  string        `json:"responsecode"`
		Status        string        `json:"status,omitempty"`
		CreatedAt     time.Time     `json:"createdat,omitempty"`
		UpdatedAt     time.Time     `json:"updatedat,omitempty"`
	}

	// Org/Billing for DB Model
	Billing struct {
		ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name       string        `json:"name"`
		Cost       float32       `json:"cost"`
		Currency   string        `json:"currency"`
		ActiveDays time.Duration `json:"activedays"`
		MaxUsers   int           `json:"maxusers"`
		LastBilled time.Time     `json:"postalcode"`
		Status     string        `json:"status"`
		Transactions
	}

	// Roles/Permissions for DB Model
	Permissions struct {
		ID    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name  string        `json:"name"`
		Level string        `json:"level"`
	}

	// Users/Roles for DB Model
	Roles struct {
		ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name          string        `json:"name"`
		DocumentType  string        `json:"documenttype"`
		DocumentStage string        `json:"documentstage"`
		Status        string        `json:"status,omitempty"`
		CreatedAt     time.Time     `json:"createdat,omitempty"`
		UpdatedAt     time.Time     `json:"updatedat,omitempty"`
		Permissions
	}

	// EmployeeContact/EmergencyContact for DB Model
	EmergencyContact struct {
		ID                bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FullName          string        `json:"firstname"`
		Address           string        `json:"address"`
		Email             string        `json:"email"`
		RelationshipType  string        `json:"relationshiptype"`
		AccommodationType string        `json:"accommodationtype"`
		Phone             string        `json:"phone"`
	}

	// Employees/EmployeeContact for DB Model
	EmployeeContact struct {
		ID                bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName         string        `json:"firstname"`
		LastName          string        `json:"lastname"`
		MiddleName        string        `json:"middlename"`
		AddressPermanent  string        `json:"addresspermanent"`
		AddressCurrent    string        `json:"addresscurrent"`
		EmailPersonal     string        `json:"emailpersonal"`
		EmailCompany      string        `json:"emailcompany"`
		Image             string        `json:"image"`
		AccommodationType string        `json:"accommodationtype"`
		PrimaryPhone      string        `json:"primaryphone"`
		SecondaryPhone    string        `json:"secondaryphone"`
		EmergencyContact
	}

	// Biodata/PersonalIdentification for DB Model
	PersonalIdentification struct {
		ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		IDType     string        `json:"idtype"`
		IDNo       string        `json:"idno"`
		VaidTill   string        `json:"validtill"`
		IssuePlace string        `json:"issueplace"`
		IssueDate  string        `json:"issuedate"`
	}

	// Biodata/HealthDetails for DB Model
	HealthDetails struct {
		ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Height         string        `json:"height"`
		Weight         string        `json:"weight"`
		EyeColor       string        `json:"eyecolor"`
		KnownAllergies string        `json:"knownallergies"`
		HealthConcerns string        `json:"healthconcerns"`
	}

	// Biodata/Education for DB Model
	Education struct {
		ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Qualification  string        `json:"qualification"`
		CGPA           float32       `json:"cgpa"`
		GraduatedOn    string        `json:"graduatedon"`
		SchoolAttended string        `json:"schoolattended"`
		HonorType      string        `json:"honortype"`
	}

	// Biodata/WorkExperience for DB Model
	WorkExperience struct {
		ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CompanyWorked  string        `json:"companyworked"`
		WorkHistoryExt time.Duration `json:"workhistoryext"`
		Designation    string        `json:"designation"`
		ResignedOn     string        `json:"resignedon"`
		Address        string        `json:"address"`
		Salary         string        `json:"salary"`
	}

	// Employees/Biodata for DB Model
	Biodata struct {
		ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		DateOfBirth    string        `json:"dateofbirth"`
		Sex            string        `json:"sex"`
		BloodGroup     string        `json:"bloodgroup"`
		MaritalStatus  string        `json:"maritalstatus"`
		DisabilityType string        `json:"disabilitytype"`
		Nationality    string        `json:"nationality"`
		StateOfOrigin  string        `json:"stateoforigin"`
		PersonalIdentification
		HealthDetails
		Education
		WorkExperience
	}

	// ExpenseClaim/ExpenseDetail for DB Model
	ExpenseDetail struct {
		ID               bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ExpenseDate      string        `json:"expensedate"`
		ExpenseType      string        `json:"expensetype"`
		Description      string        `json:"description"`
		ClaimAmount      float32       `json:"claimamount"`
		SanctionedAmount float32       `json:"sanctionedamount"`
	}

	// Employees/ExpenseClaim for DB Model
	ExpenseClaim struct {
		ID                    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ExpApprover           bson.ObjectId `json:"empid"`
		ProjectID             bson.ObjectId `json:"projectid"`
		TaskID                bson.ObjectId `json:"taskid"`
		PayableAccount        bson.ObjectId `json:"accountid"`
		IsPaid                bool          `json:"ispaid"`
		ApprovalStatus        string        `json:"approvalstatus"`
		TotalClaimedAmount    float32       `json:"totalclaimedamount"`
		TotalSanctionedAmount float32       `json:"totalsanctionedamount"`
		TotalAmountReimbursed float32       `json:"totalamountreimbursed"`
		PostingDate           time.Time     `json:"postingdate"`
		PaymentMode           string        `json:"paymentmode"`
		Status                string        `json:"status,omitempty"`
		Remarks               string        `json:"remarks"`
		ExpenseDetail
	}

	// Employees/LeaveAllocation for DB Model
	LeaveAllocation struct {
		ID                   bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Description          string        `json:"description"`
		LeaveType            string        `json:"leavetype"`
		FromDate             string        `json:"fromdate"`
		ToDate               string        `json:"todate"`
		NewLeavesAllocated   string        `json:"newleavesallocated"`
		CarryForward         string        `json:"carryforward"`
		CarryForwardedLeaves int           `json:"carryforwardedleaves"`
		TotalLeavesAllocated int           `json:"totalleavesallocated"`
		AllocatedOn          time.Time     `json:"allocatedon,omitempty"`
	}

	// Employees/LeaveApplication for DB Model
	LeaveApplication struct {
		ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		LeaveApprover  bson.ObjectId `json:"empid"`
		Description    string        `json:"description"`
		LeaveType      string        `json:"leavetype"`
		FromDate       string        `json:"fromdate"`
		ToDate         string        `json:"todate"`
		LeaveBalance   int           `json:"leavebalance"`
		IsHalfDay      bool          `json:"ishalfday"`
		HalfDayDate    string        `json:"halfdaydate"`
		TotalLeaveDays int           `json:"totalleavedays"`
		Status         string        `json:"status,omitempty"`
		AppliedAt      time.Time     `json:"appliedat,omitempty"`
		UpdatedAt      time.Time     `json:"updatedat,omitempty"`
	}

	// LeaveBlockList/BlockedDates for DB Model
	BlockedDates struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		BlockDate string        `json:"blockdate"`
		Reason    string        `json:"reason"`
	}

	// LeaveBlockList/AllowedUsers for DB Model
	AllowedUsers struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		AllowUser string        `json:"allowuser"`
	}

	// Employees/LeaveBlockList for DB Model
	LeaveBlockList struct {
		ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name     string        `json:"name"`
		ApplyAll bool          `json:"applyall"`
		BlockedDates
		AllowedUsers
	}

	// HolidayList/Holidays for DB Model
	Holidays struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Description string        `json:"description"`
		HolidayDate string        `json:"holidaydate"`
	}

	// Employees/HolidayList for DB Model
	HolidayList struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		FromDate  string        `json:"fromdate"`
		ToDate    string        `json:"todate"`
		WeeklyOff bool          `json:"weeklyoff"`
		Holidays
	}

	// Appraisal/AppraisalTemplate for DB Model
	AppraisalTemplate struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Title       string        `json:"title"`
		Description string        `json:"description"`
	}

	// Appraisal/AppraisalGoalfor DB Model
	AppraisalGoal struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		KRA          string        `json:"kra"`
		PerWeightage float32       `json:"perweightage"`
		Score        float32       `json:"score"`
		ScoreEarned  float32       `json:"scoreearned"`
	}

	// Employees/Appraisal for DB Model
	Appraisal struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ForEmployee bson.ObjectId `json:"empid"`
		StartDate   string        `json:"startdate"`
		EndDate     string        `json:"enddate"`
		Remarks     string        `json:"remarks"`
		TotalScore  float32       `json:"totalscore"`
		Status      string        `json:"status,omitempty"`
		AppraisedAt time.Time     `json:"appraisedat,omitempty"`
		AppraisalTemplate
		AppraisalGoal
	}

	// Employees/Exit for DB Model
	Exit struct {
		ID                    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ResignationLetterDate string        `json:"resignationletterdate"`
		LeavingReason         string        `json:"leavingreason"`
		RelievingDate         string        `json:"relievingdate"`
		LeaveEncashed         float32       `json:"leaveencashed"`
		EncashmentDate        string        `json:"encashmentdate"`
		InterviewDate         string        `json:"interviewdate"`
		ResignationReason     string        `json:"resignationreason"`
		NewWorkPlace          string        `json:"newworkplace"`
		Feedback              string        `json:"feedback"`
		ExitedAt              time.Time     `json:"exitedat,omitempty"`
	}

	// SalaryStructure/SalaryEmployees for DB Model
	SalaryEmployees struct {
		ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
		EmployeeNo string        `json:"employeeno"`
		FromDate   bool          `json:"fromdate"`
		ToDate     bool          `json:"todate"`
		Base       float32       `json:"base"`
		Varaible   float32       `json:"variable"`
	}

	// SalaryDetail/SalaryComponent for DB Model
	SalaryComponent struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name        string        `json:"name"`
		Type        string        `json:"type"`
		Description string        `json:"description"`
	}

	// SalaryStructure/SalaryDetail for DB Model
	SalaryDetail struct {
		ID                   bson.ObjectId `bson:"_id,omitempty" json:"id"`
		IsAmountFormulaBased bool          `json:"isamountformulabased"`
		IsAmountLwpBased     bool          `json:"asamountlwpbased"`
		Amount               float32       `json:"amount"`
		DefaulAmount         float32       `json:"defaultamount"`
		SalaryComponent
	}

	// Employees/SalaryStructure for DB Model
	SalaryStructure struct {
		ID               bson.ObjectId `bson:"_id,omitempty" json:"id"`
		PayableAccount   bson.ObjectId `json:"accountid"`
		PayrollFrequency string        `json:"payrollfrequency"`
		IsActive         bool          `json:"isactive"`
		IsDefault        bool          `json:"isdefault"`
		HourRate         float32       `json:"hourrate"`
		PaymentMode      string        `json:"paymentmode"`
		Deductions       float32       `json:"deductions"`
		Earnings         float32       `json:"earnings"`
		TotalDeductions  float32       `json:"totaldeductions"`
		TotalEarnings    float32       `json:"totalearnings"`
		NetPay           float32       `json:"netpay"`
		SalaryEmployees
		SalaryDetail
	}

	// TimesheetDetail/ActivityType for DB Model
	ActivityType struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name        string        `json:"name"`
		BillingRate float32       `json:"billingrate"`
		CostingRate float32       `json:"costingrate"`
		Disabled    bool          `json:"disabled"`
	}

	// Workstation/WorkingHours for DB Model
	WorkingHours struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		StartTime string        `json:"starttime"`
		EndTime   string        `json:"endtime"`
		Enabled   bool          `json:"enabled"`
	}

	// Workstation/Operation for DB Model
	Operation struct {
		ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
		WorkstationID bson.ObjectId `json:"workstationid"`
		Description   string        `json:"description"`
	}

	// TimesheetDetail/Workstation for DB Model
	Workstation struct {
		ID                  bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name                string        `json:"name"`
		Description         string        `json:"description"`
		HourRate            float32       `json:"hourrate"`
		HourRateElectricity float32       `json:"hourrateelectricity"`
		HourRateRent        float32       `json:"hourraterent"`
		HourRateLabor       float32       `json:"hourratelabor"`
		HourRateConsumable  float32       `json:"hourrateconsumable"`
		WorkingHours
		Operation
	}

	// Timesheet/TimesheetDetail for DB Model
	TimesheetDetail struct {
		ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ProjectID     bson.ObjectId `json:"projectid"`
		TaskID        bson.ObjectId `json:"taskid"`
		FromTime      string        `json:"fromtime"`
		ToTime        string        `json:"totime"`
		CompletedQty  int           `json:"completedqty"`
		Hours         float32       `json:"hours"`
		Billable      float32       `json:"billable"`
		BillingHours  float32       `json:"billinghours"`
		BillingAmount float32       `json:"billingamount"`
		CostingAmount float32       `json:"costingamount"`
		ActivityType
		Workstation
	}

	// SalarySlip/Timesheet for DB Model
	Timesheet struct {
		ID                  bson.ObjectId `bson:"_id,omitempty" json:"id"`
		EmployeeID          bson.ObjectId `json:"empid"`
		StartDate           string        `json:"startdate"`
		EndDate             string        `json:"enddate"`
		TotalWorkingDays    int           `json:"totalworkingdays"`
		TotalHours          float32       `json:"totalhours"`
		TotalBillableHours  float32       `json:"totalbillablehours"`
		TotalCostingAmount  float32       `json:"totalcostingamount"`
		TotalBillableAmount float32       `json:"totalbillableamount"`
		PercentageBilled    float32       `json:"percentagebilled"`
		TotalBilledAmount   float32       `json:"totalbilledamount"`
		Note                string        `json:"note"`
		TimesheetDetail
	}

	// Employees/SalarySlip for DB Model
	SalarySlip struct {
		ID                    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		DepartmentID          bson.ObjectId `json:"departmentid"`
		BranchID              bson.ObjectId `json:"branchid"`
		EmployeeID            bson.ObjectId `json:"empid"`
		IsSalarySlipTimesheet bool          `json:"issalarysliptimesheet"`
		StartDate             string        `json:"startdate"`
		EndDate               string        `json:"enddate"`
		LeaveWithoutPay       string        `json:"leavewithoutpay"`
		PaymentDays           string        `json:"paymentdays"`
		Designation           string        `json:"designation"`
		GrossPay              float32       `json:"grosspay"`
		InterestAmount        float32       `json:"interestamount"`
		RoundedTotal          int           `json:"roundedtotal"`
		PostingDate           time.Time     `json:"bankaccount"`
		Timesheet
	}

	// Employees/SalarySlipTimesheet for DB Model
	SalarySlipTimesheet struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TimesheetID  bson.ObjectId `json:"timesheetid"`
		WorkingHours float32       `json:"workinghours"`
	}

	// Employees/ActivityCost for DB Model
	ActivityCost struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		EmployeeID   bson.ObjectId `json:"empid"`
		ActivityType string        `json:"activitytype"`
		BillingRate  float32       `json:"billingrate"`
		CostingRate  float32       `json:"costingrate"`
		Status       string        `json:"status,omitempty"`
	}

	// Task/TaskDependent for DB Model
	TaskDependent struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ProjectID bson.ObjectId `json:"projectid"`
		TaskID    bson.ObjectId `json:"taskid"`
		Subject   string        `json:"subject"`
	}

	// ProjectTasks/Task for DB Model
	Task struct {
		ID                 bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ProjectID          bson.ObjectId `json:"projectid"`
		Name               string        `json:"name"`
		Subject            string        `json:"subject"`
		IsMilestone        bool          `json:"ismilestone"`
		PercentProgress    float32       `json:"percentprogress"`
		Priority           string        `json:"priority"`
		ExpectedStartDate  string        `json:"expectedstartdate"`
		ExpectedEndDate    string        `json:"expectedenddate"`
		ActualStartDate    string        `json:"actualstartdate"`
		ActualEndDate      string        `json:"actualenddate"`
		TaskWeight         float32       `json:"taskweight"`
		ReviewDate         string        `json:"reviewdate"`
		Note               string        `json:"note"`
		EstimatedCosting   float32       `json:"estimatedcosting"`
		TotalCostingAmount float32       `json:"totalcostingamount"`
		TotalExpenseClaim  float32       `json:"totalexpenseclaim"`
		TotalBillingAmount float32       `json:"totalbillingamount"`
		TotalPurchaseCost  float32       `json:"totalpurchasecost"`
		TotalSalesCost     float32       `json:"totalsalescost"`
		ClosingDate        string        `json:"closingdate"`
		PercentComplete    float32       `json:"percentcomplete"`
		ActualStartTime    time.Time     `json:"actualstarttime"`
		ExpectedStartTime  time.Time     `json:"expectedstarttime"`
		Status             string        `json:"status,omitempty"`
		TaskDependent
	}

	// Project/ProjectTasks for DB Model
	ProjectTasks struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Description string        `json:"description"`
		StartDate   string        `json:"startdate"`
		EndDate     string        `json:"enddate"`
		Title       string        `json:"title"`
		Status      string        `json:"status,omitempty"`
		Task
	}

	// Employees/Project for DB Model
	Project struct {
		ID                 bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ProjectTypeID      bson.ObjectId `json:"projecttypeid"`
		Name               string        `json:"name"`
		IsActive           bool          `json:"isactive"`
		PercentComplete    float32       `json:"percentcomplete"`
		Priority           string        `json:"priority"`
		ExpectedStartDate  string        `json:"expectedstartdate"`
		ExpectedEndDate    string        `json:"expectedenddate"`
		ActualStartDate    string        `json:"actualstartdate"`
		ActualEndDate      string        `json:"actualenddate"`
		Customer           string        `json:"customer"`
		SalesOrder         string        `json:"salesorder"`
		Note               string        `json:"note"`
		EstimatedCosting   float32       `json:"estimatedcosting"`
		TotalCostingAmount float32       `json:"totalcostingamount"`
		TotalExpenseClaim  float32       `json:"totalexpenseclaim"`
		TotalBillingAmount float32       `json:"totalbillingamount"`
		TotalPurchaseCost  float32       `json:"totalpurchasecost"`
		TotalSalesCost     float32       `json:"totalsalescost"`
		GrossMargin        float32       `json:"grossmargin"`
		PercentGrossMargin float32       `json:"percentgrossmargin"`
		ActualStartTime    time.Time     `json:"actualstarttime"`
		Status             string        `json:"status,omitempty"`
		ProjectTasks
	}

	// JournalEntry/JournalAccount for DB Model
	JournalAccount struct {
		ID                      bson.ObjectId `bson:"_id,omitempty" json:"id"`
		AccountID               bson.ObjectId `json:"accountid"`
		ProjectID               bson.ObjectId `json:"projectid"`
		AccountType             string        `json:"accounttype"`
		Balance                 float32       `json:"balance"`
		AccountCurrency         string        `json:"accountcurrency"`
		DebitInAccountCurrency  string        `json:"debitinaccountcurrency"`
		CreditInAccountCurrency string        `json:"creditinaccountcurrency"`
		ReferenceType           string        `json:"referencetype"`
		ReferenceName           string        `json:"referencename"`
		WriteOffBasedOn         string        `json:"writeoffbasedon"`
		TotalAmountCurrency     string        `json:"totalamountcurrency"`
		Debit                   float32       `json:"debit"`
		Credit                  float32       `json:"credit"`
		ExchangeRate            float32       `json:"exchangerate"`
		IsAdvance               bool          `json:"isadvance"`
	}

	// Employees/JournalEntry for DB Model
	JournalEntry struct {
		ID                  bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Title               string        `json:"title"`
		VoucherType         float32       `json:"vouchertype"`
		ChequeNo            string        `json:"chequeno"`
		ChequeDate          string        `json:"chequedate"`
		BillNo              string        `json:"billno"`
		BillDate            string        `json:"billdate"`
		DueDate             string        `json:"duedate"`
		Remarks             string        `json:"remarks"`
		WriteOffBasedOn     string        `json:"writeoffbasedon"`
		TotalAmountCurrency string        `json:"totalamountcurrency"`
		TotalDebit          float32       `json:"totaldebit"`
		TotalCredit         float32       `json:"totalcredit"`
		Difference          float32       `json:"difference"`
		TotalAmount         float32       `json:"totalamount"`
		WriteOffAmount      float32       `json:"writeoffamount"`
		PostingDate         time.Time     `json:"postingdate"`
		JournalAccount
	}

	// Users/Employees for DB Model
	Employees struct {
		ID                        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		DepartmentID              bson.ObjectId `json:"departmentid"`
		BranchID                  bson.ObjectId `json:"branchid"`
		ReportsTo                 bson.ObjectId `json:"empid"`
		EmployeeNo                string        `json:"employeeno"`
		JoiningDate               string        `json:"joiningdate"`
		EmploymentType            string        `json:"employmenttype"`
		ScheduledConfirmationDate string        `json:"scheduledconfirmationdate"`
		FinalConfirmationDate     string        `json:"finalconfirmationdate"`
		ContractEndDate           string        `json:"contractenddate"`
		RetirementDate            time.Time     `json:"retirementdate"`
		Status                    string        `json:"status,omitempty"`
		BankName                  string        `json:"bankname"`
		BankAccount               int           `json:"bankaccount"`
		EmployeeContact
		Biodata
		ExpenseClaim
		LeaveAllocation
		LeaveApplication
		LeaveBlockList
		HolidayList
		Appraisal
		Exit
		SalaryStructure
		SalarySlip
		SalarySlipTimesheet
		ActivityCost
		Project
		JournalEntry
	}

	// Org/Users for DB Model
	Users struct {
		ID                bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Username          string        `json:"username"`
		TempPassword      string        `json:"temppassword,omitempty"`
		HashPassword      []byte        `json:"hashpassword,omitempty"`
		LastLogin         time.Time     `json:"lastlogin"`
		LastLoginLocation string        `json:"lastloginlocation"`
		LastLoginIP       string        `json:"lastloginip"`
		UserType          string        `json:"usertype"`
		AllowedModules    string        `json:"allowedmodules"`
		SecurityQues      string        `json:"securityques"`
		SecurityAns       string        `json:"securityans"`
		Status            string        `json:"status,omitempty"`
		CreatedAt         time.Time     `json:"createdat,omitempty"`
		UpdatedAt         time.Time     `json:"updatedat,omitempty"`
		Roles
		Employees
	}

	// Org/Departments for DB Model
	Departments struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedAt   time.Time     `json:"createdat"`
		UpdatedAt   time.Time     `json:"updatedat"`
		Status      string        `json:"status,omitempty"`
	}

	// Org/Designations for DB Model
	Designations struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// Org/SalaryModes for DB Model
	SalaryModes struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// Org/Branches for DB Model
	Branches struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// Org/LeaveType for DB Model
	LeaveType struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// Org/ExpenseClaimType for DB Model
	ExpenseClaimType struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// Org/ProjectType for DB Model
	ProjectType struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// Org/Account for DB Model
	Account struct {
		ID              bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ParentAccount   bson.ObjectId `json:"accountid"`
		Name            string        `json:"name"`
		RootType        string        `json:"roottype"`
		ReportType      string        `json:"reporttype"`
		AccountCurrency string        `json:"accountcurrency"`
		Type            string        `json:"type"`
		TaxRate         float32       `json:"taxrate"`
		BalanceType     string        `json:"balancetype"`
		IsAccountFreeze bool          `json:"isaccountfreeze"`
		IsGroup         bool          `json:"isgroup"`
		CreatedAt       time.Time     `json:"createdat"`
		UpdatedAt       time.Time     `json:"updatedat"`
		Status          string        `json:"status,omitempty"`
	}

	// Org Struct for DB Model
	Org struct {
		ID                     bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserID                 bson.ObjectId `json:"userid"`
		Name                   string        `json:"name"`
		BusinessNature         string        `json:"businessnature"`
		Modules                string        `json:"modules"`
		FinancialYearStartDate string        `json:"financialyearstartdate"`
		FinancialYearEndDate   string        `json:"financialyearenddate"`
		RegistrationNo         string        `json:"registrationno"`
		DefaultCurrency        string        `json:"defaultcurrency"`
		Status                 string        `json:"status,omitempty"`
		CreatedAt              time.Time     `json:"createdat,omitempty"`
		UpdatedAt              time.Time     `json:"updatedat,omitempty"`
		CompanyContact
		Billing
		Users []Users
		Departments
		Designations
		SalaryModes
		Branches
		LeaveType
		ExpenseClaimType
		ProjectType
		Account
	}
)
