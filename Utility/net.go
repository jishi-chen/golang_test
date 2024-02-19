package Utility

import (
	"fmt"
	"io"
	"net"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"time"
)

func ClientSocket() {
	service := "128.0.0.1"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := io.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func ServerSocket() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn) //多併發執行
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128)                          // set maxium request length to 128B to prevent flood attack
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// 定義服務結構體
type Calculator struct{}

// 定義遠程方法
func (c *Calculator) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

// 定義參數結構體
type Args struct {
	A, B int
}

func ServerRPC() {
	// 創建 Calculator 實例
	calculator := new(Calculator)

	// 註冊 Calculator 實例為 RPC 服務
	rpc.Register(calculator)

	// 監聽 TCP 連接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 1234...")

	// 接受並處理客戶端的請求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

func ClientRPC() {
	// 連接 RPC 服務
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	// 準備遠程方法的參數
	args := &Args{A: 5, B: 3}

	// 調用遠程方法
	var result int
	err = client.Call("Calculator.Add", args, &result)
	if err != nil {
		fmt.Println("Error calling remote method:", err)
		return
	}

	fmt.Println("Result:", result)
}

/* net 套件
addr := net.ParseIP("128.0.0.1")
轉化成 IP 型別

func (c *TCPConn) Write(b []byte) (int, error)
func (c *TCPConn) Read(b []byte) (int, error)
用來作為客戶端和伺服器端互動的通道

func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
取得一個TCPAddr，表示一個 TCP 的地址資訊

func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
建立一個 TCP 連線，network 參數是"tcp4"、"tcp6"、"tcp"中的任意一個，laddr 表示本機地址，一般設定為 nil，raddr 表示遠端的服務地址

func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
func (l *TCPListener) Accept() (Conn, error)
監聽客戶端連線的請求

func (c *TCPConn) SetReadDeadline(t time.Time) error
func (c *TCPConn) SetWriteDeadline(t time.Time) error
設定建立連線的超時時間

func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
設定 keepAlive 屬性

func ResolveUDPAddr(net, addr string) (*UDPAddr, os.Error)
func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err os.Error)
func ListenUDP(net string, laddr *UDPAddr) (c *UDPConn, err os.Error)
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err os.Error)
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (n int, err os.Error)
UDP Socket

go get golang.org/x/net/websocket
由官方維護的 go.net 子套件, 實現 WebSocket
*/
