package service

import (
	"backend/common/producer"
	"backend/onlineconversation/internal/client"
	"backend/onlineconversation/internal/dto"
	"backend/onlineconversation/internal/projection"
	"backend/onlineconversation/internal/repository"
	"context"

	"github.com/google/uuid"
)

type Service interface {
	CreateConversation(ctx context.Context, memberId uuid.UUID, req dto.CreateConversationRequest) (map[string]uuid.UUID, error)
	UpdateConversation(ctx context.Context, memberId uuid.UUID, req dto.UpdateConversationRequest) error
	DeleteConversation(ctx context.Context, memberId, conversationId uuid.UUID) error
	BanParticipant(ctx context.Context, modId, conversationId, banId uuid.UUID) error
	GetConversationDetail(ctx context.Context, conversationId, memberId uuid.UUID) (*dto.OnlineConversationDetailResponse, error)
	RegisterConversation(ctx context.Context, memberId, conversationId uuid.UUID) error
	DeregisterConversation(ctx context.Context, memberId, conversationId uuid.UUID) error
	ScheduleNotification(ctx context.Context, memberId, conversationId uuid.UUID) error
	CancelNotification(ctx context.Context, memberId, conversationId uuid.UUID) error
	ReportConversation(ctx context.Context, conversationId, memberId uuid.UUID) error
	ManageReport(ctx context.Context, conversationId uuid.UUID) error
}

type service struct {
	repository repository.Repository
	producer   producer.SyncProducer
	moderation client.ModerationClient
}

// the calls of a conversation, its websocket and TURN credentials, are my-server's chat's
func NewService(r repository.Repository, p producer.SyncProducer, m client.ModerationClient) Service {
	s := &service{
		repository: r,
		producer:   p,
		moderation: m,
	}

	return s
}

func about(d projection.Detail) string {
	aboutRaw := []rune(d.Novel + d.Play + d.Poem + d.ShortStory + d.Film + d.WrittenBy)
	if len(aboutRaw) > 6 {
		return string(aboutRaw[:6]) + "..."
	}
	return string(aboutRaw)
}
