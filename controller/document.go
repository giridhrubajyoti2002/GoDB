package controller

import (
	"encoding/json"
	"godb/service"
	"io"
	"net/http"
	"strings"
)

func parseRequestBody(r *http.Request) (string, string, string, map[string]interface{}, error) {
	bodyContent, err := io.ReadAll(r.Body)
	if err != nil {
		return "", "", "", nil, err
	}
	defer r.Body.Close()
	var bodyMap map[string]interface{}
	if err := json.Unmarshal(bodyContent, &bodyMap); err != nil {
		return "", "", "", nil, err
	}
	cluster, collection, _id := "", "", ""
	value, ok := bodyMap["cluster"]
	if ok {
		cluster = value.(string)
	}
	value, ok = bodyMap["collection"]
	if ok {
		collection = value.(string)
	}
	value, ok = bodyMap["_id"]
	if ok {
		_id = value.(string)
	}
	var bodyObj interface{}
	for k := range bodyMap {
		if k != "cluster" && k != "collection" && k != "_id" {
			bodyObj = bodyMap[k]
		}
	}
	cluster, collection = strings.ToLower(cluster), strings.ToLower(collection)
	if bodyObj != nil {
		return cluster, collection, _id, bodyObj.(map[string]interface{}), nil
	}
	return cluster, collection, _id, nil, nil
}

func InsertDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("POST method required"))
		return
	}
	statusCode, responseType, responseBody := http.StatusOK, "", "Document inserted"
	cluster, collection, _, bodyObj, err := parseRequestBody(r)
	if err != nil {
		statusCode, responseBody = http.StatusInternalServerError, err.Error()
	} else if len(cluster) == 0 || len(collection) == 0 || bodyObj == nil {
		message := "Parameters("
		if len(cluster) == 0 {
			message += " cluster,"
		}
		if len(collection) == 0 {
			message += " collection,"
		}
		if bodyObj == nil {
			message += " document"
		}
		message += ") missing !"
		statusCode, responseType, responseBody = http.StatusBadRequest, "", message
	} else {
		document, err := service.InsertDocument(cluster, collection, bodyObj)
		statusCode, responseType, responseBody = http.StatusOK, "json", document
		if err != nil {
			statusCode, responseType, responseBody = http.StatusBadRequest, "", err.Error()
		}
	}
	if responseType == "json" {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(responseBody))
}
func UpdateDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("PATCH method required"))
		return
	}
	statusCode, responseType, responseBody := http.StatusOK, "", "Document updated"
	cluster, collection, _id, bodyObj, err := parseRequestBody(r)
	if err != nil {
		statusCode, responseBody = http.StatusInternalServerError, err.Error()
	} else if len(cluster) == 0 || len(collection) == 0 || len(_id) == 0 || bodyObj == nil {
		message := "Parameters("
		if len(cluster) == 0 {
			message += " cluster,"
		}
		if len(collection) == 0 {
			message += " collection,"
		}
		if len(_id) == 0 {
			message += " _id,"
		}
		if bodyObj == nil {
			message += " document"
		}
		message += ") missing !"
		statusCode, responseType, responseBody = http.StatusBadRequest, "", message
	} else {
		document, err := service.UpdateDocument(cluster, collection, _id, bodyObj)
		statusCode, responseType, responseBody = http.StatusOK, "json", document
		if err != nil {
			statusCode, responseType, responseBody = http.StatusBadRequest, "", err.Error()
		}
	}
	if responseType == "json" {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(responseBody))
}
func FetchDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("GET method required"))
		return
	}
	statusCode, responseType, responseBody := http.StatusOK, "", "Document fetched"
	cluster := r.FormValue("cluster")
	collection := r.FormValue("collection")
	_id := r.FormValue("_id")
	if len(cluster) == 0 || len(collection) == 0 || _id == "" {
		message := "Parameters("
		if len(cluster) == 0 {
			message += " cluster,"
		}
		if len(collection) == 0 {
			message += " collection,"
		}
		if _id == "" {
			message += " _id"
		}
		message += ") missing !"
		statusCode, responseType, responseBody = http.StatusBadRequest, "", message
	} else {
		document, err := service.FetchDocument(cluster, collection, _id)
		statusCode, responseType, responseBody = http.StatusOK, "json", document
		if err != nil {
			statusCode, responseType, responseBody = http.StatusBadRequest, "", err.Error()
		}
	}
	if responseType == "json" {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(responseBody))
}
func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("DELETE method required"))
		return
	}
	statusCode, responseBody := http.StatusOK, "Document deleted"
	cluster := r.FormValue("cluster")
	collection := r.FormValue("collection")
	_id := r.FormValue("_id")
	if len(cluster) == 0 || len(collection) == 0 || len(_id) == 0 {
		message := "Parameters("
		if len(cluster) == 0 {
			message += " cluster,"
		}
		if len(collection) == 0 {
			message += " collection,"
		}
		if len(_id) == 0 {
			message += " _id,"
		}
		message += ") missing !"
		statusCode, responseBody = http.StatusBadRequest, message
	} else if err := service.DeleteDocument(cluster, collection, _id); err != nil {
		statusCode, responseBody = http.StatusBadRequest, err.Error()
	}
	w.WriteHeader(statusCode)
	w.Write([]byte(responseBody))
}
