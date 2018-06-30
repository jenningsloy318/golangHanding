package main

import (
    "fmt"
    "unicode/utf8"
)
func printBytes(s string) {
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
}

func printChars(s string) {
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%c ",s[i])
    }
}
func printRuneChars(s string) {
    runes := []rune(s)
    for i:= 0; i < len(runes); i++ {
        fmt.Printf("%c ",runes[i])
    }
}
func printCharsAndBytes(s string) {
    for index, value := range s {
        fmt.Printf("type of value is %T\n",value)
        fmt.Printf("%c starts at byte %d\n", value, index)
    }
}
func length(s string) {  
    fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}
func mutate(s []rune) string {  
    s[0] = 'a' 
    return string(s)
}
func main() {
    // print string
    name := "Hello World"
    fmt.Println(name)

    //print each character  in byte code
    printBytes(name)
    // print each character 
    fmt.Printf("\n")
    printChars(name)
    newname := "Señor"
    // print each character  with rune
    fmt.Printf("\n")
    printRuneChars(newname)

    // 使用for range 遍历， 如果使用for range进行遍历，会自动将字符转换为rune类型
    fmt.Printf("\n")
    newname2 := "Señor"
    printCharsAndBytes(newname2) 

    //用字节切片构造字符串
    byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
    str := string(byteSlice)
    fmt.Println(str)
    //用rune切片构造字符串
    runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str2 := string(runeSlice)
    fmt.Println(str2)

    //字符串的长度
    word1 := "Señor" 
    length(word1)
    word2 := "Pets"
    length(word2)
    //修改字符串
    h := "hello"
    fmt.Println(mutate([]rune(h)))
    fmt.Print(h)
}
