package repository

import (
	"gorm.io/gorm"

	"github.com/Ilmyrat1822/subs/internal/models"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(sub *models.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *SubscriptionRepository) GetByID(id int) (*models.Subscription, error) {
	var sub models.Subscription
	err := r.db.First(&sub, id).Error
	if err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *SubscriptionRepository) Update(sub *models.Subscription) error {
	return r.db.Save(sub).Error
}

func (r *SubscriptionRepository) Delete(id int) error {
	return r.db.Delete(&models.Subscription{}, id).Error
}

func (r *SubscriptionRepository) List(userID, serviceName string) ([]models.Subscription, error) {
	var subs []models.Subscription

	query := r.db.Model(&models.Subscription{})

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		query = query.Where("service_name ILIKE ?", "%"+serviceName+"%")
	}

	err := query.Order("created_at DESC").Find(&subs).Error
	return subs, err
}

func (r *SubscriptionRepository) CalculateTotalCost(
	startDate, endDate, userID, serviceName string,
) (int, int, error) {

	type result struct {
		Total int
		Count int
	}

	var res result

	query := r.db.Model(&models.Subscription{}).
		Select("COALESCE(SUM(price),0) as total, COUNT(*) as count").
		Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", endDate, startDate)

	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		query = query.Where("service_name ILIKE ?", "%"+serviceName+"%")
	}

	err := query.Scan(&res).Error
	return res.Total, res.Count, err
}
