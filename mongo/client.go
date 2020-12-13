package mongo

import (
	"encoding/json"
	mgo "gopkg.in/mgo.v2"
	"os"
	"time"
)

var Db *mgo.Database

type Config struct {
	ConnectionUrl    string `json:"connectionUrl"`
	DatabaseName     string `json:"databaseName"`
}

func Connect(connectionUrl string,databaseName string) error {
	info := &mgo.DialInfo{
		Addrs:    []string{connectionUrl},
		Timeout:  5 * time.Second,
		Database: databaseName,
		Username: "",
		Password: "",
	}

	session, err := mgo.DialWithInfo(info)

	if err != nil {
		return err
	} else {
		Db = session.DB(databaseName)
		return nil
	}
}

func LoadConfiguration() error{
	var (
		err error
		configFile *os.File
	)

	config:=Config{}
	configFile, err = os.Open("config.qa.json")
	defer configFile.Close()

	if err != nil {
		return err
	}
	
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	if err = Connect(config.ConnectionUrl,config.DatabaseName); err != nil {
		return err
	} else {
		return nil
	}
}