package dao

import "Douban/model"

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}

	row := db.QueryRow("SELECT id, password FROM user WHERE username = ?", username)
	if row.Err() != nil {
		return user, row.Err()
	}

	err := row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func InsertUser(user model.User) error {
	_, err := db.Exec("INSERT INTO user(username,password)" + "values(?, ?);",user.Username,user.Password)
	return err
}