package db

import (
	"os"
	"sync"
)

const SchemaFileName = "schema.go"
const StorageDir = "storage"

type Mutex struct {
	dbMutex    *sync.Mutex
	cluster    map[string]*sync.Mutex
	collection map[string]map[string]*sync.Mutex
	document   map[string]map[string]map[string]*sync.Mutex
}
type DbDriver struct {
	mutex *Mutex
}

var driver *DbDriver

func GetDriver() *DbDriver {
	if driver == nil {
		driver = &DbDriver{}
		driver.mutex = &Mutex{dbMutex: &sync.Mutex{}}
		driver.mutex.cluster = make(map[string]*sync.Mutex)
		driver.mutex.collection = make(map[string]map[string]*sync.Mutex)
		driver.mutex.document = make(map[string]map[string]map[string]*sync.Mutex)
	}
	return driver
}
func (d *DbDriver) GetOrCreateMutex(keys ...string) *sync.Mutex {
	d.mutex.dbMutex.Lock()
	defer d.mutex.dbMutex.Unlock()
	size := len(keys)
	var mutex *sync.Mutex
	if size == 3 {
		if d.mutex.document[keys[0]] == nil {
			d.mutex.document[keys[0]] = make(map[string]map[string]*sync.Mutex)
		}
		if d.mutex.document[keys[0]][keys[1]] == nil {
			d.mutex.document[keys[0]][keys[1]] = make(map[string]*sync.Mutex)
		}
		if d.mutex.document[keys[0]][keys[1]][keys[2]] == nil {
			d.mutex.document[keys[0]][keys[1]][keys[2]] = &sync.Mutex{}
		}
		mutex = d.mutex.document[keys[0]][keys[1]][keys[2]]
	} else if size == 2 {
		if d.mutex.collection[keys[0]] == nil {
			d.mutex.collection[keys[0]] = make(map[string]*sync.Mutex)
		}
		if d.mutex.collection[keys[0]][keys[1]] == nil {
			d.mutex.collection[keys[0]][keys[1]] = &sync.Mutex{}
		}
		mutex = d.mutex.collection[keys[0]][keys[1]]
	} else if size == 1 {
		if d.mutex.cluster[keys[0]] == nil {
			d.mutex.cluster[keys[0]] = &sync.Mutex{}
		}
		mutex = d.mutex.cluster[keys[0]]
	}
	return mutex
}

func (d *DbDriver) DirExist(dirPath string) bool {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) || err != nil {
		return false
	}
	return true
}
func (d *DbDriver) CreateDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0777)
}
func (d *DbDriver) DeleteDir(dirPath string) error {
	err := os.RemoveAll(dirPath)
	return err
}

func (d *DbDriver) FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) || err != nil {
		return false
	}
	return true
}
func (d *DbDriver) WriteFile(filePath string, content []byte) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	return err
}
func (d *DbDriver) FetchFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	var content = make([]byte, fileInfo.Size())
	_, err = file.Read(content)
	return content, err
}
func (d *DbDriver) DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

func (d *DbDriver) Write(filePath string, prefix string, json []byte) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	fileInfo, _ := os.Stat(filePath)
	_, err = file.WriteAt(json, fileInfo.Size()-2)
	_, err = file.WriteAt([]byte{3}, fileInfo.Size()) // TODO
	return nil
}

//	func (d *Driver) WriteSchema(schemaPath string, collection string, schema []byte) error {
//		file, err := os.OpenFile(schemaPath, os.O_APPEND, os.ModeAppend)
//		if err != nil {
//			return err
//		}
//		defer file.Close()
//		fileInfo, _ := os.Stat(schemaPath)
//		size := fileInfo.Size()
//		if bLen, err := file.WriteAt([]byte("type "+collection+" struct"), size); err == nil {
//			size += int64(bLen)
//		}
//		if _, err := file.WriteAt(schema, size); err != nil {
//			return err
//		}
//		return nil
//	}
