package Utility

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Peson struct {
	Servers []Server
	Name    string `json:"name"`
	Age     int    `json:"age"`
	City    string `json:"city"`
	Email   string `json:"email,omitempty"`
}

func JsonDecode() {
	var person Peson
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}], "name":"John Doe","age":30,"city":"New York"}`
	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Decoded Person: %+v\n", person)
}

func JsonDecode2(b []byte) {
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	/*
		f：接口變數，可能包含任何型別的值。
		.()：類型斷言的語法。
		map[string]interface{}：目標類型，希望將 f 轉換為這種類型。
	*/

	if m, ok := f.(map[string]interface{}); ok {
		// 使用 m（轉換成功）
		fmt.Println(m["Name"])

		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
	} else {
		// 處理類型不匹配的情況
		fmt.Println("f 不是 map[string]interface{} 型別")
	}
}

func JsonEncode() {
	type ServerData struct {
		// ID 不會匯出到 JSON 中
		ID int `json:"-"`

		// ServerName2 的值會進行二次 JSON 編碼
		ServerName  string `json:"serverName"`
		ServerName2 string `json:"serverName2,string"`
		ServerName3 string
		// 如果 ServerIP 為空，則不輸出到 JSON 串中
		ServerIP string `json:"serverIP,omitempty"`
	}
	type ServerDataSlice struct {
		Servers []ServerData
	}

	var s ServerDataSlice
	s.Servers = append(s.Servers, ServerData{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, ServerData{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	server := ServerData{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    ``,
	}
	// 編碼結構為 JSON 字符串
	jsonData, err := json.Marshal(server)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	count, err := os.Stdout.Write(jsonData)
	fmt.Println(count)

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

/*
JSON 物件只支援 string 作為 key，所以要編碼一個 map，那麼必須是 map[string]T 這種型別(T 是 Go 語言中任意的型別)
Channel, complex 和 function 是不能被編碼成 JSON 的
巢狀的資料是不能編碼的，不然會讓 JSON 編碼進入無窮遞迴
指標在編碼的時候會輸出指標指向的內容，而空指標會輸出 null
*/
