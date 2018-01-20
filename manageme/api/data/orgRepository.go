package data

/**
 *
 * @author Sika Kay
 * @date 03/01/18
 *
 */
import (
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
