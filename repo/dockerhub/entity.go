package dockerhub

type Response struct {
	Results interface{} `json:"results"`
	Total   int         `json:"total"`
}

type SearchRequest struct {
	Query      string `json:"query"`
	Categories string `json:"categories"`
	OpenSource string `json:"open_source"`
	Official   string `json:"official"`
	From       string `json:"from"`
	Size       string `json:"size"`
	RetryPages []int  `json:"retry_pages"`
}

type SearchItem struct {
	ID string `json:"id"`
}

type ImageCategoryData struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
