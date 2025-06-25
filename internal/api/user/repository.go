package user

import "github.com/jmoiron/sqlx"

type UserRepository interface {
	GetUserByID(id string) (*User, error)
	GetUserGreetingByUserID(id string) (*UserGreeting, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByID(id string) (*User, error) {
	user := &User{}
	query := "SELECT  user_id, name, image, password_hash, pin_hash, created_at, updated_at FROM users WHERE user_id = ?"
	err := r.db.Get(user, query, id)
	return user, err
}

func (r *userRepository) GetUserGreetingByUserID(id string) (*UserGreeting, error) {
	greeting := &UserGreeting{}
	query := "SELECT  user_id, greeting, created_at FROM user_greetings WHERE user_id = ?"
	err := r.db.Get(greeting, query, id)
	return greeting, err
}
