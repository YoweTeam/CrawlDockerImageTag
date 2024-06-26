package dto

type SearchRequest struct {
	Query      string `json:"query"`
	Categories string `json:"categories"`
	OpenSource string `json:"open_source"`
	Official   string `json:"official"`
	From       string `json:"from"`
	Size       string `json:"size"`
}
