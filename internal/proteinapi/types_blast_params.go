package proteinapi

// Blast Params for Submission
type BlastParams struct {
	Cmd      string
	Query    string
	Database string
	Program  string
	Format   string
}

type Blastp struct {
	BlastOutput2 []struct {
		Report struct {
			Program      string `json:"program"`
			Version      string `json:"version"`
			Reference    string `json:"reference"`
			SearchTarget struct {
				Db string `json:"db"`
			} `json:"search_target"`
			Params struct {
				Matrix    string `json:"matrix"`
				Expect    int    `json:"expect"`
				GapOpen   int    `json:"gap_open"`
				GapExtend int    `json:"gap_extend"`
				Filter    string `json:"filter"`
				Cbs       int    `json:"cbs"`
			} `json:"params"`
			Results struct {
				Search struct {
					QueryID    string `json:"query_id"`
					QueryTitle string `json:"query_title"`
					QueryLen   int    `json:"query_len"`
					Hits       []struct {
						Num         int `json:"num"`
						Description []struct {
							ID        string `json:"id"`
							Accession string `json:"accession"`
							Title     string `json:"title"`
							Taxid     int    `json:"taxid"`
							Sciname   string `json:"sciname"`
						} `json:"description"`
						Len  int `json:"len"`
						Hsps []struct {
							Num       int     `json:"num"`
							BitScore  float64 `json:"bit_score"`
							Score     int     `json:"score"`
							Evalue    float64     `json:"evalue"`
							Identity  int     `json:"identity"`
							Positive  int     `json:"positive"`
							QueryFrom int     `json:"query_from"`
							QueryTo   int     `json:"query_to"`
							HitFrom   int     `json:"hit_from"`
							HitTo     int     `json:"hit_to"`
							AlignLen  int     `json:"align_len"`
							Gaps      int     `json:"gaps"`
							Qseq      string  `json:"qseq"`
							Hseq      string  `json:"hseq"`
							Midline   string  `json:"midline"`
						} `json:"hsps"`
					} `json:"hits"`
					Stat struct {
						DbNum    int     `json:"db_num"`
						DbLen    int     `json:"db_len"`
						HspLen   int     `json:"hsp_len"`
						EffSpace int64   `json:"eff_space"`
						Kappa    float64 `json:"kappa"`
						Lambda   float64 `json:"lambda"`
						Entropy  float64 `json:"entropy"`
					} `json:"stat"`
				} `json:"search"`
			} `json:"results"`
		} `json:"report"`
	} `json:"BlastOutput2"`
}
