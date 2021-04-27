package helper

import (
	"context"

	"github.com/palantir/stacktrace"
	"gorm.io/gorm"
)

type Transactioner interface {
	Run(ctx context.Context, fun func(ormer *gorm.DB) error) error
}

type transactioner struct {
	ormerConstructor func() *gorm.DB
}

func (t *transactioner) Run(ctx context.Context, fun func(ormer *gorm.DB) error) error {
	ormer := t.ormerConstructor()

	return run(ctx, ormer, fun)
}

func run(ctx context.Context, ormer *gorm.DB, fun func(txnOrmer *gorm.DB) error) error {
	isFinished := false

	txn := ormer.Begin()
	defer func() {
		if !isFinished {
			//log.WithContext(ctx).Error("Rolling back transaction in defer")
			err := txn.Rollback()
			if err != nil {
				//log.WithContext(ctx).Error("Failed to rolling back transaction in transactioner defer")
			}
		}
	}()

	err := fun(txn)
	if err != nil {
		//log.WithContext(ctx).Error("Rolling back transaction")

		txn = txn.Rollback()
		isFinished = true
		if txn.Error != nil {
			//log.WithContext(ctx).Error(err.Error())
			return stacktrace.Propagate(txn.Error, "Failed to rollback transaction")
		}
		return err
	}

	txn = txn.Commit()
	isFinished = true // Once you try to commit/rollback, dont try it again in defer regardless of its result
	if txn.Error != nil {
		return stacktrace.Propagate(txn.Error, "Failed to commit transaction")
	}
	return nil
}

func NewTransactioner(ormerConstructor func() *gorm.DB) Transactioner {
	return &transactioner{ormerConstructor: ormerConstructor}
}
