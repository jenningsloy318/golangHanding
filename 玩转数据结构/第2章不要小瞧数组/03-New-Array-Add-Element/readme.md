
Note:

##1. 值类型和引用类型

这一节除了定义的新的 **`array结构体`** 以外，还加深了golang **值类型**和**引用类型**的理解， 包括数组和结构体都是值类型，故通过方法来改变这个结构体的字段值时，应该使用的是**struct变量的指针**，而**非struct变量**

故如果修改 `a`内的值的时候，使用指针，就会达到同时更改`a.data` 和`a.size`  
  ```
  func (a *Array) Add(index int, element int)  
  ```

如果不使用指针，以下函数没有办法更改`a.size`,但是还是可以更改`a.data`,由于`a.data`是slice，slice是引用类型，底层是指针，故可以更改`a.data`，但是不能更改`a.size`, 由于单纯的int变量也是属于值类型的
```
func (a Array) Add(index int, element int)  
```


- 值类型：所有像`int`、`float`、`bool`、`string`、 `struct`, `array`这些类型都属于值类型，使用这些类型的变量直接指向存在内存中的值，值类型的变量的值存储在栈中。当使用等号=将一个变量的值赋给另一个变量时，如 j = i ,实际上是在内存中将 i 的值进行了拷贝。可以通过 &i 获取变量 i 的内存地址
- 引用类型：像`slice`,`map`,`channel`,`指针(point)`，`函数(function)`都是引用类型。比如`slice`自己不拥有任何数据。它只是底层`array`的一种表示,对切片所做的任何修改都会反映在底层数组中,如果多个`slice`都是引用同一个`array`,任意一个`slice`的更改都会改动底层的`array`。


