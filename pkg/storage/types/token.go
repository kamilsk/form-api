package types

import (
	"time"

	"github.com/kamilsk/form-api/pkg/domain"
)

// Token TODO
type Token struct {
	ID        domain.ID  `db:"id"`
	UserID    domain.ID  `db:"user_id"`
	ExpiredAt *time.Time `db:"expired_at"`
	CreatedAt time.Time  `db:"created_at"`
	User      *User      `db:"-"`
}
