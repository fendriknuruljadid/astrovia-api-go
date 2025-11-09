package handlers

import (
    "github.com/gin-gonic/gin"
    "astrovia-api-go/internal/services/v1/user/models"
    "astrovia-api-go/internal/services/v1/user/repository"
    "astrovia-api-go/internal/services/v1/user/dto"
    "net/http"
    "astrovia-api-go/internal/packages/response"
    "github.com/go-playground/validator/v10"
    "astrovia-api-go/internal/packages/errors"
    "astrovia-api-go/internal/packages/utils"
    
)

var validate = validator.New()

func CreateUser(c *gin.Context) {
    var req dto.CreateDTO

    if err := c.ShouldBindJSON(&req); err != nil {
        // Deteksi apakah ini validation error dari validator bawaan Gin
        if ve, ok := err.(validator.ValidationErrors); ok {
            c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
            return
        }
    
        // Kalau bukan validasi, baru fallback ke invalid JSON
        c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
        return
    }
    user := models.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: req.Password,
    }

    if err := repository.CreateUser(&user); err != nil {
        c.Error(errors.NewInternal(utils.ParseDBError(err)))
        return
    }

    res := dto.ToResponseDTO(&user)

    c.JSON(201, response.Success(201, "success", res))
}

func GetUsers(c *gin.Context) {
    users, err := repository.GetUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    res := dto.ToResponseDTOs(users)
    c.JSON(200, response.Success(200, "success", res))
}


func GetUserByID(c *gin.Context) {
    id := c.Param("id")
    user, err := repository.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    res := dto.ToResponseDTO(user)
    c.JSON(200, response.Success(200, "success", res))
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    
    var req dto.UpdateDTO
    if err := c.ShouldBindJSON(&req); err != nil {
        // Deteksi apakah ini validation error dari validator bawaan Gin
        if ve, ok := err.(validator.ValidationErrors); ok {
            c.JSON(402, response.Error(402, "validation failed", utils.ValidationErrors(ve)))
            return
        }
    
        // Kalau bukan validasi, baru fallback ke invalid JSON
        c.JSON(400, response.Error(400, "invalid request payload", err.Error()))
        return
    }

    user, err := repository.GetUserByID(id)
    if err != nil {
        c.Error(errors.NewNotFound("user not found"))
        return
    }

    // Apply updates safely
    if req.Name != nil {
        user.Name = *req.Name
    }
    if req.Email != nil {
        user.Email = *req.Email
    }
    if req.Password != nil {
        user.Password = *req.Password
    }

    if err := repository.UpdateUser(user); err != nil {
        c.Error(errors.NewInternal(utils.ParseDBError(err)))
        return
    }

    c.JSON(200, response.Success(200, "updated", dto.ToResponseDTO(user)))
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")

    // Cek dulu apakah user ada
    user, err := repository.GetUserByID(id)
    if err != nil || user == nil {
        c.Error(errors.NewNotFound("user not found"))
        return
    }

    // Jika ada, lanjut hapus
    if err := repository.DeleteUser(id); err != nil {
        c.Error(errors.NewInternal(utils.ParseDBError(err)))
        return
    }

    c.JSON(200, response.Success(200, "deleted successfully", gin.H{
        "deleted": true,
        "id":      id,
    }))
}

