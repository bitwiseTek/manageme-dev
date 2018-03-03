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

//RoleRepository Struct for Mongo Persistence
type RoleRepository struct {
	C *mgo.Collection
}

//AddRole persists Role associated with PermID
func (r *RoleRepository) AddRoleByOrgID(orgID string) (role models.Role, err error) {
	objID := bson.NewObjectId()
	role.ID = objID
	role.OrgID = bson.ObjectIdHex(orgID)
	role.Status = "Active"
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()

	err = r.C.Insert(&role)
	return
}

//GetRoles gets all roles
func (r *RoleRepository) GetRoles() []models.Role {
	var roles []models.Role
	iter := r.C.Find(nil).Iter()
	result := models.Role{}
	for iter.Next(&result) {
		roles = append(roles, result)
	}
	return roles
}

//GetRoleByID gets role associated with an ID
func (r *RoleRepository) GetRoleByID(id string) (role models.Role, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&role)
	return
}

//GetRolesByOrgID gets role associated with a PermID
func (r *RoleRepository) GetRolesByOrgID(orgID string) []models.Role {
	var roles []models.Role
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Role{}
	for iter.Next(&result) {
		//permissions := result.Permissions
		/*for _, p := range permissions {
			p = append(p, result)
		}*/
	}
	return roles
}

//EditRoleByID edits role associated with an ID
func (r *RoleRepository) EditRoleByID(role *models.Role) error {
	err := r.C.Update(bson.M{"_id": role.ID},
		bson.M{"$set": bson.M{
			"name":          role.Name,
			"documenttype":  role.DocumentType,
			"documentstage": role.DocumentStage,
			"status":        role.Status,
			"updatedat":     time.Now(),
		}})
	return err
}

//DeleteRoleById deletes role out of the system by Id
func (r *RoleRepository) DeleteRoleById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//AddPermissionByOrgID persists Permission associated with OrgID
func (r *RoleRepository) AddPermission(orgID string) (perm models.Permission, err error) {
	objID := bson.NewObjectId()
	perm.ID = objID
	perm.OrgID = bson.ObjectIdHex(orgID)
	perm.Status = "Active"
	perm.CreatedAt = time.Now()
	perm.UpdatedAt = time.Now()

	err = r.C.Insert(&perm)
	return
}

//GetPermissions gets all permissions
func (r *RoleRepository) GetPermissions() []models.Permission {
	var perms []models.Permission
	iter := r.C.Find(nil).Iter()
	result := models.Permission{}
	for iter.Next(&result) {
		perms = append(perms, result)
	}
	return perms
}

//GetPermissionByID gets perm associated with an ID
func (r *RoleRepository) GetPermByID(id string) (perms models.Permission, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&perms)
	return
}

//GetPermissionsByOrgID gets permissions associated with an OrgID
func (r *RoleRepository) GetPermissionsByOrgID(orgID string) []models.Permission {
	var perms []models.Permission
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Permission{}
	for iter.Next(&result) {
		perms = append(perms, result)
	}
	return perms
}

//EditPermissionByID edits permission associated with an ID
func (r *RoleRepository) EditPermissionByID(perm *models.Permission) error {
	err := r.C.Update(bson.M{"_id": perm.ID},
		bson.M{"$set": bson.M{
			"name":      perm.Name,
			"level":     perm.Level,
			"status":    perm.Status,
			"updatedat": time.Now(),
		}})
	return err
}

//DeletePermissionById deletes permission out of the system by Id
func (r *RoleRepository) DeletePermissionById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
