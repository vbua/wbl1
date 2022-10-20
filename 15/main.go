package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

/*
проблемы:
1. глобальная переменная - плохой тон, так как возникает больший риск допуска ошибок и становится сложнее тестировать.
2. при слайсе строки побайтово может обрезаться символ, который занимает большой байта. надо перевести в слайс рун и
тогда уже обрезать

P.S. слайс рун будет ссылаться на полный массив, но когда мы переведем его в строку, то связь будет утеряна.
можно убедиться в этом если сделать fmt.Println(cap([]rune(justString))), то увидим капасити не 1024, а меньше.
*/

func createHugeString(i int) string {
	runes := make([]rune, i)
	for l := 0; l < i; l++ {
		runes[l] = 'Ж'
	}
	return string(runes)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	runes := []rune(v)
	justString := string(runes[:100])
	return justString
}

func main() {
	someFunc()
}
