package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func MergeAndMarshal(s1, s2 any) (string, error) {
	merged, err := mergeStructs(s1, s2)
	if err != nil {
		return "", err
	}
	jsonBytes, err := json.Marshal(merged)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func mergeStructs(s1, s2 any) (any, error) {
	v1 := reflect.ValueOf(s1)
	v2 := reflect.ValueOf(s2)

	if v1.Kind() == reflect.Pointer {
		v1 = v1.Elem()
	}
	if v2.Kind() == reflect.Pointer {
		v2 = v2.Elem()
	}

	if v1.Kind() != reflect.Struct || v2.Kind() != reflect.Struct {
		return nil, fmt.Errorf("both arguments must be structs")
	}

	t1 := v1.Type()
	t2 := v2.Type()

	var fields []reflect.StructField
	var fieldsSet = make(map[string]reflect.Type)
	for i := 0; i < t1.NumField(); i++ {
		field := t1.Field(i)
		fields = append(fields, field)
		fieldsSet[field.Name] = field.Type
	}

	for i := 0; i < t2.NumField(); i++ {
		field := t2.Field(i)
		if existingType, ok := fieldsSet[field.Name]; ok {
			if existingType != field.Type {
				return nil, fmt.Errorf("field %s has conflicting types: %s and %s", field.Name, existingType, field.Type)
			}
		} else {
			fields = append(fields, field)
			fieldsSet[field.Name] = field.Type
		}
	}

	newType := reflect.StructOf(fields)

	newValue := reflect.New(newType).Elem()

	for i := 0; i < t1.NumField(); i++ {
		valueField := newValue.FieldByName(t1.Field(i).Name)
		if valueField.IsValid() {
			valueField.Set(v1.Field(i))
		}
	}

	for i := 0; i < t2.NumField(); i++ {
		valueField := newValue.FieldByName(t2.Field(i).Name)
		if valueField.IsValid() {
			valueField.Set(v2.Field(i))
		}
	}

	return newValue.Interface(), nil
}
