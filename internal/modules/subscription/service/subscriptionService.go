package service

import (
	"errors"
	"fmt"

	"github.com/Ilmyrat1822/subs/internal/models"
	"github.com/Ilmyrat1822/subs/internal/modules/subscription/dtos"
	"github.com/Ilmyrat1822/subs/internal/modules/subscription/repository"
	"gorm.io/gorm"
)

type SubscriptionService interface {
	Create(req dtos.CreateSubscriptionRequest) (*models.Subscription, error)
	Get(id int) (*models.Subscription, error)
	List(
		userID, serviceName string,
		limit, offset int,
	) ([]models.Subscription, *dtos.PaginationMeta, error)
	Update(id int, req dtos.UpdateSubscriptionRequest) (*models.Subscription, error)
	Delete(id int) error
	GetTotalCost(
		startDate, endDate, userID, serviceName string,
	) (*dtos.TotalCostResponse, error)
}

type subscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(
	repo repository.SubscriptionRepository,
) SubscriptionService {
	return &subscriptionService{repo: repo}
}

func (s *subscriptionService) Create(req dtos.CreateSubscriptionRequest) (*models.Subscription, error) {
	sub := &models.Subscription{
		ServiceName: req.ServiceName,
		Price:       req.Price,
		UserID:      req.UserID,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}
	return sub, s.repo.Create(sub)
}

func (s *subscriptionService) Get(id int) (*models.Subscription, error) {
	sub, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSubscriptionNotFound
		}
		return nil, err
	}
	return sub, nil
}

const (
	defaultLimit = 20
	maxLimit     = 100
)

func (s *subscriptionService) List(userID, serviceName string, limit, offset int) ([]models.Subscription, *dtos.PaginationMeta, error) {

	if limit <= 0 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}
	if offset < 0 {
		offset = 0
	}

	subs, total, err := s.repo.List(userID, serviceName, limit, offset)
	if err != nil {
		return nil, nil, err
	}

	meta := &dtos.PaginationMeta{
		Limit:  limit,
		Offset: offset,
		Total:  int(total),
	}

	return subs, meta, nil
}

func (s *subscriptionService) Update(id int, req dtos.UpdateSubscriptionRequest) (*models.Subscription, error) {

	sub, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSubscriptionNotFound
		}
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
	updated, err := s.repo.Update(sub)
	if err != nil {
		return nil, err
	}

	if !updated {
		return nil, fmt.Errorf("subscription not found")
	}

	return sub, nil

}

var ErrSubscriptionNotFound = errors.New("subscription not found")

func (s *subscriptionService) Delete(id int) error {
	found, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	if !found {
		return ErrSubscriptionNotFound
	}
	return nil
}

func (s *subscriptionService) GetTotalCost(startDate, endDate, userID, serviceName string) (*dtos.TotalCostResponse, error) {
	if startDate == "" || endDate == "" {
		return nil, fmt.Errorf("start_date and end_date are required")
	}

	result, err := s.repo.GetTotalCost(
		startDate,
		endDate,
		userID,
		serviceName,
	)
	if err != nil {
		return nil, err
	}

	return &dtos.TotalCostResponse{
		Total: int64(result.Total),
		Count: int64(result.Count),
	}, nil
}
