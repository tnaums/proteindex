package proteinapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) CheckBlast(rid string) (Blastp, error) {
	fmt.Printf("Retrieving rid %s...\n", rid)
	baseURL := "https://blast.ncbi.nlm.nih.gov/Blast.cgi"
	blastResp := Blastp{}
	url := fmt.Sprintf(baseURL+"?RID=%s&CMD=%s&FORMAT_TYPE=%s", rid, "GET", "JSON2_S")

	if val, ok := c.cache.Get(url); ok {
		fmt.Println("found cache entry!")
		err := json.Unmarshal(val, &blastResp)
		if err != nil {
			return Blastp{}, err
		}
		return blastResp, nil
	}

	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Blastp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Blastp{}, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Blastp{}, err
	}

	err = json.Unmarshal(b, &blastResp)
	if err != nil {
		return Blastp{}, err
	}

	c.cache.Add(url, b)
	return blastResp, nil
}
