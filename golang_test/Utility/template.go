package Utility

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type User struct {
	UserName string
	Email    string
}

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

//使用 template 套件來進行範本處理，使用類似Parse、ParseFile、Execute等方法從檔案或者字串載入範本
//範本透過 {{}} 來包含需要在渲染時被替換的欄位

func TemplateHandler() {
	t := template.New("some template") //建立一個範本
	//t, _ = t.ParseFiles("tmpl/welcome.html") //解析範本檔案
	t, _ = t.Parse("hello {{.UserName}} {{.Email}}!")
	user := User{
		UserName: "Willy",
		Email:    "@gmail.com",
	} //取得當前使用者資訊
	t.Execute(os.Stdout, user) //執行範本的 merger 操作

	/*
		{{range .Emails}}...{{end}}：這是一個 range 操作，用於遍歷數據切片 .Emails，在循環中使用 {{.}} 替換為實際的郵件地址。 {{.}} 會代表當前迭代的元素
		{{with .Friends}}...{{end}}：這是一個 with 操作，用於處理 .Friends 這個可能為空的字段。如果 .Friends 不為空，則進入循環，使用 {{.Fname}} 替換為實際的朋友名字。
	*/
	t, _ = t.Parse(`hello {{.UserName}}!
            {{range .Emails}}
                an email {{.}}
            {{end}}
            {{with .Friends}}
            {{range .}}
                my friend name is {{.Fname}}
            {{end}}
            {{end}}
            `)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{{Fname: "minux.ma"}, {Fname: "xushiwei"}}}
	t.Execute(os.Stdout, p)
}

func TemplateHandlerTest() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不會輸出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不為空的 pipeline if demo: {{if `anything`}} 我有內容，我會輸出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if 部分 {{else}} else 部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)

	//函式Must:檢測範本是否正確，例如大括號是否匹配，註釋是否正確的關閉，變數是否正確的書寫
	tOk := template.New("first")
	template.Must(tOk.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed OK.")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("The second one parsed OK.")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }"))
}

func TemplateMain() {

	// 定義一個模板函式
	var funcMap = template.FuncMap{
		"double": func(x int) int {
			return x * 2
		},
	}
	// 創建一個模板
	tmpl, err := template.New("example").Funcs(funcMap).Parse("Double of {{.}} is {{double .}}.\n")
	if err != nil {
		panic(err)
	}

	// 用一個具體的值來執行模板
	data := 5
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": FuncMapExample})
	t, _ = t.Parse(`hello {{.UserName}}!
                {{range .Emails}}
                    an emails {{.|emailDeal}}
                {{end}}
                {{with .Friends}}
                {{range .}}
                    my friend name is {{.Fname}}
                {{end}}
                {{end}}
                `)
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

	//測試巢狀範本
	s1, _ := template.ParseFiles("template/header.tmpl", "template/content.tmpl", "template/footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)
}

// 每一個範本函式都有一個唯一值的名字，然後與一個 Go 函式關聯，透過如下的方式來關聯
func FuncMapExample(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return (substrs[0] + " at " + substrs[1])
}

/*
以下模板用途相同，不同的模板代碼使用了不同的語法結構，目的都是將字符串 "output" 格式化成帶引號的形式，即 '"output"'
$x 為範本變數，pipelines注入：{{. | html}} 即轉化為 html 的實體
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
{{with $x := "output"}}{{printf "%q" $x}}{{end}}
{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
*/
