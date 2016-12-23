package secret_agent

import (
	"github.com/ncruz8991/golang-web-dev/001_prereq/hands_on_2/person"
	"fmt"
)

func GetInstance(realName, disguiseName string) SecretAgent {
	return SecretAgent{realName, person.GetInstance(disguiseName)}
}

// Embedding Person lets SecretAgent "inherit" from Person. Specifically, this inherits
// Pspeak() from Person.
// Inheritance works differently in Go--calling PSpeak() from a secret agent will call the PSpeak() function on the
// embedded Person object.
type SecretAgent struct {
	realName string
	person.Person
}

func (sa SecretAgent) SASpeak() {
	fmt.Println("I am actually a secret agent named " + sa.realName)
}