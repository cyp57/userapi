package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/cyp57/userapi/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOneDocument(col string, query primitive.M) (bool, error) {

	result, err := Database.Collection(col).DeleteOne(context.TODO(), query)
	if err != nil {
		return false, err
	}
	if result.DeletedCount != 1 {
		return false, errors.New("DeletedCount : " + fmt.Sprint(result.DeletedCount))
	}

	utils.Debug("DeletedCount :" + fmt.Sprint(result.DeletedCount))
	return true, nil
}
