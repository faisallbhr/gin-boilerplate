package structs

type MetaParams struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limir"`
	Search string `json:"search"`
	SortBy string `json:"sort_by"`
	Order  string `json:"order"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type Sort struct {
	By    string `json:"by"`
	Order string `json:"order"`
}

type Meta struct {
	Search     string     `json:"search,omitempty"`
	Sort       Sort       `json:"sort"`
	Pagination Pagination `json:"pagination"`
}
