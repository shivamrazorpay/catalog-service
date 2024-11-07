package internal

type Service struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	LatestVersion string   `json:"latest_version"`
	Versions      []string `json:"versions,omitempty"`
	CreatedAt     int64    `json:"created_at"`
	UpdatedAt     int64    `json:"updated_at"`
}

type Filter struct {
	Search     Search     `json:"search"`
	Pagination Pagination `json:"pagination"`
	Sort       Sort       `json:"sort"`
}

type Search struct {
	Column string `json:"column"`
	Value  string `json:"value"`
}

type Sort struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}

type Pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ServiceVersionResponse struct {
	Versions []string `json:"versions"`
	Count    int      `json:"count"`
}
