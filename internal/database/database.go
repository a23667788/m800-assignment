package database

import (
	"context"
	"log"

	"github.com/a23667788/m800-assignment/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDb(mongoDb config.MongoDb) (*Mongo, error) {
	connectionString := mongoDb.ConnectionString
	dbName := mongoDb.Name

	log.Println(mongoDb.ConnectionString)
	log.Println(mongoDb.Name)

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &Mongo{client: client, db: db}, nil
}

func (m *Mongo) Close() {
	if m.client != nil {
		m.client.Disconnect(context.Background())
	}
}

func (m *Mongo) Insert(collectionName string, document interface{}) error {
	collection := m.db.Collection(collectionName)

	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) Query(collectionName string, filter interface{}) ([]interface{}, error) {
	collection := m.db.Collection(collectionName)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []interface{}

	for cursor.Next(context.TODO()) {
		var result interface{}
		if err := cursor.Decode(&result); err != nil {
			log.Println("Error decoding:", err)
			continue
		}
		results = append(results, result)
	}

	return results, nil
}

func (m *Mongo) QueryAll(collectionName string) ([]interface{}, error) {
	collection := m.db.Collection(collectionName)
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []interface{}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Println("Error decoding", err)
	}

	return results, nil
}

func (m *Mongo) QueryEvent(collectionName string, filter interface{}) ([]RawEvent, error) {
	collection := m.db.Collection(collectionName)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []RawEvent

	for cursor.Next(context.TODO()) {
		var result RawEvent
		if err := cursor.Decode(&result); err != nil {
			log.Println("Error decoding:", err)
			continue
		}
		results = append(results, result)
	}

	return results, nil
}

func (m *Mongo) QueryAllEvent(collectionName string) ([]RawEvent, error) {
	collection := m.db.Collection(collectionName)
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []RawEvent

	for cursor.Next(context.TODO()) {
		var result RawEvent
		if err := cursor.Decode(&result); err != nil {
			log.Println("Error decoding:", err)
			continue
		}
		results = append(results, result)
	}

	return results, nil
}
