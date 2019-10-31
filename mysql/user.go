package mysql

type User struct {
	ID       int64  `json:"id" xorm:"pk autoincr 'id'"`
	Username string `json:"username" xorm:"notnull"`
	Password string `json:"password" xorm:"notnull"`
}

func GetAdminByUsername(username string) (*User, error) {
	var admin User
	admin.Username = username
	has, err := engine.Get(&admin)
	if !has {
		return nil, err
	}
	return &admin, err
}

func GetAdminByUid(uid int64) (*User, error) {
	var admin User
	has, err := engine.ID(uid).Get(&admin)
	if !has {
		return nil, err
	}
	return &admin, err
}

func UpdatePasswordByUid(uid int64, password string) error {
	var admin User
	admin.Password = password
	_, err := engine.Id(uid).Update(admin)
	return err
}
