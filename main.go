package main

import (
	"fmt"
	"godb/controller"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/cluster/create", controller.CreateCluster)
	mux.HandleFunc("/cluster/delete", controller.DeleteCluster)
	mux.HandleFunc("/collection/create", controller.CreateCollection)
	mux.HandleFunc("/collection/delete", controller.DeleteCollection)
	mux.HandleFunc("/document/insert", controller.InsertDocument)
	mux.HandleFunc("/document/fetch", controller.FetchDocument)
	mux.HandleFunc("/document/update", controller.UpdateDocument)
	mux.HandleFunc("/document/delete", controller.DeleteDocument)

	host, port := "localhost", "8080"

	server := http.Server{
		Addr:    host + ":" + port,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

}
