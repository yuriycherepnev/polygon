package main

func main() {
	var fn func()

	fn()
}

/*
попытка запустить nil переменную - будет паника
Чтоб исправить,
надо функции что то присвоить:
fn = func() {
    fmt.Println(1)
}
*/
