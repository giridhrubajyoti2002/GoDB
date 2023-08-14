package controller

import (
	"godb/service"
	"net/http"
)

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("POST method required"))
		return
	}
	cluster := r.FormValue("cluster")
	collection := r.FormValue("collection")
	statusCode, responseBody := http.StatusOK, "Collection created successfully"
	if len(cluster) == 0 || len(collection) == 0 {
		message := "Parameters("
		if len(cluster) == 0 {
			message += " cluster,"
		}
		if len(collection) == 0 {
			message += " collection"
		}
		message += ") missing !"
		statusCode, responseBody = http.StatusBadRequest, message
	} else if err := service.CreateCollection(cluster, collection); err != nil {
		statusCode, responseBody = http.StatusBadRequest, err.Error()
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(responseBody))
}
func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("DELETE method required"))
		return
	}
	cluster := r.FormValue("cluster")
	collection := r.FormValue("collection")
	statusCode, responseBody := http.StatusOK, "Collection deleted successfully"
	if len(cluster) == 0 || len(collection) == 0 {
		message := "Parameters("
		if len(cluster) == 0 {
			message += " cluster,"
		}
		if len(collection) == 0 {
			message += " collection"
		}
		message += ") missing !"
		statusCode, responseBody = http.StatusBadRequest, message
	} else if err := service.DeleteCollection(cluster, collection); err != nil {
		statusCode, responseBody = http.StatusBadRequest, err.Error()
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(responseBody))
}
