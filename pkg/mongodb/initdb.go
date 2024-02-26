package mongodb

import (
	"context"

	"time"

	"github.com/cyp57/userapi/cnst"
	lrlog "github.com/cyp57/userapi/pkg/logrus"
	"github.com/cyp57/userapi/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func dbConnection(c *mongo.Database) {
	Database = c
}

func MongoDbConnect() {

	dbHost := utils.GetYaml(cnst.DBHost)
	dbName := utils.GetYaml(cnst.DBName)
	userDb := utils.GetYaml(cnst.DBUser)
	passDb := utils.GetYaml(cnst.DBPassword)

	ls := &lrlog.LrlogObj{Data: bson.M{"dbHost": dbHost,
		"dbName": dbName, "userDb": userDb, "passDb": passDb}, Txt: "MongoDbConnect()", Level: logrus.DebugLevel}
	ls.Print()

	// // Database Config
	credential := options.Credential{
		Username: userDb,
		Password: passDb,
	}

	clientOptions := options.Client().ApplyURI(dbHost).SetAuth(credential)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		ls := &lrlog.LrlogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
		ls.Print()
	}

	//Cancel context to avoid memory leak
	defer cancel()
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		ls := &lrlog.LrlogObj{Data: nil, Txt: err.Error(), Level: logrus.FatalLevel}
		ls.Print()
	} else {
		ls := &lrlog.LrlogObj{Data: nil, Txt: "DB Connected!", Level: logrus.DebugLevel}
		ls.Print()

	}

	db := client.Database(dbName)
	dbConnection(db)

}
