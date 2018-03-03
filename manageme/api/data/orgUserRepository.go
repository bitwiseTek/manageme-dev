package data

/**
 *
 * @author Sika Kay
 * @date 17/01/18
 *
 */
import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"time"

	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
)

//OrgUserRepository Struct for Mongo Persistence
type OrgUserRepository struct {
	C *mgo.Collection
}

//AddUser persists OrgUser associated with OrgID, RoleID
func (r *OrgUserRepository) AddUserByOrgID(orgID string, roleID string) (user models.OrgUser, err error) {
	objID := bson.NewObjectId()
	user.ID = objID
	user.OrgID = bson.ObjectIdHex(orgID)
	user.RoleID = bson.ObjectIdHex(roleID)
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.TempPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	user.TempPassword = ""
	user.Status = "Active"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = r.C.Insert(&user)
	return
}

//GetUsers gets all org users
func (r *OrgUserRepository) GetOrgUsers() []models.OrgUser {
	var users []models.OrgUser
	iter := r.C.Find(nil).Iter()
	result := models.OrgUser{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

//GetUserByID gets org user associated with an ID
func (r *OrgUserRepository) GetOrgUserByID(id string) (users models.OrgUser, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&users)
	return
}

//GetUsersByOrgID gets org user associated with an OrgID
func (r *OrgUserRepository) GetUsersByOrgID(orgID string) []models.OrgUser {
	var users []models.OrgUser
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.OrgUser{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (r *OrgUserRepository) DeleteOrgUserById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//EditOrgUserByID edits org user associated with an ID
func (r *OrgUserRepository) EditOrgUserByID(user *models.OrgUser) error {
	err := r.C.Update(bson.M{"_id": user.ID},
		bson.M{"$set": bson.M{
			"roleid":       user.RoleID,
			"username":     user.Username,
			"temppassword": user.TempPassword,
			"usertype":     user.UserType,
			"status":       user.Status,
			"updatedat":    time.Now(),
		}})
	return err
}

//EditOrgUserByOrgUserID edits org user associated with an ID
func (r *OrgUserRepository) EditOrgUserByOrgUserID(user *models.OrgUser) error {
	err := r.C.Update(bson.M{"_id": user.ID},
		bson.M{"$set": bson.M{
			"temppassword": user.TempPassword,
			"securityques": user.SecurityQues,
			"securityans":  user.SecurityAns,
			"updatedat":    time.Now(),
		}})
	return err
}

//OrgUserSignIn logs org users into the system
func (r *OrgUserRepository) SignIn(user *models.OrgUser) (u models.OrgUser, err error) {
	err = r.C.Find(bson.M{"username": user.Username}).One(&u)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.TempPassword))
	if err != nil {
		u = models.OrgUser{}
		u.LastLogin = time.Now()
	}
	return
}
