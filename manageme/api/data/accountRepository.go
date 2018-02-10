package data

/**
 *
 * @author Sika Kay
 * @date 17/01/18
 *
 */
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"time"

	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
)

//AccountRepository Struct for Mongo Persistence
type AccountRepository struct {
	C *mgo.Collection
}

//AddAccount persists Account associated with OrgID
func (r *AccountRepository) AddAcctByOrgID(orgID string) (acct models.Account, err error) {
	objID := bson.NewObjectId()
	acct.ID = objID
	acct.OrgID = bson.ObjectIdHex(orgID)
	acct.Status = "Active"
	acct.CreatedAt = time.Now()

	err = r.C.Insert(&acct)
	return
}

//AddAccount persists Account associated with OrgID, AcctID
func (r *AccountRepository) AddChildAcctByOrgID(orgID string, acctID string) (acct models.Account, err error) {
	objID := bson.NewObjectId()
	acct.ID = objID
	acct.OrgID = bson.ObjectIdHex(orgID)
	acct.ParentAccount = bson.ObjectIdHex(acctID)
	acct.Status = "Active"
	acct.CreatedAt = time.Now()

	err = r.C.Insert(&acct)
	return
}

//GetAcctByID gets account associated with an ID
func (r *AccountRepository) GetAcctByID(id string) (acct models.Account, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&acct)
	return
}

//GetAccounts gets all accounts
func (r *AccountRepository) GetAccounts() []models.Account {
	var accts []models.Account
	iter := r.C.Find(nil).Iter()
	result := models.Account{}
	for iter.Next(&result) {
		accts = append(accts, result)
	}
	return accts
}

//GetAcctsByOrgID gets accts associated with an OrgID
func (r *AccountRepository) GetAcctsByOrgID(orgID string) []models.Account {
	var accts []models.Account
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Account{}
	for iter.Next(&result) {
		accts = append(accts, result)
	}
	return accts
}

//GetAcctsByAcctID gets accts associated with an AcctID
func (r *AccountRepository) GetAcctsByAcctID(acctID string) []models.Account {
	var accts []models.Account
	acctid := bson.ObjectIdHex(acctID)
	iter := r.C.Find(bson.M{"accountid": acctid}).Iter()
	result := models.Account{}
	for iter.Next(&result) {
		accts = append(accts, result)
	}
	return accts
}

//EditAccountByID edits account associated with an ID
func (r *AccountRepository) EditAccountByID(acct *models.Account) error {
	err := r.C.Update(bson.M{"_id": acct.ID},
		bson.M{"$set": bson.M{
			"name": 			acct.Name,
			"roottype":   		acct.RootType,
			"reporttype":       acct.ReportType,
			"accountcurrency":	acct.AccountCurrency,
			"type":				acct.Type,
			"taxrate":			acct.TaxRate,
			"balancetype":		acct.BalanceType,
			"isaccountfreeze":	acct.IsAccountFreeze,
			"isgroup":			acct.IsGroup,
			"updatedat":        time.Now(),
		}})
	return err
}

//EditChildAccountByID edits child account associated with an ID
func (r *AccountRepository) EditChildAccountByID(acct *models.Account) error {
	err := r.C.Update(bson.M{"accountid": acct.ParentAccount},
		bson.M{"$set": bson.M{
			"name": 			acct.Name,
			"roottype":   		acct.RootType,
			"reporttype":       acct.ReportType,
			"accountcurrency":	acct.AccountCurrency,
			"type":				acct.Type,
			"taxrate":			acct.TaxRate,
			"balancetype":		acct.BalanceType,
			"isaccountfreeze":	acct.IsAccountFreeze,
			"isgroup":			acct.IsGroup,
			"updatedat":        time.Now(),
		}})
	return err
}

//DeleteAccountById deletes account out of the system by Id
func (r *AccountRepository) DeleteAccountById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//AddJournalAccountID persists Journal Account associated with AccountID, ProjectID
func (r *AccountRepository) AddJournalAcctByAccountID(accountID string, projectID string) (acct models.JournalAccount, err error) {
	objID := bson.NewObjectId()
	acct.ID = objID
	acct.AccountID = bson.ObjectIdHex(accountID)
	acct.ProjectID = bson.ObjectIdHex(projectID)
	acct.CreatedAt = time.Now()
	acct.UpdatedAt = time.Now()

	err = r.C.Insert(&acct)
	return
}

//AddJournalEntry persists Journal Entry associated with AccountID
func (r *AccountRepository) AddJournalEntryByAcctID(accountID string) (entry models.JournalEntry, err error) {
	objID := bson.NewObjectId()
	entry.ID = objID
	entry.JournalAccountID = bson.ObjectIdHex(accountID)
	entry.PostingDate = time.Now()
	entry.CreatedAt = time.Now()
	entry.UpdatedAt = time.Now()
	
	err = r.C.Insert(&entry)
	return
}

