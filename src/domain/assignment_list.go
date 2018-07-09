package domain

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type AssigmentList struct {
	List []Assigment
}

func (this *AssigmentList) Run() error {
	for _, assigenment := range this.List {
		assigenment.Run()
	}

	return nil
}

func (this *AssigmentList) Load(fileName string) error {
	path := "./conf/" + fileName
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return err
	}

	err = yaml.Unmarshal(content, &this.List)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return err
	}

	log.Print(this.List)
	return nil
}
