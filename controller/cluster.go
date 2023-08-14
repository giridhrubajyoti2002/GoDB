package controller

import (
	"godb/repository"
	"net/http"
	"strings"
)

func CreateCluster(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("POST method required"))
		return
	}
	cluster := r.FormValue("cluster")
	statusCode, responseBody := http.StatusAccepted, "Cluster created successfully"
	if len(cluster) == 0 {
		statusCode, responseBody = http.StatusBadRequest, "Cluster name missing !"
	} else if err := repository.CreateCluster(cluster); err != nil {
		statusCode, responseBody = http.StatusBadRequest, err.Error()
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(responseBody))
}
func DeleteCluster(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("DELETE method required"))
		return
	}
	cluster := r.FormValue("cluster")
	statusCode, responseBody := http.StatusAccepted, "Cluster deleted successfully"
	if len(strings.TrimSpace(cluster)) == 0 {
		statusCode, responseBody = http.StatusBadRequest, "Cluster name missing !"
	} else if err := repository.DeleteCluster(cluster); err != nil {
		statusCode, responseBody = http.StatusBadRequest, err.Error()
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(responseBody))
}
