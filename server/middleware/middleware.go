package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"os"
	"github.com/jeetnation/todoApp.git/models"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init(){
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv(){
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("Error in Loading the ENV file")
	}
}

func createDBInstance(){
	connectionString := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	CollName := os.Getenv("DB_COLLECTION_NAME")

	clientOptions := options.Client().ApplyURL(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err!=nil{
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("connected to mongoDB:")

	client.Database(dbName).Collection(CollName)
	fmt.Println("Collection Instance Created:")
}


func GetAllTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content.Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTask()
	json.NewEncoder(w).Encode(payload)
}



func CreateTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content.Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	var task models.TodoList
	json.NewDecoder(r.Body).Decode(&task)
	insertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func TaskComplete(task string){
	w.Header().Set("content.Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params)
}

func UndoTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content.Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars()
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content.Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneTask(params["id"])
}

func DeletAllTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content.Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count:= deleteAllTask()
	json.NewEncoder(w).Encode(count)

}

func getAllTask() []primitive.M{
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()){
		var result bson.M
		e := cur.Decode(&result)
		if e!=nil{
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err!=nil{
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}

func taskComplete(task string){
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"_status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Modified Count: ", result.ModifiedCount)
}

func insertOneTask(task models.TodoList){
	insertResult, err := collection.InsertOne(context.Background(), task)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
}

func undoTask(task string){
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"_status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Modified Count: ", result.ModifiedCount)
}

func deleteOneTask(task string){
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Deleted Document: ", deleteResult.DeletedCount)
}

func deleteAllTask() int64{
	d, err : = collection.DeleteMany(context.Background(), bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Deleted Document: ", d.DeletedCount)
	return d.DeletedCount
}