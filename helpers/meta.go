package helpers

import (
	"strings"

	"gorm.io/gorm"

	"github.com/faisallbhr/gin-boilerplate/structs"
)

func pagination(db *gorm.DB, params structs.MetaParams) *gorm.DB {
	return db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
}

func search(db *gorm.DB, search string, fields []string) *gorm.DB {
	if search != "" && len(fields) > 0 {
		db = db.Where(fields[0]+" LIKE ?", "%"+search+"%")
		for _, field := range fields[1:] {
			db = db.Or(field+" LIKE ?", "%"+search+"%")
		}
	}
	return db
}

func sort(db *gorm.DB, sortBy string, order string, allowedFields []string) *gorm.DB {
	if sortBy == "" {
		return db.Order("id desc")
	}

	valid := false
	for _, f := range allowedFields {
		if f == sortBy {
			valid = true
			break
		}
	}

	if !valid {
		return db.Order("id desc")
	}

	order = strings.ToUpper(order)
	if order != "ASC" && order != "DESC" {
		order = "ASC"
	}

	return db.Order(sortBy + " " + order)
}

func Meta(db *gorm.DB, model any, params structs.MetaParams, searchFields []string, sortFields []string) (*structs.Meta, error) {
	query := db.Model(model)
	query = search(query, params.Search, searchFields)
	query = sort(query, params.SortBy, params.Order, sortFields)

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	totalPages := int((total + int64(params.Limit) - 1) / int64(params.Limit))

	paginatedQuery := pagination(query, params)
	if err := paginatedQuery.Find(model).Error; err != nil {
		return nil, err
	}

	meta := &structs.Meta{
		Search: params.Search,
		Sort: structs.Sort{
			By:    params.SortBy,
			Order: strings.ToLower(params.Order),
		},
		Pagination: structs.Pagination{
			Page:       params.Page,
			Limit:      params.Limit,
			Total:      int(total),
			TotalPages: totalPages,
		},
	}

	return meta, nil
}
