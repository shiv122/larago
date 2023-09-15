package repositories

import (
	"github.com/shiv122/go-test/app/models"
	"github.com/shiv122/go-test/config"
)

type UserRepository struct{}

type PaginatedUsers struct {
	Users      []models.User `json:"data"`
	Total      int64         `json:"total"`
	Page       int           `json:"page"`
	PageSize   int           `json:"pageSize"`
	TotalPages int           `json:"totalPages"`
	HasNext    bool          `json:"hasNext"`
	HasPrev    bool          `json:"hasPrev"`
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) GetPaginatedUsers(page int, pageSize int) (PaginatedUsers, error) {
	var users []models.User
	offset := (page - 1) * pageSize

	// Fetch users and total count
	result := config.DB.Limit(pageSize).Offset(offset).Find(&users)
	if result.Error != nil {
		return PaginatedUsers{}, result.Error
	}

	// Get total count of users
	var totalRecords int64
	config.DB.Model(&models.User{}).Count(&totalRecords)

	// Calculate total pages
	totalPages := int(totalRecords) / pageSize
	if totalRecords%int64(pageSize) > 0 {
		totalPages++
	}

	// Determine if there is a next page
	hasNext := page < totalPages

	// Determine if there is a previous page
	hasPrev := page > 1

	return PaginatedUsers{
		Users:      users,
		Total:      totalRecords,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}, nil
}
