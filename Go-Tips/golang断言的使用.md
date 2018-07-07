 golang断言的使用（Type Assertion） 


 golang的语言中提供了断言的功能。golang中的所有程序都实现了`interface{}`的接口，这意味着，所有的类型如`string,int,int64`甚至是自定义的`struct`类型都就此拥有了`interface{}`的接口，这种做法和java中的Object类型比较类似。

 ## 类型断言有以下几种形式：
1. 直接断言使用
```go
var a interface{}

fmt.Println("Where are you,Jonny?", a.(string))
```

但是如果断言失败一般会导致panic的发生。所以为了防止panic的发生，我们需要在断言前进行一定的判断
```go
value, ok := a.(string)
```

如果断言失败，那么`ok`的值将会是`false`,但是如果断言成功`ok`的值将会是`true`,同时`value`将会得到所期待的正确的值。示例：
```go
value, ok := a.(string)
if !ok {
    fmt.Println("It's not ok for type string")
    return
}
fmt.Println("The value is ", value)

```
 另外也可以配合switch语句进行判断： 
 ```go
 var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T", t)       // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```