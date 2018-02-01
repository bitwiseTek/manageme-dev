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

	// Org Struct for DB Model
	Org struct {
		ID                     bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserID                 bson.ObjectId `json:"userid"`
		Name                   string        `json:"name"`
		BusinessNature         string        `json:"businessnature"`
		Modules                []string      `json:"modules"`
		StaffStrength          string        `json:"staffstrength"`
		FinancialYearStartDate string        `json:"financialyearstartdate"`
		FinancialYearEndDate   string        `json:"financialyearenddate"`
		RegistrationNo         string        `json:"registrationno"`
		DefaultCurrency        string        `json:"defaultcurrency"`
		CompanyName            string        `json:"companyname"`
		Address                string        `json:"address"`
		State                  string        `json:"state"`
		Country                string        `json:"country"`
		Logo                   string        `json:"logo"`
		PostalCode             string        `json:"postalcode"`
		Phone                  string        `json:"phone"`
		Website                string        `json:"website"`
		Status                 string        `json:"status,omitempty"`
		CreatedAt              time.Time     `json:"createdat,omitempty"`
		UpdatedAt              time.Time     `json:"updatedat,omitempty"`
	}

	// Permission Struct for DB Model
	Permission struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID     bson.ObjectId `json:"orgid"`
		Name      string        `json:"name"`
		Level     string        `json:"level"`
		Status    string        `json:"status,omitempty"`
		CreatedAt time.Time     `json:"createdat,omitempty"`
		UpdatedAt time.Time     `json:"updatedat,omitempty"`
	}

	// Role Struct for DB Model
	Role struct {
		ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID         bson.ObjectId `json:"orgid"`
		Permissions   []Permission  `json:"permissions"`
		Name          string        `json:"name"`
		DocumentType  string        `json:"documenttype"`
		DocumentStage string        `json:"documentstage"`
		Status        string        `json:"status,omitempty"`
		CreatedAt     time.Time     `json:"createdat,omitempty"`
		UpdatedAt     time.Time     `json:"updatedat,omitempty"`
	}

	// OrgUser Struct for DB Model
	OrgUser struct {
		ID                bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID             bson.ObjectId `json:"orgid"`
		RoleID            bson.ObjectId `json:"roleid"`
		Username          string        `json:"username"`
		TempPassword      string        `json:"temppassword,omitempty"`
		HashPassword      []byte        `json:"hashpassword,omitempty"`
		LastLogin         time.Time     `json:"lastlogin"`
		LastLoginLocation string        `json:"lastloginlocation"`
		LastLoginIP       string        `json:"lastloginip"`
		UserType          string        `json:"usertype"`
		AllowedModules    []string      `json:"allowedmodules"`
		SecurityQues      string        `json:"securityques"`
		SecurityAns       string        `json:"securityans"`
		Status            string        `json:"status,omitempty"`
		CreatedAt         time.Time     `json:"createdat,omitempty"`
		UpdatedAt         time.Time     `json:"updatedat,omitempty"`
	}

	// Employee Struct for DB Model
	Employee struct {
		ID                        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		DepartmentID              bson.ObjectId `json:"departmentid"`
		BranchID                  bson.ObjectId `json:"branchid"`
		OrgUserID                 bson.ObjectId `json:"orguserid"`
		ReportsTo                 bson.ObjectId `json:"mgrid,omitempty"`
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
		FirstName                 string        `json:"firstname"`
		LastName                  string        `json:"lastname"`
		MiddleName                string        `json:"middlename"`
		AddressPermanent          string        `json:"addresspermanent"`
		AddressCurrent            string        `json:"addresscurrent"`
		EmailPersonal             string        `json:"emailpersonal"`
		EmailCompany              string        `json:"emailcompany"`
		Image                     string        `json:"image"`
		AccommodationType         string        `json:"accommodationtype"`
		PrimaryPhone              string        `json:"primaryphone"`
		SecondaryPhone            string        `json:"secondaryphone"`
		EmFullName                string        `json:"emfirstname"`
		EmAddress                 string        `json:"emaddress"`
		EmEmail                   string        `json:"ememail"`
		EmRelationshipType        string        `json:"emrelationshiptype"`
		EmAccommodationType       string        `json:"emaccommodationtype"`
		EmPhone                   string        `json:"emphone"`
		CreatedAt                 time.Time     `json:"createdat"`
		UpdatedAt                 time.Time     `json:"updatedat"`
	}

	// Account Struct for DB Model
	Account struct {
		ID              bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ParentAccount   bson.ObjectId `json:"accountid,omitempty"`
		OrgID           bson.ObjectId `json:"orgid"`
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

	// ExpenseClaim Struct for DB Model
	ExpenseClaim struct {
		ID                    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID                 bson.ObjectId `json:"orgid"`
		ExpenseType           bson.ObjectId `json:"exptypeid"`
		ExpApprover           bson.ObjectId `json:"mgrid"`
		ExpApplier            bson.ObjectId `json:"empid"`
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
		ExpenseDate           string        `json:"expensedate"`
		Description           string        `json:"description"`
		ClaimAmount           float32       `json:"claimamount"`
		SanctionedAmount      float32       `json:"sanctionedamount"`
		Remarks               string        `json:"remarks"`
		CreatedAt             time.Time     `json:"createdat"`
		UpdatedAt             time.Time     `json:"updatedat"`
	}

	// LeaveAllocation Struct for DB Model
	LeaveAllocation struct {
		ID                   bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID                bson.ObjectId `json:"orgid"`
		LeaveType            bson.ObjectId `json:"leavetypeid"`
		LeaveAllocator       bson.ObjectId `json:"mgrid"`
		LeaveReceiver        bson.ObjectId `json:"empid"`
		Description          string        `json:"description"`
		FromDate             string        `json:"fromdate"`
		ToDate               string        `json:"todate"`
		NewLeavesAllocated   string        `json:"newleavesallocated"`
		CarryForward         string        `json:"carryforward"`
		CarryForwardedLeaves int           `json:"carryforwardedleaves"`
		TotalLeavesAllocated int           `json:"totalleavesallocated"`
		AllocatedOn          time.Time     `json:"allocatedon,omitempty"`
		CreatedAt            time.Time     `json:"createdat"`
		UpdatedAt            time.Time     `json:"updatedat"`
	}

	// LeaveApplication Struct for DB Model
	LeaveApplication struct {
		ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID          bson.ObjectId `json:"orgid"`
		LeaveType      bson.ObjectId `json:"leavetypeid"`
		LeaveApprover  bson.ObjectId `json:"mgrid"`
		LeaveApplier   bson.ObjectId `json:"empid"`
		Description    string        `json:"description"`
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

	// LeaveBlockList Struct for DB Model
	LeaveBlockList struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID     bson.ObjectId `json:"orgid"`
		Name      string        `json:"name"`
		BlockDate string        `json:"blockdate"`
		Reason    string        `json:"reason"`
		ApplyAll  bool          `json:"applyall"`
		AllowUser string        `json:"allowuser"`
	}

	// HolidayList Struct for DB Model
	HolidayList struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID       bson.ObjectId `json:"orgid"`
		Name        string        `json:"name"`
		FromDate    string        `json:"fromdate"`
		ToDate      string        `json:"todate"`
		WeeklyOff   bool          `json:"weeklyoff"`
		Description string        `json:"description"`
		HolidayDate string        `json:"holidaydate"`
	}

	// Appraisal Struct for DB Model
	Appraisal struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID        bson.ObjectId `json:"orgid"`
		ForEmployee  bson.ObjectId `json:"empid"`
		ApGenerator  bson.ObjectId `json:"mgrid"`
		StartDate    string        `json:"startdate"`
		EndDate      string        `json:"enddate"`
		Remarks      string        `json:"remarks"`
		TotalScore   float32       `json:"totalscore"`
		Title        string        `json:"title"`
		Description  string        `json:"description"`
		KRA          string        `json:"kra"`
		PerWeightage float32       `json:"perweightage"`
		Score        float32       `json:"score"`
		ScoreEarned  float32       `json:"scoreearned"`
		Status       string        `json:"status,omitempty"`
		AppraisedAt  time.Time     `json:"appraisedat,omitempty"`
	}

	// Exit Struct for DB Model
	Exit struct {
		ID                    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID                 bson.ObjectId `json:"orgid"`
		ForEmployee           bson.ObjectId `json:"empid"`
		OverseenBy            bson.ObjectId `json:"mgrid"`
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

	// Biodata Struct for DB Model
	Biodata struct {
		ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		EmployeeID     bson.ObjectId `json:"empid"`
		DateOfBirth    string        `json:"dateofbirth"`
		Sex            string        `json:"sex"`
		BloodGroup     string        `json:"bloodgroup"`
		MaritalStatus  string        `json:"maritalstatus"`
		DisabilityType string        `json:"disabilitytype"`
		Nationality    string        `json:"nationality"`
		StateOfOrigin  string        `json:"stateoforigin"`
	}

	// PersonalIdentification Struct for DB Model
	PersonalIdentification struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		BioID       bson.ObjectId `json:"bioid"`
		Attachments []string      `json:"attachments"`
		IDType      string        `json:"idtype"`
		IDNo        string        `json:"idno"`
		ValidTill   string        `json:"validtill"`
		IssuePlace  string        `json:"issueplace"`
		IssueDate   string        `json:"issuedate"`
	}

	// HealthDetail Struct for DB Model
	HealthDetail struct {
		ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		BioID          bson.ObjectId `json:"bioid"`
		Attachments    []string      `json:"attachments"`
		Height         string        `json:"height"`
		Weight         string        `json:"weight"`
		EyeColor       string        `json:"eyecolor"`
		KnownAllergies string        `json:"knownallergies"`
		HealthConcerns string        `json:"healthconcerns"`
	}

	// Education Struct for DB Model
	Education struct {
		ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		BioID          bson.ObjectId `json:"bioid"`
		Attachments    []string      `json:"attachments"`
		Qualification  string        `json:"qualification"`
		CGPA           float32       `json:"cgpa"`
		GraduatedOn    string        `json:"graduatedon"`
		SchoolAttended string        `json:"schoolattended"`
		HonorType      string        `json:"honortype"`
	}

	// WorkExperience Struct for DB Model
	WorkExperience struct {
		ID             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		BioID          bson.ObjectId `json:"bioid"`
		Attachments    []string      `json:"attachments"`
		CompanyWorked  string        `json:"companyworked"`
		WorkHistoryExt time.Duration `json:"workhistoryext"`
		Designation    string        `json:"designation"`
		ResignedOn     string        `json:"resignedon"`
		Address        string        `json:"address"`
		Salary         string        `json:"salary"`
	}

	// SalaryEmployee Struct for DB Model
	SalaryEmployee struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID     bson.ObjectId `json:"orgid"`
		Employees []Employee    `json:"employees"`
		FromDate  bool          `json:"fromdate"`
		ToDate    bool          `json:"todate"`
		Base      float32       `json:"base"`
		Variable  float32       `json:"variable"`
	}

	// SalaryComponent Struct for DB Model
	SalaryComponent struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID       bson.ObjectId `json:"orgid"`
		Name        string        `json:"name"`
		Type        string        `json:"type"`
		Description string        `json:"description"`
	}

	// SalaryStructure Struct for DB Model
	SalaryStructure struct {
		ID                   bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID                bson.ObjectId `json:"orgid"`
		PayableAccount       bson.ObjectId `json:"accountid"`
		GeneratedBy          bson.ObjectId `json:"mgrid"`
		SalaryEmployeeID     bson.ObjectId `json:"salempid"`
		SalaryComponentID    bson.ObjectId `json:"salcompid"`
		PayrollFrequency     string        `json:"payrollfrequency"`
		IsActive             bool          `json:"isactive"`
		IsDefault            bool          `json:"isdefault"`
		HourRate             float32       `json:"hourrate"`
		PaymentMode          string        `json:"paymentmode"`
		Deductions           float32       `json:"deductions"`
		Earnings             float32       `json:"earnings"`
		TotalDeductions      float32       `json:"totaldeductions"`
		TotalEarnings        float32       `json:"totalearnings"`
		NetPay               float32       `json:"netpay"`
		IsAmountFormulaBased bool          `json:"isamountformulabased"`
		IsAmountLwpBased     bool          `json:"isamountlwpbased"`
		Amount               float32       `json:"amount"`
		DefaultAmount        float32       `json:"defaultamount"`
		CreatedAt            time.Time     `json:"createdat"`
		UpdatedAt            time.Time     `json:"updatedat"`
	}

	// ActivityType Struct for DB Model
	ActivityType struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID       bson.ObjectId `json:"orgid"`
		Name        string        `json:"name"`
		BillingRate float32       `json:"billingrate"`
		CostingRate float32       `json:"costingrate"`
		Disabled    bool          `json:"disabled"`
		CreatedAt   time.Time     `json:"createdat"`
		UpdatedAt   time.Time     `json:"updatedat"`
	}

	// WorkingHour Struct for DB Model
	WorkingHour struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID     bson.ObjectId `json:"orgid"`
		StartTime string        `json:"starttime"`
		EndTime   string        `json:"endtime"`
		Enabled   bool          `json:"enabled"`
	}

	// Operation Struct for DB Model
	Operation struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID       bson.ObjectId `json:"orgid"`
		Description string        `json:"description"`
	}

	// Workstation Struct for DB Model
	Workstation struct {
		ID                  bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OperationID         bson.ObjectId `json:"workstationid"`
		WorkingHourID       bson.ObjectId `json:"workhourid"`
		OrgID               bson.ObjectId `json:"orgid"`
		Name                string        `json:"name"`
		Description         string        `json:"description"`
		HourRate            float32       `json:"hourrate"`
		HourRateElectricity float32       `json:"hourrateelectricity"`
		HourRateRent        float32       `json:"hourraterent"`
		HourRateLabor       float32       `json:"hourratelabor"`
		HourRateConsumable  float32       `json:"hourrateconsumable"`
	}

	// Timesheet Struct for DB Model
	Timesheet struct {
		ID                  bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID               bson.ObjectId `json:"orgid"`
		ForEmployee         bson.ObjectId `json:"empid"`
		ApprovedBy          bson.ObjectId `json:"mgrid"`
		ProjectID           bson.ObjectId `json:"projectid,omitempty"`
		ActivityTypeID      bson.ObjectId `json:"activitytypeid"`
		WorkStationID       bson.ObjectId `json:"workstationid"`
		StartDate           string        `json:"startdate"`
		EndDate             string        `json:"enddate"`
		FromTime            string        `json:"fromtime"`
		ToTime              string        `json:"totime"`
		CompletedQty        int           `json:"completedqty"`
		Hours               float32       `json:"hours"`
		Billable            float32       `json:"billable"`
		BillingHours        float32       `json:"billinghours"`
		BillingAmount       float32       `json:"billingamount"`
		CostingAmount       float32       `json:"costingamount"`
		TotalWorkingDays    int           `json:"totalworkingdays"`
		TotalHours          float32       `json:"totalhours"`
		TotalBillableHours  float32       `json:"totalbillablehours"`
		TotalCostingAmount  float32       `json:"totalcostingamount"`
		TotalBillableAmount float32       `json:"totalbillableamount"`
		PercentageBilled    float32       `json:"percentagebilled"`
		TotalBilledAmount   float32       `json:"totalbilledamount"`
		Note                string        `json:"note"`
		CreatedAt           time.Time     `json:"createdat"`
		UpdatedAt           time.Time     `json:"updatedat"`
	}

	// SalarySlip Struct for DB Model
	SalarySlip struct {
		ID                    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID                 bson.ObjectId `json:"orgid"`
		DepartmentID          bson.ObjectId `json:"departmentid"`
		BranchID              bson.ObjectId `json:"branchid"`
		ForEmployee           bson.ObjectId `json:"empid"`
		ApprovedBy            bson.ObjectId `json:"mgrid"`
		TimesheetID           bson.ObjectId `json:"timesheetid"`
		IsSalarySlipTimesheet bool          `json:"issalarysliptimesheet"`
		WorkingHours          float32       `json:"workinghours"`
		StartDate             string        `json:"startdate"`
		EndDate               string        `json:"enddate"`
		LeaveWithoutPay       string        `json:"leavewithoutpay"`
		PaymentDays           string        `json:"paymentdays"`
		Designation           string        `json:"designation"`
		GrossPay              float32       `json:"grosspay"`
		InterestAmount        float32       `json:"interestamount"`
		RoundedTotal          int           `json:"roundedtotal"`
		PostingDate           time.Time     `json:"bankaccount"`
		CreatedAt             time.Time     `json:"createdat"`
		UpdatedAt             time.Time     `json:"updatedat"`
	}

	// ActivityCost Struct for DB Model
	ActivityCost struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID        bson.ObjectId `json:"orgid"`
		EmployeeID   bson.ObjectId `json:"empid"`
		ActivityType bson.ObjectId `json:"activitytypeid"`
		BillingRate  float32       `json:"billingrate"`
		CostingRate  float32       `json:"costingrate"`
		CreatedAt    time.Time     `json:"createdat"`
		UpdatedAt    time.Time     `json:"updatedat"`
		Status       string        `json:"status,omitempty"`
	}

	// Billing Struct for DB Model
	Billing struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID       bson.ObjectId `json:"orgid"`
		Name        string        `json:"name"`
		Cost        float32       `json:"cost"`
		Currency    string        `json:"currency"`
		BillingFreq string        `json:"billingfreq"`
		ActiveDays  time.Duration `json:"activedays"`
		MaxUsers    int           `json:"maxusers"`
		LastBilled  time.Time     `json:"postalcode"`
		UpdatedAt   time.Time     `json:"updatedat"`
		Status      string        `json:"status"`
	}

	// Transaction Struct for DB Model
	Transaction struct {
		ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
		BillingID     bson.ObjectId `json:"billingid"`
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

	// Department Struct for DB Model
	Department struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID       bson.ObjectId `json:"orgid"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedAt   time.Time     `json:"createdat"`
		UpdatedAt   time.Time     `json:"updatedat"`
		Status      string        `json:"status,omitempty"`
	}

	// Designation Struct for DB Model
	Designation struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID     bson.ObjectId `json:"orgid"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// SalaryMode Struct for DB Model
	SalaryMode struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID     bson.ObjectId `json:"orgid"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// Branch Struct for DB Model
	Branch struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID     bson.ObjectId `json:"orgid"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// LeaveType Struct for DB Model
	LeaveType struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID     bson.ObjectId `json:"orgid"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// ExpenseClaimType Struct for DB Model
	ExpenseClaimType struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID     bson.ObjectId `json:"orgid"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// ProjectType Struct for DB Model
	ProjectType struct {
		ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID     bson.ObjectId `json:"orgid"`
		Name      string        `json:"name"`
		CreatedAt time.Time     `json:"createdat"`
		UpdatedAt time.Time     `json:"updatedat"`
		Status    string        `json:"status,omitempty"`
	}

	// Task Struct for DB Model
	Task struct {
		ID                 bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID              bson.ObjectId `json:"orgid"`
		TaskDependent      bson.ObjectId `json:"taskid,omitempty"`
		InitiatedBy        bson.ObjectId `json:"empid"`
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
		CreatedAt          time.Time     `json:"createdat"`
		UpdatedAt          time.Time     `json:"updatedat"`
		Status             string        `json:"status,omitempty"`
	}

	// ProjectTask Struct for DB Model
	ProjectTask struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID       bson.ObjectId `json:"orgid"`
		Tasks       []Task        `json:"tasks"`
		Description string        `json:"description"`
		StartDate   string        `json:"startdate"`
		EndDate     string        `json:"enddate"`
		Title       string        `json:"title"`
		Status      string        `json:"status,omitempty"`
	}

	// Project Struct for DB Model
	Project struct {
		ID                 bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID              bson.ObjectId `json:"orgid"`
		Manager            bson.ObjectId `json:"empid"`
		ProjectTypeID      bson.ObjectId `json:"projecttypeid"`
		ProjectTaskID      bson.ObjectId `json:"projecttaskid"`
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
		CreatedAt          time.Time     `json:"createdat"`
		UpdatedAt          time.Time     `json:"updatedat"`
		Status             string        `json:"status,omitempty"`
	}

	// JournalAccount Struct for DB Model
	JournalAccount struct {
		ID                      bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID                   bson.ObjectId `json:"orgid"`
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
		CreatedAt               time.Time     `json:"createdat"`
		UpdatedAt               time.Time     `json:"updatedat"`
	}

	// JournalEntry Struct for DB Model
	JournalEntry struct {
		ID                  bson.ObjectId `bson:"_id,omitempty" json:"id"`
		OrgID               bson.ObjectId `json:"orgid"`
		JournalAccountID    bson.ObjectId `json:"journalaccountid"`
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
		CreatedAt           time.Time     `json:"createdat"`
		UpdatedAt           time.Time     `json:"updatedat"`
	}
)
