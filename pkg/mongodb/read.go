package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOneDocument(col string, filter primitive.M, projection primitive.M) (result primitive.M, err error) {

	var opts *options.FindOneOptions
	if projection != nil {
		opts = options.FindOne().SetProjection(projection)
	} else {
		opts = options.FindOne().SetProjection(bson.M{"_id": 0})
	}
	
	err = Database.Collection(col).FindOne(context.TODO(), filter,opts).Decode(&result)

	// Prints a message if no documents are matched or if any
	// other errors occur during the operation
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		} else {
			return nil, err
		}
	}

	return result, nil
}

func AggregateDocument(col string, pipeline []primitive.M) (result []primitive.M, err error) {

	cursor, err := Database.Collection(col).Aggregate(context.TODO(), pipeline)
	if err != nil {

		log.Println(err.Error())
		return nil, err
	}

	if err = cursor.All(context.TODO(), &result); err != nil {

		log.Println(err.Error())
		return nil, err
	}

	defer cursor.Close(context.TODO())

	return result, nil
}

// limit = input limit
// skip = limit * (input offset - 1)
func FindDocument(col string, filter primitive.M, projection, sort primitive.M, skip int64, limit int64) (result []primitive.M, err error) {

	opts := options.Find()
	if projection != nil {
		opts.SetProjection(projection)
	}
	if projection != nil {
		opts.SetSort(sort)
	}

	opts.SetSkip(skip)
	opts.SetLimit(limit)

	cursor, err := Database.Collection(col).Find(context.TODO(), filter, opts)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer cursor.Close(context.TODO())

	return result, nil
}
