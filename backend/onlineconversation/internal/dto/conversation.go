package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateConversationRequest struct {
	Novel         string    `json:"novel"`
	ShortStory    string    `json:"shortStory"`
	Poem          string    `json:"poem"`
	Play          string    `json:"play"`
	Film          string    `json:"film"`
	WrittenBy     string    `json:"writtenBy"`
	Description   string    `json:"description"`
	Capacity      int       `json:"capacity"`
	Time          time.Time `json:"time"`
	LengthMinutes int       `json:"lengthMinutes"`
}

type UpdateConversationRequest struct {
	Id            uuid.UUID `json:"id"`
	Novel         string    `json:"novel"`
	ShortStory    string    `json:"shortStory"`
	Poem          string    `json:"poem"`
	Play          string    `json:"play"`
	Film          string    `json:"film"`
	WrittenBy     string    `json:"writtenBy"`
	Description   string    `json:"description"`
	Capacity      int       `json:"capacity"`
	Time          time.Time `json:"time"`
	LengthMinutes int       `json:"lengthMinutes"`
}

// counted in characters, not bytes, so a Korean text gets the same room as an English one
const (
	maxTitleChars       = 50
	maxDescriptionChars = 500
)

type OnlineConversationDetailResponse struct {
	Novel                   string    `json:"novel,omitempty"`
	ShortStory              string    `json:"shortStory,omitempty"`
	Poem                    string    `json:"poem,omitempty"`
	Play                    string    `json:"play,omitempty"`
	Film                    string    `json:"film,omitempty"`
	WrittenBy               string    `json:"writtenBy"`
	Description             string    `json:"description,omitempty"`
	Capacity                int       `json:"capacity"`
	Time                    time.Time `json:"time"`
	LengthMinutes           int       `json:"lengthMinutes"`
	UpdatedAt               time.Time `json:"updatedAt,omitzero"`
	CanEnter                bool      `json:"canEnter"`
	IsRegistrant            bool      `json:"isRegistrant"`
	IsModerator             bool      `json:"isModerator"`
	IsNotificationScheduled bool      `json:"isNotificationScheduled"`
}

type OnlineConversationFeedResponse struct {
	Id         uuid.UUID `json:"id"`
	Novel      string    `json:"novel,omitempty"`
	ShortStory string    `json:"shortStory,omitempty"`
	Poem       string    `json:"poem,omitempty"`
	Play       string    `json:"play,omitempty"`
	Film       string    `json:"film,omitempty"`
	WrittenBy  string    `json:"writtenBy"`
	Time       time.Time `json:"time"`
}

type BanParticipantRequest struct {
	ConversationId uuid.UUID `json:"conversationId"`
	BanId          uuid.UUID `json:"banId"`
}

// Verdict is the moderation result of a reported conversation
type Verdict struct {
	Reason    string `json:"reason"`
	Category  string `json:"category"`
	Violation bool   `json:"violation"`
}
