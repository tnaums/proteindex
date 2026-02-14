package proteinapi

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func(c *Client) SubmitBlast(query string) string {
	params := BlastParams{
		Cmd:      "Put",
		Query:    query,
		Database: "swissprot",
		Program:  "blastp",
		Format:   "JSON2",
	}

	baseURL := "https://blast.ncbi.nlm.nih.gov/Blast.cgi"

	url := fmt.Sprintf(baseURL+"?QUERY=%s&DATABASE=%s&PROGRAM=%s&CMD=%s&FORMAT_TYPE=%s", params.Query, params.Database, params.Program, params.Cmd, params.Format)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting response")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
	}

	rid := extractRid(doc)
	c.cache.AddRid(query, rid)
	c.cache.PrintRids()
	return rid
}

// Extract the RID value from the blast request HTML response.
func extractRid(n *html.Node) string {
	returnValue := ""
	if n.Type == html.ElementNode && n.Data == "form" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if strings.HasPrefix(c.Data, "QBlast") {
				parts := strings.Fields(c.Data)
				fmt.Printf("In extractRid: %s\n", parts[3])
				return parts[3]
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		returnValue = extractRid(c)
		if returnValue != "" {
			break
		}
	}
	return returnValue
}
