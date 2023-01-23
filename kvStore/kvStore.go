package kvStore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type KeyValueStore struct {
	data map[string]string
}

func (kvs *KeyValueStore) Add(key, value string) {
	kvs.data[key] = value
}

func (kvs *KeyValueStore) Get(key string) (string, bool) {
	message, ok := kvs.data[key]
	return message, ok
}

func (kvs *KeyValueStore) Delete(key string) {
	delete(kvs.data, key)
}

func (kvs *KeyValueStore) Edit(key, value string) {
	kvs.data[key] = value
}

func (kvs *KeyValueStore) PrintList() {
	for key, value := range kvs.data {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}

func (kvs *KeyValueStore) GetList() map[string]string {
	return kvs.data
}

func (kvs *KeyValueStore) DeleteAll() {
	kvs.data = make(map[string]string)
}

func (kvs *KeyValueStore) Save(fileName string) error {
	// Convert the map to json
	jsonData, err := json.Marshal(kvs.data)
	if err != nil {
		return err
	}

	// Write json data to file
	return ioutil.WriteFile(fileName, jsonData, 0644)
}

func (kvs *KeyValueStore) Load(fileName string) error {
	// Read data from file
	jsonData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	// Unmarshal json data to map
	err = json.Unmarshal(jsonData, &kvs.data)
	return err
}
