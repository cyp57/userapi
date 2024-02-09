package mongodb

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteOneDocument(col string, query primitive.M) (bool, error) {

	result, err := Database.Collection(col).DeleteOne(context.TODO(), query)
	fmt.Println("DeleteOneDocument result :", result)
	fmt.Println("DeleteOneDocument err :", err)
	if err != nil {
		return false, err
	}
	if result.DeletedCount != 1 {
		return false, errors.New("delete count : " + fmt.Sprint(result.DeletedCount))
	}
	fmt.Println("result count", result.DeletedCount)
	return true, nil
}
