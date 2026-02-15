package routes

import "github.com/gofiber/fiber/v2"

// =================== Dummy User model untuk Swagger ===================

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// =================== Dummy Handlers untuk Swagger ===================

// @Summary Get all users
// @Description Get list of users
// @Tags Users
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Success 200 {array} routes.User
// @Router /v1/users [get]
func GetUsersHandler(c *fiber.Ctx) error { return nil }

// @Summary Check users exist by email
// @Description Check users exist
// @Tags Users
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Accept json
// @Param user body routes.User true "User info"
// @Success 200 {array} routes.User
// @Router /v1/users/check [post]
func UserCheckHandler(c *fiber.Ctx) error { return nil }

// @Summary Create password for user
// @Description create password
// @Tags Users
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Accept json
// @Param user body routes.User true "User info"
// @Success 200 {array} routes.User
// @Router /v1/users/create-password [post]
func CreatePasswordHandler(c *fiber.Ctx) error { return nil }

// @Summary Verify OTP
// @Description verify otp
// @Tags Users
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Accept json
// @Param user body routes.User true "User info"
// @Success 200 {array} routes.User
// @Router /v1/users/verify-verification [post]
func VerifyVerificationHandler(c *fiber.Ctx) error { return nil }

// @Summary Resend OTP
// @Description resend OTP
// @Tags Users
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Accept json
// @Param user body routes.User true "User info"
// @Success 200 {array} routes.User
// @Router /v1/users/resend-verification [post]
func ResendOTPHandler(c *fiber.Ctx) error { return nil }

// @Summary Get user by ID
// @Description Get single user
// @Tags Users
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} routes.User
// @Router /v1/users/{id} [get]
func GetUserHandler(c *fiber.Ctx) error { return nil }

// @Summary Create user
// @Description Create new user
// @Tags Users
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
// @Tags Users
// @Security BearerAuth
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
// @Tags Users
// @Security BearerAuth
// @Security X-Signature
// @Security X-Timestamp
// @Param id path string true "User ID"
// @Success 200 {object} map[string]bool
// @Router /v1/users/{id} [delete]
func DeleteUserHandler(c *fiber.Ctx) error { return nil }
