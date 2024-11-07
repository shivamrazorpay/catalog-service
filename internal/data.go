package internal

var DataBase = map[string]interface{}{
	"Services": []Service{
		{
			ID:            "1",
			Name:          "Locate Us",
			Description:   "Locate Us Service",
			LatestVersion: "1.0.1",
			Versions:      []string{"1.0.0", "1.0.1"},
			CreatedAt:     1612137600,
			UpdatedAt:     1612137600,
		},
		{
			ID:            "2",
			Name:          "Collect Money",
			Description:   "Collect Money Service Collects Money",
			LatestVersion: "3.0.1",
			Versions:      []string{"3.0.0", "3.0.1"},
			CreatedAt:     1612137600,
			UpdatedAt:     1612137600,
		},
	},
}
