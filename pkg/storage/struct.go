package storage

import "time"

// Account TODO
type Account struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Users     []*User    `db:"-"`
}

// User TODO
type User struct {
	ID        string     `db:"id"`
	AccountID string     `db:"account_id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Account   *Account   `db:"-"`
	Tokens    []*Token   `db:"-"`
}

// Token TODO
type Token struct {
	ID        string     `db:"id"`
	UserID    string     `db:"user_id"`
	ExpiredAt *time.Time `db:"expired_at"`
	CreatedAt time.Time  `db:"created_at"`
	User      *User      `db:"-"`
}

// Schema TODO
type Schema struct {
	ID         string     `db:"id"`
	AccountID  string     `db:"account_id"`
	Language   string     `db:"language"`
	Title      string     `db:"title"`
	Definition string     `db:"definition"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
	Account    *Account   `db:"-"`
}

// Template TODO
type Template struct {
	ID         string     `db:"id"`
	AccountID  string     `db:"account_id"`
	Title      string     `db:"title"`
	Definition string     `db:"definition"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
	Account    *Account   `db:"-"`
}

// Input TODO
type Input struct {
	ID        string    `db:"id"`
	SchemaID  string    `db:"schema_id"`
	Data      []byte    `db:"data"`
	CreatedAt time.Time `db:"created_at"`
	Schema    *Schema   `db:"-"`
}

// Log TODO
type Log struct {
	ID         uint64    `db:"id"`
	AccountID  string    `db:"account_id"`
	SchemaID   string    `db:"schema_id"`
	InputID    string    `db:"input_id"`
	TemplateID *string   `db:"template_id"`
	Identifier string    `db:"identifier"`
	Code       uint16    `db:"code"`
	Context    []byte    `db:"context"`
	CreatedAt  time.Time `db:"created_at"`
	Account    *Account  `db:"-"`
	Schema     *Schema   `db:"-"`
	Input      *Input    `db:"-"`
	Template   *Template `db:"-"`
}