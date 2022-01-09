package entity

type User struct {
	Id         int
	LineId     string
	OfficeCode int
	AreaCode   int
}

func NewUser(lineId string, officeCode int, areaCode int) (user *User) { // userをreturnする
	user = new(User) // newがポインタを返す
	user.Id = 0      //(*user).Id = 0 これをしたい
	user.LineId = lineId
	user.OfficeCode = officeCode
	user.AreaCode = areaCode
	return
}
