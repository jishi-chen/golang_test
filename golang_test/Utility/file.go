package Utility

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func CreateMultipartPostFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}                 //創建了一個新的 bytes.Buffer 實例，這是一個實現了 io.Writer 和 io.Reader 介面的類型
	bodyWriter := multipart.NewWriter(bodyBuf) //multipart.Writer 用於構建 multipart 表單

	//關鍵的一步操作
	//CreateFormFile 方法用於創建一個新的表單數據部分，並返回一個 io.Writer 用於將文件的內容寫入這個部分，第一個參數是表單字段的名稱，第二個參數是文件的名稱
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//開啟檔案控制代碼操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}
type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	//XMLName    xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func ReadXml() {
	file, err := os.Open("template/server.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
}

func WriteXml() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

func DirectoryAdd(path string, perm os.FileMode) {
	os.Mkdir(path, perm)
	os.MkdirAll(path, perm)
}

func FileDirectoryRemove(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll(path)
}
func FileBasic() {
	// 創建一個新文件，如果文件已存在，則截斷為空
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 在文件中寫入一些內容
	_, err = file.WriteString("Hello, world!\n")
	if err != nil {
		log.Fatal(err)
	}

	existingFile, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer existingFile.Close()

	// 獲取文件描述符
	fd := existingFile.Fd()

	// 使用 os.NewFile 包裝現有文件描述符
	wrappedFile := os.NewFile(fd, "wrappedExampleFile")

	// 在包裝的文件中進行讀取操作
	data := make([]byte, 100)
	n, err := wrappedFile.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印讀取的內容
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])
}

/* os 套件

func Mkdir(name string, perm FileMode) error
建立名稱為 name 的目錄，許可權設定是 perm，例如 0777
func MkdirAll(path string, perm FileMode) error
根據 path 建立多階層子目錄，例如 astaxie/test1/test2。
func Remove(name string) error
刪除名稱為 name 的目錄，當目錄下有檔案或者其他目錄時會出錯
func RemoveAll(path string) error
根據 path 刪除多階層子目錄，如果 path 是單個名稱，那麼該目錄下的子目錄全部刪除。

func Create(name string) (file *File, err Error)
根據提供的檔名建立新的檔案，回傳一個檔案物件，預設許可權是 0666 的檔案，回傳的檔案物件是可讀寫的。
func NewFile(fd uintptr, name string) *File
根據檔案描述符建立相應的檔案，回傳一個檔案物件

func Open(name string) (file *File, err Error)
該方法開啟一個名稱為 name 的檔案，但是是隻讀方式，內部實現其實呼叫了 OpenFile。
func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
開啟名稱為 name 的檔案，flag 是開啟的方式，只讀、讀寫等，perm 是許可權

func (file *File) Write(b []byte) (n int, err Error)
寫入 byte 型別的資訊到檔案
func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
在指定位置開始寫入 byte 型別的資訊
func (file *File) WriteString(s string) (ret int, err Error)
寫入 string 資訊到檔案

func (file *File) Read(b []byte) (n int, err Error)
讀取資料到 b 中
func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
從 off 開始讀取資料到 b 中

func Remove(name string) Error
呼叫該函式就可以刪除檔名為 name 的檔案

*os.File 對象:
os.Stdout 標準輸出
os.Stdin 標準輸入
os.Stderr 標準錯誤

*/
