package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

//golang 封装 mongodb

type mongoClient struct {
	database   string
	table      string
	client     *mongo.Client
	collection *mongo.Collection
}

func NewClient(host string, port int, database string, table string) *mongoClient {
	client, err := mongo.Connect(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
	if err != nil {
		panic(fmt.Sprintf("client error: %s", err))
	}
	return &mongoClient{
		database: database,
		table:    table,
		client:   client,
	}
}
func (m *mongoClient) Collection() *mongo.Collection {
	m.collection = m.client.Database(m.database).Collection(m.table)
	return m.collection
}
func (m *mongoClient) InsertOne(ctx context.Context, data any) (*mongo.InsertOneResult, error) {
	result, err := m.collection.InsertOne(ctx, data)
	return result, err
}
func (m *mongoClient) InsertMany(ctx context.Context, data any) (*mongo.InsertManyResult, error) {
	return m.collection.InsertMany(ctx, data)
}
func (m *mongoClient) Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
	if filter == nil {
		filter = bson.D{}
	}
	return m.collection.Find(ctx, filter)
}

type user struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
	Lan  string `bson:"lan"`
}

func main() {
	client := NewClient("localhost", 27017, "demo", "my")
	collection := client.Collection()
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, bson.D{{"name", "王五"}, {"age", 111}, {"lan", "java"}})
	if err != nil {
		panic(fmt.Sprintf("insert one error %s", err))
	}
	fmt.Println(result.InsertedID)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(fmt.Sprintf("find error: %s", err))
	}
	var rs []user
	if err = cur.All(ctx, &rs); err != nil {
		panic(fmt.Sprintf("parse find error: %s", err))
	}
	fmt.Printf("%#v\n", rs)
}
