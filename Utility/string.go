package Utility

/* strings 套件

func Contains(s, substr string) bool
字串 s 中是否包含 substr，回傳 bool 值

func Join(a []string, sep string) string
字串連結，把 slice a 透過 sep 連結起來

func Index(s, sep string) int
在字串 s 中查詢 sep 所在的位置，回傳位置值，找不到回傳-1

func Repeat(s string, count int) string
重複 s 字串 count 次，最後回傳重複的字串

func Replace(s, old, new string, n int) string
在 s 字串中，把 old 字串替換為 new 字串，n 表示替換的次數，小於 0 表示全部替換

func Split(s, sep string) []string
把 s 字串按照 sep 分割，回傳 slice

func Trim(s string, cutset string) string
在 s 字串的頭部和尾部去除 cutset 指定的字串

func Fields(s string) []string
去除 s 字串的空格符，並且按照空格分割回傳 slice

strconv：字串轉化的函式
包括 Append, Format, Parse

*/
