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
	Q                  = new(Query)
	Advertisement      *advertisement
	Device             *device
	DeviceProduct      *deviceProduct
	Driver             *driver
	DriverCart         *driverCart
	DriverOrderForm    *driverOrderForm
	Factory            *factory
	FactoryProduct     *factoryProduct
	LoginInfo          *loginInfo
	OrderProduct       *orderProduct
	Passenger          *passenger
	PassengerCart      *passengerCart
	PassengerOrderForm *passengerOrderForm
	Platform           *platform
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Advertisement = &Q.Advertisement
	Device = &Q.Device
	DeviceProduct = &Q.DeviceProduct
	Driver = &Q.Driver
	DriverCart = &Q.DriverCart
	DriverOrderForm = &Q.DriverOrderForm
	Factory = &Q.Factory
	FactoryProduct = &Q.FactoryProduct
	LoginInfo = &Q.LoginInfo
	OrderProduct = &Q.OrderProduct
	Passenger = &Q.Passenger
	PassengerCart = &Q.PassengerCart
	PassengerOrderForm = &Q.PassengerOrderForm
	Platform = &Q.Platform
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                 db,
		Advertisement:      newAdvertisement(db, opts...),

		Device:             newDevice(db, opts...),
		DeviceProduct:      newDeviceProduct(db, opts...),
		Driver:             newDriver(db, opts...),
		DriverCart:         newDriverCart(db, opts...),
		DriverOrderForm:    newDriverOrderForm(db, opts...),
		Factory:            newFactory(db, opts...),
		FactoryProduct:     newFactoryProduct(db, opts...),
		LoginInfo:          newLoginInfo(db, opts...),
		OrderProduct:       newOrderProduct(db, opts...),
		Passenger:          newPassenger(db, opts...),
		PassengerCart:      newPassengerCart(db, opts...),
		PassengerOrderForm: newPassengerOrderForm(db, opts...),
		Platform:           newPlatform(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Advertisement      advertisement

	Device             device
	DeviceProduct      deviceProduct
	Driver             driver
	DriverCart         driverCart
	DriverOrderForm    driverOrderForm
	Factory            factory
	FactoryProduct     factoryProduct
	LoginInfo          loginInfo
	OrderProduct       orderProduct
	Passenger          passenger
	PassengerCart      passengerCart
	PassengerOrderForm passengerOrderForm
	Platform           platform
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		Advertisement:      q.Advertisement.clone(db),
		Device:             q.Device.clone(db),
		DeviceProduct:      q.DeviceProduct.clone(db),
		Driver:             q.Driver.clone(db),
		DriverCart:         q.DriverCart.clone(db),
		DriverOrderForm:    q.DriverOrderForm.clone(db),
		Factory:            q.Factory.clone(db),
		FactoryProduct:     q.FactoryProduct.clone(db),
		LoginInfo:          q.LoginInfo.clone(db),
		OrderProduct:       q.OrderProduct.clone(db),
		Passenger:          q.Passenger.clone(db),
		PassengerCart:      q.PassengerCart.clone(db),
		PassengerOrderForm: q.PassengerOrderForm.clone(db),
		Platform:           q.Platform.clone(db),
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
		db:                 db,
		Advertisement:      q.Advertisement.replaceDB(db),
		Device:             q.Device.replaceDB(db),
		DeviceProduct:      q.DeviceProduct.replaceDB(db),
		Driver:             q.Driver.replaceDB(db),
		DriverCart:         q.DriverCart.replaceDB(db),
		DriverOrderForm:    q.DriverOrderForm.replaceDB(db),
		Factory:            q.Factory.replaceDB(db),
		FactoryProduct:     q.FactoryProduct.replaceDB(db),
		LoginInfo:          q.LoginInfo.replaceDB(db),
		OrderProduct:       q.OrderProduct.replaceDB(db),
		Passenger:          q.Passenger.replaceDB(db),
		PassengerCart:      q.PassengerCart.replaceDB(db),
		PassengerOrderForm: q.PassengerOrderForm.replaceDB(db),
		Platform:           q.Platform.replaceDB(db),
	}
}

type queryCtx struct {
	Advertisement      IAdvertisementDo
	Device             IDeviceDo
	DeviceProduct      IDeviceProductDo
	Driver             IDriverDo
	DriverCart         IDriverCartDo
	DriverOrderForm    IDriverOrderFormDo
	Factory            IFactoryDo
	FactoryProduct     IFactoryProductDo
	LoginInfo          ILoginInfoDo
	OrderProduct       IOrderProductDo
	Passenger          IPassengerDo
	PassengerCart      IPassengerCartDo
	PassengerOrderForm IPassengerOrderFormDo
	Platform           IPlatformDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Advertisement:      q.Advertisement.WithContext(ctx),
		Device:             q.Device.WithContext(ctx),
		DeviceProduct:      q.DeviceProduct.WithContext(ctx),
		Driver:             q.Driver.WithContext(ctx),
		DriverCart:         q.DriverCart.WithContext(ctx),
		DriverOrderForm:    q.DriverOrderForm.WithContext(ctx),
		Factory:            q.Factory.WithContext(ctx),
		FactoryProduct:     q.FactoryProduct.WithContext(ctx),
		LoginInfo:          q.LoginInfo.WithContext(ctx),
		OrderProduct:       q.OrderProduct.WithContext(ctx),
		Passenger:          q.Passenger.WithContext(ctx),
		PassengerCart:      q.PassengerCart.WithContext(ctx),
		PassengerOrderForm: q.PassengerOrderForm.WithContext(ctx),
		Platform:           q.Platform.WithContext(ctx),
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
