package service

import (
	"godb/repository"
)

func CreateCluster(cluster string) error {
	return repository.CreateCluster(cluster)
}
func DeleteCluster(cluster string) error {
	return repository.DeleteCluster(cluster)
}
