package services

import (
	"gopkg.in/mgo.v2"
)

type (
	LoginSrv struct{}
	User     struct {
		Name     string `json:"name" bson:"name"`
		Password string `json:"password" bson:"password"`
	}
)

var (
	mgoSession = newMongoClient()
)

const (
	Users = "users"
)

func newMongoClient() *mgo.Session {
	session, err := mgo.Dial("mongodb://root:123456@121.196.11.0:27017/first_website")
	if err != nil {
		panic(err)
	}
	return session
}
func (srv LoginSrv) Save(user User) (err error) {
	return mgoSession.DB("").C(Users).Insert(user)
}
