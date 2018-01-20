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

//BillingRepository Struct for Mongo Persistence
type BillingRepository struct {
	C *mgo.Collection
}

//AddBilling persists Billing associated with a UserID
func (r *BillingRepository) AddBillingByOrgID(orgID string) (billing models.Billing, err error) {
	objID := bson.NewObjectId()
	billing.ID = objID
	billing.OrgID = bson.ObjectIdHex(orgID)
	billing.Currency = "NGN"
	billing.Status = "Due"
	billing.LastBilled = time.Now()

	err = r.C.Insert(&billing)
	return
}

//GetBills gets all bills
func (r *BillingRepository) GetBillings() []models.Billing {
	var billings []models.Billing
	iter := r.C.Find(nil).Iter()
	result := models.Billing{}
	for iter.Next(&result) {
		billings = append(billings, result)
	}
	return billings
}

//GetBillByID gets bill associated with an ID
func (r *BillingRepository) GetBillingByID(id string) (billings models.Billing, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&billings)
	return
}

//GetBillingByOrgID gets bill associated with an OrgID
func (r *BillingRepository) GetBillingsByOrgID(orgID string) []models.Billing {
	var bills []models.Billing
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Billing{}
	for iter.Next(&result) {
		bills = append(bills, result)
	}
	return bills
}

//AddTransaction persists Transcaction associated with a BillingID
func (r *BillingRepository) AddTransactionByBillingID(billingID string) (tx models.Transaction, err error) {
	objID := bson.NewObjectId()
	tx.ID = objID
	tx.BillingID = bson.ObjectIdHex(billingID)
	tx.Status = "Pending"
	tx.CreatedAt = time.Now()

	err = r.C.Insert(&tx)
	return
}

//GetTransactions gets all transactions
func (r *BillingRepository) GetTransactions() []models.Transaction {
	var txs []models.Transaction
	iter := r.C.Find(nil).Iter()
	result := models.Transaction{}
	for iter.Next(&result) {
		txs = append(txs, result)
	}
	return txs
}

//GetTransactionByID gets transaction associated with an ID
func (r *BillingRepository) GetTransactionByID(id string) (txs models.Transaction, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&txs)
	return
}

//GetTransactionByBillingID gets tranaction associated with an BillID
func (r *BillingRepository) GetTransactionByBillingID(billingID string) []models.Transaction {
	var txs []models.Transaction
	billingid := bson.ObjectIdHex(billingID)
	iter := r.C.Find(bson.M{"billingid": billingid}).Iter()
	result := models.Transaction{}
	for iter.Next(&result) {
		txs = append(txs, result)
	}
	return txs
}
