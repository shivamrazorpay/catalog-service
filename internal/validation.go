package internal

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"service-catalog/common"
)

func ValidateGetServiceById(data string) error {
	err := validation.Validate(data, validation.Required)
	if err != nil {
		return errors.New(fmt.Sprintf("id %s", err.Error()))
	}
	return nil
}

func ValidateCreateService(service Service) error {
	return validation.ValidateStruct(&service,
		validation.Field(&service.Name, validation.Required),
		validation.Field(&service.Description, validation.Required, validation.Length(10, 100)),
		validation.Field(&service.LatestVersion, validation.Required),
	)
}

func ValidateUpdateService(service Service) error {
	return validation.ValidateStruct(&service,
		validation.Field(&service.Name, validation.NilOrNotEmpty),
		validation.Field(&service.Description, validation.NilOrNotEmpty, validation.Length(10, 100)),
	)
}

func ValidateFilters(filter Filter) error {

	err := validation.Validate(filter.Sort.Column, validation.In("id", "name", "created_at"))
	if err != nil {
		return errors.New(fmt.Sprintf("filter.sort.column %s", err.Error()))
	}

	err = validation.Validate(filter.Sort.Order, validation.In(common.Asc, common.Desc))
	if err != nil {
		return errors.New(fmt.Sprintf("filter.sort.order %s", err.Error()))
	}

	err = validation.Validate(filter.Pagination.Limit, validation.Min(1), validation.Max(20))
	if err != nil {
		return errors.New(fmt.Sprintf("filter.pagination.limit %s", err.Error()))
	}

	err = validation.Validate(filter.Pagination.Offset, validation.Min(0))
	if err != nil {
		return errors.New(fmt.Sprintf("filter.pagination.offset %s", err.Error()))
	}

	err = validation.Validate(filter.Search.Column, validation.In("name", "description"))
	if err != nil {
		return errors.New(fmt.Sprintf("filter.search.column %s", err.Error()))
	}

	return nil
}
