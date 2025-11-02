package handlers

import (
    "github.com/gin-gonic/gin"
    "astrovia-api-go/internal/services/user/models"
    "astrovia-api-go/internal/services/user/repository"
    "strconv"
)

func CreateUser(c *gin.Context) {
    var body models.User
    if err := c.BindJSON(&body); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    if err := repository.CreateUser(&body); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(201, body)
}

func GetUsers(c *gin.Context) {
    users, err := repository.GetUsers()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, users)
}

func GetUserByID(c *gin.Context) {
    id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    user, err := repository.GetUserByID(id)
    if err != nil {
        c.JSON(404, gin.H{"error": "Not found"})
        return
    }
    c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
    id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    var body models.User

    if err := c.BindJSON(&body); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    body.ID = id

    if err := repository.UpdateUser(&body); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, body)
}

func DeleteUser(c *gin.Context) {
    id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    if err := repository.DeleteUser(id); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"deleted": true})
}
