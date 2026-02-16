package proteinapi

import (
	"bytes"
	//	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (c *Client) CheckBlast(protein, query, rid string) error {
	fmt.Printf("Retrieving rid %s...\n", rid)
	baseURL := "https://blast.ncbi.nlm.nih.gov/Blast.cgi"
	//	blastResp := Blastp{}
	url := fmt.Sprintf(baseURL+"?RID=%s&CMD=%s&FORMAT_TYPE=%s", rid, "GET", "JSON2_S")

	for {
		fmt.Println("Waiting 20 seconds...")
		duration := time.Duration(20) * time.Second
		time.Sleep(duration)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		prefix := []byte{'<', '!', 'D', 'O', 'C'}
		if bytes.HasPrefix(b, prefix) {
			fmt.Println("Results not yet available...")
		} else {
			c.cache.Add(protein, b)
			fmt.Println("Added to cache!")
			return nil
		}

	}
}
