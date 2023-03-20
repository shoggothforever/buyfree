// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"buyfree/repo/model"
)

func newDriverOrderForm(db *gorm.DB, opts ...gen.DOOption) driverOrderForm {
	_driverOrderForm := driverOrderForm{}

	_driverOrderForm.driverOrderFormDo.UseDB(db, opts...)
	_driverOrderForm.driverOrderFormDo.UseModel(&model.DriverOrderForm{})

	tableName := _driverOrderForm.driverOrderFormDo.TableName()
	_driverOrderForm.ALL = field.NewAsterisk(tableName)
	_driverOrderForm.FactoryID = field.NewInt64(tableName, "factory_id")
	_driverOrderForm.DriverID = field.NewInt64(tableName, "driver_id")
	_driverOrderForm.CarID = field.NewString(tableName, "car_id")
	_driverOrderForm.Comment = field.NewString(tableName, "comment")
	_driverOrderForm.GetTime = field.NewTime(tableName, "get_time")
	_driverOrderForm.OrderID = field.NewString(tableName, "order_id")
	_driverOrderForm.Cost = field.NewInt64(tableName, "cost")
	_driverOrderForm.State = field.NewInt(tableName, "state")
	_driverOrderForm.Location = field.NewString(tableName, "location")
	_driverOrderForm.Placetime = field.NewTime(tableName, "placetime")
	_driverOrderForm.Paytime = field.NewTime(tableName, "paytime")
	_driverOrderForm.ProductInfo = driverOrderFormHasManyProductInfo{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ProductInfo", "model.OrderProduct"),
	}

	_driverOrderForm.fillFieldMap()

	return _driverOrderForm
}

type driverOrderForm struct {
	driverOrderFormDo

	ALL         field.Asterisk
	FactoryID   field.Int64
	DriverID    field.Int64
	CarID       field.String
	Comment     field.String
	GetTime     field.Time
	OrderID     field.String
	Cost        field.Int64
	State       field.Int
	Location    field.String
	Placetime   field.Time
	Paytime     field.Time
	ProductInfo driverOrderFormHasManyProductInfo

	fieldMap map[string]field.Expr
}

func (d driverOrderForm) Table(newTableName string) *driverOrderForm {
	d.driverOrderFormDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d driverOrderForm) As(alias string) *driverOrderForm {
	d.driverOrderFormDo.DO = *(d.driverOrderFormDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *driverOrderForm) updateTableName(table string) *driverOrderForm {
	d.ALL = field.NewAsterisk(table)
	d.FactoryID = field.NewInt64(table, "factory_id")
	d.DriverID = field.NewInt64(table, "driver_id")
	d.CarID = field.NewString(table, "car_id")
	d.Comment = field.NewString(table, "comment")
	d.GetTime = field.NewTime(table, "get_time")
	d.OrderID = field.NewString(table, "order_id")
	d.Cost = field.NewInt64(table, "cost")
	d.State = field.NewInt(table, "state")
	d.Location = field.NewString(table, "location")
	d.Placetime = field.NewTime(table, "placetime")
	d.Paytime = field.NewTime(table, "paytime")

	d.fillFieldMap()

	return d
}

func (d *driverOrderForm) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *driverOrderForm) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 12)
	d.fieldMap["factory_id"] = d.FactoryID
	d.fieldMap["driver_id"] = d.DriverID
	d.fieldMap["car_id"] = d.CarID
	d.fieldMap["comment"] = d.Comment
	d.fieldMap["get_time"] = d.GetTime
	d.fieldMap["order_id"] = d.OrderID
	d.fieldMap["cost"] = d.Cost
	d.fieldMap["state"] = d.State
	d.fieldMap["location"] = d.Location
	d.fieldMap["placetime"] = d.Placetime
	d.fieldMap["paytime"] = d.Paytime

}

func (d driverOrderForm) clone(db *gorm.DB) driverOrderForm {
	d.driverOrderFormDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d driverOrderForm) replaceDB(db *gorm.DB) driverOrderForm {
	d.driverOrderFormDo.ReplaceDB(db)
	return d
}

type driverOrderFormHasManyProductInfo struct {
	db *gorm.DB

	field.RelationField
}

