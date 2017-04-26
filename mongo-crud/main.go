package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

const (
	host = "ds153730.mlab.com:53730"
	databaseName = "comolocco"
	username = "ercinakcay"
	password = "123456"
)


type Player struct {
	Name   string    `bson:"name"`
	Points uint8     `bson:"points"`
	Place  uint8     `bson:"place"`
}

func main() {
	player := Player {
		Name: "Dave",
		Points: 10,
		Place: 2,
	}

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{host},
		Username: username,
		Password: password,
		Database: databaseName,
	})
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Printf("Connected to %v\n", session.LiveServers())


	coll := session.DB(databaseName).C("Player")
	if err := coll.Insert(player); err != nil {
		panic(err)
	}
	fmt.Println("Document inserted successfully!")
}

