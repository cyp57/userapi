package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOneDocument(col string, filter primitive.M) (result primitive.M, err error) {

	err = Database.Collection(col).FindOne(context.TODO(), filter).Decode(&result)

	// Prints a message if no documents are matched or if any
	// other errors occur during the operation
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		} else {
			return
		}
	}

	return result, nil
}


func AggregateDocument(col string, pipeline []primitive.M) (result primitive.M , err error){


	cursor , err := Database.Collection(col).Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Println(err.Error())
		return nil , err
	}

	if err = cursor.All(context.TODO(), &result); err != nil {
		log.Println(err.Error())
	return nil , err
	}

defer cursor.Close(context.TODO())

	return result , nil
}



// limit = input limit
//skip = limit * (input offset - 1)
func FindDocument(col string ,filter primitive.M ,projection , sort primitive.M , skip int64 , limit int64) (result []primitive.M , err error) {

opts := options.Find()
if projection != nil {
	opts.SetProjection(projection)
}
if projection != nil { 
	opts.SetSort(sort)
}

opts.SetSkip(skip) 
opts.SetLimit(limit)

cursor , err := Database.Collection(col).Find(context.TODO(), filter,opts)
if err != nil {
	log.Println(err.Error())
	return nil , err
}

if err = cursor.All(context.TODO(), &result); err != nil {
	log.Println(err.Error())
	return nil , err
}

defer cursor.Close(context.TODO())

	return result , nil
}