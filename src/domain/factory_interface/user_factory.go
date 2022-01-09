// entityでは判別つかないルールを実装するためのとろ
// 依存関係逆転の法則

package factory_interface

import "amechan/src/domain/entity"

type UserFactory interface {
	Create(string, int, int) (*entity.User, error)
}
