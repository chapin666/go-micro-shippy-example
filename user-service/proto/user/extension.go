package go_micro_srv_user

import (
	"log"
	uuid "github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
)

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("create uuid error: %v\n", err)
	}
	return scope.SetColumn("Id", uuid.String())
}