package models

import (
	"context"
	"encoding/json"
	"errors"
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
// Constants for number of rows and columns
const (
	NumRows    = 10
	NumColumns = 10
)

type Movie struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
	Description string `json:"description"`
	Duration time.Duration `json:"duration"`
	AvailableSeats  [NumRows][NumColumns]bool `json:"available_seats" bson:"available_seats"`
	Active bool	`json:"active"`
}

/* -------------------------------- Helpers ----------------------------------*/

func (m *Movie) Prepare(r *http.Request) error {
	// Custom jackpot type to unmarshell time correctly
	type parseType struct {
		ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Name string `json:"name"`
		Description string `json:"description"`
		Duration string	`json:"duration"`
		AvailableSeats  [NumRows][NumColumns]bool `json:"available_seats" bson:"available_seats"`
		Active bool	`json:"active"`
	}

	var res parseType

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return err
	}

	// Parse string duration
	parsedDuration, err := time.ParseDuration(res.Duration)
	if err != nil{
		return err
	}

	// Parsing success
	m.ID = res.ID
	m.Name = res.Name
	m.Description = res.Description
	m.Duration = parsedDuration
	m.AvailableSeats = res.AvailableSeats
	m.Active = true
	
	return nil
}

// Before updating movie struct
func (m *Movie) BeforeUpdate(r *http.Request) error {
    decoder := json.NewDecoder(r.Body)
    var updateFields map[string]interface{}
    if err := decoder.Decode(&updateFields); err != nil {
        return err
    }

    // Loop through the updateFields and update the corresponding fields in the Movie object
    for key, value := range updateFields {
        switch key {
        case "name":
            m.Name = value.(string)
		case "description":
            m.Description = value.(string)
        case "duration":
            // Parse string duration
			parsedDuration, err := time.ParseDuration(value.(string))
				if err != nil{
					return err
				}
				m.Duration = parsedDuration
        case "active":
            m.Active = value.(bool)
		case "available_seats":
			m.AvailableSeats = value.([10][10]bool)
        }
    }

    return nil
}


/* -------------------------------- DB Actions ----------------------------------*/

// Insert 1 record
func (m *Movie) InsertOneMovie(db *mongo.Database) (*Movie, error) {
	collection := db.Collection(os.Getenv("MONGO_DB_MOVIE_COL_NAME"))
	
	// Calling mongoDB insert one action
	inserted, err := collection.InsertOne(context.Background(), m);

	// In case inserted operation fails
	if err != nil {
		return &Movie{}, err
	}

	// Success
	fmt.Println("Inserted 1 Movie to Movie Theater DB with ID: ", inserted.InsertedID)

	return m, nil
}

// Get one record
func (m *Movie) GetMovieByID(movieID string, db *mongo.Database) (*Movie, error) {
	// Establish connection pointer
	collection := db.Collection(os.Getenv("MONGO_DB_MOVIE_COL_NAME"))
	
	// String ID to mongoDB object ID
	id, err := primitive.ObjectIDFromHex(movieID)

	// In case of ID conversation fails
	if err != nil {
		return &Movie{}, err
	}

	// Establish filter 
	filter := bson.M{"_id": id}
	
	// Find & Decode result to m struct
	err = collection.FindOne(context.Background(), filter).Decode(&m)

	// In case of get record failed
	if err != nil {
		return &Movie{}, err
	}

	return m, err
}

// Get all movies
func (m *Movie) GetAllMovies(db *mongo.Database) ([]Movie, error) {
	collection := db.Collection(os.Getenv("MONGO_DB_MOVIE_COL_NAME"))

	// Getting cursor to all documents in DB
	cur, err := collection.Find(context.Background(), bson.D{{}})

	// In case fetching all records fails
	if err != nil {
		return []Movie{}, err
	}

	var movies []Movie // Will hold all fetched movies from DB

	// Loop through all records and save in local var
	for cur.Next(context.Background()) {
		var movie Movie // Will hold current pointed db document
		err := cur.Decode(&movie)

		// In case decode fails
		if err != nil {
			log.Fatalf("There was an error trying to decode the record, err: %s", err)
		}

		movies = append(movies, movie)
	}
	return movies, nil
}


/* Delete one NFT */
func (nft *Movie) DeleteMovieByID(movieID string, db *mongo.Database) (int64, error) {
	var err error

	collection := db.Collection(os.Getenv("MONGO_DB_MOVIE_COL_NAME"))
	id, err := primitive.ObjectIDFromHex(movieID)

	// In case of id convertion failed
	if err != nil{
		return 0, err
	}

	// In case of id converation successful
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	//In case error deleting from db
	if err != nil{
		return 0, err
	}

	// In case 0 documents deleted, means no such document found
	if deleteCount.DeletedCount == 0 {
		// Construct error string to printed out and send as error
		errorMessage := fmt.Sprintf("%d Movie was deleted, incorrect ID number, no such document ID: %s", deleteCount.DeletedCount, movieID)
		fmt.Printf(errorMessage)
		err := errors.New(errorMessage)		

		return deleteCount.DeletedCount, err
	}

	fmt.Printf("%d Movie was deleted, ID is: %s ", deleteCount.DeletedCount, movieID)

	return deleteCount.DeletedCount, nil
}

/* Edit one movie */
func (movie *Movie) EditMovieByID(movieID string, db *mongo.Database) (int64, error) {
	var err error

	collection := db.Collection(os.Getenv("MONGO_DB_MOVIE_COL_NAME"))
	id, err := primitive.ObjectIDFromHex(movieID)

	// In case of id convertion failed
	if err != nil{
		return 0, err
	}

	// Convert the movie struct to a map to use it as the updateFields
	updateFields := map[string]interface{}{
		"name": movie.Name,
		"duration":  movie.Duration,
		"active":    movie.Active,
	}

	// In case of id converation successful
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateFields}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	// In case of update record failed
	if err != nil {
		return result.ModifiedCount, err	
	}

	// In case of update record success
	fmt.Println("Modified one record success, count: ", result.ModifiedCount)
	
	return result.ModifiedCount, nil
}