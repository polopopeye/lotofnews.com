package main

import (
	"context"
	"log"

	"lotofnews.com/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//importpath

func main() {
	// models.CreateConnection()
	// result := models.ExistTable("newsorigin")
	// fmt.Println(result)
	//
	// models.Ping()
	// models.CloseConnection()

	//-------------
	c := models.GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to mongo database", err)
	} else {
		log.Println("Connected to mongodb!")
	}
	// 	//one := ReturnOneHero(c, bson.M{})
	// 	hero := models.ReturnOneHero(c, bson.M{"name": "Vision"})
	// 	log.Println(hero.Name, hero.Alias, hero.Signed)
	// //	log.Println("--------------------------")
	// 	newhero := models.Hero{Name: "Stephen Strange", Alias: "Doctor Strange", Signed: true}
	// 	insertedID := models.InsertNewHero(c, newhero)
	// 	log.Println(insertedID)
	// 	newhero = models.ReturnOneHero(c, bson.M{"alias": "Doctor Strange"})
	// 	log.Println(newhero.Name, newhero.Alias, newhero.Signed)
	// 	log.Println("--------------------------")

	//heroes := ReturnAllHeroes(c, bson.M{"signed": true})
	articleshome := models.LastArticles(c, bson.M{})
	for _, article := range articleshome {
		log.Println(article.Lang, article.Titulo, article.Content, article.Fecha, article.Source, article.Cat)
	}

}

// UID, Lang, Titulo, Content, Fecha, Source
//
//
// Categorias, Ocio(37), Viajes(7),
//
// UID va a ser lang+nombrecat+UIDpost -
// UID, Lang, NombreCat, UIDpost, Fecha,
//
//
// UID, Lang, NombreCat, UIDpost, Fecha,
