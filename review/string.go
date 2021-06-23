package main

import (
	"bytes"
	"fmt"
)

func main() {
	// go语言中字符串是一个不可改变的字节序列, string的内部结构是一个指向值的指针ptr和长度值len
	s := "hello"
	s = "world"
	//第i个字节并不一定是字符串的第i个字符，因为对非ASCII字符的UTF8编码会是2个及以上字节
	var c rune = rune(s[0])

	//字符串的值是永远不可改变的，s只是持有了一个新的字符串，t仍然保有之前的字符串
	t := s
	s += "changed"
	// bytes包还提供了Buffer类型用于字节slice的缓存。

	var buf bytes.Buffer
	for _, v := range s {
		fmt.Fprintf(&buf, "%d", v)
	}

	fmt.Println(s, c, t, buf.String())
}
