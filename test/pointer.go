package test

type IUserRepository interface {
	InsertAdUserInfoWithDefaults(username string) (*User, error)
}
type UserRepository struct {
	Name string
}

var SingletonUserRepository IUserRepository

//这种形式可以声明该实例
func GetUserRepository() IUserRepository {
	if SingletonUserRepository == nil {
		SingletonUserRepository = NewDefaultUserRepository()
	}
	return SingletonUserRepository
}
func NewDefaultUserRepository() IUserRepository {
	return &UserRepository{}
}

/*在这里，(m *UserRepository) 是接收者（receiver）声明，它指明了这个方法是绑定到 UserRepository 类型的指针上的。
这意味着，你需要一个 UserRepository 类型的实例（更准确地是指向 UserRepository 实例的指针），才能调用这个方法。
例子：如果有一个 UserRepository 的实例 ur，你可以这样调用这个方法：ur.InsertAdUserInfoWithDefaults("username")。*/
func (m *UserRepository) InsertAdUserInfoWithDefaults(username string) (*User, error) {
	user := &User{
		Username:      m.Name,
		Role:          "admin",
		Authorization: "22222",
	}

	return user, nil
}
