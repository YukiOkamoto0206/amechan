package repository_interface

import "amechan/src/domain/entity"

// Userに関するDBチックなことのメソッドをまとめる
type UserRepository interface {
	Insert(*entity.User) error
	GetUserFromLineId(string) (*entity.User, error)
}
