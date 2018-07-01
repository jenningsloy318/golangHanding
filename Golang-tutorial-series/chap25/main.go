package main  
import (  
    "fmt"
    "sync"
    )
    
var x  = 0  
var x2  = 0  

func increment(wg *sync.WaitGroup, m *sync.Mutex) {  
    m.Lock()
    x = x + 1
    m.Unlock()
    wg.Done()   
}

func increment2(wg *sync.WaitGroup, ch chan bool) {  
    ch <- true
    x2 = x2 + 1
    <- ch
    wg.Done()   
}
func main() {  

    var w sync.WaitGroup
    var m sync.Mutex
    for i := 0; i < 1000; i++ {
        w.Add(1)        
        go increment(&w, &m)
    }
    w.Wait()
    fmt.Println("final value of x is (using Mutex) ", x)
    // use chan 

    var w2 sync.WaitGroup

    ch := make(chan bool, 1)
    for i := 0; i < 1000; i++ {
        w2.Add(1)        
        go increment2(&w2, ch)
    }
    w2.Wait()
    fmt.Println("final value of x is (using chan)", x)
}