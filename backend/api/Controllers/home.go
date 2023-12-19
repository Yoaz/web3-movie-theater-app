package controllers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	resp := "<h1>Welcome To Movie-Theater API<h1>"
	w.Write([]byte(resp))
	return
}