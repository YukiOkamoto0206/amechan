// 実際のコード
package usecase

import (
	"amechan/src/domain/entity"
	"amechan/src/domain/repository_interface"
)

type RegisterUserUseCase struct {
	UserRepository repository_interface.UserRepository
}

func NewRegisterUserUseCase(repository repository_interface.UserRepository) (registerUserUseCase *RegisterUserUseCase) {
	registerUserUseCase = new(RegisterUserUseCase)
	registerUserUseCase.UserRepository = repository
	return
}

// Invoke実際に行う
func (uc *RegisterUserUseCase) Invoke(lineId string, officeCode int, areaCode int) (err error) {
	// TODO:既にLINE_IDでユーザーが登録されてあるかチェックする
	// entityの情報をもとにチェックしたいけど、DBでチェックしないといけないからdomain層のdomain_serviceでする
	// TODO:新しくuser(entityの)を作る
	user := entity.NewUser(lineId, officeCode, areaCode)
	// TODO:DBに保存する
	err = uc.UserRepository.Insert(user)
	return
}
