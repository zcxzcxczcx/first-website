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

func init() {
	session, err := mgo.Dial("server1.example.com,server2.example.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()
}
func (srv LoginSrv) Save(user User) {

}
