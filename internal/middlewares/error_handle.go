package middlewares

import (
    "log"
    "github.com/gin-gonic/gin"
    "app/internal/packages/errors"
    "app/internal/packages/response"
)

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("panic recovered: %v", r)

                c.JSON(500, response.Error(
                    500, 
                    "internal server error",
                    r,
                ))
                c.Abort()
            }
        }()

        c.Next()

        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err

            switch e := err.(type) {
            case *errors.AppError:
                c.JSON(e.Code, response.Error(e.Code, e.Message, e.Detail))
            default:
                c.JSON(500, response.Error(
                    500,
                    "internal server error",
                    e.Error(),
                ))
            }

            c.Abort()
        }
    }
}
