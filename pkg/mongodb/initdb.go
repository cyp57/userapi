package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cyp57/user-api/cnst"

	"github.com/cyp57/user-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

// DBConnection ..
func dbConnection(c *mongo.Database) {
	Database = c
}

// Connect is for get mongo driver connection
func MongoDbConnect() {

	connectionString := utils.GetYaml(cnst.DBHost)
	dbName := utils.GetYaml(cnst.DBName)
	userDb := utils.GetYaml(cnst.DBUser)
	passDb := utils.GetYaml(cnst.DBPassword)
	fmt.Println("connectionString :", connectionString)
	fmt.Println("dbName :", dbName)
	fmt.Println("userDb :", userDb)
	fmt.Println("passDb :", passDb)
	// // Database Config
	credential := options.Credential{
		Username: userDb,
		Password: passDb,
	}

	clientOptions := options.Client().ApplyURI(connectionString).SetAuth(credential)
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	
	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//Cancel context to avoid memory leak
	defer cancel()
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		//log.Println("Connected!")
		log.Println("DB Connected!")
	}
	// defer client.Disconnect(ctx)

	// //Connect to the database
	db := client.Database(dbName)
	dbConnection(db)


}

