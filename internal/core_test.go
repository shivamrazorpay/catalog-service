package internal

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

// Mock Database
var mockServices = []Service{
	{ID: "BbItY3V9rV", Name: "Service One", Description: "First service", LatestVersion: "v1.0", Versions: []string{"v1.0", "v1.1"}},
	{ID: "6Zt5mpdDsQ", Name: "Service Two", Description: "Second service", LatestVersion: "v2.0", Versions: []string{"v2.0"}},
	{ID: "PQt5mpd111", Name: "Contact Us", Description: "Contact Us Service", LatestVersion: "v1.0", Versions: []string{"v1.0"}},
}

// Mock the database structure for testing
func init() {
	DataBase = map[string]interface{}{
		"Services": mockServices,
	}
}

func TestGetServicesCore(t *testing.T) {
	tests := []struct {
		name       string
		want       []Service
		filter     Filter
		wantErr    bool
		errMessage string
	}{
		{
			name:    "Valid services",
			filter:  Filter{},
			want:    mockServices,
			wantErr: false,
		},
		{
			name: "InValid Limit",
			filter: Filter{
				Pagination: Pagination{
					Limit:  100,
					Offset: 0,
				},
			},
			wantErr:    true,
			errMessage: "filter.pagination.limit must be no greater than 20",
		},
		{
			name: "VeryBig Offset",
			filter: Filter{
				Pagination: Pagination{
					Limit:  10,
					Offset: 100,
				},
			},
			wantErr: false,
			want:    []Service{},
		},
		{
			name: "Valid Limit and Offset",
			filter: Filter{
				Pagination: Pagination{
					Limit:  2,
					Offset: 1,
				},
			},
			want:    mockServices[1:],
			wantErr: false,
		},
		{
			name: "Valid Asc Sort by Name",
			filter: Filter{
				Sort: Sort{
					Column: "name",
					Order:  "asc",
				},
			},
			want:    []Service{mockServices[2], mockServices[0], mockServices[1]},
			wantErr: false,
		},
		{
			name: "Valid Desc Sort by Name",
			filter: Filter{
				Sort: Sort{
					Column: "name",
					Order:  "desc",
				},
			},
			want:    []Service{mockServices[1], mockServices[0], mockServices[2]},
			wantErr: false,
		},
		{
			name: "Invalid Sort Column",
			filter: Filter{
				Sort: Sort{
					Column: "invalid",
					Order:  "asc",
				},
			},
			want:       []Service{},
			wantErr:    true,
			errMessage: "filter.sort.column must be a valid value",
		},
		{
			name: "Invalid Sort Order",
			filter: Filter{
				Sort: Sort{
					Column: "name",
					Order:  "invalid",
				},
			},
			want:       []Service{},
			wantErr:    true,
			errMessage: "filter.sort.order must be a valid value",
		},
		{
			name: "Valid Search by Name",
			filter: Filter{
				Search: Search{
					Column: "name",
					Value:  "Service",
				},
			},
			want:    []Service{mockServices[0], mockServices[1]},
			wantErr: false,
		},
		{
			name: "Valid Search by Description: No match",
			filter: Filter{
				Search: Search{
					Column: "description",
					Value:  "test",
				},
			},
			want:    []Service{},
			wantErr: false,
		},
		{
			name: "Invalid Search Column",
			filter: Filter{
				Search: Search{
					Column: "invalid",
					Value:  "Service",
				},
			},
			want:       []Service{},
			wantErr:    true,
			errMessage: "filter.search.column must be a valid value",
		},
		{
			name: "Valid with Search, Sort and Pagination",
			filter: Filter{
				Search: Search{
					Column: "name",
					Value:  "Service",
				},
				Sort: Sort{
					Column: "name",
					Order:  "asc",
				},
				Pagination: Pagination{
					Limit:  1,
					Offset: 0,
				},
			},
			want:    []Service{mockServices[0]},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetServicesCore(tt.filter)

			if err != nil {
				assert.Equal(t, err.Error(), tt.errMessage)
			} else {
				assert.Equal(t, len(*got), len(tt.want))
				for i, service := range tt.want {
					assert.Equal(t, service.Name, tt.want[i].Name)
					assert.Equal(t, service.Description, tt.want[i].Description)
					assert.Equal(t, service.LatestVersion, tt.want[i].LatestVersion)
				}
			}
		})
	}
}

