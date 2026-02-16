package proteinapi

import (
	"encoding/json"
	"errors"
	//	"fmt"
)

// Catch Protein
func (c *Client) CatchProtein(name string) (Blastp, error){

	if val, ok := c.cache.Get(name); ok {
		//		fmt.Printf("%s\n\n", val)
		proteinData := Blastp{}
		err := json.Unmarshal(val, &proteinData)
		if err != nil {
			return Blastp{}, err
		}

		return proteinData, nil
	}
	return Blastp{}, errors.New("Protein not in cache.")
}
