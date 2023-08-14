# GoDB

A NoSQL Database written in Go.

## Key Features

- Mutex is used for synchronization
- JSON files are used to store documents

## How To Use

- Clone the repository : `https://github.com/giridhrubajyoti2002/GoDB.git`
- Open GoDB project in any IDE
- Open terminal and run below command :
  - Install required packages : `go get`
  - Start the server : `go run main.go`
- Open any tools like Postman or Insomnia to make HTTP requests to the APIs
- Default server address : `http://localhost:8080`
- You can change the server `host` and `port` in `main.go` file

## APIs

#### Cluster

- /cluster/create
  - Definition : Creates a cluster/database
  - Request Method : POST
  - Query Params:
    - cluster : cluster_name
- /cluster/delete
  - Definition : Deletes the cluster/database
  - Request Method : DELETE
  - Query Params :
    - cluster : cluster_name

#### Collection

- /collection/create
  - Definition : Creates a collection
  - Request Method : POST
  - Query Params :
    - cluster : cluster_name
    - collection : collection_name
- /collection/delete
  - Definition : Deletes a collection
  - Request Method : DELETE
  - Query Params :
    - cluster : cluster_name
    - collection : collection_name

##### Document

- /document/insert
  - Definition : Inserts a document
  - Request Method : POST
  - Request Body : <br>
  ```
  {
      "cluster" : "cluster_name"
      "collection" : "collection_name"
      "document" : {
          "key1" : "value1",
          "key2" : {                                                // Embedded Objects
              "key3" : "value3"
          }
          "key4" : "Object(cbdae53a-fc29-4ce4-a57e-454c40cd5753)"    // Object(ObjectId), Mapped Object, Object should be in same cluster
      }
  }
  ```
  - Response Body : <br>
  ```
  {
      "__createdAt" : "Time()"
      "_id" : "object_id"
      "key1" : "value1",
      "key2" : {                                                // Embedded Objects
          "key3" : "value3"
      }
      "key4" : "Object(cbdae53a-fc29-4ce4-a57e-454c40cd5753)"    // Object(ObjectId), Mapped Object
  }
  ```
- /document/fetch
  - Definition : Fetch a document
  - Request Method : GET
  - Query Params :
    - cluster : cluster_name
    - collection : collection_name
    - \_id : object_id
  - Response Body : <br>
  ```
  {
      "__createdAt" : "Time()"
      "__updatedAt" : "Time()"
      "_id" : "object_id"
      "key1" : "value1",
      "key2" : {                                                // Embedded Objects
          "key3" : "value3"
      }
      "key4" : "Object(cbdae53a-fc29-4ce4-a57e-454c40cd5753)"    // Object(ObjectId), Mapped Object
  }
  ```
- /document/update
  - Definition : Updates a document
  - Request Method : PATCH
  - Request Body : <br>
  ```
  {
      "cluster" : "cluster_name"
      "collection" : "collection_name"
      "_id" : "object_id"
      "document" : {
          "key1" : "value1",
          "key2" : {                                                // Embedded Objects
              "key3" : "value3"
          }
          "key4" : "Object(cbdae53a-fc29-4ce4-a57e-454c40cd5753)"    // Object(ObjectId), Mapped Object, Object should be in same cluster
      }
  }
  ```
  - Response Body : <br>
  ```
  {
      "__createdAt" : "Time()"
      "__updatedAt" : "Time()"
      "_id" : "object_id"
      "key1" : "value1",
      "key2" : {                                                // Embedded Objects
          "key3" : "value3"
      }
      "key4" : "Object(cbdae53a-fc29-4ce4-a57e-454c40cd5753)"    // Object(ObjectId), Mapped Object
  }
  ```
- /document/delete
  - Definition : Deletes a document
  - Request Method : DELETE
  - Query Params :
    - cluster : cluster_name
    - collection : "collection_name
    - \_id : object_id
