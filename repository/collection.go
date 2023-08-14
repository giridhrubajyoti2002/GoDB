package repository

import (
	"fmt"
	"godb/db"
	"path/filepath"
)

func CreateCollection(cluster string, collection string) error {
	collectionMutex := Driver.GetOrCreateMutex(cluster, collection)
	collectionMutex.Lock()
	defer collectionMutex.Unlock()
	if !Driver.DirExist(filepath.Join(db.StorageDir, cluster)) {
		return fmt.Errorf("Cluster: '%s' Not Found", cluster)
	} else if Driver.DirExist(filepath.Join(db.StorageDir, cluster, collection)) {
		return fmt.Errorf("Collection: '%s' already exists in Cluster: '%s'", collection, cluster)
	}
	return Driver.CreateDir(filepath.Join(db.StorageDir, cluster, collection))
}
func DeleteCollection(cluster string, collection string) error {
	collectionMutex := Driver.GetOrCreateMutex(cluster, collection)
	collectionMutex.Lock()
	defer collectionMutex.Unlock()
	if !Driver.DirExist(filepath.Join(db.StorageDir, cluster)) {
		return fmt.Errorf("Cluster: '%s' Not Found", cluster)
	} else if !Driver.DirExist(filepath.Join(db.StorageDir, cluster, collection)) {
		return fmt.Errorf("Collection: '%s' Not Found in Cluster: '%s'", collection, cluster)
	}
	return Driver.DeleteDir(filepath.Join(db.StorageDir, cluster, collection))
}
