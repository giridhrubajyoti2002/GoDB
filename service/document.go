package service

import (
	"encoding/json"
	"godb/repository"
	"time"

	"github.com/google/uuid"
)

func InsertDocument(cluster string, collection string, bodyObj map[string]interface{}) (string, error) {
	if err := repository.FindObjectDocuments(cluster, bodyObj); err != nil {
		return "", err
	}
	_id := uuid.New()
	bodyObj["_id"] = _id
	bodyObj["__createdAt"] = time.Now()
	content, err := json.MarshalIndent(bodyObj, "", "\t")
	if err != nil {
		return "", err
	}
	err = repository.InsertDocument(cluster, collection, _id.String(), content)
	return string(content), err
}
func FetchDocument(cluster string, collection string, _id string) (string, error) {
	content, err := repository.FetchDocument(cluster, collection, _id)
	return string(content), err
}
func UpdateDocument(cluster string, collection string, _id string, bodyObj map[string]interface{}) (string, error) {
	if err := repository.FindObjectDocuments(cluster, bodyObj); err != nil {
		return "", err
	}
	content, err := repository.FetchDocument(cluster, collection, _id)
	if err != nil {
		return "", err
	}
	documentMap := make(map[string]interface{})
	err = json.Unmarshal(content, &documentMap)
	delete(bodyObj, "_id")
	delete(bodyObj, "__createdAt")
	for k, v := range bodyObj {
		if len(v.(string)) > 0 {
			documentMap[k] = v
		}
	}
	documentMap["__updatedAt"] = time.Now()
	content, err = json.MarshalIndent(documentMap, "", "\t")
	if err != nil {
		return "", err
	}
	err = repository.UpdateDocument(cluster, collection, _id, content)
	return string(content), err
}
func DeleteDocument(cluster string, collection string, _id string) error {
	return repository.DeleteDocument(cluster, collection, _id)
}
