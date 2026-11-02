package common

import (
	"encoding/json"

	"github.com/google/uuid"
)

type OnlineConversationSignal struct {
	FromId []byte          `json:"fromId"`
	ToIds  [][]byte        `json:"toIds"`
	Signal json.RawMessage `json:"signal,omitempty"`
}

type NotificationScheduling struct {
	ScheduledTime int64          `json:"scheduledTime,omitempty"`
	PartitionId   uuid.UUID      `json:"partitionId"`
	KeyId         uuid.UUID      `json:"keyId"`
	Contents      map[int]string `json:"contents"`
	Type          string         `json:"type,omitempty"`
}

// RetryEvent is sent to "exponential-backoff-retry", the job re-sends Value to Topic with Key and Headers
// plus a "partitionId" header holding PartitionId. A follow-up failure only needs PartitionId and Reason,
// []byte fields are base64 in json to match byte[] of the job
type RetryEvent struct {
	PartitionId string            `json:"partitionId"` // "<group>:<topic>:<partition>:<offset>" of the first failed record
	Reason      string            `json:"reason"`
	Backoff     int64             `json:"backoff,omitempty"`
	Multiplier  int64             `json:"multiplier,omitempty"`
	Cap         int64             `json:"cap,omitempty"`
	MaxFailure  int               `json:"maxFailure,omitempty"`
	Topic       string            `json:"topic,omitempty"`
	Key         []byte            `json:"key,omitempty"`
	Headers     map[string][]byte `json:"headers,omitempty"` // original headers without "partitionId"
	Value       []byte            `json:"value,omitempty"`
}

type ConversationRequest struct {
	Id uuid.UUID `json:"id"`
}
