package repository

import (
	"gorm.io/gorm"

	"github.com/Ilmyrat1822/subs/internal/models"
)

type SubscriptionRepository interface {
	Create(sub *models.Subscription) error
	GetByID(id int) (*models.Subscription, error)
	Update(sub *models.Subscription) error
	Delete(id int) (bool, error)
	List(userID, serviceName string, limit, offset int) ([]models.Subscription, int64, error)
	GetTotalCost(startDate, endDate, userID, serviceName string) (*TotalAggResult, error)
}
type subscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) SubscriptionRepository {
	return &subscriptionRepository{db: db}
}

func (r *subscriptionRepository) Create(sub *models.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *subscriptionRepository) GetByID(id int) (*models.Subscription, error) {
	var sub models.Subscription
	err := r.db.First(&sub, id).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *subscriptionRepository) Update(sub *models.Subscription) error {
	return r.db.Save(sub).Error
}

func (r *subscriptionRepository) Delete(id int) (bool, error) {
	res := r.db.Delete(&models.Subscription{}, id)

	if res.Error != nil {
		return false, res.Error
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (r *subscriptionRepository) List(userID, serviceName string, limit, offset int) ([]models.Subscription, int64, error) {

	var subs []models.Subscription
	var total int64

	query := r.db.Model(&models.Subscription{})

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		query = query.Where("service_name ILIKE ?", "%"+serviceName+"%")
	}

	// count before limit
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&subs).Error

	return subs, total, err
}

type TotalAggResult struct {
	Total int64
	Count int64
}

func (r *subscriptionRepository) GetTotalCost(startDate, endDate, userID, serviceName string) (*TotalAggResult, error) {

	var res TotalAggResult

	q := r.db.
		Model(&models.Subscription{}).
		Select(`
			COALESCE(SUM(price), 0) AS total,
			COUNT(*) AS count
		`).
		Where("start_date >= ? AND start_date <= ?", startDate, endDate)

	if userID != "" {
		q = q.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		q = q.Where("service_name = ?", serviceName)
	}

	if err := q.Scan(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
