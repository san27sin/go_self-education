/*
86. Горутины
Горутина — это легковесный поток, управляемый средой выполнения Go.

go f(x, y, z)

запускает новую горутину, выполняющую

f(x, y, z)

Вычисление f, x, y и z происходит в текущей горутине, а выполнение f — в новой горутине.

Горутины выполняются в одном адресном пространстве, поэтому доступ к общей памяти должен быть синхронизирован.
Пакет sync предоставляет полезные примитивы, хотя они нечасто понадобятся вам в Go, поскольку существуют другие примитивы.
(Смотрите следующий слайд.)
*/

package lessons

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
    }
}

func RunLesson86() {
    // go say("world")
    // time.Sleep(150 * time.Millisecond) // успевает напечатать только одну, потому что main заканчивает свою работу раньше
    // say("hello")
    var wg sync.WaitGroup

    wg.Add(1)
    go say("a", &wg)
    wg.Add(1)
    go say("b", &wg)
    wg.Add(1)
    go say("c", &wg)
    wg.Wait()
}