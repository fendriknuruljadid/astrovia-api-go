
package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy User model untuk Swagger ===================


type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

// =================== Dummy Handlers untuk Swagger ===================

// @Summary Get all users
// @Description Get list of users
// @Tags users
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.User
// @Router /v1/users [get]
func GetUsersHandler(c *fiber.Ctx) error { return nil }

// @Summary Get user by ID
// @Description Get single user
// @Tags users
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} routes.User
// @Router /v1/users/{id} [get]
func GetUserHandler(c *fiber.Ctx) error { return nil }

// @Summary Create user
// @Description Create new user
// @Tags users
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param user body routes.User true "User info"
// @Success 201 {object} routes.User
// @Router /v1/users [post]
func CreateUserHandler(c *fiber.Ctx) error { return nil }

// @Summary Update user
// @Description Update user by ID
// @Tags users
// @Security X-Signature
// @Security X-Timestamp
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body routes.User true "User info"
// @Success 200 {object} routes.User
// @Router /v1/users/{id} [put]
func UpdateUserHandler(c *fiber.Ctx) error { return nil }

// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Security X-Signature
// @Security X-Timestamp
// @Param id path string true "User ID"
// @Success 200 {object} map[string]bool
// @Router /v1/users/{id} [delete]
func DeleteUserHandler(c *fiber.Ctx) error { return nil }