package person

import "fmt"

func GetInstance(name string) Person {
	return Person{name}
}

type Person struct {
	name string
}

func (p Person) PSpeak() {
	fmt.Println("Hello I am a normal person named " + p.name)
}