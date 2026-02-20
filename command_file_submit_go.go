package main

import (
	"io/ioutil"
)

func commandFileSubmitGo(cfg *config, args ...string) error {
	data, err := ioutil.ReadFile("sequences/Afca.pep")
	if err != nil {
		return err
	}
	content := string(data)
	go commandSubmit(cfg, args[0], content)
	return nil

}
