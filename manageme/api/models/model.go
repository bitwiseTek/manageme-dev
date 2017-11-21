package models

/**
 *
 * @author Sika Kay
 * @date 21/07/17
 *
 */

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	//User Struct for DB Model
	User struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name         string        `json:"name"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
		DOB          string        `json:"dob"`
		Phone        []string      `json:"phone"`
		ProfileImage string        `json:"profileimage"`
		Status       string        `json:"status,omitempty"`
		CreatedAt    time.Time     `json:"createdat,omitempty"`
		UpdatedAt    time.Time     `json:"updatedat,omitempty"`
	}

	//Transaction Struct for DB Model
	Transaction struct {
		ID            bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserID        bson.ObjectId `json:"userid"`
		RefCode       string        `json:"refcode"`
		Type          string        `json:"type"`
		Statement     string        `json:"statement"`
		PaymentRef    string        `json:"paymentref"`
		PaymentMethod string        `json:"paymentmethod"`
		ResponseCode  string        `json:"responsecode"`
		Status        string        `json:"status,omitempty"`
		CreatedAt     time.Time     `json:"createdat,omitempty"`
		UpdatedAt     time.Time     `json:"updatedat,omitempty"`
	}
)
