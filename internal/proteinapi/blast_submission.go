package proteinapi

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func (c *Client) SubmitBlast(name, query string) (string, error) {

	if _, ok := c.cache.Get(query); ok {
		//		fmt.Println("blastp results already in cache")
		return "foundit", nil
	}

	params := BlastParams{
		Cmd:      "Put",
		Query:    query,
		Database: "swissprot",
		Program:  "blastp",
		Format:   "JSON2",
	}

	baseURL := "https://blast.ncbi.nlm.nih.gov/Blast.cgi"

	url := fmt.Sprintf(baseURL+"?QUERY=%s&DATABASE=%s&PROGRAM=%s&CMD=%s&FORMAT_TYPE=%s", params.Query, params.Database, params.Program, params.Cmd, params.Format)

	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		//		fmt.Println("Error getting response")
		return "", err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		//		fmt.Println("Error parsing HTML:", err)
		return "", err
	}

	rid := extractRid(doc)
	c.cache.AddRid(query, rid)
	c.cache.PrintRids()
	return rid, nil
}

// Extract the RID value from the blast request HTML response.
func extractRid(n *html.Node) string {
	returnValue := ""
	if n.Type == html.ElementNode && n.Data == "form" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if strings.HasPrefix(c.Data, "QBlast") {
				parts := strings.Fields(c.Data)
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
