package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const urldatabase string = "mongodb://localhost:27017"
const databasemongo string = "lotofnewsgo"
const collection string = "heroes"

type Hero struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Signed bool   `json:"signed"`
}

//structs by Type
type Articles struct {
	Uid      string   `json:"_id" bson:"_id,omitempty"`
	Lang     string   `json:"Lang"`
	Titulo   string   `json:"Titulo"`
	Content  string   `json:"Content"`
	Fecha    int64    `json:"Fecha"`
	Source   string   `json:"Source"`
	Cat      []string `json:"Cat"`
	Img      string   `json:"Img"`
	Featured bool     `json:"Featured"`
	//published bool `json:"published"` //another option to take in mind like a filter not implemented yet
}

//function to call to list LastArticles created per page of input given in every call.
//options on const
const articlestoload int64 = 9
const collectionarticles string = "OriginArticles"

//end options start func
func LastArticles(client *mongo.Client, filter bson.M) []*Articles {

	var listarticles []*Articles
	collection := client.Database(databasemongo).Collection(collectionarticles)
	findOptions := options.Find()
	findOptions.SetLimit(articlestoload)

	findOptions.SetSort(bson.M{"_id": -1})

	lastart, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for lastart.Next(context.TODO()) {
		var articles Articles
		err = lastart.Decode(&articles)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		listarticles = append(listarticles, &articles)
	}

	return listarticles
}

//end func lastarticles by page with filter

func ReturnAllHeroes(client *mongo.Client, filter bson.M) []*Hero {
	var heroes []*Hero
	collection := client.Database(databasemongo).Collection(collection)
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var hero Hero
		err = cur.Decode(&hero)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		heroes = append(heroes, &hero)
	}
	return heroes
}

func ReturnOneHero(client *mongo.Client, filter bson.M) Hero {
	var hero Hero
	collection := client.Database(databasemongo).Collection(collection)
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&hero)
	return hero
}

func InsertNewHero(client *mongo.Client, hero Hero) interface{} {
	collection := client.Database(databasemongo).Collection(collection)
	insertResult, err := collection.InsertOne(context.TODO(), hero)
	if err != nil {
		log.Fatalln("Error on inserting new Hero", err)
	}
	return insertResult.InsertedID
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(urldatabase)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
