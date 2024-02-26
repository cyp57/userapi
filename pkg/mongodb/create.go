package mongodb

import (
	"context"

	"github.com/cyp57/userapi/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOneDocument(collectionName string, data primitive.M, prefixId string) (string, error) {
	infoId := ""
	if !utils.IsEmptyString(prefixId) { // custom generate field id
		infoId = utils.GenerateOid(prefixId)
	} else { // default
		infoId = utils.GenerateOid("SH")
	}

	data["id"] = infoId

	_, err := Database.Collection(collectionName).InsertOne(context.TODO(), data)
	if err != nil {
		return "", err
	}

	return infoId, nil
}
