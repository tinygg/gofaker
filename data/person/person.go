package person

import (
	"github.com/tinygg/gofaker/data/person/en_US"
	"github.com/tinygg/gofaker/data/person/zh_CN"
)

func Provider(lang string) map[string][]string {
	switch lang {
	case "ZH_CN":
		return zh_CN.Person
	case "EN_US":
		return en_US.Person
	default:
		return en_US.Person
	}

}
