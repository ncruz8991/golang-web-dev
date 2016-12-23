package main

import (
	"github.com/ncruz8991/golang-web-dev/001_prereq/hands_on_2/secret_agent"
	"github.com/ncruz8991/golang-web-dev/001_prereq/hands_on_2/person"
)

func main() {
	human := person.GetInstance("Alexander Hamilton")
	spy := secret_agent.GetInstance("Hercules Mulligan", "John Smith")

	human.PSpeak()
	spy.PSpeak()
	spy.SASpeak()

}
