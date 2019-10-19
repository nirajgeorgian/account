package db

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"github.com/nirajgeorgian/account/src/model"
)

func (db *Database) ValidateUsername(ctx context.Context, username string) (bool, error) {
	var accountORM []*model.AccountORM
	if err := db.First(&accountORM, "username = ?", username).Error; err != nil {
		return false, errors.New("error fetching account")
	}

	if len(accountORM) > 0 {
		return true, nil
	}

	return false, nil
}

func (db *Database) ValidateEmail(ctx context.Context, email string) (bool, error) {
	var accountORM []*model.AccountORM
	if err := db.First(&accountORM, "email = ?", email).Error; err != nil {
		return false, errors.New("error fetching account")
	}

	if len(accountORM) > 0 {
		return true, nil
	}

	return false, nil
}

// CreateAccount :- create's an account
func (db *Database) CreateAccount(ctx context.Context, in *model.Account) (*model.Account, error) {
	tx := db.Begin()
	defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()
	if err := tx.Error; err != nil {
    return nil, err
  }

	accountORM, err := in.ToORM(ctx)
	if err != nil {
		return nil, errors.New("error converting input to ORM")
	}

	if err = tx.Create(&accountORM).Error; err != nil {
		return nil, errors.New("error creating job")
	}

	account, err := accountORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("error converting back to PB")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error commiting job create commit")
	}

	return &account, nil
}

func (db *Database) ReadAccount(ctx context.Context, accountId string) (*model.Account, error) {
	tx := db.Begin()
	defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()
	if err := tx.Error; err != nil {
    return nil, err
  }

	var accountORM model.AccountORM

	if accountId == "" {
		return nil, errors.New("no account found for this profile")
	}

	if err := tx.First(&accountORM, "account_id = ?", accountId).Error; err != nil {
		return nil, errors.New("error fetching account")
	}

	account, err := accountORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("error converting back to PB")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error fetching account details commit")
	}

	return &account, nil
}

func (db *Database) UpdateAccount(ctx context.Context, in *model.Account) (*model.Account, error) {
	tx := db.Begin()
	defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()
	if err := tx.Error; err != nil {
    return nil, err
  }

	accountORM, err := in.ToORM(ctx)
	if err != nil {
		return nil, errors.New("error converting input to ORM")
	}

	if accountORM.AccountId == "" {
		return nil, errors.New("please pass account ID")
	}

	if err := tx.First(&accountORM, "account_id = ?", accountORM.AccountId).Error; err != nil {
		return nil, errors.New("error fetching account")
	}

	if err := tx.Model(&accountORM).Updates(model.Account{Description: in.Description, Username: in.Username, Email: in.Email}).Error; err != nil {
		return nil, errors.New("error updating account")
	}

	account, err := accountORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("error converting back to PB")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error commiting job update commit")
	}

	return &account, nil
}

// // ReadAccount :- read account profile
// func (db *Database) ReadAccount(ctx context.Context, accountID string) error {
// 	account := &model.Account{}
// 	db.First(&account, accountID)
//
// 	if account.GetAccountId() == "" {
// 		return errors.New("no account found for this profile")
// 	}
//
// 	return nil
// }
//
// // UpdateAccount :- update any existing account profile
// func (db *Database) UpdateAccount(ctx context.Context, in *model.Account) error {
// 	if in.GetAccountId() == "" {
// 		return errors.New("empty account id")
// 	}
//
// 	var originalORM model.AccountORM
// 	// var originalAccount model.Account
//
// 	tx := db.Begin()
// 	success := false
// 	defer func() {
// 		if success {
// 			if err := tx.Commit().Error; err == nil {
// 				return
// 			}
// 		}
// 		tx.Rollback()
// 	}()
//
// 	if err := tx.First(&originalORM, "account_id", in.GetAccountId()).Error; err != nil {
// 		return errors.New("failed finding account")
// 	}
//
// 	// originalAccount, err := originalORM.ToPB(ctx)
// 	// if err != nil {
// 	// 	return errors.New("failed transforming ORM to Proto for account")
// 	// }
//
// 	return nil
// }
