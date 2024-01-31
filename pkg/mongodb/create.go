package mongodb

import (
	"context"
	"fmt"

	"github.com/cyp57/user-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertDocument inserts a document into the specified MongoDB collection
func InsertOneDocument(collectionName string, data primitive.M, prefixId string) (string, error) {
	infoId := ""
	if !utils.IsEmptyString(prefixId) {
		infoId = utils.GenerateOid(prefixId)
	} else { // default
		infoId = utils.GenerateOid("SH")
	}

	data["id"] = infoId
	fmt.Println("Database  = =", Database)
	_, err := Database.Collection(collectionName).InsertOne(context.TODO(), data)
	if err != nil {
		return "", err
	}

	return infoId, nil
}
