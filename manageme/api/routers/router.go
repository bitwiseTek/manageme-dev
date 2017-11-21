package routers

/**
 *  
 * @author Sika Kay
 * @date 20/07/17
 *
 */

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	//Users Routes
	router = SetUserRoutes(router)

	//Transactions Routes
	router = SetTransactionRoutes(router)

}
