package internal

import (
	"sort"
)

// SortServicesByName sorts the services slice by Name in ascending or descending order
func SortServicesByName(services []Service, ascending bool) {
	sort.Slice(services, func(i, j int) bool {
		if ascending {
			return services[i].Name < services[j].Name
		}
		return services[i].Name > services[j].Name
	})
}

// SortServicesByCreatedAt sorts the services slice by CreatedAt in ascending or descending order
func SortServicesByCreatedAt(services []Service, ascending bool) {
	sort.Slice(services, func(i, j int) bool {
		if ascending {
			return services[i].CreatedAt < services[j].CreatedAt
		}
		return services[i].CreatedAt > services[j].CreatedAt
	})
}

// SortServicesByIdAsc sorts the service slice by ID in ascending order
func SortServicesByIdAsc(services []Service, ascending bool) {
	sort.Slice(services, func(i, j int) bool {
		if ascending {
			return services[i].ID < services[j].ID
		}
		return services[i].ID > services[j].ID
	})
}

func setDefaultValues(filter *Filter) {
	if filter.Sort.Column == "" {
		filter.Sort.Column = "created_at"
	}

	if filter.Sort.Order == "" {
		filter.Sort.Column = "asc"
	}

	if filter.Pagination.Limit == 0 {
		filter.Pagination.Limit = 10
	}
}
