package structs

// Структура для хранения данных о пользователе.
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UserPatch struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

// Структура для хранения статистики по пользователю.
type UserStat struct {
	UserName    string
	TotalAmount float32
	AvgPrice    float32
}

type UserService interface {
	GetUsers() ([]User, error)

	CreateUser(*UserPatch) error

	UpdateUser(int, *User, UserPatch) error

	DeleteUser(int) error

	GetUserStat(int) ([]UserStat, error)
}
