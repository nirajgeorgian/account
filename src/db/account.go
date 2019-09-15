package db

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"

	"github.com/nirajgeorgian/account/src/model"
)

// CreateAccount :- create's an account
func (db *Database) CreateAccount(ctx context.Context, in *model.Account) (*model.Account, error) {
	fmt.Println("database: CreateAccount")
	tx := db.Begin()
	accountORM, err := in.ToORM(ctx)

	uuid := uuid.NewV4()
	accountORM.AccountId = uuid.String()

	if err != nil {
		return nil, errors.New("error converting input to ORM")
	}

	// hash password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(accountORM.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error hashing password: %v", err))
	}
	accountORM.PasswordHash = string(hashedPass)

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
		return nil, errors.New("error commiting job create coommit")
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
