package dto

import "astrovia-api-go/internal/services/user/models"

type CreateDTO struct {
    Name     string `json:"name" binding:"required,min=3"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type UpdateDTO struct {
    Name     *string `json:"name,omitempty" binding:"omitempty,min=3"`
    Email    *string `json:"email,omitempty" binding:"omitempty,email"`
    Password *string `json:"password,omitempty" binding:"omitempty,min=6"`
}

type ResponseDTO struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

func ToResponseDTO(u *models.User) ResponseDTO {
    return ResponseDTO{
        ID:        u.ID,
        Name:      u.Name,
        Email:     u.Email,
        CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
        UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
    }
}

func ToResponseDTOs(users []models.User) []ResponseDTO {
    res := make([]ResponseDTO, len(users))
    for i, u := range users {
        res[i] = ResponseDTO{
            ID:        u.ID,
            Name:      u.Name,
            Email:     u.Email,
            CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
            UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
        }
    }
    return res
}

