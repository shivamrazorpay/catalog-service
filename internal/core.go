package internal

import (
	"errors"
	"service-catalog/common"
	"strings"
	"time"
)

func GetServicesCore(filter Filter) (*[]Service, error) {
	err := ValidateFilters(filter)
	if err != nil {
		return nil, err
	}

	setDefaultValues(&filter)

	services := make([]Service, 0)

	// Filtering
	for _, service := range DataBase["Services"].([]Service) {
		if filter.Search.Column == "" {
			services = append(services, service)
			continue
		}

		switch filter.Search.Column {
		case "name":
			if strings.Contains(service.Name, filter.Search.Value) {
				services = append(services, service)
			}
		case "description":
			if strings.Contains(service.Description, filter.Search.Value) {
				services = append(services, service)
			}
		}
	}

	if len(services) == 0 {
		return &services, nil
	}

	// Pagination
	tempService := make([]Service, 0)
	for i, service := range services {
		if i >= filter.Pagination.Offset && i < filter.Pagination.Offset+filter.Pagination.Limit {
			tempService = append(tempService, service)
		}
	}
	services = tempService

	if len(services) == 0 {
		return &services, nil
	}

	// Sorting
	switch filter.Sort.Column {
	case "id":
		SortServicesByIdAsc(services, filter.Sort.Order == "asc")
	case "name":
		SortServicesByName(services, filter.Sort.Order == "asc")
	case "created_at":
		SortServicesByCreatedAt(services, filter.Sort.Order == "asc")
	}

	return &services, nil
}

func GetServiceByIdCore(serviceId string) (Service, error) {

	// Validate serviceId
	err := ValidateGetServiceById(serviceId)
	if err != nil {
		return Service{}, err
	}

	for _, service := range DataBase["Services"].([]Service) {
		if service.ID == serviceId {
			return service, nil
		}
	}

	return Service{}, errors.New("service not found")
}

func GetServiceVersionsCore(serviceId string) (ServiceVersionResponse, error) {
	// Validate serviceId
	err := ValidateGetServiceById(serviceId)
	if err != nil {
		return ServiceVersionResponse{}, err
	}

	for _, service := range DataBase["Services"].([]Service) {
		if service.ID == serviceId {
			return ServiceVersionResponse{
				Versions: service.Versions,
				Count:    len(service.Versions),
			}, nil
		}
	}

	return ServiceVersionResponse{}, errors.New("service not found")
}

func CreateServiceCore(service *Service) error {
	err := ValidateCreateService(*service)
	if err != nil {
		return err
	}

	service.ID, _ = common.UniqueId()
	service.Versions = []string{service.LatestVersion}
	service.CreatedAt = time.Now().Unix()
	service.UpdatedAt = time.Now().Unix()
	DataBase["Services"] = append(DataBase["Services"].([]Service), *service)
	return nil
}

func UpdateServiceCore(serviceId string, service *Service) error {
	err := ValidateUpdateService(*service)
	if err != nil {
		return err
	}

	for _, s := range DataBase["Services"].([]Service) {
		if s.ID == serviceId {

			// if the latest version is not the latest version, add it to the versions
			if s.LatestVersion != service.LatestVersion {
				s.Versions = append(s.Versions, service.LatestVersion)
				s.LatestVersion = service.LatestVersion
			}
			s.Name = service.Name
			s.Description = service.Description
			s.UpdatedAt = time.Now().Unix()
			return nil
		}
	}

	return errors.New("service not found")
}
