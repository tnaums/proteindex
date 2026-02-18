package proteinapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (c *Client) CheckBlast(protein, query, rid string) error {
	fmt.Printf("\nRetrieving rid %s...\n", rid)
	fmt.Print("Proteindex > ")	
	baseURL := "https://blast.ncbi.nlm.nih.gov/Blast.cgi"
	url := fmt.Sprintf(baseURL+"?RID=%s&CMD=%s&FORMAT_TYPE=%s", rid, "GET", "JSON2_S")

	for {
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
			continue
		} else {
			c.cache.Add(protein, b)
			return nil
		}

	}
}
