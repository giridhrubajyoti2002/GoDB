package repository

import (
	"fmt"
	"godb/db"
	"path/filepath"
)

var Driver = db.GetDriver()

func CreateCluster(cluster string) error {
	collectionMutex := Driver.GetOrCreateMutex(cluster)
	collectionMutex.Lock()
	defer collectionMutex.Unlock()
	if Driver.DirExist(filepath.Join(db.StorageDir, cluster)) {
		return fmt.Errorf("Cluster: '%s' already exists", cluster)
	}
	return Driver.CreateDir(filepath.Join(db.StorageDir, cluster))
}
func DeleteCluster(cluster string) error {
	collectionMutex := Driver.GetOrCreateMutex(cluster)
	collectionMutex.Lock()
	defer collectionMutex.Unlock()
	if !Driver.DirExist(filepath.Join(db.StorageDir, cluster)) {
		return fmt.Errorf("Cluster: '%s' Not Found", cluster)
	}
	return Driver.DeleteDir(filepath.Join(db.StorageDir, cluster))
}
