package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlanType int32

const (
	PlanType_Trial      PlanType = 0
	PlanType_Premium    PlanType = 1
	PlanType_Enterprise PlanType = 2
)

var (
	PlanType_name = map[int32]string{
		0: "trial",
		1: "premium",
		2: "enterprise",
	}
	PlanType_value = map[string]int32{
		"trial":      0,
		"premium":    1,
		"enterprise": 2,
	}
)

type PlanStatus int32

const (
	PlanStatus_Active   PlanStatus = 0
	PlanStatus_Expired  PlanStatus = 1
	PlanStatus_Canceled PlanStatus = 2
)

var (
	PlanStatus_name = map[int32]string{
		0: "active",
		1: "expired",
		2: "canceled",
	}
	PlanStatus_value = map[string]int32{
		"active":   0,
		"expired":  1,
		"canceled": 2,
	}
)

// Subscription represents a user's subscription plan
type Subscription struct {
	ID           uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	UserID       uuid.UUID  `gorm:"type:char(36);not null;uniqueIndex" json:"user_id"`
	Subscription PlanType   `gorm:"type:varchar(20);not null" json:"plan_type"` // trial, premium, enterprise
	Status       PlanStatus `gorm:"type:varchar(20);not null" json:"status"`    // active, expired, cancelled
	StartDate    time.Time  `gorm:"not null" json:"start_date"`
	EndDate      time.Time  `gorm:"not null" json:"end_date"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// BeforeCreate hook
func (s *Subscription) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

// IsActive checks if subscription is currently active
func (s *Subscription) IsActive() bool {
	now := time.Now()
	return s.Status == PlanStatus_Active && now.After(s.StartDate) && now.Before(s.EndDate)
}

// DaysRemaining calculates days until subscription expires
func (s *Subscription) DaysRemaining() int {
	if !s.IsActive() {
		return 0
	}
	duration := time.Until(s.EndDate)
	return int(duration.Hours() / 24)
}

type NotificationEventType int32

const (
	NotificationType_TrialExpiring       NotificationEventType = 0
	NotificationType_TrialExpired        NotificationEventType = 1
	NotificationType_SubscriptionRenewed NotificationEventType = 2
)

var (
	NotificationType_name = map[int32]string{
		0: "trial_expiring",
		1: "trial_expired",
		2: "subscription_renewed",
	}
	NotificationType_value = map[string]int32{
		"trial_expiring":       0,
		"trial_expired":        1,
		"subscription_renewed": 2,
	}
)

type NotificationEvent struct {
	Type NotificationEventType
	Data *Subscription
}

// Notification represents a notification sent to users
type Notification struct {
	ID        uuid.UUID         `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID         `gorm:"type:char(36);not null;index" json:"user_id"`
	Type      NotificationEvent `gorm:"type:varchar(50);not null" json:"type"` // trial_expiring, trial_expired, etc.
	Message   string            `gorm:"type:text;not null" json:"message"`
	Read      bool              `gorm:"default:false" json:"read"`
	SentAt    time.Time         `gorm:"not null" json:"sent_at"`
	CreatedAt time.Time         `json:"created_at"`
}

// BeforeCreate hook
func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	if n.SentAt.IsZero() {
		n.SentAt = time.Now()
	}
	return nil
}
