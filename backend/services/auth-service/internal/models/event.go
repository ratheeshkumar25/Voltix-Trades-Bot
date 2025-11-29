package models

// WbSocket events
const (
	//General
	EventBadRequest          EventType = "bad_request"
	EventInternalServerError EventType = "internal_server_error"
	EventNotFound            EventType = "not_found"
	EventUnauthorized        EventType = "unauthorized"
	EventForbidden           EventType = "forbidden"

	// Operations
	EventOperationCreate    EventType = "operation_create"
	EventOperationDelete    EventType = "operation_delete"
	EventOperationDeleteAll EventType = "operation_delete_all"

	// Sessions orders
	EventSessionLogin        EventType = "session_login"
	EventSessionLogout       EventType = "session_logout"
	EventSessionLogoutAll    EventType = "session_logout_all"
	EventSessionClientLogout EventType = "session_client_logout"
)

type EventAction int32

const (
	EventAction_Nothing = 0
	EventAction_Trade   = 1
	EventAction_Notice  = 2
)

// Enum value maps for EventAction
var (
	EventActiom_Name = map[int32]string{
		0: "nothing",
		1: "trade",
		2: "notice",
	}
	EventActiom_Values = map[string]int32{
		"nothing": 0,
		"trade":   1,
		"notice":  2,
	}
)

type SeverityLevel int32

const (
	SeverityLevel_general = 0
	SeverityLevel_danger  = 1
	SeverityLevel_exterem = 2
)

// Enum value maps for SeverityLevel.
var (
	SeverityLevel_name = map[int32]string{
		0: "general",
		1: "danger",
		2: "exterem",
	}
	SeverityLevel_value = map[string]int32{
		"general": 0,
		"danger":  1,
		"exterem": 2,
	}
)

type EventSource int32

const (
	EventSource_core       = 0
	EventSource_server     = 1
	EventSource_device     = 2
	EventSource_webapi     = 3
	EventSource_automation = 4
)

// Enum value maps for EventSource.
var (
	EventSource_name = map[int32]string{
		0: "core",
		1: "server",
		2: "device",
		3: "webapi",
		4: "automation",
	}
	EventSource_value = map[string]int32{
		"core":       0,
		"server":     1,
		"device":     2,
		"webapi":     3,
		"automation": 4, // automation script
	}
)

type SystemEvent struct {
	Id        int32         `gorm:"primaryKey;autoIncrement:true;column:id" json:"id"`
	AccountId int32         `gorm:"column:account_id" json:"account_id"`
	Source    EventSource   `gorm:"column:source" json:"source"`
	Desc      string        `gorm:"column:desc" json:"desc"`
	Action    EventAction   `gorm:"column:action" json:"action"`
	Severity  SeverityLevel `gorm:"column:severity" json:"severity"`
}

type EventType string

type MessageType int32

const (
	JsonMessage   MessageType = 0 // json
	TextMessage   MessageType = 1 // text
	BinaryMessage MessageType = 2 // binary
	PingMessage   MessageType = 3 // binary
	PongMessage   MessageType = 4 // binary
)

// Websocket Event
type Event struct {
	Type      EventType   `json:"type"`
	SessionId string      `json:"session_id"`
	Payload   interface{} `json:"payload"`
	Format    MessageType `json:"-"`
}

// Websocket Error payload Event
type ErrorPayload struct {
	Message string    `json:"message"`
	Reason  EventType `json:"reason"`
}
