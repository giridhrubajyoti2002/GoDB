package repository

import (
	"fmt"
	"godb/db"
	"io/fs"
	"path/filepath"
	"strings"
)

func validateObjectId(value string) error {
	size := len(value)
	if value[size-1:] != ")" {
		return fmt.Errorf("Invalid value: %s. ')' missing at end", value)
	} else if len(value[7:size-1]) != 36 {
		return fmt.Errorf("ObjectId must be of 36 characters. Got '%d' characters", size-8)
	}
	return nil
}
func FindObjectDocuments(cluster string, bodyObj map[string]interface{}) error {
	if !Driver.DirExist(filepath.Join(db.StorageDir, cluster)) {
		return fmt.Errorf("Cluster: '%s' Not Found", cluster)
	}
	for _, v := range bodyObj {
		value := strings.ToLower(v.(string))
		size := len(value)
		if size >= 7 && value[0:7] == "object(" {
			if err := validateObjectId(value); err != nil {
				return err
			}
			_id := value[7 : size-1]
			fileName := _id + ".json"
			fileFound := false
			err := filepath.WalkDir(filepath.Join(db.StorageDir, cluster), func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				} else if !d.IsDir() && d.Name() == fileName {
					fileFound = true
					return fmt.Errorf("File Found")
				}
				return nil
			})
			if !fileFound {
				if err == nil {
					return fmt.Errorf("Object Not Found with _id: '%s' in Cluster: '%s'", _id, cluster)
				}
				return err
			}
		}
	}
	return nil
}

func InsertDocument(cluster string, collection string, _id string, content []byte) error {
	collectionMutex := Driver.GetOrCreateMutex(cluster, collection, _id)
	collectionMutex.Lock()
	defer collectionMutex.Unlock()
	if !Driver.DirExist(filepath.Join(db.StorageDir, cluster)) {
		return fmt.Errorf("Cluster: '%s' Not Found", cluster)
	} else if !Driver.DirExist(filepath.Join(db.StorageDir, cluster, collection)) {
		return fmt.Errorf("Collection: '%s' Not Found in Cluster: '%s'", collection, cluster)
	}
	return Driver.WriteFile(filepath.Join(db.StorageDir, cluster, collection, _id)+".json", content)
}
func FetchDocument(cluster string, collection string, _id string) ([]byte, error) {
	collectionMutex := Driver.GetOrCreateMutex(cluster, collection, _id)
	collectionMutex.Lock()
	defer collectionMutex.Unlock()
	if !Driver.DirExist(filepath.Join(db.StorageDir, cluster)) {
		return nil, fmt.Errorf("Cluster: '%s' Not Found", cluster)
	} else if !Driver.DirExist(filepath.Join(db.StorageDir, cluster, collection)) {
		return nil, fmt.Errorf("Collection: '%s' Not Found in Cluster: '%s'", collection, cluster)
	}
	return Driver.FetchFile(filepath.Join(db.StorageDir, cluster, collection, _id) + ".json")
}
func UpdateDocument(cluster string, collection string, _id string, content []byte) error {
	collectionMutex := Driver.GetOrCreateMutex(cluster, collection, _id)
	collectionMutex.Lock()
	defer collectionMutex.Unlock()
	if !Driver.DirExist(filepath.Join(db.StorageDir, cluster)) {
		return fmt.Errorf("Cluster: '%s' Not Found", cluster)
	} else if !Driver.DirExist(filepath.Join(db.StorageDir, cluster, collection)) {
		return fmt.Errorf("Collection: '%s' Not Found in Cluster: '%s'", collection, cluster)
	}
	return Driver.WriteFile(filepath.Join(db.StorageDir, cluster, collection, _id)+".json", content)
}
func DeleteDocument(cluster string, collection string, _id string) error {
	collectionMutex := Driver.GetOrCreateMutex(cluster, collection, _id)
	collectionMutex.Lock()
	defer collectionMutex.Unlock()
	if !Driver.DirExist(filepath.Join(db.StorageDir, cluster)) {
		return fmt.Errorf("Cluster: '%s' Not Found", cluster)
	} else if !Driver.DirExist(filepath.Join(db.StorageDir, cluster, collection)) {
		return fmt.Errorf("Collection: '%s' Not Found in Cluster: '%s'", collection, cluster)
	}
	return Driver.DeleteFile(filepath.Join(db.StorageDir, cluster, collection) + ".json")
}
