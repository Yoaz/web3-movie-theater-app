package controllers

//Initializing server routes
func (server *Server) InitRoutes() { 

	//Movie Routes
	server.Router.HandleFunc("/", Home)
	server.Router.HandleFunc("/api/movies", server.GetAllMovies).Methods("GET")
	server.Router.HandleFunc("/api/movie", server.CreateOneMovie).Methods("POST")
	server.Router.HandleFunc("/api/movie/{id}", server.GetMovieByID).Methods("GET")
	server.Router.HandleFunc("/api/movie/{id}", server.DeleteMovieByID).Methods("POST")
	server.Router.HandleFunc("/api/movie/{id}", server.UpdateMovieByID).Methods("PATCH")
	
	//Schedule Routes
	server.Router.HandleFunc("/api/schedule/{date}", server.GetMoviesByDate).Methods("GET")
	server.Router.HandleFunc("/api/schedule", server.CreateOneSchedule).Methods("POST")
	server.Router.HandleFunc("/api/schedules", server.GetAllSchedules).Methods("GET")
}