package repository

import (
	"database/sql"
	"ingressos-api/dto"
)

func CreateUserRepository(db *sql.DB, user *dto.CreateUserDTO) error {
	query := `INSERT INTO usuario (nome, email) VALUES ($1, $2)`
	_, err := db.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsersRepository(db *sql.DB) ([]dto.ResponseUserDTO, error) {
	rows, err := db.Query("SELECT id, nome, email FROM usuario")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []dto.ResponseUserDTO
	for rows.Next() {
		var user dto.ResponseUserDTO
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(db *sql.DB, id int) (*dto.ResponseUserDTO, error) {
	var user dto.ResponseUserDTO
	err := db.QueryRow("SELECT id, nome, email FROM usuario WHERE id = $1", id).
		Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
func DeleteUserById(db *sql.DB, id int) error {
	query := "DELETE from usuario WHERE id = $1"
	_, err := db.Exec(query, id)
	return err
}
