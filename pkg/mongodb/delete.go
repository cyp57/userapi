package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/cyp57/user-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOneDocument(col string, query primitive.M) (bool, error) {

	result, err := Database.Collection(col).DeleteOne(context.TODO(), query)
	// fmt.Println("DeleteOneDocument result :", result)
	// fmt.Println("DeleteOneDocument err :", err)
	if err != nil {
		return false, err
	}
	if result.DeletedCount != 1 {
		return false, errors.New("DeletedCount : " + fmt.Sprint(result.DeletedCount))
	}

	utils.Debug("DeletedCount :" + fmt.Sprint(result.DeletedCount))
	return true, nil
}
