package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const urldatabase string = "mongodb://localhost:27017"
const database string = "lotofnewsgo"
const collection string = "heroes"

type Hero struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Signed bool   `json:"signed"`
}

func ReturnAllHeroes(client *mongo.Client, filter bson.M) []*Hero {
	var heroes []*Hero
	collection := client.Database(database).Collection(collection)
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
	collection := client.Database(database).Collection(collection)
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&hero)
	return hero
}

func InsertNewHero(client *mongo.Client, hero Hero) interface{} {
	collection := client.Database(database).Collection(collection)
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

func main() {
	c := GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	//one := ReturnOneHero(c, bson.M{})
	hero := ReturnOneHero(c, bson.M{"name": "Vision"})
	log.Println(hero.Name, hero.Alias, hero.Signed)
	log.Println("--------------------------")
	newhero := Hero{Name: "Stephen Strange", Alias: "Doctor Strange", Signed: true}
	insertedID := InsertNewHero(c, newhero)
	log.Println(insertedID)
	newhero = ReturnOneHero(c, bson.M{"alias": "Doctor Strange"})
	log.Println(newhero.Name, newhero.Alias, newhero.Signed)
	log.Println("--------------------------")

	//heroes := ReturnAllHeroes(c, bson.M{"signed": true})
	heroes := ReturnAllHeroes(c, bson.M{})
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed)
	}
}
