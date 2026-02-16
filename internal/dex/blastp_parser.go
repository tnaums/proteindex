package dex

import (
	"bufio"
	"fmt"
	"os"
)

func ParseBlastp(b Protein) {
	for _, hit := range b.Blast.BlastOutput2[0].Report.Results.Search.Hits {
		fmt.Println()
		fmt.Printf("Hit %d\n", hit.Num)

		for _, d := range hit.Description {
			fmt.Println("Description:")
			fmt.Printf("    ID: %s\n", d.ID)
			fmt.Printf("    Accession: %s\n", d.Accession)
			fmt.Printf("    Title: %s\n", d.Title)
			fmt.Printf("    Taxid: %d\n", d.Taxid)
			fmt.Printf("    %s\n", d.Sciname)
		}
		fmt.Printf("Length: %d\n", hit.Len)
		for _, hsp := range hit.Hsps {
			fmt.Printf("Hsp: %d\n", hsp.Num)
			fmt.Printf("    Bitscore %e\n", hsp.BitScore)
			fmt.Printf("    Score %d\n", hsp.Score)
			fmt.Printf("    Evalue %f\n", hsp.Evalue)
			fmt.Printf("    Identity %d\n", hsp.Identity)
			fmt.Printf("    Posititve %d\n", hsp.Positive)
			fmt.Printf("    Query %d -> %d\n", hsp.QueryFrom, hsp.QueryTo)
			fmt.Printf("    Hit %d -> %d\n", hsp.HitFrom, hsp.HitTo)
			fmt.Printf("    AlignLen %d\n", hsp.AlignLen)
			fmt.Printf("    Gaps %d\n", hsp.Gaps)
			printAlignment(hsp.Qseq, hsp.Hseq, hsp.Midline)

		}
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("press Enter to continue: ")
		for scanner.Scan() {
			scanner.Text()
			break
		}
	}

}

func printAlignment(q, h, m string) {
	query := ""
	hit := ""
	midline := ""
	for idx, x := range q {
		if idx == 0 {
			query += string(x)
			hit += string(h[0])
			midline += string(m[0])
			continue
		}
		if idx%70 == 0 {
			fmt.Println()
			fmt.Printf("    %s\n", query)
			fmt.Printf("    %s\n", midline)
			fmt.Printf("    %s\n", hit)

			query = string(x)
			hit = string(h[idx])
			midline = string(m[idx])
			continue
		}
		query += string(x)
		hit += string(h[idx])
		midline += string(m[idx])
	}
	fmt.Println()
	fmt.Printf("    %s\n", query)
	fmt.Printf("    %s\n", midline)
	fmt.Printf("    %s\n", hit)

}
