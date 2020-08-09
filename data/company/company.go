package company

import (
	"github.com/tinygg/gofaker/data/company/en_US"
	"github.com/tinygg/gofaker/data/company/zh_CN"
)

func Provider(lang string) map[string][]string {
	switch lang {
	case "zh_CN":
		return zh_CN.Company
	case "en_US":
		return en_US.Company
	default:
		return en_US.Company
	}

}
