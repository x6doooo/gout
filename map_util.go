package gout

import (
    "reflect"
    "fmt"
    "errors"
    "strings"
)
func snake2camel(str string) string {
    arr := strings.Split(str, "_")
    for i, v := range arr {
        arr[i] = strings.Title(v)
    }
    return strings.Join(arr, "")
}

func SetField(obj interface{}, name string, value interface{}) error {
    structValue := reflect.ValueOf(obj).Elem()

    var findByName = func(s string) bool {
        if s == name {
            return true
        }
        if s == snake2camel(name) {
            return true
        }
        return false
    }

    structFieldValue := structValue.FieldByNameFunc(findByName)

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

/**
    stict 是否严格匹配
        true的时候，每次setvalue都必须成功 否则返回error
        false的时候，如果每次setvalue都报错，才会返回error，否则视为正常
 */
func Map2Struct(oneMap map[string]interface{}, oneStruct interface{}, strict bool) error {
    errCount := 0
    for k, v := range oneMap {
        err := SetField(oneStruct, k, v)
        if err != nil {
            if strict {
                return err
            } else {
                errCount += 1
            }
        }
    }
    if !strict && errCount == len(oneMap) {
        return errors.New("convert failed")
    }
    return nil
}
