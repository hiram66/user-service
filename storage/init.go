package storage

import (
	"fmt"
	"github.com/hiram66/user-service/configs"
	"github.com/hiram66/user-service/storage/mongodb"
	"os"
)

func init() {
	connErr := mongodb.Connect(configs.MongoURI, configs.MongoDbName)
	if connErr != nil {
		fmt.Printf("mongodb connection error %v", connErr)
		os.Exit(1)
	}
	UserStorage = mongodb.NewUserStorage()
}
