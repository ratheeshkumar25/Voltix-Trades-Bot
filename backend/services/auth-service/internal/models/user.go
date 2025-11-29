package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	UserRole_USER  = 0
	UserRole_ADMIN = 1
)

var (
	UserRole_name = map[int32]string{
		0: "user",
		1: "admin",
	}
	UserRole_value = map[string]int32{
		"user":  0,
		"admin": 1,
	}
)

type ExchangeType string

const (
	ExchangeType_Binance = 0
	ExchangeType_Ctrader = 1
	ExchangeType_Mt5     = 2
)

var (
	ExchangeType_name = map[int32]string{
		0: "binance",
		1: "ctrader",
		2: "mt5",
	}
	ExchangeType_value = map[string]int32{
		"binance": 0,
		"ctrader": 1,
		"mt5":     2,
	}
)

// User represents a platform user
type User struct {
	ID           uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Email        string         `gorm:"uniqueIndex;not null" json:"email"`
	GoogleID     string         `gorm:"uniqueIndex" json:"google_id,omitempty"`
	PasswordHash string         `gorm:"column:password_hash" json:"-"`
	Role         UserRole       `gorm:"type:varchar(20);default:'user'" json:"role"` // user, admin
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hook to generate UUID
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

// ExchangeCredential stores encrypted exchange API credentials
type ExchangeCredential struct {
	ID                 uuid.UUID    `gorm:"type:char(36);primaryKey" json:"id"`
	UserID             uuid.UUID    `gorm:"type:char(36);not null;index" json:"user_id"`
	ExchangeType       ExchangeType `gorm:"type:varchar(50);not null" json:"exchange_type"` // binance, mt5, ctrader
	APIKeyEncrypted    string       `gorm:"type:text" json:"-"`
	APISecretEncrypted string       `gorm:"type:text" json:"-"`
	IsActive           bool         `gorm:"default:true" json:"is_active"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}

// BeforeCreate hook for ExchangeCredential
func (ec *ExchangeCredential) BeforeCreate(tx *gorm.DB) error {
	if ec.ID == uuid.Nil {
		ec.ID = uuid.New()
	}
	return nil
}

// Session represents a user session
type Session struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;index" json:"user_id"`
	Token     string    `gorm:"type:text;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

// BeforeCreate hook for Session
func (s *Session) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}
