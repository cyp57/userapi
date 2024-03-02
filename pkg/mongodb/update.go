package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// filter : condition
// update : data
// arrayfilter : filter for array case
func UpdateDocument(collectionName string, filter primitive.M, update primitive.M, arrayfilter []interface{}) (primitive.M, error) {

	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)

	defer cancel()

	arrayfilterOpts := options.ArrayFilters{
		Filters: arrayfilter,
	} 

	opts := &options.FindOneAndUpdateOptions{}

	opts.SetUpsert(true)
	if arrayfilter != nil {
		opts.SetArrayFilters(arrayfilterOpts)
	}

	var updatedDocument primitive.M

	err := Database.Collection(collectionName).FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDocument)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		return nil, err
	}

	return updatedDocument, nil
}
