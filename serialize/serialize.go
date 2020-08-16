package serialize

import (
	"encoding/gob"
	"encoding/json"
	"log"
	"os"
)

func MapToStruct(m interface{}, i interface{}) error {
	mapString, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		log.Fatal(errMarshal)
		return errMarshal
	}
	errUnmarshal := json.Unmarshal(mapString, i)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
		return errUnmarshal
	}
	return nil
}

func WriteGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	encoder.Encode(object)
	return nil
}

func ReadGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(object)
	return nil
}
