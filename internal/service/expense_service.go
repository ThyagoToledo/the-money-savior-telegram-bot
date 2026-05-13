package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"money-telegram-bot/internal/models"
	"money-telegram-bot/internal/repository"
	"money-telegram-bot/internal/utils"
)

type ExpenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) GetByID(ctx context.Context, userID, expenseID string) (*models.Expense, error) {
	return s.repo.FindByID(ctx, userID, expenseID)
}
func (s *ExpenseService) ListByUser(ctx context.Context, userID string) ([]models.Expense, error) {
	return s.repo.FindByUser(ctx, userID)
}

func (s *ExpenseService) DeleteAllExpenses(ctx context.Context, userID int64) error {
	return s.repo.DeleteAllExpenses(ctx, userID)
}

func (s *ExpenseService) DeleteExpense(ctx context.Context, userID int64, expenseID string) error {
	return s.repo.DeleteByID(ctx, strconv.FormatInt(userID, 10), expenseID)
}

func (s *ExpenseService) CreateExpense(
	ctx context.Context,
	userID int64,
	chatID int64,
	username string,
	amount float64,
	categoryInput string,
	methodInput string,
) (*models.Expense, error) {

	nowUtc := time.Now().UTC()
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return nil, err
	}

	brazilCreationTime := nowUtc.In(loc)

	expense := &models.Expense{
		UserID:    userID,
		ChatID:    chatID,
		Username:  username,
		Amount:    amount,
		Category:  utils.FormatTitle(categoryInput),
		Method:    utils.NormalizeMethod(methodInput),
		CreatedAt: brazilCreationTime,
	}

	if err := s.repo.Save(ctx, expense); err != nil {
		return nil, err
	}

	return expense, nil
}
