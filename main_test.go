package gofaker

import (
	"log"
	"testing"
)

func TestAaa(t *testing.T) {
	for i := 0; i < 100; i++ {
		info := GetFuncLookup("csv")
		m := map[string][]string{
			"rowcount": {"100000"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Call(&m, info)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
