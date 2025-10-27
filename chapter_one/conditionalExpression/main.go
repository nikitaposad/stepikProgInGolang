/*
Условные выражения в ЯП включают операции отношения и логические операции.
Они позволяют создавать условия,которые возвращают логические значение типа
bool:true(условие истинно) or false(условие ложное).

Операции отношения.
Операции отношения позволяют сравнить два значения.
В языке Golang доступны следующие операции отношения:
- == : Операция "равно".Возвращает true,если оба операнда равны, и false
если они не равны:

package main
import "fmt"

func main(){
	var a int = 8
	var b int = 3
	var c bool = a == b
	fmt.Println(c) //false
}
*/
/*
- >: Операция "больше чем".Возвращает true, если первый операнд больше второго
и false,если первый операнд меньше или равен второму:

package main
import "fmt"

func main(){
	var a int = 8
	var b int = 3
	var c bool = a > b
	fmt.Println(c) //true
}

- <: Операция "меньше чем".Возвращает true,если первый операнд меньше второго,
и false,если первый операнд больше или равен второму:

package main
import "fmt"

func main(){
	var a int = 8
	var b int = 3
	var c bool = a < b
	fmt.Println(c) //false
}
*/
/*
- <=: Операция "меньше или равно".Возвращает true,если первый операнд
меньше или равен второму,и false,если первый операнд больше второго:

package main
import "fmt"

func main(){
	var a int = 8
	var b int = 3
	var c bool = a <= b
	fmt.Println(c) //false
}
*/
/*
- >=: Операция "больше или равно".Возвращает true,если первый операнд
больше или равен второму,и false,если первый операнд меньше второго:

package main
import "fmt"

func main(){
	var a int = 8
	var b int = 3
	var c bool = a >= b
	fmt.Println(c) //true
}
*/
/*
- !=: Операция "не равно".Возвращает true,если первый операнд не равен второму,
и false,если они равны:

package main
import "fmt"

func main(){
	var a int = 8
	var b int = 3
	var c bool = a != b //true
	var d bool = a != 8 //false
	fmt.Println(c,d)
}
*/