func (a driverOrderFormHasManyProductInfo) Where(conds ...field.Expr) *driverOrderFormHasManyProductInfo {
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

func (a driverOrderFormHasManyProductInfo) WithContext(ctx context.Context) *driverOrderFormHasManyProductInfo {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a driverOrderFormHasManyProductInfo) Model(m *model.DriverOrderForm) *driverOrderFormHasManyProductInfoTx {
	return &driverOrderFormHasManyProductInfoTx{a.db.Model(m).Association(a.Name())}
}

type driverOrderFormHasManyProductInfoTx struct{ tx *gorm.Association }

func (a driverOrderFormHasManyProductInfoTx) Find() (result []*model.OrderProduct, err error) {
	return result, a.tx.Find(&result)
}

func (a driverOrderFormHasManyProductInfoTx) Append(values ...*model.OrderProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a driverOrderFormHasManyProductInfoTx) Replace(values ...*model.OrderProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a driverOrderFormHasManyProductInfoTx) Delete(values ...*model.OrderProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a driverOrderFormHasManyProductInfoTx) Clear() error {
	return a.tx.Clear()
}

func (a driverOrderFormHasManyProductInfoTx) Count() int64 {
	return a.tx.Count()
}

type driverOrderFormDo struct{ gen.DO }

type IDriverOrderFormDo interface {
	gen.SubQuery
	Debug() IDriverOrderFormDo
	WithContext(ctx context.Context) IDriverOrderFormDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDriverOrderFormDo
	WriteDB() IDriverOrderFormDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDriverOrderFormDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDriverOrderFormDo
	Not(conds ...gen.Condition) IDriverOrderFormDo
	Or(conds ...gen.Condition) IDriverOrderFormDo
	Select(conds ...field.Expr) IDriverOrderFormDo
	Where(conds ...gen.Condition) IDriverOrderFormDo
	Order(conds ...field.Expr) IDriverOrderFormDo
	Distinct(cols ...field.Expr) IDriverOrderFormDo
	Omit(cols ...field.Expr) IDriverOrderFormDo
	Join(table schema.Tabler, on ...field.Expr) IDriverOrderFormDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDriverOrderFormDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDriverOrderFormDo
	Group(cols ...field.Expr) IDriverOrderFormDo
	Having(conds ...gen.Condition) IDriverOrderFormDo
	Limit(limit int) IDriverOrderFormDo
	Offset(offset int) IDriverOrderFormDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDriverOrderFormDo
	Unscoped() IDriverOrderFormDo
	Create(values ...*model.DriverOrderForm) error
	CreateInBatches(values []*model.DriverOrderForm, batchSize int) error
	Save(values ...*model.DriverOrderForm) error
	First() (*model.DriverOrderForm, error)
	Take() (*model.DriverOrderForm, error)
	Last() (*model.DriverOrderForm, error)
	Find() ([]*model.DriverOrderForm, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DriverOrderForm, err error)
	FindInBatches(result *[]*model.DriverOrderForm, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DriverOrderForm) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDriverOrderFormDo
	Assign(attrs ...field.AssignExpr) IDriverOrderFormDo
	Joins(fields ...field.RelationField) IDriverOrderFormDo
	Preload(fields ...field.RelationField) IDriverOrderFormDo
	FirstOrInit() (*model.DriverOrderForm, error)
	FirstOrCreate() (*model.DriverOrderForm, error)
	FindByPage(offset int, limit int) (result []*model.DriverOrderForm, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDriverOrderFormDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FGetAllOrderFormsFrom(id int64) (result []model.DriverOrderForm, err error)
	FGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.DriverOrderForm, err error)
	DGetAllOrderFormsFrom(id int64) (result []model.DriverOrderForm, err error)
	DGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.DriverOrderForm, err error)
	PGetAllOrderFormsFrom(id int64) (result []model.DriverOrderForm, err error)
	PGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.DriverOrderForm, err error)
}

// sql(SELECT * FROM @@table where @id =(SELECT factory_id from driver_order_forms where factory_id=@id))
func (d driverOrderFormDo) FGetAllOrderFormsFrom(id int64) (result []model.DriverOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM driver_order_forms where ? =(SELECT factory_id from driver_order_forms where factory_id=?) ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(SELECT * FROM @@table where state=@mode and @id =(SELECT factory_id from driver_order_forms where factory_id=@id))
func (d driverOrderFormDo) FGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.DriverOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, mode)
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM driver_order_forms where state=? and ? =(SELECT factory_id from driver_order_forms where factory_id=?) ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(SELECT * FROM @@table where @id =(SELECT id from drivers where id=@id))
func (d driverOrderFormDo) DGetAllOrderFormsFrom(id int64) (result []model.DriverOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM driver_order_forms where ? =(SELECT id from drivers where id=?) ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(SELECT * FROM @@table where state=@mode and @id =(SELECT id from driver_order_forms where factory_id=@id))
func (d driverOrderFormDo) DGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.DriverOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, mode)
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM driver_order_forms where state=? and ? =(SELECT id from driver_order_forms where factory_id=?) ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(SELECT * FROM @@table where @id =(SELECT id from passengers where id=@id))
func (d driverOrderFormDo) PGetAllOrderFormsFrom(id int64) (result []model.DriverOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM driver_order_forms where ? =(SELECT id from passengers where id=?) ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(SELECT * FROM @@table where statestate=@mode and @id =(SELECT factory_id from driver_order_forms where factory_id=@id))
func (d driverOrderFormDo) PGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.DriverOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, mode)
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM driver_order_forms where statestate=? and ? =(SELECT factory_id from driver_order_forms where factory_id=?) ")

	var executeSQL *gorm.DB
	executeSQL = d.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (d driverOrderFormDo) Debug() IDriverOrderFormDo {
	return d.withDO(d.DO.Debug())
}

