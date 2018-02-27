package routers

/**
 *
 * @author Sika Kay
 * @date 22/11/17
 *
 */

import (
	"github.com/gorilla/mux"
)

//InitRoutes initializes routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	//Users Routes
	router = SetUserRoutes(router)

	//Billings Routes
	router = SetBillingRoutes(router)

	//Accounts Routes
	router = setAccountRoutes(router)

	//Employees Routes
	router = setEmployeeRoutes(router)

	//Settings Routes
	router = setSettingRoutes(router)

	//Orgs Routes
	router = setOrgRoutes(router)

	//OrgUsers Routes
	router = setOrgUserRoutes(router)

	//Projects Routes
	router = setProjectRoutes(router)

	//Roles Routes
	router = setRoleRoutes(router)

}