//GetAcctByID gets journal account associated with an ID
func (r *AccountRepository) GetJournalAcctByID(id string) (acct models.JournalAccount, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&acct)
	return
}

//GetJournalAccounts gets all accounts
func (r *AccountRepository) GetJournalAccounts() []models.JournalAccount {
	var accts []models.JournalAccount
	iter := r.C.Find(nil).Iter()
	result := models.JournalAccount{}
	for iter.Next(&result) {
		accts = append(accts, result)
	}
	return accts
}

//GetJournalAcctsByOrgID gets accts associated with an OrgID
func (r *AccountRepository) GetJournalAccountsByOrgID(orgID string) []models.JournalAccount {
	var accts []models.JournalAccount
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.JournalAccount{}
	for iter.Next(&result) {
		accts = append(accts, result)
	}
	return accts
}

//GetJournalAcctsByAcctID gets accts associated with an AccountID
func (r *AccountRepository) GetJournalAccountsByAcctID(accountID string) []models.JournalAccount {
	var accts []models.JournalAccount
	accountid := bson.ObjectIdHex(accountID)
	iter := r.C.Find(bson.M{"accountid": accountid}).Iter()
	result := models.JournalAccount{}
	for iter.Next(&result) {
		accts = append(accts, result)
	}
	return accts
}

//EditJournalAccountByID edits journal account associated with an ID
func (r *AccountRepository) EditJournalAccountByID(acct *models.JournalAccount) error {
	err := r.C.Update(bson.M{"_id": acct.ID},
		bson.M{"$set": bson.M{
			"accounttype": 				acct.AccountType,
			"balance":   				acct.Balance,
			"debitinaccountcurrency":   acct.DebitInAccountCurrency,
			"accountcurrency":			acct.AccountCurrency,
			"creditinaccountcurrency":	acct.CreditInAccountCurrency,
			"referencetype":			acct.ReferenceType,
			"referencename":			acct.ReferenceName,
			"writeoffbasedon":			acct.WriteOffBasedOn,
			"debit":					acct.Debit,
			"credit":        			acct.Credit,
			"exchangerate":				acct.ExchangeRate,	
			"isadvance":				acct.IsAdvance,		
			"totalamountcurrency":		acct.TotalAmountCurrency,	
			"updatedat":				time.Now(),
		}})
	return err
}

//DeleteJournalAccountById deletes journal account out of the system by Id
func (r *AccountRepository) DeleteJournalAccountById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//GetJournalEntryByID gets journal entry associated with an ID
func (r *AccountRepository) GetEntryByID(id string) (entry models.JournalEntry, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&entry)
	return
}

//GetEntries gets all entries
func (r *AccountRepository) GetEntries() []models.JournalEntry {
	var entries []models.JournalEntry
	iter := r.C.Find(nil).Iter()
	result := models.JournalEntry{}
	for iter.Next(&result) {
		entries = append(entries, result)
	}
	return entries
}

//GetJournalEntiriesByOrgID gets entries associated with an OrgID
func (r *AccountRepository) GetJournalEntriesByOrgID(orgID string) []models.JournalEntry {
	var entries []models.JournalEntry
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.JournalEntry{}
	for iter.Next(&result) {
		entries = append(entries, result)
	}
	return entries
}

//GetJournalEntiriesByJAcctID gets entries associated with a JournalAcctID
func (r *AccountRepository) GetJournalEntriesByAcctID(orgID string) []models.JournalEntry {
	var entries []models.JournalEntry
	journalaccountid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"journalaccountid": journalaccountid}).Iter()
	result := models.JournalEntry{}
	for iter.Next(&result) {
		entries = append(entries, result)
	}
	return entries
}

//EditJournalEntryByID edits journal entry associated with an ID
func (r *AccountRepository) EditJournalEntryByID(entry *models.JournalEntry) error {
	err := r.C.Update(bson.M{"_id": entry.ID},
		bson.M{"$set": bson.M{
			"title": 				entry.Title,
			"vouchertype":   		entry.VoucherType,
			"chequeno":   			entry.ChequeNo,
			"chequedate":			entry.ChequeDate,
			"billno":				entry.BillNo,
			"billdate":				entry.BillDate,
			"duedate":				entry.DueDate,
			"remarks":				entry.Remarks,
			"writeoffbasedon":		entry.WriteOffBasedOn,
			"totaldebit":        	entry.TotalDebit,
			"totalcredit":			entry.TotalCredit,	
			"difference":			entry.Difference,	
			"totalamount":			entry.TotalAmount,
			"writeoffamount":		entry.WriteOffAmount,
			"totalamountcurrency":	entry.TotalAmountCurrency,
			"updatedat":			time.Now(),
		}})
	return err
}

//DeleteJournalEntryById deletes journal entry out of the system by Id
func (r *AccountRepository) DeleteJournalEntryById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
