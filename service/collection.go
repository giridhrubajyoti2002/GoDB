package service

import (
	"godb/repository"
)

func CreateCollection(cluster string, collection string) error {
	return repository.CreateCollection(cluster, collection)
}
func DeleteCollection(cluster string, collection string) error {
	return repository.CreateCollection(cluster, collection)
}
