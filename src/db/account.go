package db

import (
	"fmt"

	"go.opencensus.io/trace"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/satori/go.uuid"

	"github.com/nirajgeorgian/account/src/model"
)

func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}

func (db *Database) ValidateUsername(ctx context.Context, username string) (bool, error) {
	ctx, span := trace.StartSpan(ctx, "account.grpc.db.Auth")
	defer span.End()

	var accountORM []*model.AccountORM
	if err := db.First(&accountORM, "username = ?", username).Error; err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return false, err
	}

	if len(accountORM) > 0 {
		return true, nil
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("query", "ValidateUsername"),
	}, "account src/db")

	return false, nil
}

func (db *Database) ValidateEmail(ctx context.Context, email string) (bool, error) {
	ctx, span := trace.StartSpan(ctx, "account.grpc.db.ValidateEmail")
	defer span.End()

	var accountORM []*model.AccountORM
	if err := db.First(&accountORM, "email = ?", email).Error; err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		err := status.Error(codes.NotFound, "no account found")
		return false, err
	}

	if len(accountORM) > 0 {
		return true, nil
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("query", "ValidateEmail"),
	}, "account src/db")

	return false, nil
}

// CreateAccount :- create's an account
func (db *Database) CreateAccount(ctx context.Context, in *model.Account) (*model.Account, error) {
	ctx, span := trace.StartSpan(ctx, "account.grpc.db.CreateAccount")
	defer span.End()

	tx := db.Begin()
	defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()
	if err := tx.Error; err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		err := status.Error(codes.Internal, "database query failed")
    return nil, err
  }

	accountORM, err := in.ToORM(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		err := status.Error(codes.DataLoss, "error converting input to ORM")
		return nil, err
	}

	u1 := uuid.NewV4()
	accountORM.AccountId = u1.String()

	// hash password
	hashedPass, err := HashPassword(accountORM.PasswordHash)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		err := status.Error(codes.DataLoss, fmt.Sprintf("error hashing password: %v", err))
		return nil, err
	}
	accountORM.PasswordHash = hashedPass

	if err = tx.Create(&accountORM).Error; err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		err := status.Error(codes.DataLoss, "error converting input to ORM")
		return nil, err
	}

	account, err := accountORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		err := status.Error(codes.DataLoss, "error converting input to PB")
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		err := status.Error(codes.Internal, "database insertion failed")
		return nil, err
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("query", "CreateAccount"),
	}, "account src/db")

	return &account, nil
}

func (db *Database) Auth(ctx context.Context, in *model.Account) (*model.Account, error) {
	ctx, span := trace.StartSpan(ctx, "account.grpc.db.Auth")
	defer span.End()

	accountORM, err := in.ToORM(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("error converting input to ORM")
	}

	db.Where("email = ?", accountORM.Email).First(&accountORM)

	account, err := accountORM.ToPB(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("error converting back to PB")
	}
	if account.GetAccountId() == "" {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("no account found for this profile")
	}

	if matched := CheckPasswordHash(in.PasswordHash, account.PasswordHash); !matched {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("password do not matched")
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("query", "Auth"),
	}, "account src/db")

	return &account, nil
}

func (db *Database) ReadAccount(ctx context.Context, accountId string) (*model.Account, error) {
	ctx, span := trace.StartSpan(ctx, "account.grpc.db.ReadAccount")
	defer span.End()

	tx := db.Begin()
	defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()
	if err := tx.Error; err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
    return nil, err
  }

	var accountORM model.AccountORM

	if accountId == "" {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: "no account found for this profile"})
		return nil, errors.New("no account found for this profile")
	}

	if err := tx.First(&accountORM, "account_id = ?", accountId).Error; err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("error fetching account")
	}

	account, err := accountORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("error converting back to PB")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("error fetching account details commit")
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("query", "ReadAccount"),
	}, "account src/db")

	return &account, nil
}

func (db *Database) UpdateAccount(ctx context.Context, in *model.Account) (*model.Account, error) {
	ctx, span := trace.StartSpan(ctx, "account.grpc.db.UpdateAccount")
	defer span.End()

	tx := db.Begin()
	defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()
	if err := tx.Error; err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
    return nil, err
  }

	accountORM, err := in.ToORM(ctx)
	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("error converting input to ORM")
	}

	if accountORM.AccountId == "" {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("please pass account ID")
	}

	if err := tx.First(&accountORM, "account_id = ?", accountORM.AccountId).Error; err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("error fetching account")
	}

	if err := tx.Model(&accountORM).Updates(model.Account{Description: in.Description, Username: in.Username, Email: in.Email}).Error; err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("error updating account")
	}

	account, err := accountORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("error converting back to PB")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, errors.New("error commiting job update commit")
	}

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("query", "UpdateAccount"),
	}, "account src/db")

	return &account, nil
}
