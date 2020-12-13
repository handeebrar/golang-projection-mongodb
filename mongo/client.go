package mongo
import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	mgo "gopkg.in/mgo.v2"
)

var Db *mgo.Database

// Config db connection struct
type Config struct {
	ConnectionUrl    string `json:"connectionUrl"`
	DatabaseName     string `json:"databaseName"`
}

// Connect Establish a connection to database
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

//LoadConfiguration Parse the configuration file 'config.toml', and establish a connection to DB
func LoadConfiguration() {
	config:=Config{}
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	Connect(config.ConnectionUrl,config.DatabaseName)
}