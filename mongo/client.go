package mongo
import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	mgo "gopkg.in/mgo.v2"
)

var Db *mgo.Database

type Config struct {
	ConnectionUrl    string `json:"connectionUrl"`
	DatabaseName     string `json:"databaseName"`
}

func Connect(connectionUrl string,databaseName string) {
	info := &mgo.DialInfo{
		Addrs:    []string{connectionUrl},
		Timeout:  5 * time.Second,
		Database: databaseName,
		Username: "",
		Password: "",
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println(err.Error())
	}
	Db = session.DB(databaseName)
}

func LoadConfiguration() error{

	config:=Config{}
	configFile, err := os.Open("config.qa.json")
	defer configFile.Close()

	if err != nil {
		return err
	}
	
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	Connect(config.ConnectionUrl,config.DatabaseName)

	return nil
}