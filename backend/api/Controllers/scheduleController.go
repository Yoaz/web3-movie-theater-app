package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	models "github.com/yoaz/movie-theater-api/api/Models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Insert One Schedule Details to DB
func (server *Server) CreateOneSchedule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Method", "POST")

	var resp models.Response
	var Schedule models.Schedule

	//TODO: validate the request body using validator
	
	//Prepare for insertion
	err := Schedule.Prepare(r);

	// In case decoding failed
	if err != nil {
		log.Fatalf("There was an error decoding body request! %s", err)
		return
	}

	// Call mongoDB associated helper
	insertedSchedule, err := Schedule.InsertOneSchedule(server.DB)
	if err != nil{
		resp.BadResponse(w, http.StatusInternalServerError, "failed to insert Schedule to db", err)
		return
	} 

	// Craft a layout response
	resp.OKResponse(w, http.StatusOK, "success", map[string]interface{}{"data": insertedSchedule})
	return
}


// Get all movies by date
func (server *Server) GetMoviesByDate(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	params := mux.Vars(r)
	scheduleDate := params["date"]
	
	var resp models.Response
	var schedule models.Schedule

	// Parse string date to RFC3339
	expectedTimeLayout := "2006-01-02"
	parsedStringDate, err := time.Parse(expectedTimeLayout, scheduleDate)
	if err != nil {
		resp.BadResponse(w, http.StatusFailedDependency, "failed to cast srting date to RFC3339", err)
		return 
	}

	// Parse to primitive.DateTime MongoDB date type
	parsedDate := primitive.NewDateTimeFromTime(parsedStringDate)

	// Call mongoDB associated helper
	fetchedSchedule, err := schedule.GetMoviesByDate(server.DB, parsedDate)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.BadResponse(w, http.StatusBadRequest, "no Schedule records for provided date", err)
			return
		}

		resp.BadResponse(w, http.StatusInternalServerError, "failed to fetch Schedule from db", err)
		return

		}
		
	// Craft a layout resposne
	resp.OKResponse(w, http.StatusOK, "success", map[string]interface{}{"data": fetchedSchedule})
	return
}

// Get all movies
func (server *Server) GetAllSchedules(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	var resp models.Response
	var schedule models.Schedule

	schedules, err := schedule.GetAllSchedules(server.DB)
	if err != nil {
		resp.BadResponse(w, http.StatusInternalServerError, "failed to fetch Schedules from db", err)
	}

	// Craft a layout resposne
	resp.OKResponse(w, http.StatusOK, "success", map[string]interface{}{"data": schedules})
	return
}