// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"buyfree/repo/model"
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newDriver(db *gorm.DB, opts ...gen.DOOption) driver {
	_driver := driver{}

	_driver.driverDo.UseDB(db, opts...)
	_driver.driverDo.UseModel(&model.Driver{})

	tableName := _driver.driverDo.TableName()
	_driver.ALL = field.NewAsterisk(tableName)
	_driver.ID = field.NewInt64(tableName, "id")
	_driver.CreatedAt = field.NewTime(tableName, "created_at")
	_driver.UpdatedAt = field.NewTime(tableName, "updated_at")
	_driver.DeletedAt = field.NewField(tableName, "deleted_at")
	_driver.Balance = field.NewFloat64(tableName, "balance")
	_driver.Pic = field.NewString(tableName, "pic")
	_driver.Name = field.NewString(tableName, "name")
	_driver.Password = field.NewString(tableName, "password")
	_driver.Mobile = field.NewString(tableName, "mobile")
	_driver.IDCard = field.NewString(tableName, "id_card")
	_driver.Role = field.NewInt(tableName, "role")
	_driver.Level = field.NewInt(tableName, "level")
	_driver.CarID = field.NewString(tableName, "car_id")
	_driver.PlatformID = field.NewInt64(tableName, "platform_id")
	_driver.IsAuth = field.NewBool(tableName, "is_auth")
	_driver.Location = field.NewString(tableName, "location")
	_driver.Cart = driverHasOneCart{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Cart", "model.DriverCart"),
		Products: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Cart.Products", "model.OrderProduct"),
		},
	}

	_driver.Devices = driverHasManyDevices{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Devices", "model.Device"),
		Products: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Devices.Products", "model.DeviceProduct"),
		},
	}

	_driver.DriverOrderForms = driverHasManyDriverOrderForms{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("DriverOrderForms", "model.DriverOrderForm"),
		ProductInfo: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("DriverOrderForms.ProductInfo", "model.OrderProduct"),
		},
	}

	_driver.fillFieldMap()

	return _driver
}

type driver struct {
	driverDo

	ALL        field.Asterisk
	ID         field.Int64
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	Balance    field.Float64
	Pic        field.String
	Name       field.String
	Password   field.String
	Mobile     field.String
	IDCard     field.String
	Role       field.Int
	Level      field.Int
	CarID      field.String
	PlatformID field.Int64
	IsAuth     field.Bool
	Location   field.String
	Cart       driverHasOneCart

	Devices driverHasManyDevices

	DriverOrderForms driverHasManyDriverOrderForms

	fieldMap map[string]field.Expr
}

func (d driver) Table(newTableName string) *driver {
	d.driverDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d driver) As(alias string) *driver {
	d.driverDo.DO = *(d.driverDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *driver) updateTableName(table string) *driver {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewInt64(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.Balance = field.NewFloat64(table, "balance")
	d.Pic = field.NewString(table, "pic")
	d.Name = field.NewString(table, "name")
	d.Password = field.NewString(table, "password")
	d.Mobile = field.NewString(table, "mobile")
	d.IDCard = field.NewString(table, "id_card")
	d.Role = field.NewInt(table, "role")
	d.Level = field.NewInt(table, "level")
	d.CarID = field.NewString(table, "car_id")
	d.PlatformID = field.NewInt64(table, "platform_id")
	d.IsAuth = field.NewBool(table, "is_auth")
	d.Location = field.NewString(table, "location")

	d.fillFieldMap()

	return d
}

func (d *driver) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *driver) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 19)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["balance"] = d.Balance
	d.fieldMap["pic"] = d.Pic
	d.fieldMap["name"] = d.Name
	d.fieldMap["password"] = d.Password
	d.fieldMap["mobile"] = d.Mobile
	d.fieldMap["id_card"] = d.IDCard
	d.fieldMap["role"] = d.Role
	d.fieldMap["level"] = d.Level
	d.fieldMap["car_id"] = d.CarID
	d.fieldMap["platform_id"] = d.PlatformID
	d.fieldMap["is_auth"] = d.IsAuth
	d.fieldMap["location"] = d.Location

}

func (d driver) clone(db *gorm.DB) driver {
	d.driverDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d driver) replaceDB(db *gorm.DB) driver {
	d.driverDo.ReplaceDB(db)
	return d
}

