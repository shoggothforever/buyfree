// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q             = new(Query)
	DEVICE        *dEVICE
	DeviceProduct *deviceProduct
	Driver        *driver
	Factory       *factory
	LoginInfo     *loginInfo
	Passenger     *passenger
	Platform      *platform
	User          *user
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	DEVICE = &Q.DEVICE
	DeviceProduct = &Q.DeviceProduct
	Driver = &Q.Driver
	Factory = &Q.Factory
	LoginInfo = &Q.LoginInfo
	Passenger = &Q.Passenger
	Platform = &Q.Platform
	User = &Q.User
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:            db,
		DEVICE:        newDEVICE(db, opts...),
		DeviceProduct: newDeviceProduct(db, opts...),
		Driver:        newDriver(db, opts...),
		Factory:       newFactory(db, opts...),
		LoginInfo:     newLoginInfo(db, opts...),
		Passenger:     newPassenger(db, opts...),
		Platform:      newPlatform(db, opts...),
		User:          newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	DEVICE        dEVICE
	DeviceProduct deviceProduct
	Driver        driver
	Factory       factory
	LoginInfo     loginInfo
	Passenger     passenger
	Platform      platform
	User          user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		DEVICE:        q.DEVICE.clone(db),
		DeviceProduct: q.DeviceProduct.clone(db),
		Driver:        q.Driver.clone(db),
		Factory:       q.Factory.clone(db),
		LoginInfo:     q.LoginInfo.clone(db),
		Passenger:     q.Passenger.clone(db),
		Platform:      q.Platform.clone(db),
		User:          q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		DEVICE:        q.DEVICE.replaceDB(db),
		DeviceProduct: q.DeviceProduct.replaceDB(db),
		Driver:        q.Driver.replaceDB(db),
		Factory:       q.Factory.replaceDB(db),
		LoginInfo:     q.LoginInfo.replaceDB(db),
		Passenger:     q.Passenger.replaceDB(db),
		Platform:      q.Platform.replaceDB(db),
		User:          q.User.replaceDB(db),
	}
}

type queryCtx struct {
	DEVICE        IDEVICEDo
	DeviceProduct IDeviceProductDo
	Driver        IDriverDo
	Factory       IFactoryDo
	LoginInfo     ILoginInfoDo
	Passenger     IPassengerDo
	Platform      IPlatformDo
	User          IUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		DEVICE:        q.DEVICE.WithContext(ctx),
		DeviceProduct: q.DeviceProduct.WithContext(ctx),
		Driver:        q.Driver.WithContext(ctx),
		Factory:       q.Factory.WithContext(ctx),
		LoginInfo:     q.LoginInfo.WithContext(ctx),
		Passenger:     q.Passenger.WithContext(ctx),
		Platform:      q.Platform.WithContext(ctx),
		User:          q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
