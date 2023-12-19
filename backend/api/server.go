package api

import (
	"os"

	controllers "github.com/yoaz/movie-theater-api/api/Controllers"
)


var Server = controllers.Server{}

//Run API
func APIRun(){
	//Load envoirmental variables
	controllers.LoadEnvVars()

	Server.InitDBRouter(os.Getenv("MONGO_DB_CONNECTION_STRING"), os.Getenv("MONGO_DB_NAME"))
	Server.RunServer(":" + os.Getenv("PORT"), os.Getenv("PORT"))
}	