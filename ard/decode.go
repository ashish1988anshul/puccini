package ard

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

func DecodeJson(reader io.Reader, locate bool) (Map, Locator, error) {
	data := make(Map)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&data); err == nil {
		return EnsureMap(data), nil, nil
	} else {
		return nil, nil, err
	}
}

func DecodeYaml(reader io.Reader, locate bool) (Map, Locator, error) {
	var data Map
	var locator Locator
	var node yaml.Node

	decoder := yaml.NewDecoder(reader)
	if err := decoder.Decode(&node); err == nil {
		if value, err := DecodeYamlNode(&node); err == nil {
			var ok bool
			if data, ok = value.(Map); ok {
				if locate {
					locator = NewYamlLocator(&node)
				}
			} else {
				return nil, nil, fmt.Errorf("YAML content is a \"%T\" instead of a map", value)
			}
		} else {
			return nil, nil, err
		}
	} else {
		return nil, nil, err
	}

	// We do not need to call EnsureMap because DecodeYamlNode takes care of it
	return data, locator, nil
}

func DecodeXml(reader io.Reader, locate bool) (Map, Locator, error) {
	data := make(Map)
	decoder := xml.NewDecoder(reader)
	if err := decoder.Decode(&data); err == nil {
		return EnsureMap(data), nil, nil
	} else {
		return nil, nil, err
	}
}
