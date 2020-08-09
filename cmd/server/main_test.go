package main

import (
	"fmt"
	gofaker "github.com/tinygg/faker"
	"net/url"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	var response map[string]gofaker.Info
	var statusCode int
	testRequest(&testRequestStruct{
		Testing:    t,
		Method:     "GET",
		Path:       "/list",
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if len(response) == 0 {
		t.Fatalf("Was expecting list length to be more than 0 got %d", len(response))
	}
}

func TestGetAllRequests(t *testing.T) {
	for field, info := range gofaker.FuncLookups {
		mapData := url.Values{}
		if info.Params != nil && len(info.Params) != 0 {
			// Loop through params and add fields to mapdata
			for _, p := range info.Params {
				if p.Default != "" {
					mapData.Add(p.Field, p.Default)
					continue
				}

				switch p.Type {
				case "bool":
					mapData.Add(p.Field, fmt.Sprintf("%v", gofaker.Bool()))
					break
				case "string":
					mapData.Add(p.Field, gofaker.Letter())
					break
				case "uint":
					mapData.Add(p.Field, fmt.Sprintf("%v", gofaker.Uint16()))
				case "int":
					mapData.Add(p.Field, fmt.Sprintf("%v", gofaker.Int16()))
				case "float":
					mapData.Add(p.Field, fmt.Sprintf("%v", gofaker.Float32()))
					break
				case "[]string":
					mapData.Add(p.Field, gofaker.Letter())
					mapData.Add(p.Field, gofaker.Letter())
					mapData.Add(p.Field, gofaker.Letter())
					mapData.Add(p.Field, gofaker.Letter())
					break
				case "[]int":
					mapData.Add(p.Field, fmt.Sprintf("%d", gofaker.Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", gofaker.Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", gofaker.Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", gofaker.Int8()))
					break
				case "[]Field":
					mapData.Add(p.Field, `{"name":"first_name","function":"firstname"}`)
					break
				default:
					t.Fatalf("Looking for %s but switch case doesnt have it", p.Type)
				}
			}
		}

		var statusCode int
		testRequest(&testRequestStruct{
			Testing:     t,
			Method:      "GET",
			Path:        "/" + field,
			QueryParams: mapData,
			StatusCode:  &statusCode,
		})

		if statusCode != 200 {
			t.Fatalf("Was expecting 200 got %d", statusCode)
		}
	}
}

func TestGetLookupNoParams(t *testing.T) {
	var response string
	var statusCode int
	testRequest(&testRequestStruct{
		Testing:    t,
		Method:     "GET",
		Path:       "/firstname",
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if response == "" {
		t.Fatalf("Was expecting a string with value got empty")
	}
}

func TestGetLookupWithParams(t *testing.T) {
	var response string
	var statusCode int
	testRequest(&testRequestStruct{
		Testing: t,
		Method:  "GET",
		Path:    "/password",
		QueryParams: url.Values{
			"length": []string{"5"},
		},
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if response == "" {
		t.Fatalf("Was expecting a string with value got empty")
	}

	if len(response) != 5 {
		t.Fatalf("Was expecting a string length of 5 got %d", len(response))
	}
}

func TestPostAllRequests(t *testing.T) {
	for field, info := range gofaker.FuncLookups {
		var mapData map[string][]string
		if info.Params != nil && len(info.Params) != 0 {
			// Make sure mapdata is set
			if mapData == nil {
				mapData = make(map[string][]string)
			}

			// Loop through params and add fields to mapdata
			for _, p := range info.Params {
				if p.Default != "" {
					mapData[p.Field] = []string{p.Default}
					continue
				}

				switch p.Type {
				case "bool":
					mapData[p.Field] = []string{fmt.Sprintf("%v", gofaker.Bool())}
					break
				case "string":
					mapData[p.Field] = []string{gofaker.Letter()}
					break
				case "uint":
					mapData[p.Field] = []string{fmt.Sprintf("%v", gofaker.Uint16())}
				case "int":
					mapData[p.Field] = []string{fmt.Sprintf("%v", gofaker.Int16())}
				case "float":
					mapData[p.Field] = []string{fmt.Sprintf("%v", gofaker.Float32())}
					break
				case "[]string":
					mapData[p.Field] = []string{gofaker.Letter(), gofaker.Letter(), gofaker.Letter(), gofaker.Letter()}
					break
				case "[]int":
					mapData[p.Field] = []string{fmt.Sprintf("%d", gofaker.Int8()), fmt.Sprintf("%d", gofaker.Int8()), fmt.Sprintf("%d", gofaker.Int8()), fmt.Sprintf("%d", gofaker.Int8())}
					break
				case "[]Field":
					mapData[p.Field] = []string{`{"name":"first_name","function":"firstname"}`}
					break
				default:
					t.Fatalf("Looking for %s but switch case doesnt have it", p.Type)
				}
			}
		}

		var statusCode int
		testRequest(&testRequestStruct{
			Testing:    t,
			Method:     "POST",
			Path:       "/" + field,
			Body:       mapData,
			StatusCode: &statusCode,
		})

		if statusCode != 200 {
			t.Fatalf("Was expecting 200 got %d", statusCode)
		}
	}
}

func TestPostLookupNoParams(t *testing.T) {
	var response string
	var statusCode int
	testRequest(&testRequestStruct{
		Testing:    t,
		Method:     "POST",
		Path:       "/firstname",
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if response == "" {
		t.Fatalf("Was expecting a string with value got empty")
	}
}

func TestPostLookupWithParams(t *testing.T) {
	var response string
	var statusCode int
	testRequest(&testRequestStruct{
		Testing: t,
		Method:  "POST",
		Path:    "/password",
		Body: map[string]interface{}{
			"space":  true,
			"length": "500",
		},
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if response == "" {
		t.Fatalf("Was expecting a string with value got empty")
	}

	if len(response) != 500 {
		t.Fatalf("Was expecting a string length of 500 got %d", len(response))
	}

	if !strings.Contains(response, " ") {
		t.Fatal("Was expecting a space in the password and didnt get it")
	}
}

func TestPostLookupWithParamsArray(t *testing.T) {
	var response []string
	var statusCode int
	testRequest(&testRequestStruct{
		Testing: t,
		Method:  "POST",
		Path:    "/shufflestrings",
		Body: map[string]interface{}{
			"strs": []string{"a", "b", "c", "d", "e", "f"},
		},
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if len(response) != 6 {
		t.Fatalf("Was expecting a array length of 6 got %d", len(response))
	}
}
