package Utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/scrypt"
)

func Encrypt() {
	//普通方法:單向雜湊
	h := sha256.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Println()

	h = sha1.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
	fmt.Println()

	//進階方法:加入salt
	h = md5.New()
	io.WriteString(h, "需要加密的密碼")
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))
	//指定兩個 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"
	salt := "#EDC$RFV"

	//salt1+使用者名稱+salt2+MD5 拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)
	fmt.Printf("%x", h.Sum(nil))
	fmt.Println()

	//專業方法:使用scrypt 方案
	dk, err := scrypt.Key([]byte("some password"), []byte(salt), 16384, 8, 1, 32)
	checkError(err)
	fmt.Printf("%x", dk)
}

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func AESCrypto() {
	var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	plaintext := []byte("My name is Astaxie")
	key_text := "astaxie12798akljzmknm.ahkjkljl;k" //參數 key 必須是 16、24 或者 32 位的[]byte

	// 建立加密演算法 aes
	c, err := aes.NewCipher([]byte(key_text)) //AES-128, AES-192 或 AES-256 演算法
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key_text), err)
		os.Exit(-1)
	}
	//加密字串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	// 解密字串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)
}

/*
預防 CSRF:
1、正確使用 GET,POST 和 Cookie；
mux.Get("/user/:uid", getuser)
mux.Post("/user/:uid", modifyuser)
限制對資源的存取方法

2、在非 GET 請求中增加偽隨機數；
產生隨機數 token, 輸出 token, 驗證 token
h := md5.New()


防禦 XSS：

1.過濾特殊字元
避免 XSS 的方法之一主要是將使用者所提供的內容進行過濾，Go 語言提供了 HTML 的過濾函式：
text/template 套件下面的 HTMLEscapeString、JSEscapeString 等函式

2.使用 HTTP 頭指定型別
`w.Header().Set("Content-Type","text/javascript")`
這樣就可以讓瀏覽器解析 javascript 程式碼，而不會是 html 輸出。


過濾資料：
strconv 套件下面的字串轉化相關函式，因為從 Request 中的r.Form回傳的是字串，而有些時候我們需要將之轉化成整/浮點數，Atoi、ParseBool、ParseFloat、ParseInt等函式就可以派上用場了。
string 套件下面的一些過濾函式Trim、ToLower、ToTitle等函式，能夠幫助我們按照指定的格式取得資訊。
regexp 套件用來處理一些複雜的需求，例如判定輸入是否是 Email、生日之類別。


預防 SQL 注入:
1.嚴格限制 Web 應用的資料庫的操作許可權，給此使用者提供僅僅能夠滿足其工作的最低許可權，從而最大限度的減少注入攻擊對資料庫的危害。
2.檢查輸入的資料是否具有所期望的資料格式，嚴格限制變數的型別，例如使用 regexp 套件進行一些匹配處理，或者使用 strconv 套件對字串轉化成其他基本型別的資料進行判斷。
3.對進入資料庫的特殊字元（'"\尖括號&*;等）進行轉義處理，或編碼轉換。Go 的text/template套件裡面的 HTMLEscapeString 函式可以對字串進行轉義處理。
4.所有的查詢語句建議使用資料庫提供的參數化查詢介面，參數化的語句使用參數而不是將使用者輸入變數嵌入到 SQL 語句中，即不要直接拼接 SQL 語句。例如使用database/sql裡面的查詢函式 Prepare 和Query，或者Exec(query string, args ...interface{})。
5.在應用釋出之前建議使用專業的 SQL 注入檢測工具進行檢測，以及時修補被發現的 SQL 注入漏洞。網上有很多這方面的開源工具，例如 sqlmap、SQLninja 等。
6.避免網站顯示出 SQL 錯誤資訊，比如型別錯誤、欄位不匹配等，把程式碼裡的 SQL 語句暴露出來，以防止攻擊者利用這些錯誤資訊進行 SQL 注入。


單向雜湊有兩個特性：
1）同一個密碼進行單向雜湊，得到的總是唯一確定的摘要。
2）計算速度快。隨著技術進步，一秒鐘能夠完成數十億次單向雜湊計算。

*/
