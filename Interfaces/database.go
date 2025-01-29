package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Using interfaces in Go to interact with different databases
is a common and powerful pattern. This allows for database
abstraction—your application code can remain database-agnostic and
interact with multiple database implementations (e.g., MySQL, PostgreSQL, MongoDB)
through a common interface.
*/

func databaseInitilization() {
	var db Database

	// Use MySQL
	db = &MySQLDatabase{}
	err := db.Connect("user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		fmt.Println("Error connecting to MySQL:", err)
		return
	}
	fmt.Println("Connected to MySQL")

	db.Create("users", map[string]interface{}{"name": "John Doe", "age": 30})
	result, _ := db.Read("users", "1")
	fmt.Println("MySQL Read:", result)

	// Use MongoDB
	db = &MongoDBDatabase{}
	err = db.Connect("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	fmt.Println("Connected to MongoDB")

	db.Create("users", map[string]interface{}{"_id": "1", "name": "John Doe", "age": 30})
	result, _ = db.Read("users", "1")
	fmt.Println("MongoDB Read:", result)
}

type Database interface {
	Connect(connectionString string) error
	Create(table string, data map[string]interface{}) error
	Read(table string, id string) (map[string]interface{}, error)
	Update(table string, id string, data map[string]interface{}) error
	Delete(table string, id string) error
}

//MYSQL

type MySQLDatabase struct {
	db *sql.DB
}

func (m *MySQLDatabase) Connect(connectionString string) error {
	var err error
	m.db, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	return m.db.Ping() // Ensure the connection is valid
}

func (m *MySQLDatabase) Create(table string, data map[string]interface{}) error {
	query := fmt.Sprintf("INSERT INTO %s (name, age) VALUES (?, ?)", table)
	_, err := m.db.Exec(query, data["name"], data["age"])
	return err
}

func (m *MySQLDatabase) Read(table string, id string) (map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT name, age FROM %s WHERE id = ?", table)
	row := m.db.QueryRow(query, id)

	var name string
	var age int
	err := row.Scan(&name, &age)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"name": name, "age": age}, nil
}

func (m *MySQLDatabase) Update(table string, id string, data map[string]interface{}) error {
	query := fmt.Sprintf("UPDATE %s SET name = ?, age = ? WHERE id = ?", table)
	_, err := m.db.Exec(query, data["name"], data["age"], id)
	return err
}

func (m *MySQLDatabase) Delete(table string, id string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", table)
	_, err := m.db.Exec(query, id)
	return err
}

//MongoDB

type MongoDBDatabase struct {
	client *mongo.Client
}

func (m *MongoDBDatabase) Connect(connectionString string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return err
	}
	m.client = client
	return nil
}

func (m *MongoDBDatabase) Create(table string, data map[string]interface{}) error {
	collection := m.client.Database("test").Collection(table)
	_, err := collection.InsertOne(context.TODO(), data)
	return err
}

func (m *MongoDBDatabase) Read(table string, id string) (map[string]interface{}, error) {
	collection := m.client.Database("test").Collection(table)
	var result map[string]interface{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	return result, err
}

func (m *MongoDBDatabase) Update(table string, id string, data map[string]interface{}) error {
	collection := m.client.Database("test").Collection(table)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": data}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (m *MongoDBDatabase) Delete(table string, id string) error {
	collection := m.client.Database("test").Collection(table)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}
