package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
	//`json:""`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//DeleteMovie
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies := append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	//File test data using chatgpt )):
	//movies = append(movies, Movie{ID:"1",Isbn: "2345",Title: "Movie one",Director: &Director{Firstname: "Peter", Latname: "Jackson"}})

	movies = append(movies, Movie{ID: "1", Isbn: "2345", Title: "Movie one", Director: &Director{Firstname: "Peter", Lastname: "Jackson"}})
	movies = append(movies, Movie{ID: "2", Isbn: "6789", Title: "Movie two", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "3", Isbn: "1234", Title: "Movie three", Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"}})
	movies = append(movies, Movie{ID: "4", Isbn: "5678", Title: "Movie four", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	movies = append(movies, Movie{ID: "5", Isbn: "9012", Title: "Movie five", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	movies = append(movies, Movie{ID: "6", Isbn: "3456", Title: "Movie six", Director: &Director{Firstname: "Martin", Lastname: "Scorsese"}})
	movies = append(movies, Movie{ID: "7", Isbn: "7890", Title: "Movie seven", Director: &Director{Firstname: "Stanley", Lastname: "Kubrick"}})
	movies = append(movies, Movie{ID: "8", Isbn: "2341", Title: "Movie eight", Director: &Director{Firstname: "Alfred", Lastname: "Hitchcock"}})
	movies = append(movies, Movie{ID: "9", Isbn: "6781", Title: "Movie nine", Director: &Director{Firstname: "Ridley", Lastname: "Scott"}})
	movies = append(movies, Movie{ID: "10", Isbn: "0123", Title: "Movie ten", Director: &Director{Firstname: "Francis", Lastname: "Ford Coppola"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Startting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
