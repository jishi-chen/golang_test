package guide

/*break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var*/

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"test/mymath"
	"time"
)

type VerTex struct {
	x, y int
}
type Vertex struct {
	Lat, Long float64
}

func Lesson1() {
	var bookName, bookNo, isEnable = "A Little Prince", 10, true
	const World = "世界"
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)
	price1, price2, price3 := 100, 200, 350
	f := float64(price1)
	var (
		ToBe   bool   = false
		MaxInt uint64 = 1<<64 - 1
	)

	fmt.Println(mymath.Subtract(6, 6), mymath.Add(6, 6))
	fmt.Println(math.Pi, rand.Intn(10), math.Sqrt(16))
	fmt.Println(swap("Apple", "Banana"))
	fmt.Println(split(17))
	fmt.Println(bookName, bookNo, isEnable)
	fmt.Println(price1, price2, price3)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", f, f)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum, sum2)
	fmt.Println(pow(3, 2, 10))

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

}

func Lesson2() {

	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 透過指針讀取 i 的值
	*p = 21         // 透過指針設置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 透過指針對 j 進行運算
	fmt.Println(j) // 查看 j 的值

	v := VerTex{5, 10} //struct
	v.x = 4
	fmt.Println(v)

	// 映射
	var m = make(map[string]Vertex) //初始化(make)一個map m, 其鍵為string ,值為Vertex
	m["Location1"] = Vertex{40.68433, -74.39967}
	m["Location2"] = Vertex{37.774929, -122.419416}

	m2 := map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	// 獲取 map 中的值
	fmt.Println(m["Location1"]) // 输出：{40.68433 -74.39967}
	fmt.Println(m["Location2"]) // 输出：{37.774929 -122.419416}
	fmt.Println(m2["Bell Labs"])

	m["Location1"] = Vertex{40, -74} //改值
	delete(m, "Location1")           //刪除
	value, ok := m["Location1"]      //存在
	fmt.Println("The value:", value, "Present?", ok)

	//函數
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	//閉包(Closure) 內部狀態會保留，因為他們是閉包，每次調用都能記住上一次調用後的狀態。
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	vertex := Vertex{3, 4}
	vertex.Scale(10)
	fmt.Println(vertex.Abs())
}

type Human struct {
	name  string
	age   int
	phone string
}
type Element interface{}
type List []Element

// 嵌入 interface
type TestInterface interface {
	sort.Interface
	io.ReadWriter
	Push(x interface{})
	Pop() interface{}
}

func (h Human) String() string {
	return "<" + h.name + " - " + strconv.Itoa(h.age) + " years - ✆ " + h.phone + ">"
}

func Lesson3() {
	fmt.Println("Start")
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is :", Bob.String())

	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Human{"Dennis", 70, "123456789"}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Human); ok {
			fmt.Printf("list[%d] is a Human and its value is %s\n", index, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Human:
			fmt.Printf("list[%d] is a Human and its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type", index)
		}
	}
	//`element.(type)`語法不能在 switch 外的任何邏輯裡面使用，如果你要在 switch 外面判斷一個型別就使用`comma-ok`
	var x float64 = 3.4
	p := reflect.ValueOf(x)
	p2 := reflect.ValueOf(&x)
	v := p2.Elem()
	v.SetFloat(7.1)
	fmt.Println("type:", p.Type())
	fmt.Println("kind is float64:", p.Kind() == reflect.Float64)
	fmt.Println("value:", p.Float())
	fmt.Println("value:", v.Float())
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 定義了一個名為Abs() 的方法，它關連到 Vertex 結構體。這個方法接收一個 Vertex 類型的接收者 v。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

// 定義了一個名為Scale() 的方法，它關連到 Vertex 結構體，接收一個指向 Vertex 類型的指針作為接收者 v
func (v *Vertex) Scale(f float64) {
	v.Lat = v.Lat * f
	v.Long = v.Long * f
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func slice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)

	c[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	rr := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(rr)

	/////////////////////////////

	arr := [5]string{"a", "b", "c", "d"}
	fmt.Println(arr)      // [a b c d _]
	fmt.Println(len(arr)) // 5

	s1 := arr[1:3]
	fmt.Println(s1)      // [b c]
	fmt.Println(len(s1)) // 2
	fmt.Println(cap(s1)) // 4 [b c d _]

	s2 := s1[1:]         // [c]
	fmt.Println(len(s2)) // 1
	fmt.Println(cap(s2)) // 3 [c d e]

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	a := make([]int, 5)    // len(a)=5
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	b = b[:cap(b)]         // len(b)=5, cap(b)=5
	b = b[1:]              // len(b)=4, cap(b)=4

	a = append(a, 2, 3, 4) // 添加元素
	fmt.Println(a)

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128} //遍歷切片 range
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func TicTacToe() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func GoToExample() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Trying operation at (%d, %d)\n", i, j)
			if i == 2 && j == 1 {
				// 模擬錯誤發生
				fmt.Println("Error: Operation failed!")
				goto handleError
			}
		}
	}

	fmt.Println("Operation successful!")
	return

handleError:
	fmt.Println("Handling error...")
	// 處理錯誤
}

func routine() {
	a := []int{7, 2, 8, -9, 4, 0}

	//channel 接收和傳送資料都是阻塞的
	ci := make(chan int)
	go sum(a[:len(a)/2], ci)
	go sum(a[len(a)/2:], ci)
	//阻塞：此接收操作會等待 goroutine 完成並向 ci 傳入值，因此主程序會一直等待直到兩個 goroutine 完成。
	x, y := <-ci, <-ci // receive from ci
	//同理任何傳送（ch<-5）將會被阻塞，直到資料被讀出
	ch := make(chan int, 3) //Buffered Channels
	ch <- 1
	ch <- 2
	close(ch) //關閉 channel 之後就無法再發送任何資料
	for i := range ch {
		fmt.Println(i)
	}
	fmt.Println(x, y, x+y)

	//select 預設是阻塞的，只有當監聽的 channel 中有傳送或接收可以進行時才會執行，當多個 channel 都準備好的時候，select 會隨機選擇其中一個執行。
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(5 * time.Second): //超時
				fmt.Println("timeout")
				o <- true
				break
				//default:
				// 當 c 阻塞的時候執行這裡
			}
		}
	}()
	c <- 5
	<-o
}

func testGoroutine() {
	//用來設定可以平行計算的 CPU 核數的最大值，並回傳之前的值。
	runtime.GOMAXPROCS(2)
	// 使用 WaitGroup 等待 goroutine 完成
	var wg sync.WaitGroup
	//計數器
	wg.Add(3)

	cs := make(chan string)

	go say("Hello", &wg, cs)
	go say("Hi", &wg, cs)
	say("From Main", &wg, cs)
	// 等待所有 goroutine 完成
	wg.Wait()
}

func say(s string, wg *sync.WaitGroup, c chan string) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		//主動讓出執行權限，使其他goroutine 有機會執行
		runtime.Gosched()
		fmt.Println(s)
		c <- s
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}
