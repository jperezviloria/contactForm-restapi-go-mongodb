package repository

import (
	"context"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2"
	"log"
	"time"
)

const (
	DBName = "golang"
	CName  = "users"
)

var session *mgo.Session

func createSession() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
}

func getSession() *mgo.Session {
	if session == nil {
		createSession()
	}
	return session
}

func InitData() {
	createSession()
}

func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI("http://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

}
