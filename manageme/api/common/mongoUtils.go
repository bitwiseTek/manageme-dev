package common

/**
 *
 * @author Sika Kay
 * @date 22/11/17
 *
 */

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

//GetSession initializes a Mongo Session
func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost},
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}

func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}

func addIndexes() {
	var err error
	userIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	orgUserIndex := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	orgIndex := mgo.Index{
		Key:        []string{"userid"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	empIndex := mgo.Index{
		Key:        []string{"orguserid"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	accIndex := mgo.Index{
		Key:        []string{"orgid"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}
	taskIndex := mgo.Index{
		Key:        []string{"orgid"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	projectIndex := mgo.Index{
		Key:        []string{"orgid"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	txIndex := mgo.Index{
		Key:        []string{"billingid"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}

	session := GetSession().Copy()
	defer session.Close()
	userCol := session.DB(AppConfig.Database).C("users")
	orgUserCol := session.DB(AppConfig.Database).C("orgusers")
	orgCol := session.DB(AppConfig.Database).C("orgs")
	empCol := session.DB(AppConfig.Database).C("employees")
	accCol := session.DB(AppConfig.Database).C("accounts")
	taskCol := session.DB(AppConfig.Database).C("tasks")
	projectCol := session.DB(AppConfig.Database).C("projects")
	txCol := session.DB(AppConfig.Database).C("transactions")

	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	err = orgUserCol.EnsureIndex(orgUserIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	err = orgCol.EnsureIndex(orgIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	err = empCol.EnsureIndex(empIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	err = accCol.EnsureIndex(accIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	err = taskCol.EnsureIndex(taskIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}

	err = projectCol.EnsureIndex(projectIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	
	err = txCol.EnsureIndex(txIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
}
