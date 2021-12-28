package main

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

import "fmt"

// Структура Human с полем Name и методом Sleep()
type Human struct {
	Name string
}

func (h Human) Sleep() {
	fmt.Println(h.Name + ": ...zzZzZZ")
}

// Структура Dog с полем Name и методом Speak()
type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println(d.Name + ": Bark!Bark!Bark!")
}

// Структура Action, в которую встроены Human и Dog
type Action struct {
	Human
	Dog
}

//  main() в разных файлах одного пакета только в рамках этого задания
func main() {
	h := Human{
		Name: "John",
	}
	d := Dog{
		Name: "Booch",
	}

	a := Action{
		Human: Human{
			Name: "Paul",
		},
		Dog: Dog{
			Name: "Champ",
		},
	}

	// Вызов методов Sleep() и Speak() у структур Human и Dog
	h.Sleep()
	d.Speak()
	// Output:
	//		John: ...zzZzZZ
	//		Booch: Bark!Bark!Bark!

	// У a Action можно вызывать методы встроенных структур
	a.Sleep()
	a.Speak()
	// Output:
	//		Paul: ...zzZzZZ
	//		Champ: Bark!Bark!Bark!
}
