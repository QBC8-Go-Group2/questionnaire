package mapper

import (
	"fmt"

	"github.com/QBC8-Go-Group2/questionnaire/internal/user/domain"
	"github.com/QBC8-Go-Group2/questionnaire/pkg/adapter/storage/types"
)

func UserStorage2Domain(user types.User) domain.User {
	return domain.User{
		ID:        domain.UserDbID(user.ID),
		CreatedAT: user.CreatedAt,
		UserID:    domain.UserID(hastUserID(user.ID, user.Email)),
		Email:     user.Email,
		Password:  user.Password,
		NatId:     user.NatID,
		Role:      domain.RoleType(user.Role),
	}
}

func UserDomain2Storage(user domain.User) types.User {
	return types.User{
		ID:        uint(user.ID),
		CreatedAt: user.CreatedAT,
		UserID:    string(user.UserID), // Add this line to map UserID
		Email:     user.Email,
		Password:  user.Password,
		NatID:     user.NatId,
		Role:      uint8(user.Role),
	}
}

func hastUserID(id uint, email string) string {
	return fmt.Sprintf("%d_%s", id, email)
}
