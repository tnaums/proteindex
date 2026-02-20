package proteinapi

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func (c *Client) SubmitBlast(protein, query string) (string, error) {

	if _, ok := c.cache.Get(protein); ok {
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
	baseURL, _ := url.Parse("https://blast.ncbi.nlm.nih.gov/Blast.cgi")
	//	baseURL := "https://blast.ncbi.nlm.nih.gov/Blast.cgi"
	paramsUrl := url.Values{}
	paramsUrl.Add("QUERY", params.Query)
	paramsUrl.Add("DATABASE", params.Database)
	paramsUrl.Add("PROGRAM", params.Program)
	paramsUrl.Add("CMD", params.Cmd)
	paramsUrl.Add("FORMAT", params.Format)
	baseURL.RawQuery = paramsUrl.Encode()
	//fmt.Println(baseURL.String()) // Output: https://example.com?key=value
	//	url := fmt.Sprintf(baseURL+"?QUERY=%s&DATABASE=%s&PROGRAM=%s&CMD=%s&FORMAT_TYPE=%s", params.Query, params.Database, params.Program, params.Cmd, params.Format)
	url := baseURL.String()
	fmt.Printf("url is: %s\n\n", url)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	rid := extractRid(doc)
	//	c.cache.AddRid(protein, rid)
	//	c.cache.PrintRids()
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
