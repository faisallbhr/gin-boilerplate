package controllers

import (
	"net/http"
	"strconv"

	"github.com/faisallbhr/gin-boilerplate/database"
	"github.com/faisallbhr/gin-boilerplate/helpers"
	"github.com/faisallbhr/gin-boilerplate/models"
	"github.com/faisallbhr/gin-boilerplate/presenters"
	"github.com/faisallbhr/gin-boilerplate/requests"
	"github.com/faisallbhr/gin-boilerplate/structs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(c *gin.Context) {
	userId := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, userId).Error; err != nil {
		helpers.ResponseError(c, "User not found", http.StatusNotFound, nil)
		return
	}

	helpers.ResponseSuccess(c, presenters.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, "User found", http.StatusOK, nil)
}

func GetUsers(c *gin.Context) {
	var users []models.User

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limir", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	params := structs.MetaParams{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
		SortBy: c.DefaultQuery("sort_by", "id"),
		Order:  c.DefaultQuery("order", "desc"),
	}

	searchFields := []string{"name", "email"}
	sortFields := []string{"id", "name", "email", "created_at"}

	meta, err := helpers.Meta(database.DB, &users, params, searchFields, sortFields)
	if err != nil {
		helpers.ResponseError(c, "Error getting users", http.StatusInternalServerError, nil)
		return
	}

	formatted := presenters.FormatUsers(users)

	helpers.ResponseSuccess(c, formatted, "Users found", http.StatusOK, meta)
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, userId).Error; err != nil {
		helpers.ResponseError(c, "User not found", http.StatusNotFound, nil)
		return
	}

	req := requests.UpdateUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ResponseError(c, "Failed to update user", http.StatusBadRequest, helpers.TranslateErrorMessage(err, req))
		return
	}

	user.Name = req.Name
	user.Email = req.Email

	if err := database.DB.Save(&user).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {
			helpers.ResponseError(c, "Failed to update user", http.StatusConflict, map[string]string{
				"email": "Email already exists",
			})
			return
		} else {
			helpers.ResponseError(c, "Error updating user", http.StatusInternalServerError, helpers.TranslateErrorMessage(err, req))
			return
		}
	}

	helpers.ResponseSuccess(c, presenters.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, "User updated successfully", http.StatusOK, nil)
}

func UpdatePassword(c *gin.Context) {
	userId := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, userId).Error; err != nil {
		helpers.ResponseError(c, "User not found", http.StatusNotFound, nil)
		return
	}

	req := requests.UpdatePasswordRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ResponseError(c, "Failed to update password", http.StatusBadRequest, helpers.TranslateErrorMessage(err, req))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.CurrentPassword)); err != nil {
		helpers.ResponseError(c, "Failed to update password", http.StatusUnprocessableEntity, map[string]string{
			"current_password": "Invalid current password",
		})
		return
	}

	if req.CurrentPassword == req.NewPassword {
		helpers.ResponseError(c, "Failed to update password", http.StatusUnprocessableEntity, map[string]string{
			"new_password": "New password must be different from current password",
		})
		return
	}

	user.Password = helpers.HashPassword(req.NewPassword)

	if err := database.DB.Save(&user).Error; err != nil {
		helpers.ResponseError(c, "Error updating password", http.StatusInternalServerError, helpers.TranslateErrorMessage(err, req))
		return
	}

	helpers.ResponseSuccess(c, presenters.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}, "Password updated successfully", http.StatusOK, nil)
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, userId).Error; err != nil {
		helpers.ResponseError(c, "User not found", http.StatusNotFound, nil)
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		helpers.ResponseError(c, "Error deleting user", http.StatusInternalServerError, helpers.TranslateErrorMessage(err, user))
		return
	}

	helpers.ResponseSuccess(c, nil, "User deleted successfully", http.StatusOK, nil)
}
