package db

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"github.com/nirajgeorgian/account/src/model"
)

// CreateAccount :- create's an account
func (db *Database) CreateAccount(ctx context.Context, in *model.Account) error {
	tx := db.Begin()
	accountORM, err := in.ToORM(ctx)
	if err != nil {
		return errors.New("error converting input to ORM")
	}
	if err = tx.Create(&accountORM).Error; err != nil {
		return errors.New("error creating job")
	}

	_, err = accountORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		return errors.New("error converting back to PB")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New("error commiting job create coommit")
	}

	return nil
}

func (db *Database) ReadAccount(ctx context.Context, accountId string) (*model.Account, error) {
	tx := db.Begin()
	var accountORM model.AccountORM

	if err := tx.First(&accountORM, "account_id = ?", accountId).Error; err != nil {
		return nil, errors.New("error fetching email")
	}

	account, err := accountORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("error converting back to PB")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, errors.New("error commiting job create coommit")
	}

	return &account, nil
}

func (db *Database) UpdateAccount(ctx context.Context, in *model.Account) error {
	tx := db.Begin()

	var accountORM model.AccountORM
	// var originalAccount model.Account

	// accountORM, err = in.ToORM(ctx)
	// if err != nil {
	// 	return errors.New("error converting input to ORM")
	// }

	if err := tx.First(&accountORM, "account_id = ?", in.GetAccountId()).Error; err != nil {
		return errors.New("failed finding account")
	}

	if err := tx.Model(&accountORM).Updates(accountORM).Error; err != nil {
		return errors.New("error updating account")
	}

	_, err := accountORM.ToPB(ctx)
	if err != nil {
		tx.Rollback()
		return errors.New("error converting back to PB")
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New("error commiting job create coommit")
	}

	return nil
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
