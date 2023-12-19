package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/* -------------------------------- Types ----------------------------------*/

// Schedule represents a scheduled showing of movies for a specific date.
type Schedule struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Date           time.Time          `json:"date" bson:"date"`
	MovieSeats	[]Movie   `json:"movies" bson:"movies"`
}

/* -------------------------------- Helpers ----------------------------------*/

func (s Schedule) Prepare(r *http.Request) error {
	// Custom schedule type to unmarshel time correctly
	type parseType struct {
		ID primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
		Date string `json:"date" bson:"date"`
		MovieSeats	[]Movie   `json:"movies_schedule" bson:"movies_schedule"`
	}

	var res parseType

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return err
	}

	// Parse string date to RFC3339
	parsedDate, err := time.Parse(time.RFC3339, res.Date)
	if err != nil {
		return err
	}

	// Parsing success
	s.ID = res.ID
	s.Date = parsedDate
	s.MovieSeats = res.MovieSeats

	return nil
}


/* -------------------------------- DB Actions ----------------------------------*/


// Get movies by date provided
func (s *Schedule) GetMoviesByDate(db *mongo.Database, scheduleDate primitive.DateTime) (*Schedule, error){
	var err error
	
	collection := db.Collection(os.Getenv("MONGO_DB_SCHEDULE_COL_NAME"))

	filter := bson.M{"date": scheduleDate} 

	err = collection.FindOne(context.Background(), filter).Decode(&s)

	// In case of no such schedule in the provided date in DB
	if err == mongo.ErrNoDocuments {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return s, nil
}


// Insert one schedule to database
func (s *Schedule) InsertOneSchedule(db *mongo.Database) (*Schedule, error) {
	var err error

	collection := db.Collection(os.Getenv("MONGO_DB_SCHEDULE_COL_NAME"))
	inserted, err := collection.InsertOne(context.Background(), s)

	// In case inserted operation failed
	if err != nil {
		return &Schedule{}, err
	}

	//In case of success
	fmt.Println("Inserted 1 Schedule details with ID: ", inserted.InsertedID)
	
	return s, nil 
}

// Get all movies
func (s *Schedule) GetAllSchedules(db *mongo.Database) ([]Schedule, error) {
	collection := db.Collection(os.Getenv("MONGO_DB_SCHEDULE_COL_NAME"))

	// Getting cursor to all documents in DB
	cur, err := collection.Find(context.Background(), bson.D{{}})

	// In case fetching all records fails
	if err != nil {
		return []Schedule{}, err
	}

	var schedules []Schedule // Will hold all fetched movies from DB

	// Loop through all records and save in local var
	for cur.Next(context.Background()) {
		var schedule Schedule // Will hold current pointed db document
		err := cur.Decode(&schedule)

		// In case decode fails
		if err != nil {
			log.Fatalf("There was an error trying to decode the record, err: %s", err)
		}

		schedules = append(schedules, schedule)
	}
	return schedules, nil
}