package gout

import (
    "reflect"
    "fmt"
    "errors"
)

func SetField(obj interface{}, name string, value interface{}) error {
    structValue := reflect.ValueOf(obj).Elem()
    structFieldValue := structValue.FieldByName(name)

    if !structFieldValue.IsValid() {
        return fmt.Errorf("No such field: %s in obj", name)
    }

    if !structFieldValue.CanSet() {
        return fmt.Errorf("Cannot set %s field value", name)
    }

    structFieldType := structFieldValue.Type()
    val := reflect.ValueOf(value)
    if structFieldType != val.Type() {
        return errors.New("Provided value type didn't match obj field type")
    }

    structFieldValue.Set(val)
    return nil
}

func Map2Struct(oneMap map[string]interface{}, oneStruct interface{}) error {
    for k, v := range oneMap {
        err := SetField(oneStruct, k, v)
        if err != nil {
            return err
        }
    }
    return
}
