package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Struct2Map(obj interface{}) (map[string]interface{}, error) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Ptr {
		return nil, errors.New("needs a pointer to a value")
	} else if objValue.Elem().Kind() == reflect.Ptr {
		return nil, errors.New("a pointer to a pointer is not allowed")
	} else if objValue.Elem().Kind() != reflect.Struct {
		return nil, errors.New("needs a pointer to struct")
	}

	m := make(map[string]interface{})
	elemValue := objValue.Elem()
	for i := 0; i < elemValue.NumField(); i++ {
		value := elemValue.Field(i)
		name := elemValue.Type().Field(i).Name
		m[name] = value.Interface()
	}
	return m, nil
}

func SetField(obj interface{}, name string, value interface{}) error {
	elemValue := reflect.ValueOf(obj).Elem()
	fieldValue := elemValue.FieldByName(name)

	if !fieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}
	if !fieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	if value != nil {
		switch fieldValue.Kind() {
		case reflect.Int8:
			val, err := strconv.Atoi(value.(string))
			if err != nil {
				return errors.New("provided value type didn't match type string")
			}
			v := reflect.ValueOf(int8(val))
			fieldValue.Set(v)
		case reflect.Int16:
			val, err := strconv.Atoi(value.(string))
			if err != nil {
				return errors.New("provided value type didn't match type string")
			}
			v := reflect.ValueOf(int16(val))
			fieldValue.Set(v)
		case reflect.Int32:
			val, err := strconv.Atoi(value.(string))
			if err != nil {
				return errors.New("provided value type didn't match type string")
			}
			v := reflect.ValueOf(int32(val))
			fieldValue.Set(v)
		case reflect.Int64:
			val, err := strconv.Atoi(value.(string))
			if err != nil {
				return errors.New("provided value type didn't match type string")
			}
			v := reflect.ValueOf(int64(val))
			fieldValue.Set(v)
		case reflect.Int:
			val, err := strconv.Atoi(value.(string))
			if err != nil {
				return errors.New("provided value type didn't match type string")
			}
			v := reflect.ValueOf(int(val))
			fieldValue.Set(v)
		case reflect.Uint64:
			val, err := strconv.Atoi(value.(string))
			if err != nil {
				return errors.New("provided value type didn't match type string")
			}
			v := reflect.ValueOf(uint64(val))
			fieldValue.Set(v)
		case reflect.Uint32:
			val, err := strconv.Atoi(value.(string))
			if err != nil {
				return errors.New("provided value type didn't match type string")
			}
			v := reflect.ValueOf(uint32(val))
			fieldValue.Set(v)
		case reflect.String:
			val := reflect.ValueOf(value)
			fieldValue.Set(val)
		case reflect.Bool:
			val, err := strconv.ParseBool(value.(string))
			if err != nil {
				return errors.New("provided value type didn't match type string")
			}
			v := reflect.ValueOf(val)
			fieldValue.Set(v)
		default:
			return errors.New("provided value type didn't match obj field type")
		}
	}
	return nil
}
