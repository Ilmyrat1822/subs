package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/Ilmyrat1822/subs/internal/modules/subscription/dtos"
	"github.com/Ilmyrat1822/subs/internal/modules/subscription/service"
)

type SubscriptionHandler struct {
	service *service.SubscriptionService
}

func NewSubscriptionHandler(service *service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

// CreateSubscription godoc
// @Summary Create subscription
// @Description Create a new subscription
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body dtos.CreateSubscriptionRequest true "Subscription data"
// @Success 201 {object} models.Subscription
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/subs [post]
func (h *SubscriptionHandler) Create(c echo.Context) error {
	var req dtos.CreateSubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Error: err.Error()})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Error: err.Error()})
	}

	sub, err := h.service.Create(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, sub)
}

// GetSubscription godoc
// @Summary Get subscription
// @Description Get subscription by ID
// @Tags subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} models.Subscription
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/subs/{id} [get]
func (h *SubscriptionHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Error: "invalid id"})
	}

	sub, err := h.service.Get(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dtos.ErrorResponse{Error: "not found"})
	}

	return c.JSON(http.StatusOK, sub)
}

// ListSubscriptions godoc
// @Summary List subscriptions
// @Description List all subscriptions with optional filters
// @Tags subscriptions
// @Produce json
// @Param user_id query string false "User ID (UUID)"
// @Param service_name query string false "Service name"
// @Success 200 {array} models.Subscription
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/subs/list [get]
func (h *SubscriptionHandler) List(c echo.Context) error {
	userID := c.QueryParam("user_id")
	serviceName := c.QueryParam("service_name")

	subs, err := h.service.List(userID, serviceName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, subs)
}

// UpdateSubscription godoc
// @Summary Update subscription
// @Description Update subscription by ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Param subscription body dtos.UpdateSubscriptionRequest true "Updated subscription data"
// @Success 200 {object} models.Subscription
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/subs/{id} [put]
func (h *SubscriptionHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Error: "invalid id"})
	}

	var req dtos.UpdateSubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Error: err.Error()})
	}

	sub, err := h.service.Update(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, sub)
}

// DeleteSubscription godoc
// @Summary Delete subscription
// @Description Delete subscription by ID
// @Tags subscriptions
// @Param id path int true "Subscription ID"
// @Success 204
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/subs/{id} [delete]
func (h *SubscriptionHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Error: "invalid id"})
	}

	if err := h.service.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, dtos.ErrorResponse{Error: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

// GetTotalCost godoc
// @Summary Get total cost
// @Description Calculate total cost of subscriptions for a period
// @Tags subscriptions
// @Produce json
// @Param start_date query string true "Start date (MM-YYYY)"
// @Param end_date query string true "End date (MM-YYYY)"
// @Param user_id query string false "User ID (UUID)"
// @Param service_name query string false "Service name"
// @Success 200 {object} dtos.TotalCostResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /api/subs/total [get]
func (h *SubscriptionHandler) TotalCost(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")

	if startDate == "" || endDate == "" {
		return c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Error: "start_date and end_date are required",
		})
	}

	resp, err := h.service.GetTotalCost(
		startDate,
		endDate,
		c.QueryParam("user_id"),
		c.QueryParam("service_name"),
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}
