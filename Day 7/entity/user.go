package entity

type User struct {
	ID       int    `json:"id" gorm:"primary_key; auto_increment; not_null"`
	Username string `json:"username" gorm:"not_null"`
	Email    string `json:"email" gorm:"not_null"`
	Phone    string `json:"phone" gorm:"not_null"`
}

// create	user
type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// get user
type GetUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// update	user
type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// delete user
type DeleteUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
