package controllers

import (
	"net/http"

	"github.com/faisallbhr/gin-boilerplate/database"
	"github.com/faisallbhr/gin-boilerplate/helpers"
	"github.com/faisallbhr/gin-boilerplate/models"
	"github.com/faisallbhr/gin-boilerplate/presenters"
	"github.com/faisallbhr/gin-boilerplate/requests"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Me(c *gin.Context) {
	userId := c.GetUint("user_id")

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

func Register(c *gin.Context) {
	req := requests.RegisterRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ResponseError(c, "Failed to register", http.StatusBadRequest, helpers.TranslateErrorMessage(err, req))
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {
			helpers.ResponseError(c, "Failed to register", http.StatusConflict, map[string]string{
				"email": "Email already exists",
			})
			return
		} else {
			helpers.ResponseError(c, "Error creating user", http.StatusInternalServerError, helpers.TranslateErrorMessage(err, req))
			return
		}
	}

	token := helpers.GenerateToken(user.Id)
	refreshToken := helpers.GenerateRefreshToken(user.Id)

	if token == "" || refreshToken == "" {
		helpers.ResponseError(c, "Failed to generate token", http.StatusInternalServerError, nil)
		return
	}

	helpers.ResponseSuccess(c, presenters.UserResponse{
		Id:           user.Id,
		Name:         user.Name,
		Email:        user.Email,
		Token:        &token,
		RefreshToken: &refreshToken,
	}, "User registered successfully", http.StatusCreated, nil)
}

func Login(c *gin.Context) {
	req := requests.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ResponseError(c, "Failed to login", http.StatusBadRequest, helpers.TranslateErrorMessage(err, req))
		return
	}

	var user models.User

	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		helpers.ResponseError(c, "User not found", http.StatusNotFound, nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		helpers.ResponseError(c, "Invalid credentials", http.StatusUnauthorized, nil)
		return
	}

	token := helpers.GenerateToken(user.Id)
	refreshToken := helpers.GenerateRefreshToken(user.Id)

	if token == "" || refreshToken == "" {
		helpers.ResponseError(c, "Failed to generate token", http.StatusInternalServerError, nil)
		return
	}

	helpers.ResponseSuccess(c, presenters.UserResponse{
		Id:           user.Id,
		Name:         user.Name,
		Email:        user.Email,
		Token:        &token,
		RefreshToken: &refreshToken,
	}, "Login successful", http.StatusOK, nil)
}

func RefreshToken(c *gin.Context) {
	req := requests.RefreshTokenRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ResponseError(c, "Failed to refresh token", http.StatusBadRequest, helpers.TranslateErrorMessage(err, req))
		return
	}

	claims, err := helpers.VerifyToken(req.RefreshToken)
	if err != nil {
		helpers.ResponseError(c, "Invalid or expired refresh token", http.StatusUnauthorized, nil)
		return
	}

	var user models.User

	if err := database.DB.First(&user, claims.UserId).Error; err != nil {
		helpers.ResponseError(c, "User not found", http.StatusNotFound, nil)
		return
	}

	token := helpers.GenerateToken(user.Id)
	refreshToken := helpers.GenerateRefreshToken(user.Id)

	if token == "" || refreshToken == "" {
		helpers.ResponseError(c, "Failed to generate token", http.StatusInternalServerError, nil)
		return
	}

	helpers.ResponseSuccess(c, presenters.UserResponse{
		Id:           user.Id,
		Name:         user.Name,
		Email:        user.Email,
		Token:        &token,
		RefreshToken: &refreshToken,
	}, "Token refreshed successfully", http.StatusOK, nil)
}
