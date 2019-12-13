#   35. 读取文件

文件读取是所有编程语言中最常见的操作之一。本教程我们会学习如何使用 Go 读取文件。

本教程分为如下小节。

- 将整个文件读取到内存
  - 使用绝对文件路径
  - 使用命令行标记来传递文件路径
  - 将文件绑定在二进制文件中
- 分块读取文件
- 逐行读取文件

## 将整个文件读取到内存
将整个文件读取到内存是最基本的文件操作之一。这需要使用 `ioutil` 包中的 `ReadFile` 函数。

让我们在 Go 程序所在的目录中，读取一个文件。我已经在当前目录中创建了一个文本文件 `test.txt`，我们会使用 Go 程序 来读取它。`test.txt` 包含文本 `“Hello World. Welcome to file handling in Go”`。


接下来我们来看看代码。

```go{.line-numbers}
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:", string(data))
}
```

该程序会输出：
```sh
Contents of file: Hello World. Welcome to file handling in Go.
```

如果在其他位置运行这个程序（没有这个文件)，会打印下面的错误。
```sh
File reading error open test.txt: The system cannot find the file specified.
```


由于在运行二进制文件的位置上没有找到 test.txt，因此程序会报错，提示无法找到指定的文件。

有三种方法可以解决这个问题。

1. 使用绝对文件路径
2. 使用命令行标记来传递文件路径

让我们来依次介绍。
### 1. 使用绝对文件路径

要解决问题，最简单的方法就是传入绝对文件路径。我已经修改了程序，把路径改成了绝对路径。

```go{.line-numbers}
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    data, err := ioutil.ReadFile("/home/jenningsl/development/projects/golang-notes/Golang-tutorial-series/chap35/test.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:", string(data))
}
```

现在可以在任何位置上运行程序，打印出 `test.txt` 的内容。

例如，可以在我的家目录运行。

该程序打印出了 test.txt 的内容。

看似这是一个简单的方法，但它的缺点是：文件必须放在程序指定的路径中，否则就会出错。

### 2. 使用命令行标记来传递文件路径
另一种解决方案是使用命令行标记来传递文件路径。使用 `flag` 包，我们可以从输入的命令行获取到文件路径，接着读取文件内容。

首先我们来看看 `flag` 包是如何工作的。`flag` 包有一个名为 `String` 的函数。该函数接收三个参数。第一个参数是标记名，第二个是默认值，第三个是标记的简短描述。

让我们来编写程序，从命令行读取文件名。
```go{.line-numbers}

package main
import (
    "flag"
    "fmt"
)

func main() {
    fptr := flag.String("fpath", "test.txt", "file path to read from")
    flag.Parse()
    fmt.Println("value of fpath is", *fptr)
}
```

在上述程序中第 8 行，通过 `String` 函数，创建了一个字符串标记，名称是 `fpath`，默认值是 `test.txt`，描述为 `file path to read from`。这个函数返回存储 flag 值的字符串变量的地址。

在程序访问 flag 之前，必须先调用 `flag.Parse()`。

在第 10 行，程序会打印出 flag 值。

使用下面命令运行程序。
```sh
example2 -fpath=/path-of-file/test.txt
```
我们传入 `/path-of-file/test.txt`，赋值给了 fpath 标记。

该程序输出：
```
value of fpath is /path-of-file/test.txt
```
这是因为 fpath 的默认值是 `test.txt`。

现在我们知道如何从命令行读取文件路径了，让我们继续完成我们的文件读取程序。

```go{.line-numbers}
package main
import (
    "flag"
    "fmt"
    "io/ioutil"
)

func main() {
    fptr := flag.String("fpath", "test.txt", "file path to read from")
    flag.Parse()
    data, err := ioutil.ReadFile(*fptr)
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:", string(data))
}
```

在上述程序里，命令行传入文件路径，程序读取了该文件的内容。使用下面命令运行该程序。
```
example3 -fpath=/path-of-file/test.txt
```
请将 `/path-of-file/ `替换为 `test.txt` 的真实路径。该程序将打印：
```
Contents of file: Hello World. Welcome to file handling in Go.
```



## 分块读取文件
在前面的章节，我们学习了如何把整个文件读取到内存。当文件非常大时，尤其在 RAM 存储量不足的情况下，把整个文件都读入内存是没有意义的。更好的方法是分块读取文件。这可以使用 bufio 包来完成。

让我们来编写一个程序，以 3 个字节的块为单位读取 test.txt 文件。如下所示，替换内容。

```go{.line-numbers}
package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
)

func main() {
    fptr := flag.String("fpath", "test.txt", "file path to read from")
    flag.Parse()

    f, err := os.Open(*fptr)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = f.Close(); err != nil {
            log.Fatal(err)
        }
    }()
    r := bufio.NewReader(f)
    b := make([]byte, 3)
    for {
        _, err := r.Read(b)
        if err != nil {
            fmt.Println("Error reading file:", err)
            break
        }
        fmt.Println(string(b))
    }
}
```

在上述程序的第 15 行，我们使用命令行标记传递的路径，打开文件。

在第 19 行，我们延迟了文件的关闭操作。

在上面程序的第 24 行，我们新建了一个缓冲读取器`（buffered reader）`。在下一行，我们创建了长度和容量为 `3` 的字节切片，程序会把文件的字节读取到切片中。

第 27 行的 `Read` 方法会读取 `len(b)` 个字节（达到 3 字节），并返回所读取的字节数。当到达文件最后时，它会返回一个 EOF 错误。程序的其他地方比较简单，不做解释。

如果我们运行程序会得到以下输出：
```{.line-number}
Hel
lo
Wor
ld.
 We
lco
me
to
fil
e h
and
lin
g i
n G
o.
Error reading file: EOF
```


## 逐行读取文件
本节我们讨论如何使用 Go 逐行读取文件。这可以使用 `bufio` 来实现。

请将 `test.txt` 替换为以下内容。
```
Hello World. Welcome to file handling in Go.
This is the second line of the file.
We have reached the end of the file.
```
逐行读取文件涉及到以下步骤。

1. 打开文件；
2. 在文件上新建一个 scanner；
3. 扫描文件并且逐行读取。
```go{.line-numbers}
package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
)

func main() {
    fptr := flag.String("fpath", "test.txt", "file path to read from")
    flag.Parse()

    f, err := os.Open(*fptr)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = f.Close(); err != nil {
        log.Fatal(err)
    }
    }()
    s := bufio.NewScanner(f)
    for s.Scan() {
        fmt.Println(s.Text())
    }
    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}
```

在上述程序的第 15 行，我们用命令行标记传入的路径，打开文件。在第 24 行，我们用文件创建了一个新的 `scanner`。第 25 行的 `Scan()` 方法读取文件的下一行，如果可以读取，就可以使用 Text() 方法。

当 Scan 返回 false 时，除非已经到达文件末尾（此时 `Err()` 返回 `nil`），否则 `Err()`就会返回扫描过程中出现的错误。

程序会输出：
```
Hello World. Welcome to file handling in Go.
This is the second line of the file.
We have reached the end of the file.
```

