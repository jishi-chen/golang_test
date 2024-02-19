package Utility

import (
	"fmt"
	"time"
)

var locales map[string]map[string]string

func GlobalSources() {
	locales = make(map[string]map[string]string, 2)
	en := make(map[string]string, 10)
	locales["en"] = en
	cn := make(map[string]string, 10)
	locales["zh-CN"] = cn

	en["pea"] = "pea"
	en["bean"] = "bean"
	cn["pea"] = "豌豆"
	cn["bean"] = "毛豆"
	en["how old"] = "I am %d years old"
	cn["how old"] = "我今年%d 歲了"
	en["time_zone"] = "America/Chicago"
	cn["time_zone"] = "Asia/Shanghai"
	en["date_format"] = "%Y-%m-%d %H:%M:%S"
	cn["date_format"] = "%Y 年%m 月%d 日 %H 時%M 分%S 秒"

	lang := "zh-CN"
	loc, _ := time.LoadLocation(getData(lang, "time_zone")) //取得時區
	t := time.Now()
	t = t.In(loc) //轉換時區
	fmt.Println(getData(lang, "pea"))
	fmt.Println(getData(lang, "bean"))
	fmt.Printf(getData(lang, "how old"), 30)
	fmt.Println()
	fmt.Println(t)
	fmt.Println(date(getData(lang, "date_format"), t))
}

func getData(locale, key string) string {
	if v, ok := locales[locale]; ok {
		if v2, ok := v[key]; ok {
			return v2
		}
	}
	return ""
}
func date(fomate string, t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	return fmt.Sprintf("%d 年%d 月%d 日 %02d 時%02d 分%02d 秒\n", year, int(month), day, hour, min, sec)
}

/*
本地化日期和時間
1.時區問題
2.格式問題
*/

// GO 語言預設採用"UTF-8"編碼集
/*
func SetLocale(w http.ResponseWriter, r *http.Request) {
	if r.Host == "www.asta.com" {
		i18n.SetLocale("en")
	} else if r.Host == "www.asta.cn" {
		i18n.SetLocale("zh-CN")
	} else if r.Host == "www.asta.tw" {
		i18n.SetLocale("zh-TW")
	}

	prefix := strings.Split(r.Host, ".")

	if prefix[0] == "en" {
		i18n.SetLocale("en")
	} else if prefix[0] == "cn" {
		i18n.SetLocale("zh-CN")
	} else if prefix[0] == "tw" {
		i18n.SetLocale("zh-TW")
	}

	AL := r.Header.Get("Accept-Language")
	if AL == "en" {
		i18n.SetLocale("en")
	} else if AL == "zh-CN" {
		i18n.SetLocale("zh-CN")
	} else if AL == "zh-TW" {
		i18n.SetLocale("zh-TW")
	}
}
*/