func (d driverOrderFormDo) WithContext(ctx context.Context) IDriverOrderFormDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d driverOrderFormDo) ReadDB() IDriverOrderFormDo {
	return d.Clauses(dbresolver.Read)
}

func (d driverOrderFormDo) WriteDB() IDriverOrderFormDo {
	return d.Clauses(dbresolver.Write)
}

func (d driverOrderFormDo) Session(config *gorm.Session) IDriverOrderFormDo {
	return d.withDO(d.DO.Session(config))
}

func (d driverOrderFormDo) Clauses(conds ...clause.Expression) IDriverOrderFormDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d driverOrderFormDo) Returning(value interface{}, columns ...string) IDriverOrderFormDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d driverOrderFormDo) Not(conds ...gen.Condition) IDriverOrderFormDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d driverOrderFormDo) Or(conds ...gen.Condition) IDriverOrderFormDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d driverOrderFormDo) Select(conds ...field.Expr) IDriverOrderFormDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d driverOrderFormDo) Where(conds ...gen.Condition) IDriverOrderFormDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d driverOrderFormDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDriverOrderFormDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d driverOrderFormDo) Order(conds ...field.Expr) IDriverOrderFormDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d driverOrderFormDo) Distinct(cols ...field.Expr) IDriverOrderFormDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d driverOrderFormDo) Omit(cols ...field.Expr) IDriverOrderFormDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d driverOrderFormDo) Join(table schema.Tabler, on ...field.Expr) IDriverOrderFormDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d driverOrderFormDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDriverOrderFormDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d driverOrderFormDo) RightJoin(table schema.Tabler, on ...field.Expr) IDriverOrderFormDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d driverOrderFormDo) Group(cols ...field.Expr) IDriverOrderFormDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d driverOrderFormDo) Having(conds ...gen.Condition) IDriverOrderFormDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d driverOrderFormDo) Limit(limit int) IDriverOrderFormDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d driverOrderFormDo) Offset(offset int) IDriverOrderFormDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d driverOrderFormDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDriverOrderFormDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d driverOrderFormDo) Unscoped() IDriverOrderFormDo {
	return d.withDO(d.DO.Unscoped())
}

func (d driverOrderFormDo) Create(values ...*model.DriverOrderForm) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d driverOrderFormDo) CreateInBatches(values []*model.DriverOrderForm, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d driverOrderFormDo) Save(values ...*model.DriverOrderForm) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d driverOrderFormDo) First() (*model.DriverOrderForm, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DriverOrderForm), nil
	}
}

func (d driverOrderFormDo) Take() (*model.DriverOrderForm, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DriverOrderForm), nil
	}
}

func (d driverOrderFormDo) Last() (*model.DriverOrderForm, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DriverOrderForm), nil
	}
}

func (d driverOrderFormDo) Find() ([]*model.DriverOrderForm, error) {
	result, err := d.DO.Find()
	return result.([]*model.DriverOrderForm), err
}

func (d driverOrderFormDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DriverOrderForm, err error) {
	buf := make([]*model.DriverOrderForm, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d driverOrderFormDo) FindInBatches(result *[]*model.DriverOrderForm, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d driverOrderFormDo) Attrs(attrs ...field.AssignExpr) IDriverOrderFormDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d driverOrderFormDo) Assign(attrs ...field.AssignExpr) IDriverOrderFormDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d driverOrderFormDo) Joins(fields ...field.RelationField) IDriverOrderFormDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d driverOrderFormDo) Preload(fields ...field.RelationField) IDriverOrderFormDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d driverOrderFormDo) FirstOrInit() (*model.DriverOrderForm, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DriverOrderForm), nil
	}
}

func (d driverOrderFormDo) FirstOrCreate() (*model.DriverOrderForm, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DriverOrderForm), nil
	}
}

func (d driverOrderFormDo) FindByPage(offset int, limit int) (result []*model.DriverOrderForm, count int64, err error) {
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

func (d driverOrderFormDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d driverOrderFormDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d driverOrderFormDo) Delete(models ...*model.DriverOrderForm) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *driverOrderFormDo) withDO(do gen.Dao) *driverOrderFormDo {
	d.DO = *do.(*gen.DO)
	return d
}
