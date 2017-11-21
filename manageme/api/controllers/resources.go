package controllers

/**
 *
 * @author Sika Kay
 * @date 21/07/17
 *
 */
import (
	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
)

type (
	//UserResource For Post/Put/Get - /api/v1/users/{id}
	UserResource struct {
		Data models.User `json:"data"`
	}

	//UsersResource For Post/Put/Get - /api/v1/users/
	UsersResource struct {
		Data []models.User `json:"data"`
	}

	//SignInResource For Post/Put/Get - /api/v1/users/signin
	SignInResource struct {
		Data SignInModel `json:"data"`
	}

	//AuthUserResource Response For Authorized Users - /api/v1/users/signin
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	//TransactionResource For Post/Put/Get - /api/v1/transactions/{id}
	TransactionResource struct {
		Data models.Transaction `json:"data"`
	}

	//TransactionsResource For Post/Put/Get - /api/v1/transactions/{id}
	TransactionsResource struct {
		Data []models.Transaction `json:"data"`
	}

	//SignInModel For /users/signin
	SignInModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//AuthUserModel For Authorized User with Access Token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)
