package do

type DataChainDo struct {
	Results []string `json:"results"`
	Page    int      `json:"page"`
	Total   int      `json:"total"`
}
