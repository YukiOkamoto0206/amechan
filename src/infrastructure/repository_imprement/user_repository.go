// DBの作られた関数を実行しているファイル

package repository_imprement

import (
	"amechan/src/domain/entity"
	"amechan/src/infrastructure"
)

type UserRepository struct {
	Database infrastructure.Mysql
}

func NewUserRepository(database infrastructure.Mysql) (userRepository *UserRepository) {
	userRepository = new(UserRepository)
	userRepository.Database = database
	return
}

// SaveはUserRepositoryのメソッドであることを(userRepo *UserRepository)で指定
func (userRepo *UserRepository) Create(user *entity.User) (err error) {
	db, err := userRepo.Database.Connect()
	if err != nil {
		return
	}

	query := `
INSERT INTO users(line_id, office_code, area_code)
VALUE (?,?,?);
`
	// MustExecはクエリを実行
	result := db.MustExec(query, user.LineId, user.OfficeCode, user.AreaCode)
	_, err = result.RowsAffected()
	return
}
