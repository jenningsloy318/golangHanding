# 3. 变量

## 变量是什么
变量指定了某存储单元（Memory Location）的名称，该存储单元会存储特定类型的值。在 Go 中，有多种语法用于声明变量。

## 声明单个变量
  ```var name type``` 是声明单个变量的语法。

  ```go
  package main

  import "fmt"

  func main() {
      var age int // 变量声明
      fmt.Println("my age is", age)
  }
  ```
  语句` var age int` 声明了一个 int 类型的变量，名字为 age。我们还没有给该变量赋值。如果变量未被赋值，Go 会自动地将其初始化，赋值该变量类型的零值（Zero Value）。本例中 age 就被赋值为 0。如果你运行该程序，你会看到如下输出：


  ```my age is 0```
  变量可以赋值为本类型的任何值。上一程序中的 age 可以赋值为任何整型值（Integer Value）。

  ```go
  package main

  import "fmt"

  func main() {
      var age int // 变量声明
      fmt.Println("my age is", age)
      age = 29 // 赋值
      fmt.Println("my age is", age)
      age = 54 // 赋值
      fmt.Println("my new age is", age)
  }
  ```

  上面的程序会有如下输出：
  ```
  my age is  0  
  my age is 29  
  my new age is 54
  ```

##  声明变量并初始化

  声明变量的同时可以给定初始值。 `var name type = initialvalue` 的语法用于声明变量并初始化。
  ```go
  package main

  import "fmt"

  func main() {
      var age int = 29 // 声明变量并初始化

      fmt.Println("my age is", age)
  }
  ```
  在上面的程序中，age 是具有初始值 29 的 int 类型变量。如果你运行上面的程序，你可以看见下面的输出，证实 age 已经被初始化为 29。

  `my age is 29`

  ## 类型推断（Type Inference）
  如果变量有初始值，那么 Go 能够自动推断具有初始值的变量的类型。因此，如果变量有初始值，就可以在变量声明中省略 type。

  如果变量声明的语法是 `var name = initialvalue`，Go 能够根据初始值自动推断变量的类型。

  在下面的例子中，你可以看到在第 6 行，我们省略了变量 `age` 的 `int` 类型，Go 依然推断出了它是 int 类型。
  ```go
  package main

  import "fmt"

  func main() {
      var age = 29 // 可以推断类型

      fmt.Println("my age is", age)
  }
  ```
## 声明多个变量

  Go 能够通过一条语句声明多个变量。

  声明多个变量的语法是 `var name1, name2 type = initialvalue1, initialvalue2`。
  ```go
  package main

  import "fmt"

  func main() {
      var width, height int = 100, 50 // 声明多个变量

      fmt.Println("width is", width, "height is", heigh)
  }
  ```
  上述程序将在标准输出打印` width is 100 height is 50`。

  你可能已经想到，如果 `width` 和 `height` 省略了初始化，它们的初始值将赋值为 0。
  ```go
  package main

  import "fmt"

  func main() {  
      var width, height int
      fmt.Println("width is", width, "height is", height)
      width = 100
      height = 50
      fmt.Println("new width is", width, "new height is ", height)
  }
  ```
  上面的程序将会打印：
  ```shell
  width is 0 height is 0  
  new width is 100 new height is  50
  ```