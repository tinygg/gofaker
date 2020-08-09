package address

import (
	"github.com/tinygg/gofaker/data/address/en_US"
	"github.com/tinygg/gofaker/data/address/zh_CN"
)

func Provider(lang string) map[string][]string {
	switch lang {
	case "zh_CN":
		return zh_CN.Address
	case "en_US":
		return en_US.Address
	default:
		return en_US.Address
	}

}
