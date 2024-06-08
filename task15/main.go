package main

import (
	"fmt"
)

/*
1. Использование глобальной переменной justString. Недостатками использования глобальных переменных являются:
- неконтролируемый жизненный цикл
- невозможность иметь два экземпляра без изменения кода, использующего их
- усложнение читаемости кода из-за неявных связей
- неконтролируемый доступ

2. Когда в исходном коде мы присваиваем глобальной переменной justString значение v[:100]



*/


func main() {
    var justString string = SomeFunc()
    fmt.Print(justString)
    
  }

  
func SomeFunc() string {
    v := createHugeString(1 << 10)
    fmt.Printf("%T", v[:100])
    return string(v[:100])  
}
  
func createHugeString(size int) string {
    return string(make([]rune, size))
}