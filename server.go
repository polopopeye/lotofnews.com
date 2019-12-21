package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ip2location/ip2location-go"

	"lotofnews.com/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Homeinterface struct {
	Nombre   string
	FechaNac int
	Email    string
	Titulo   string
	Desc     string
	Fecha    string
	Pais     string
}

type homedata struct {
	Featuredarticleshome []*models.Articles
	Articleshome         []*models.Articles
	Datos                Homeinterface
}

//structs by Type

var Inicio = template.Must(template.New("Inicio").ParseGlob("templates/*.html"))
var ip = "95.169.231.127"

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	ip2location.Open("./IP-COUNTRY.BIN")
	results := ip2location.Get_all(ip)
	current_time := time.Now().Local()
	datos := Homeinterface{"Kenneth", 1996, "kennethsuarez@gmx.com", "Home", "Desc", current_time.Format("02-01-2006"), results.Country_long}
	Inicio.ExecuteTemplate(w, "test", datos)
}

func Home(w http.ResponseWriter, r *http.Request) {

	c := models.GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to mongo database", err)
	}
	ip2location.Open("./IP-COUNTRY.BIN")
	results := ip2location.Get_all(ip)
	current_time := time.Now().Local()
	datos := Homeinterface{"Kenneth", 1996, "kennethsuarez@gmx.com", "Home", "Desc", current_time.Format("02-01-2006"), results.Country_long}
	articleshome := models.LastArticles(c, bson.M{})
	featuredarticleshome := models.LastArticles(c, bson.M{"Featured": true})

	datos2 := homedata{featuredarticleshome, articleshome, datos}

	//println(articleshome)
	// for _, article := range articleshome {
	// 	log.Println(article.Lang, article.Titulo, article.Content, article.Fecha, article.Source, article.Cat)
	// }

	w.Header().Set("Content-Type", "text/html")

	Inicio.ExecuteTemplate(w, "home", datos2)

}

func Cathome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cat := vars["cat"]

	w.Header().Set("Content-Type", "text/html")
	ip2location.Open("./IP-COUNTRY.BIN")
	results := ip2location.Get_all(ip)
	current_time := time.Now().Local()
	datos := Homeinterface{"Kenneth", 1996, "kennethsuarez@gmx.com", cat, "Desc", current_time.Format("02-01-2006"), results.Country_long}

	Inicio.ExecuteTemplate(w, "cat", datos)
}

func News(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cat := vars["cat"]
	id := vars["id"]

	w.Header().Set("Content-Type", "text/html")
	ip2location.Open("./IP-COUNTRY.BIN")
	results := ip2location.Get_all(ip)
	current_time := time.Now().Local()
	datos := Homeinterface{"Kenneth", 1996, "kennethsuarez@gmx.com", cat + " " + id, id, current_time.Format("02-01-2006"), results.Country_long}

	Inicio.ExecuteTemplate(w, "singlenew", datos)

}

func main() {
	staticFiles := http.FileServer(http.Dir("assets"))
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/test", Test).Methods("GET")
	r.Handle("/assets/{*}/{*}", http.StripPrefix("/assets/", staticFiles))
	//start cat and single
	r.HandleFunc("/{cat}", Cathome).Methods("GET")
	r.HandleFunc("/{cat}/{id:[0-9]+}", News).Methods("GET")
	// Bind to a port and pass our router in
	log.Println("Servidor Ejecuntandose y sirviendo en puerto 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