type driverHasOneCart struct {
	db *gorm.DB

	field.RelationField

	Products struct {
		field.RelationField
	}
}

func (a driverHasOneCart) Where(conds ...field.Expr) *driverHasOneCart {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a driverHasOneCart) WithContext(ctx context.Context) *driverHasOneCart {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a driverHasOneCart) Model(m *model.Driver) *driverHasOneCartTx {
	return &driverHasOneCartTx{a.db.Model(m).Association(a.Name())}
}

type driverHasOneCartTx struct{ tx *gorm.Association }

func (a driverHasOneCartTx) Find() (result *model.DriverCart, err error) {
	return result, a.tx.Find(&result)
}

func (a driverHasOneCartTx) Append(values ...*model.DriverCart) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a driverHasOneCartTx) Replace(values ...*model.DriverCart) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a driverHasOneCartTx) Delete(values ...*model.DriverCart) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a driverHasOneCartTx) Clear() error {
	return a.tx.Clear()
}

func (a driverHasOneCartTx) Count() int64 {
	return a.tx.Count()
}

type driverHasManyDevices struct {
	db *gorm.DB

	field.RelationField

	Products struct {
		field.RelationField
	}
}

func (a driverHasManyDevices) Where(conds ...field.Expr) *driverHasManyDevices {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a driverHasManyDevices) WithContext(ctx context.Context) *driverHasManyDevices {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a driverHasManyDevices) Model(m *model.Driver) *driverHasManyDevicesTx {
	return &driverHasManyDevicesTx{a.db.Model(m).Association(a.Name())}
}

type driverHasManyDevicesTx struct{ tx *gorm.Association }

func (a driverHasManyDevicesTx) Find() (result []*model.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a driverHasManyDevicesTx) Append(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a driverHasManyDevicesTx) Replace(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a driverHasManyDevicesTx) Delete(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a driverHasManyDevicesTx) Clear() error {
	return a.tx.Clear()
}

func (a driverHasManyDevicesTx) Count() int64 {
	return a.tx.Count()
}

type driverHasManyDriverOrderForms struct {
	db *gorm.DB

	field.RelationField

	ProductInfo struct {
		field.RelationField
	}
}

func (a driverHasManyDriverOrderForms) Where(conds ...field.Expr) *driverHasManyDriverOrderForms {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a driverHasManyDriverOrderForms) WithContext(ctx context.Context) *driverHasManyDriverOrderForms {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a driverHasManyDriverOrderForms) Model(m *model.Driver) *driverHasManyDriverOrderFormsTx {
	return &driverHasManyDriverOrderFormsTx{a.db.Model(m).Association(a.Name())}
}

type driverHasManyDriverOrderFormsTx struct{ tx *gorm.Association }

func (a driverHasManyDriverOrderFormsTx) Find() (result []*model.DriverOrderForm, err error) {
	return result, a.tx.Find(&result)
}

func (a driverHasManyDriverOrderFormsTx) Append(values ...*model.DriverOrderForm) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a driverHasManyDriverOrderFormsTx) Replace(values ...*model.DriverOrderForm) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a driverHasManyDriverOrderFormsTx) Delete(values ...*model.DriverOrderForm) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a driverHasManyDriverOrderFormsTx) Clear() error {
	return a.tx.Clear()
}

func (a driverHasManyDriverOrderFormsTx) Count() int64 {
	return a.tx.Count()
}

type driverDo struct{ gen.DO }

