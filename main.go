package main

import (
	"fmt"
	"sort"
	"strconv"
	"test/mymath"
)

func main() {
	//guide.Lesson2()
	//guide.TicTacToe()
	//guide.GoToExample()
	//Utility.GetTime()
	//Web.StartConnect()
	//Utility.StartConnDatabase()
	//Utility.WriteXml()
	//b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	//Utility.JsonDecode2(b)
	//Utility.JsonEncode()
	//Utility.RegAdvanced()
	//Utility.TemplateHandler()
	//Utility.TemplateMain()
	//Utility.FileDirectoryRemove("example.txt")
	//Utility.Encrypt()
	//Utility.GlobalSources()

	// 創建一個 UserController 實例
	//userController := Utility.UserController{}

	// 使用介面來調用方法
	//userInfo := userController.GetUser(123)

	f2input := []int{5, 4, 3, 2, 1}
	f3input := []int{0x30a, 0x30b, 0x30c, 0x30d, 0x30e}
	fmt.Println(f())
	fmt.Println(f2(f2input))
	fmt.Println(f3(f3input))
}

func f() string {
	return "hello world"
}

func f2(input []int) string {
	sort.Ints(input)
	result := ""
	for i, num := range input {
		if i > 0 {
			result += " "
		}
		result += strconv.Itoa(num)
	}
	return result
}

func f3(input []int) int {
	cardSet := new(mymath.CardSet)
	cardSet.Cards = input
	cardSet.Sort()
	cardSet.GetSetType()
	return cardSet.Type
}
