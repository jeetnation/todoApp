package middleware

import{
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
}

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

func GetAllTask(w http.ResponseWriter, r *http.Request){}

func CreateTask(w http.ResponseWriter, r *http.Request){}

func TaskComplete(w http.ResponseWriter, r *http.Request){}

func UndoTask(w http.ResponseWriter, r *http.Request){}

func DeleteTask(w http.ResponseWriter, r *http.Request){}

func DeletAllTask(w http.ResponseWriter, r *http.Request){}