type IDriverDo interface {
	gen.SubQuery
	Debug() IDriverDo
	WithContext(ctx context.Context) IDriverDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDriverDo
	WriteDB() IDriverDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDriverDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDriverDo
	Not(conds ...gen.Condition) IDriverDo
	Or(conds ...gen.Condition) IDriverDo
	Select(conds ...field.Expr) IDriverDo
	Where(conds ...gen.Condition) IDriverDo
	Order(conds ...field.Expr) IDriverDo
	Distinct(cols ...field.Expr) IDriverDo
	Omit(cols ...field.Expr) IDriverDo
	Join(table schema.Tabler, on ...field.Expr) IDriverDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDriverDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDriverDo
	Group(cols ...field.Expr) IDriverDo
	Having(conds ...gen.Condition) IDriverDo
	Limit(limit int) IDriverDo
	Offset(offset int) IDriverDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDriverDo
	Unscoped() IDriverDo
	Create(values ...*model.Driver) error
	CreateInBatches(values []*model.Driver, batchSize int) error
	Save(values ...*model.Driver) error
	First() (*model.Driver, error)
	Take() (*model.Driver, error)
	Last() (*model.Driver, error)
	Find() ([]*model.Driver, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Driver, err error)
	FindInBatches(result *[]*model.Driver, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Driver) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDriverDo
	Assign(attrs ...field.AssignExpr) IDriverDo
	Joins(fields ...field.RelationField) IDriverDo
	Preload(fields ...field.RelationField) IDriverDo
	FirstOrInit() (*model.Driver, error)
	FirstOrCreate() (*model.Driver, error)
	FindByPage(offset int, limit int) (result []*model.Driver, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDriverDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id int64) (result model.Driver, err error)
	GetByName(id int64) (result model.Driver, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (d driverDo) GetByID(id int64) (result model.Driver, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM drivers WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE name=@id
func (d driverDo) GetByName(id int64) (result model.Driver, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM drivers WHERE name=? ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (d driverDo) Debug() IDriverDo {
	return d.withDO(d.DO.Debug())
}

func (d driverDo) WithContext(ctx context.Context) IDriverDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d driverDo) ReadDB() IDriverDo {
	return d.Clauses(dbresolver.Read)
}

func (d driverDo) WriteDB() IDriverDo {
	return d.Clauses(dbresolver.Write)
}

func (d driverDo) Session(config *gorm.Session) IDriverDo {
	return d.withDO(d.DO.Session(config))
}

func (d driverDo) Clauses(conds ...clause.Expression) IDriverDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d driverDo) Returning(value interface{}, columns ...string) IDriverDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d driverDo) Not(conds ...gen.Condition) IDriverDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d driverDo) Or(conds ...gen.Condition) IDriverDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d driverDo) Select(conds ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d driverDo) Where(conds ...gen.Condition) IDriverDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d driverDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDriverDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d driverDo) Order(conds ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d driverDo) Distinct(cols ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d driverDo) Omit(cols ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d driverDo) Join(table schema.Tabler, on ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d driverDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDriverDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d driverDo) RightJoin(table schema.Tabler, on ...field.Expr) IDriverDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d driverDo) Group(cols ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d driverDo) Having(conds ...gen.Condition) IDriverDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d driverDo) Limit(limit int) IDriverDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d driverDo) Offset(offset int) IDriverDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d driverDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDriverDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d driverDo) Unscoped() IDriverDo {
	return d.withDO(d.DO.Unscoped())
}

func (d driverDo) Create(values ...*model.Driver) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d driverDo) CreateInBatches(values []*model.Driver, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d driverDo) Save(values ...*model.Driver) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d driverDo) First() (*model.Driver, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Driver), nil
	}
}

func (d driverDo) Take() (*model.Driver, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Driver), nil
	}
}

func (d driverDo) Last() (*model.Driver, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Driver), nil
	}
}

func (d driverDo) Find() ([]*model.Driver, error) {
	result, err := d.DO.Find()
	return result.([]*model.Driver), err
}

func (d driverDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Driver, err error) {
	buf := make([]*model.Driver, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d driverDo) FindInBatches(result *[]*model.Driver, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d driverDo) Attrs(attrs ...field.AssignExpr) IDriverDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d driverDo) Assign(attrs ...field.AssignExpr) IDriverDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d driverDo) Joins(fields ...field.RelationField) IDriverDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d driverDo) Preload(fields ...field.RelationField) IDriverDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d driverDo) FirstOrInit() (*model.Driver, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Driver), nil
	}
}

func (d driverDo) FirstOrCreate() (*model.Driver, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Driver), nil
	}
}

func (d driverDo) FindByPage(offset int, limit int) (result []*model.Driver, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d driverDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d driverDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d driverDo) Delete(models ...*model.Driver) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *driverDo) withDO(do gen.Dao) *driverDo {
	d.DO = *do.(*gen.DO)
	return d
}
