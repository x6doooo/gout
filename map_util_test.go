package gout

import (
    "testing"
)

func Test_snake2camel(t *testing.T) {
    strs := [][]string {
        []string{"test", "Test"},
        []string{"test_abc", "TestAbc"},
        []string{"TzAc", "TzAc"},
    }
    for _, strPair := range strs {
        if snake2camel(strPair[0]) != strPair[1] {
            t.Error("failed")
        }
    }
}

func TestSetField(t *testing.T) {
    obj := struct {
        A string
        TestField float64
    }{}

    var err error
    err = SetField(&obj, "a", "12")
    if err != nil {
        t.Error("failed")
    }
    err = SetField(&obj, "test_field", 12.1)
    if err != nil {
        t.Error("failed")
    }

}

func TestMap2Struct(t *testing.T) {
    obj := struct {
        A string
        TestField float64
    }{}

    testMap := map[string]interface{}{
        "a": "123",
        "test_field": 11.1,
        "abc": true,
    }

    var err error

    err = Map2Struct(testMap, &obj, false)
    if err != nil {
        t.Error("failed")
    }

    err = Map2Struct(testMap, &obj, true)
    if err == nil {
        t.Error("failed")
    }

}

