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
func (r *AccountRepository) AddAcctByOrgID(orgID string, acctID string) (acct models.Account, err error) {
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

//AddJournalAccountID persists Journal Account associated with AccountID, ProjectID
func (r *AccountRepository) AddJournalAcctByAccountID(accountID string, projectID string) (acct models.JournalAccount, err error) {
	objID := bson.NewObjectId()
	acct.ID = objID
	acct.accountID = bson.ObjectIdHex(accountID)
	acct.projectID = bson.ObjectIdHex(projectID)

	err = r.C.Insert(&acct)
	return
}

//AddJournalEntry persists Journal Entry associated with AccountID
func (r *AccountRepository) AddJournalEntryByAcctID(accountID string) (entry models.JournalEntry, err error) {
	objID := bson.NewObjectId()
	entry.ID = objID
	entry.JournalAccountID = bson.ObjectIdHex(accountID)
	entry.PostingDate = time.Now()

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
func (r *AccountRepository) GetJournalAccountsByOrgID() []models.JournalAccount {
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
func (r *AccountRepository) GetJournalAccountsByAcctID() []models.JournalAccount {
	var accts []models.JournalAccount
	accountid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"accountid": accountid}).Iter()
	result := models.JournalAccount{}
	for iter.Next(&result) {
		accts = append(accts, result)
	}
	return accts
}

//GetJournalEntryByID gets journal entry associated with an ID
func (r *AccountRepository) GetEntryByID(id string) (entry models.JournalEntry, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&entry)
	return
}

//GetEntries gets all entries
func (r *AccountRepository) GetEntries() []models.JournalEntry {
	var entry []models.JournalEntry
	iter := r.C.Find(nil).Iter()
	result := models.JournalEntry{}
	for iter.Next(&result) {
		entries = append(accts, result)
	}
	return entries
}

//GetJournalEntiriesByOrgID gets entries associated with an OrgID
func (r *AccountRepository) GetJournalEntriesByOrgID() []models.JournalEntry {
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
func (r *AccountRepository) GetJournalEntriesByJAcctID() []models.JournalEntry {
	var entries []models.JournalEntry
	journalaccountid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"journalaccountid": journalaccountid}).Iter()
	result := models.JournalEntry{}
	for iter.Next(&result) {
		entries = append(entries, result)
	}
	return entries
}