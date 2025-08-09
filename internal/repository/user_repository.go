package repository

import (
	"context"
	"database/sql"
	"dione-backend/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type userRepository struct {
	db *goqu.Database
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", db),
	}
}

func (u userRepository) FindAll(ctx context.Context, filter domain.UserFilter) (result []domain.User, err error) {
	dataset := u.db.From("users")
	if filter.Limit > 0 {
		dataset = dataset.Limit(uint(filter.Limit))
	} else {
		dataset = dataset.Limit(50)
	}

	if filter.Cursor != "" {
		if cursorUUID, e := uuid.Parse(filter.Cursor); e == nil {
			dataset = dataset.Where(goqu.C("id").Gt(cursorUUID))
		}
	}

	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (u userRepository) FindByID(ctx context.Context, id string) (result domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (u userRepository) FindByIDs(ctx context.Context, ids []string) (result []domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("id").In(ids))
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (u userRepository) FindByEmail(ctx context.Context, email string) (result domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.C("email").Eq(email))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (u userRepository) Save(ctx context.Context, user *domain.User) error {
	executor := u.db.Insert("users").Rows(user).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (u userRepository) Update(ctx context.Context, user *domain.User) error {
	executor := u.db.Update("users").Where(goqu.C("id").Eq(user.ID)).Set(user).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (u userRepository) Delete(ctx context.Context, id string) error {
	executor := u.db.From("users").Where(goqu.C("id").Eq(id)).Delete().Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
