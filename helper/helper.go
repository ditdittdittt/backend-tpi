package helper

import (
	"github.com/ditdittdittt/backend-tpi/entities"
)

func ValidatePermission(permissionList []entities.Permission, permissionNeeded string) bool {
	for _, permission := range permissionList {
		if permissionNeeded == permission.Name {
			return true
		}
	}
	return false
}