func TestGetServiceByIdCore(t *testing.T) {
	tests := []struct {
		name       string
		serviceId  string
		want       Service
		wantErr    bool
		errMessage string
	}{
		{
			name:      "Valid service ID",
			serviceId: "6Zt5mpdDsQ",
			want:      mockServices[1],
			wantErr:   false,
		},
		{
			name:       "Non-existent service ID",
			serviceId:  "6ZtXXXdDsQ",
			want:       Service{},
			wantErr:    true,
			errMessage: "service not found",
		},
		{
			name:       "Empty service ID",
			serviceId:  "",
			want:       Service{},
			wantErr:    true,
			errMessage: "id cannot be blank",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetServiceByIdCore(tt.serviceId)

			if err != nil {
				assert.Equal(t, err.Error(), tt.errMessage)
			} else {
				assert.Equal(t, got.Name, tt.want.Name)
				assert.Equal(t, got.Description, tt.want.Description)
				assert.Equal(t, got.LatestVersion, tt.want.LatestVersion)
			}
		})
	}
}

func TestGetServiceVersionsCore(t *testing.T) {
	tests := []struct {
		name       string
		serviceId  string
		want       ServiceVersionResponse
		wantErr    bool
		errMessage string
	}{
		{
			name:      "Valid service ID",
			serviceId: "6Zt5mpdDsQ",
			want:      ServiceVersionResponse{Versions: []string{"v2.0"}, Count: 1},
			wantErr:   false,
		},
		{
			name:       "Non-existent service ID",
			serviceId:  "6ZtXXXdDsQ",
			want:       ServiceVersionResponse{},
			wantErr:    true,
			errMessage: "service not found",
		},
		{
			name:       "Empty service ID",
			serviceId:  "",
			want:       ServiceVersionResponse{},
			wantErr:    true,
			errMessage: "id cannot be blank",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetServiceVersionsCore(tt.serviceId)

			if err != nil {
				assert.Equal(t, err.Error(), tt.errMessage)
			} else {
				assert.Equal(t, got.Versions, tt.want.Versions)
				assert.Equal(t, got.Count, tt.want.Count)
			}
		})
	}
}

func TestCreateServiceCore(t *testing.T) {
	tests := []struct {
		name       string
		service    Service
		wantErr    bool
		errMessage string
	}{
		{
			name:    "Valid service",
			service: Service{Name: "Service Two", Description: "Second service", LatestVersion: "v2.0", Versions: []string{"v2.0"}},
			wantErr: false,
		},
		{
			name:       "Empty service name",
			service:    Service{Name: "", Description: "Second service", LatestVersion: "v2.0", Versions: []string{"v2.0"}},
			wantErr:    true,
			errMessage: "name: cannot be blank.",
		},
		{
			name:       "Empty service description",
			service:    Service{Name: "Service Two", Description: "", LatestVersion: "v2.0", Versions: []string{"v2.0"}},
			wantErr:    true,
			errMessage: "description: cannot be blank.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateServiceCore(&tt.service)

			if err != nil {
				assert.Equal(t, err.Error(), tt.errMessage)
			}
		})
	}
}

func TestUpdateServiceCore(t *testing.T) {
	tests := []struct {
		name       string
		serviceId  string
		service    Service
		wantErr    bool
		errMessage string
	}{
		{
			name:      "Valid service",
			serviceId: "6Zt5mpdDsQ",
			service:   Service{Name: "Service Two", Description: "Second service", LatestVersion: "v2.0", Versions: []string{"v2.0"}},
			wantErr:   false,
		},
		{
			name:       "service not found",
			serviceId:  "6ZtXXXdDsQ",
			service:    Service{Name: "Service Two", Description: "Second service", LatestVersion: "v2.0", Versions: []string{"v2.0"}},
			wantErr:    true,
			errMessage: "service not found",
		},
		{
			name:       "Empty service name",
			serviceId:  "6Zt5mpdDsQ",
			service:    Service{Name: "", Description: "Second service", LatestVersion: "v2.0", Versions: []string{"v2.0"}},
			wantErr:    true,
			errMessage: "name: cannot be blank.",
		},
		{
			name:       "Empty service description",
			serviceId:  "6Zt5mpdDsQ",
			service:    Service{Name: "Service Two", Description: "", LatestVersion: "v2.0", Versions: []string{"v2.0"}},
			wantErr:    true,
			errMessage: "description: cannot be blank.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateServiceCore(tt.serviceId, &tt.service)

			if err != nil {
				assert.Equal(t, err.Error(), tt.errMessage)
			}
		})
	}
}
