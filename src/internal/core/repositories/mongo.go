package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepoImpl struct {
	client   *mongo.Client
	database string
}

func (r *MongoRepoImpl) FindAll(collection string, filter interface{}, result interface{}) error {
	c := r.client.Database(r.database).Collection(collection)

	cursor, err := c.Find(context.TODO(), filter)
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), result); err != nil {
		return err
	}

	return nil
}

func NewMongoRepo(client *mongo.Client, database string) *MongoRepoImpl {
	return &MongoRepoImpl{client: client, database: database}
}
