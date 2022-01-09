// 渡されたことだけでなく、それを超えて自分で考えてやってるとこがあるからfactoryに分けた

package factory_implement

import (
	"amechan/src/domain/entity"
	"amechan/src/infrastructure"
	"github.com/pkg/errors"
)

type UserFactory struct {
	Database infrastructure.Mysql
}

func NewUserFactory(database infrastructure.Mysql) (userFactory *UserFactory) {
	userFactory = new(UserFactory)
	userFactory.Database = database
	return
}

type userMapper struct {
	Id int `db:"id"`
}

// 既に登録されているかのチェック
func (userFactory *UserFactory) Create(lineId string, officeCode int, areCode int) (user *entity.User, err error) {
	db, err := userFactory.Database.Connect()
	if err != nil {
		return
	}

	query := `
SELECT id
FROM users
WHERE line_id = ?
`

	var userMapperVar []userMapper

	err = db.Select(&userMapperVar, query, lineId)
	if err != nil {
		return
	}

	if len(userMapperVar) != 0 {
		err = errors.New("This user already exists")
		return
	}

	user = entity.NewUser(lineId, officeCode, areCode)
	return
}
