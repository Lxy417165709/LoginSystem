package commonStruct

type Error struct {
	ForUser      error
	ForDeveloper error
}

func NewError(forUser error, forDeveloper error) *Error {
	return &Error{ForUser: forUser, ForDeveloper: forDeveloper}
}
