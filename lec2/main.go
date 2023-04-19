package main

import "fmt"

func main() {
	fmt.Printf("Привет,  я новый студент\n")
	fmt.Printf("ПРивет, привет\n Меня зовут %s\n, Мне  %d\n", "MAX", 40)
	// декларирование переменных -  I/O
	// 1. В языке го принята полустрогая статическая типизация
	// 2. Способы декларирования переменных:
	// При декларировании переменной автоматичемски инициализируется её нуовое значение
	var age int
	fmt.Print("my age correctly =", age)
	age = 39
	println()
	fmt.Println("wright age after assignment =", age)
	// декларирование и инициализация пользовательским значением
	var height int = 179
	fmt.Println("my height", height)
	// в чем полустрогость типизации ? - для явных значение можно не указывать тип, Go  догадается
	var uid = 1547
	fmt.Println("My uid- ", uid)
	// деклаиррование и иниц переменных одного типа

	var firstVar, secondVar int = 20, 30

	fmt.Println(" переменные заданы", firstVar, secondVar)
	// через бллок переменных
	var (
		personName string = "max"
		personeAge int    = 15
		personeUid        = 1548
	)
	fmt.Printf("Name: %s\n- Age %d\n - UID = %d\n", personName, personeAge, personeUid)
	// немного странного
	var a, b = 30, "vovka"
	fmt.Println(a, b)
	// Сахарок -повтоорное декларирование приводит к ошибке
	// var a = 200

	// короткое объявление

	counter := 10
	fmt.Println(counter)
	counter++
	fmt.Println(counter)

	aArg, bArg := 10, 5
	fmt.Println(aArg, bArg)
	aArg, bArg = 10, 20
	fmt.Println(aArg, bArg)

	// fmt.Println(aArg, bArg)  - работать не будет, однако если добавить новую переменную, вс езаработает
	cArg, bArg := 1000, 500
	fmt.Println(aArg, bArg, cArg)

	var word1, word2, word3, word4 = "имя", "твоё", "мне", "знакомо"

	fmt.Println(word4, word3, word2, word1)
	fmt.Println(word3, word1, word4, word2)
	fmt.Println(word2, word4, word1, word3)

}
