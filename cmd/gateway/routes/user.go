
package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy User model untuk Swagger ===================
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// =================== Dummy Handlers untuk Swagger ===================

// @Summary Get all users
// @Description Get list of users
// @Tags users
// @Produce json
// @Success 200 {array} routes.User
// @Router /users [get]
func GetUsersHandler(c *fiber.Ctx) error { return nil }

// @Summary Get user by ID
// @Description Get single user
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} routes.User
// @Router /users/{id} [get]
func GetUserHandler(c *fiber.Ctx) error { return nil }

// @Summary Create user
// @Description Create new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body routes.User true "User info"
// @Success 201 {object} routes.User
// @Router /users [post]
func CreateUserHandler(c *fiber.Ctx) error { return nil }

// @Summary Update user
// @Description Update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body routes.User true "User info"
// @Success 200 {object} routes.User
// @Router /users/{id} [put]
func UpdateUserHandler(c *fiber.Ctx) error { return nil }

// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {object} map[string]bool
// @Router /users/{id} [delete]
func DeleteUserHandler(c *fiber.Ctx) error { return nil }