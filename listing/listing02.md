Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Программа выведет:
2
1

defer выполняется после оператора return
в функции test присутствует именованое поле для вывода результата и defer может работать с ним
в функции anotherTest x находится только внутри функции и после return недостижим для defer

```
