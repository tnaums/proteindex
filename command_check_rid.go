package main

import "fmt"

func commandCheckRid(cfg *config, args ...string) error {
	fmt.Println("Checking rid...")
	blastResp, err := cfg.proteinapiClient.CheckBlast(args[0])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(blastResp)
	return nil
}
