package service

import (
	"errors"
	"fmt"
	"time"

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

	reqStart, err := parseMonth(startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start_date format, expected MM-YYYY")
	}

	reqEnd, err := parseMonth(endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end_date format, expected MM-YYYY")
	}

	if reqStart.After(reqEnd) {
		return nil, fmt.Errorf("start_date must be before end_date")
	}

	subs, err := s.repo.GetForPeriod(startDate, endDate, userID, serviceName)
	if err != nil {
		return nil, err
	}

	total := 0

	for _, sub := range subs {
		subStart, _ := parseMonth(sub.StartDate)

		subEnd := reqEnd
		if sub.EndDate != nil {
			subEnd, _ = parseMonth(*sub.EndDate)
		}

		actualStart := maxTime(reqStart, subStart)
		actualEnd := minTime(reqEnd, subEnd)

		if actualStart.After(actualEnd) {
			continue
		}

		months := monthsBetween(actualStart, actualEnd)
		total += months * sub.Price
	}

	return &dtos.TotalCostResponse{
		TotalCost: total,
		Count:     len(subs),
		Period:    fmt.Sprintf("%s to %s", startDate, endDate),
	}, nil
}
func parseMonth(value string) (time.Time, error) {
	return time.Parse("01-2006", value)
}

func monthsBetween(start, end time.Time) int {
	return (end.Year()-start.Year())*12 + int(end.Month()-start.Month()) + 1
}

func maxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

func minTime(a, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}
	return b
}
