package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func GetEnvFile(env string) string {
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := path.Join(path.Join(dirname, "/config/properties/app.config.json"))
	return filePath
}

func GetEnvDBFile(env string) string {
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := path.Join(path.Join(dirname, "/config/properties/db/env."+env+".json"))
	return filePath
}

/**
 * Validate if provided key value is in app config json
 */
func ValidateEnvVar(key string, value string) (bool, error) {

	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(path.Join(dirname, "/config/properties/app.config.json")) // file.json has the json content
	if err != nil {
		log.Fatal(err)
	}

	blob, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	doc := make(map[string]interface{})

	if err := json.Unmarshal(blob, &doc); err != nil {
		log.Fatal(err)
	}
	if apiKey, contains := doc[key]; contains && apiKey == value {
		return true, nil
	} else {
		return false, errors.New("Invalid Api Key provided")
	}
}

/**
 * Function to get env variable from config file
 * @type {[type]}
 */
func GetEnvVar(key string) (string, error) {
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(path.Join(dirname, "/config/properties/app.config.json")) // file.json has the json content
	if err != nil {
		log.Fatal(err)
	}

	blob, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	doc := make(map[string]interface{})

	if err := json.Unmarshal(blob, &doc); err != nil {
		log.Fatal(err)
	}
	if value, contains := doc[key].(string); contains {
		return value, nil
	} else {
		return "", errors.New("Unable to access key from config file ")
	}
}
