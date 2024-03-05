package Web

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"test/Utility"
	"time"
)

type MyMux struct {
}

func StartConnect() {
	mux := &MyMux{}
	http.HandleFunc("/", SayHelloName) //設定存取的路由
	http.HandleFunc("/login", Login)   //設定存取的路由
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/test", test)
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/*fmt.Println(r.Form) //這些資訊是輸出到伺服器端的列印資訊
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}*/

	if r.URL.Path == "/" {
		SayHelloName(w, r)
		return
	}
	if r.URL.Path == "/login" {
		Login(w, r)
		return
	}
	if r.URL.Path == "/upload" {
		upload(w, r)
		return
	}
	if r.URL.Path == "/test" {
		test(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func SayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello astaxie!") //這個寫入到 w 的是輸出到客戶端的
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //取得請求的方法
	if r.Method == "GET" {
		currentTime := time.Now().Unix() //取得當前時間的Unix時間戳
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 10)) //將當前時間戳轉為字符串，並寫入至MD5散列中
		//h.Sum(nil): 計算 MD5 雜湊，並將其作為一個字節切片返回。
		//fmt.Sprintf("%x", ...): 將從 MD5 雜湊得到的字節轉換為十六進位字串。
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("template/login.gtpl")
		t.Execute(w, token) //將token 注入到模板中
	} else {
		r.ParseForm() //解析參數，預設是不會解析的
		isValid, isExistFruit, isExistGender := true, false, false
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("token:", token)
			//驗證 token 的合法性
		} else {
			//不存在 token 報錯
		}
		//請求的是登入資料，那麼執行登入的邏輯判斷
		if len(r.Form["username"][0]) == 0 {
			isValid = false
			fmt.Println("帳號為空")
		}
		//轉化成數字
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			isValid = false
			fmt.Println("年齡必須為數字")
		}
		//接下來就可以判斷這個數字的大小範圍了
		if getint > 100 {
			isValid = false
			fmt.Println("年齡不可大於100")
		}

		fruit := []string{"apple", "pear", "banana"}
		gender := []string{"1", "2"}
		for _, item := range fruit {
			if item == r.Form.Get("fruit") {
				isExistFruit = true
			}
		}
		for _, item := range gender {
			if item == r.Form.Get("gender") {
				isExistGender = true
			}
		}
		/*interest := []string{"football", "basketball", "tennis"}
		a := Utility.Slice_diff(r.Form["interest"], interest)
		if a != nil {
			isValid = false
		}*/
		if isValid && isExistFruit && isExistGender {
			//fmt.Println("username:", template.HTMLEscapeString(r.FormValue("username")))
			//fmt.Println("password:", template.HTMLEscapeString(r.FormValue("password")))
			template.HTMLEscape(w, []byte(r.Form.Get("username")))
			//創建一個新的模板"foo" 並解析一個定義的字符串模板 "T" 會將參數插入至 Hello, 和 ! 之間
			t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
			err = t.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
		}
	}
}

// 處理/upload 邏輯
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //取得請求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("template/upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20) //設置了解析多部分表單數據時使用的最大內存量為 32 MB
		//(32 << 20)：這部分是一個位運算，計算為 32 左移 20 位。在這裡，<< 是左位移運算符，將 32 左移 20 位，相當於將 32 乘以 2 的 20 次方。換句話說，這表示數據的最大大小為 32 MB。
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./UploadFiles/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此處假設當前目錄下已存在 test 目錄
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //取得請求的方法
	if r.Method == "GET" {
		Utility.SetCookie(w)
		Utility.GetCookie(w, r)
	} else {

	}
}

/*
一、呼叫 Http.HandleFunc
按順序做了幾件事：
1 呼叫了 DefaultServeMux 的 HandleFunc
2 呼叫了 DefaultServeMux 的 Handle
3 往 DefaultServeMux 的 map[string]muxEntry 中增加對應的 handler 和路由規則

二、呼叫 http.ListenAndServe(":9090", nil)
按順序做了幾件事情：
1 實體化 Server
2 呼叫 Server 的 ListenAndServe()
3 呼叫 net.Listen("tcp", addr)監聽埠
4 啟動一個 for 迴圈，在迴圈體中 Accept 請求
5 對每個請求實體化一個 Conn，並且開啟一個 goroutine 為這個請求進行服務 go c.serve()
6 讀取每個請求的內容 w, err := c.readRequest()
7 判斷 handler 是否為空，如果沒有設定 handler（這個例子就沒有設定 handler），handler 就設定為 DefaultServeMux
8 呼叫 handler 的 ServeHttp
9 在這個例子中，下面就進入到 DefaultServeMux.ServeHttp
10 根據 request 選擇 handler，並且進入到這個 handler 的 ServeHTTP
    mux.handler(r).ServeHTTP(w, r)
11 選擇 handler：
A 判斷是否有路由能滿足這個 request（迴圈遍歷 ServeMux 的 muxEntry）
B 如果有路由滿足，呼叫這個路由 handler 的 ServeHTTP
C 如果沒有路由滿足，呼叫 NotFoundHandler 的 ServeHTTP
*/

/* 預防跨站指令碼
func HTMLEscape(w io.Writer, b []byte)  //把 b 進行轉義之後寫到 w
func HTMLEscapeString(s string) string  //轉義 s 之後回傳結果字串
func HTMLEscaper(args ...interface{}) string //支援多個參數一起轉義，回傳結果字串
*/

/* form 的enctype屬性
application/x-www-form-urlencoded   表示在傳送前編碼所有字元（預設）
multipart/form-data      不對字元編碼。在使用包含檔案上傳控制元件的表單時，必須使用該值。
text/plain      空格轉換為 "+" 加號，但不對特殊字元編碼。
*/
