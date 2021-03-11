package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ditdittdittt/backend-tpi/constant"
	"github.com/ditdittdittt/backend-tpi/database"
	"github.com/ditdittdittt/backend-tpi/helper"
	"github.com/ditdittdittt/backend-tpi/repository/mysql"
	"github.com/ditdittdittt/backend-tpi/services"
)

func AuthorizeJWT(function string) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := services.NewJWTAuthService().ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Token invalid"})
			return
		}

		userRepository := mysql.NewUserRepository(*database.DB)
		curUser, err := userRepository.GetByToken(token.Raw)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			return
		}

		if function != constant.Logout {
			if !helper.ValidatePermission(curUser.Role.Permission, function) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Message": "Forbidden"})
				return
			}
		}

		switch curUser.RoleID {
		case 1:
			userSuperadminRepository := mysql.NewUserSuperadminRepository(*database.DB)
			_, err := userSuperadminRepository.GetByUserID(curUser.ID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
				return
			}
		case 2:
			userDistrictRepository := mysql.NewUserDistrictRepository(*database.DB)
			curUserRole, err := userDistrictRepository.GetByUserID(curUser.ID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
				return
			}
			c.Set("districtID", curUserRole.DistrictID)
		case 3:
			userTpiRepository := mysql.NewUserTpiRepository(*database.DB)
			curUserRole, err := userTpiRepository.GetByUserID(curUser.ID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
				return
			}
			c.Set("tpiID", curUserRole.TpiID)
		}

		c.Set("userID", curUser.ID)
	}

}
