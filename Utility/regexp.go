package Utility

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func IsIP(ip string) (b bool) {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip); !m {
		return false
	}
	return true
}

func GetTime() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}

// 是否為數字
func IsNumber(s string) bool {
	if m, _ := regexp.MatchString("^[0-9]+$", s); !m {
		return false
	}
	return true
}

// 是否為中文
func IsChinese(s string) bool {
	if m, _ := regexp.MatchString("^\\p{Han}+$", s); !m {
		return false
	}
	return true
}

// 是否為英文
func IsEnglish(s string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", s); !m {
		return false
	}
	return true
}

// 是否為信箱
func IsEmail(s string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, s); !m {
		return false
	}
	return true
}

// 是否為手機號碼
func IsCellphone(s string) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, s); !m {
		return false
	}
	return true
}

// 身分證
func IsIdentityNumber(s string, t int) bool {
	//驗證 15 位身份證，15 位的是全部數字
	if m, _ := regexp.MatchString(`^(\d{15})$`, s); !m {
		return false
	}

	//驗證 18 位身份證，18 位前 17 位為數字，最後一位是校驗位，可能為數字或字元 X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, s); !m {
		return false
	}
	return true
}

func RegTest() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
		os.Exit(1)
	} else if IsNumber(os.Args[1]) {
		fmt.Println("數字")
	} else {
		fmt.Println("不是數字")
	}
}

func RegAdvanced() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
		return
	}
	//替換式
	//Compile 會解析正則表示式是否合法，如果正確，那麼就會回傳一個 Regexp
	src := string(body)
	//將 HTML 標籤全轉換成小寫
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除 STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除 SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括號內的 HTML 程式碼，並換成換行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除連續的換行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))

	/////////
	//搜尋式

	a := "I am learning Go language"
	re, _ = regexp.Compile("[a-z]{2,4}")

	//查詢符合正則的第一個
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	//查詢符合正則的所有 slice,n 小於 0 表示回傳全部符合的字串，不然就是回傳指定的長度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	//查詢符合條件的 index 位置，開始位置和結束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	//查詢符合條件的所有的 index 位置，n 同上
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	//查詢 Submatch，回傳陣列，第一個元素是匹配的全部元素，第二個元素是第一個()裡面的，第三個是第二個()裡面的
	//下面的輸出第一個元素是"am learning Go language"
	//第二個元素是" learning Go "，注意包含空格的輸出
	//第三個元素是"uage"
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	//定義和上面的 FindIndex 一樣
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	//FindAllSubmatch，查詢所有符合條件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	//FindAllSubmatchIndex，查詢所有字匹配的 index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)
}
