package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/fi1atov/go_otus_hw/hw15_go_sql/structs"
)

type UserService struct {
	db *DB
}

func NewUserService(db *DB) *UserService {
	return &UserService{db}
}

// Получение списка всех пользователей.
func (ps *UserService) GetUsers() ([]structs.User, error) {
	rows, err := ps.db.Query(`SELECT u.id, u.name, u.email, u.password FROM users u`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []structs.User
	for rows.Next() {
		var user structs.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Создание пользователя.
func (ps *UserService) CreateUser(user *structs.UserPatch) error {
	tx, err := ps.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		return err
	}

	// Пользователь создан, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Пользователь успешно создан.")
	return nil
}

// Обновление пользователя.
func (ps *UserService) UpdateUser(userID int, user *structs.User, patch structs.UserPatch) error {
	tx, err := ps.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if v := patch.Name; v != nil {
		user.Name = *v
	}

	if v := patch.Email; v != nil {
		user.Email = *v
	}

	if v := patch.Password; v != nil {
		user.Password = *v
	}

	_, err = tx.Exec(
		"UPDATE users SET name=$2, email=$3, password=$4 WHERE id=$1",
		userID, user.Name, user.Email, user.Password,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	// Пользователь обновлен, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Пользователь успешно обновлен.")
	return nil
}

// Удаление пользователя.
func (ps *UserService) DeleteUser(userID int) error {
	tx, err := ps.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		log.Println(err)
		return err
	}

	// Пользователь удален, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Пользователь успешно удален.")
	return nil
}

// Получение списка всех пользователей.
func (ps *UserService) GetUserStat(userID int) ([]structs.UserStat, error) {
	rows, err := ps.db.Query(`
	SELECT u.name, COALESCE(SUM(o.total_amount), 0) AS total_amount, COALESCE(AVG(p.price), 0) as avg_price
	from orders o
	join order_products op ON o.id = op.order_id
	join products p ON op.product_id = p.id
	join users u ON o.user_id = u.id
	WHERE u.id = $1
	group by u.name`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userStats []structs.UserStat
	for rows.Next() {
		var userstat structs.UserStat
		err = rows.Scan(&userstat.UserName, &userstat.TotalAmount, &userstat.AvgPrice)
		if err != nil {
			return nil, err
		}
		userStats = append(userStats, userstat)
	}

	return userStats, nil
}
