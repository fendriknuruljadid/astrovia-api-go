package handlers

import (
    "github.com/gin-gonic/gin"
    "astrovia-api-go/internal/services/user/models"
    "astrovia-api-go/internal/services/user/repository"
    "net/http"
    "astrovia-api-go/internal/packages/response"
    "github.com/go-playground/validator/v10"
    "astrovia-api-go/internal/packages/errors"
    "astrovia-api-go/internal/packages/utils"
)

var validate = validator.New()

func CreateUser(c *gin.Context) {
    var body models.User

    // JSON binding error
    if err := c.ShouldBindJSON(&body); err != nil {
        c.Error(errors.NewBadRequest("invalid request payload", err.Error()))
        return
    }

    // Validation error
    if err := utils.Validate.Struct(body); err != nil {
        c.Error(errors.NewBadRequest("validation failed", utils.ValidationErrors(err)))
        return
    }

    // Database insertion
    if err := repository.CreateUser(&body); err != nil {
        c.Error(errors.NewInternal(utils.ParseDBError(err)))
        return
    }

    // Success
    c.JSON(201, response.Success(
        201,
        "successfully",
        body,
    ))
}

func GetUsers(c *gin.Context) {
    users, err := repository.GetUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
    id := c.Param("id")
    user, err := repository.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var body models.User

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := validate.Struct(body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    body.ID = id

    if err := repository.UpdateUser(&body); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, body)
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if err := repository.DeleteUser(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"deleted": true})
}
