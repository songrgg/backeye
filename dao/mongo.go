package dao

import (
	"fmt"

	"github.com/songrgg/backeye/common"
	mgo "gopkg.in/mgo.v2"
)

var (
	session               *mgo.Session
	db                    *mgo.Database
	targetCollection      *mgo.Collection
	watchResultCollection *mgo.Collection
)

func InitMongo() {
	session, err := NewMongoSession(common.Config.Schedules.Address)
	if err != nil {
		panic(fmt.Sprintf("fail to start mongo: %v", err))
	}
	db = session.DB("schedule")
	targetCollection = db.C("target")
	watchResultCollection = db.C("watch_result")
}

func NewMongoSession(address string) (*mgo.Session, error) {
	if session != nil {
		return session, nil
	}

	session, err := mgo.Dial(address)
	if err != nil {
		return nil, err
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}

func CloseMongo(session *mgo.Session) {
	session.Close()
}
