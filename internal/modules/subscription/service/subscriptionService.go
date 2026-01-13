package service

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/Ilmyrat1822/subs/internal/models"
	"github.com/Ilmyrat1822/subs/internal/modules/subscription/dtos"
	"github.com/Ilmyrat1822/subs/internal/modules/subscription/repository"
)

type SubscriptionService struct {
	repo *repository.SubscriptionRepository
}

func NewSubscriptionService(repo *repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(req dtos.CreateSubscriptionRequest) (*models.Subscription, error) {
	sub := &models.Subscription{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	return sub, s.repo.Create(sub)
}

func (s *SubscriptionService) Get(id int) (*models.Subscription, error) {
	return s.repo.GetByID(id)
}

func (s *SubscriptionService) List(userID, serviceName string) ([]models.Subscription, error) {
	return s.repo.List(userID, serviceName)
}

func (s *SubscriptionService) Update(id int, req dtos.UpdateSubscriptionRequest) (*models.Subscription, error) {
	sub, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.ServiceName != nil {
		sub.ServiceName = *req.ServiceName
	}
	if req.Price != nil {
		sub.Price = *req.Price
	}
	if req.StartDate != nil {
		sub.StartDate = *req.StartDate
	}
	if req.EndDate != nil {
		sub.EndDate = req.EndDate
	}

	return sub, s.repo.Update(sub)
}

func (s *SubscriptionService) Delete(id int) error {
	err := s.repo.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("subscription not found")
	}
	return err
}

func (s *SubscriptionService) GetTotalCost(
	startDate, endDate, userID, serviceName string,
) (*dtos.TotalCostResponse, error) {

	total, count, err := s.repo.CalculateTotalCost(startDate, endDate, userID, serviceName)
	if err != nil {
		return nil, err
	}

	return &dtos.TotalCostResponse{
		TotalCost: total,
		Count:     count,
		Period:    fmt.Sprintf("%s to %s", startDate, endDate),
	}, nil
}
