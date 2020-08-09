package data

type Locale string

const (
	ZH_CN Locale = "zh_CN"
	EN_US Locale = "en_US"
	RU_RU Locale = "ru_RU"
	FR_FR Locale = "fr_FR"
	AR_PS Locale = "ar_PS"
	ES_ES Locale = "es_ES"
	FA_IR Locale = "fa_IR"
	HR_HR Locale = "hr_HR"
	HU_HU Locale = "hu_HU"
	HY_AM Locale = "hy_AM"
	PT_BR Locale = "pt_BR"
	TH_TH Locale = "th_TH"
	UK_UA Locale = "uk_UA"
)

var (
	LOCALE Locale = EN_US
)

func SetLocale(locale Locale) {
	LOCALE = locale
}
