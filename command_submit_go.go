package main

import "errors"

func commandSubmitGo(cfg *config, args ...string) error {
	if len(args) != 2 {
		return errors.New("blastp command requires two rguments: blastp <name> <sequence>")
	}
	go commandSubmit(cfg, args[0], args[1])
	return nil
}
