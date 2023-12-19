package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/yoaz/movie-theater-api/api/Models"
)

/* ------------------------------------ API External --------------------------------------*/

// Get all movies
func (server *Server) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	var resp models.Response
	var movie models.Movie

	movies, err := movie.GetAllMovies(server.DB)
	if err != nil {
		resp.BadResponse(w, http.StatusInternalServerError, "failed to fetch Movies from db", err)
	}

	// Craft a layout resposne
	resp.OKResponse(w, http.StatusOK, "success", map[string]interface{}{"data": movies})
	return
}


// Get one movie
func (server *Server) GetMovieByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	movieID := params["id"]

	var resp models.Response
	var movie models.Movie

	fetchedMovie, err := movie.GetMovieByID(movieID, server.DB)
	if err != nil {
		resp.BadResponse(w, http.StatusInternalServerError, "couldn't fetch Movie from DB", err)
		return
	}

	// Craft a layout response
	resp.OKResponse(w, http.StatusOK, "success", map[string]interface{}{"data": fetchedMovie})
	return
}


// Create a movie
func (server *Server) CreateOneMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var resp models.Response
	var movie models.Movie

	//Prepare for insertion
	err := movie.Prepare(r); 

	//TODO: Validate the request body using validator

	// In case decoding failed
	if err != nil {
		log.Fatalf("There was an error decoding body request! %s", err)
	}

	// Call mongoDB associated helper
	insertedMovie, err := movie.InsertOneMovie(server.DB)
	if err != nil{
		resp.BadResponse(w, http.StatusInternalServerError, "failed to insert Movie to db", err)
		return
	}

	// Craft a layout response
	resp.OKResponse(w, http.StatusOK, "success", map[string]interface{}{"data": insertedMovie})
	return
}

// Upodate movie by ID
func (server *Server) UpdateMovieByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	movieID := params["id"]

	var resp models.Response
	var movie models.Movie

	//Prepare for insertion
	err := movie.BeforeUpdate(r); 

	//TODO: Validate the request body using validator

	// In case decoding failed
	if err != nil {
		log.Fatalf("There was an error decoding body request! %s", err)
	}

	// Call mongoDB associated helper
	updatedMovie, err := movie.EditMovieByID(movieID, server.DB)
	if err != nil{
		resp.BadResponse(w, http.StatusInternalServerError, "failed to insert Movie to db", err)
		return
	}

	// Craft a layout response
	resp.OKResponse(w, http.StatusOK, "success", map[string]interface{}{"data": updatedMovie})
	return
}

// Upodate movie by ID
func (server *Server) DeleteMovieByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	movieID := params["id"]

	var resp models.Response
	var movie models.Movie

	// Call mongoDB associated helper
	deleteMovieCount, err := movie.DeleteMovieByID(movieID, server.DB)
	if err != nil{
		resp.BadResponse(w, http.StatusInternalServerError, "failed to delete Movie from db", err)
		return
	}

	// Craft a layout response
	dataMessage := fmt.Sprintf("%d Movie was deleted, ID is: %s ", deleteMovieCount, movieID)
	resp.OKResponse(w, http.StatusOK, "success", map[string]interface{}{"data": dataMessage})
	return
}

