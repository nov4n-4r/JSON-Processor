package util

import (
	"encoding/json"
	"errors"
)

type JsonStructure struct {
	key   string
	value string
}

var NilDestinationError = errors.New("nil destination error")

var InvalidJSONError = errors.New("invalid JSON structure")

func JsonValidation(data string) (bool, error) {

	isValid := json.Valid([]byte(data))

	return isValid, nil

}

func JsonCheck(src *[]byte, filters []JsonStructure) (bool, error) {

	if src == nil {
		return false, NilDestinationError
	}

	data := make(map[string]interface{})
	json.Unmarshal(*src, &data)

	for i := 0; i < len(filters); i++ {

		filt := filters[i]

		if !json.Valid([]byte(filt.value)) {
			return false, InvalidJSONError
		}

		if data[filt.key] != filt.value {
			return false, nil
		}

	}

	return true, nil

}

func JsonMap(src *[]byte, fields []string) error {

	if src == nil {
		return NilDestinationError
	}

	data := make(map[string]interface{})
	result := make(map[string]interface{})

	if err := json.Unmarshal(*src, &data); err != nil {
		return err
	}

	for i := 0; i < len(fields); i++ {
		result
	}

	return nil

}
