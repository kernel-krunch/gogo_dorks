package main

import (
	"encoding/json"
	"io/ioutil"
)

func writeOutputToFile(output interface{}, filename string) error {
	jsonData, err := json.Marshal(output)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func readDorksFromFile(filename string) (map[string]string, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var dorks map[string]string
	err = json.Unmarshal(fileData, &dorks)
	if err != nil {
		return nil, err
	}

	return dorks, nil
}
