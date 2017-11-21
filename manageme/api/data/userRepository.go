package data

/**
 *
 * @author Sika Kay
 * @date 22/07/17
 *
 */
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"time"

	"github.com/bitwiseTek/spottnow-dev/spottnow/api/models"
	"golang.org/x/crypto/bcrypt"
)

//UserRepository Struct for Mongo Persistence
type UserRepository struct {
	C *mgo.Collection
}

//CreateUser persists User to MongoDB
func (r *UserRepository) AddUser(user *models.User) error {
	objID := bson.NewObjectId()
	user.Id = objID
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass

	user.Password = ""
	user.Status = "Active"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = r.C.Insert(&user)
	return err
}

//SignIn
func (r *UserRepository) SignIn(user *models.User) (u models.User, err error) {
	err = r.C.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}

func (r *UserRepository) GetUsers() []models.User {
	var users []models.User
	iter := r.C.Find(nil).Iter()
	result := models.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (r *UserRepository) GetUserById(id string) (user models.User, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&user)
	return
}

func (r *UserRepository) EditUserById(user *models.User) error {
	err := r.C.Update(bson.M{"_id": user.Id},
		bson.M{"$set": bson.M{
			"name":         user.Name,
			"dob":          user.DOB,
			"phone":        user.Phone,
			"profileimage": user.ProfileImage,
			"status":       user.Status,
			"updatedat":    time.Now(),
		}})
	return err
}
