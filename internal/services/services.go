package services

import (
	"Game_Mode_Usage_Web_service/internal/storage"
	"context"
)

type ModeService struct {
	repo    *storage.RedisRepository
	webhook *WebhookService 
}


func NewModeService(repo *storage.RedisRepository, webhook *WebhookService) *ModeService {
	return &ModeService{repo: repo, webhook: webhook}
}

func (s *ModeService) GetModeCounts(ctx context.Context, areaCode string) (map[string]int, error) {
	return s.repo.GetModeCounts(ctx, areaCode)
}


func (s *ModeService) JoinMode(ctx context.Context, areaCode, mode string) error {
	if err := s.repo.UpdateModeCount(ctx, areaCode, mode, 1); err != nil {
		return err
	}

	count, err := s.repo.GetModeCount(ctx, areaCode, mode)
	if err != nil {
		return err
	}

	s.webhook.Notify(areaCode, mode, count)
	return nil
}


func (s *ModeService) LeaveMode(ctx context.Context, areaCode, mode string) error {
	if err := s.repo.UpdateModeCount(ctx, areaCode, mode, -1); err != nil {
		return err
	}

	count, err := s.repo.GetModeCount(ctx, areaCode, mode)
	if err != nil {
		return err
	}

	s.webhook.Notify(areaCode, mode, count)
	return nil
}

func (s *ModeService) Unsubscribe(areaCode, url string) {
	s.webhook.Unsubscribe(areaCode, url)
}

func (s *ModeService) Subscribe(areaCode, url string) {
	s.webhook.Subscribe(areaCode, url)
}
