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

func newPassengerOrderForm(db *gorm.DB, opts ...gen.DOOption) passengerOrderForm {
	_passengerOrderForm := passengerOrderForm{}

	_passengerOrderForm.passengerOrderFormDo.UseDB(db, opts...)
	_passengerOrderForm.passengerOrderFormDo.UseModel(&model.PassengerOrderForm{})

	tableName := _passengerOrderForm.passengerOrderFormDo.TableName()
	_passengerOrderForm.ALL = field.NewAsterisk(tableName)
	_passengerOrderForm.PassengerID = field.NewInt64(tableName, "passenger_id")
	_passengerOrderForm.DriverCarID = field.NewString(tableName, "driver_car_id")
	_passengerOrderForm.OrderID = field.NewString(tableName, "order_id")
	_passengerOrderForm.Cost = field.NewInt64(tableName, "cost")
	_passengerOrderForm.State = field.NewInt(tableName, "state")
	_passengerOrderForm.Location = field.NewString(tableName, "location")
	_passengerOrderForm.Placetime = field.NewTime(tableName, "placetime")
	_passengerOrderForm.Paytime = field.NewTime(tableName, "paytime")
	_passengerOrderForm.ProductInfo = passengerOrderFormHasManyProductInfo{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ProductInfo", "model.OrderProduct"),
	}

	_passengerOrderForm.fillFieldMap()

	return _passengerOrderForm
}

type passengerOrderForm struct {
	passengerOrderFormDo

	ALL         field.Asterisk
	PassengerID field.Int64
	DriverCarID field.String
	OrderID     field.String
	Cost        field.Int64
	State       field.Int
	Location    field.String
	Placetime   field.Time
	Paytime     field.Time
	ProductInfo passengerOrderFormHasManyProductInfo

	fieldMap map[string]field.Expr
}

func (p passengerOrderForm) Table(newTableName string) *passengerOrderForm {
	p.passengerOrderFormDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p passengerOrderForm) As(alias string) *passengerOrderForm {
	p.passengerOrderFormDo.DO = *(p.passengerOrderFormDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *passengerOrderForm) updateTableName(table string) *passengerOrderForm {
	p.ALL = field.NewAsterisk(table)
	p.PassengerID = field.NewInt64(table, "passenger_id")
	p.DriverCarID = field.NewString(table, "driver_car_id")
	p.OrderID = field.NewString(table, "order_id")
	p.Cost = field.NewInt64(table, "cost")
	p.State = field.NewInt(table, "state")
	p.Location = field.NewString(table, "location")
	p.Placetime = field.NewTime(table, "placetime")
	p.Paytime = field.NewTime(table, "paytime")

	p.fillFieldMap()

	return p
}

func (p *passengerOrderForm) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *passengerOrderForm) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["passenger_id"] = p.PassengerID
	p.fieldMap["driver_car_id"] = p.DriverCarID
	p.fieldMap["order_id"] = p.OrderID
	p.fieldMap["cost"] = p.Cost
	p.fieldMap["state"] = p.State
	p.fieldMap["location"] = p.Location
	p.fieldMap["placetime"] = p.Placetime
	p.fieldMap["paytime"] = p.Paytime

}

func (p passengerOrderForm) clone(db *gorm.DB) passengerOrderForm {
	p.passengerOrderFormDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p passengerOrderForm) replaceDB(db *gorm.DB) passengerOrderForm {
	p.passengerOrderFormDo.ReplaceDB(db)
	return p
}

type passengerOrderFormHasManyProductInfo struct {
	db *gorm.DB

	field.RelationField
}

func (a passengerOrderFormHasManyProductInfo) Where(conds ...field.Expr) *passengerOrderFormHasManyProductInfo {
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

func (a passengerOrderFormHasManyProductInfo) WithContext(ctx context.Context) *passengerOrderFormHasManyProductInfo {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a passengerOrderFormHasManyProductInfo) Model(m *model.PassengerOrderForm) *passengerOrderFormHasManyProductInfoTx {
	return &passengerOrderFormHasManyProductInfoTx{a.db.Model(m).Association(a.Name())}
}

type passengerOrderFormHasManyProductInfoTx struct{ tx *gorm.Association }

func (a passengerOrderFormHasManyProductInfoTx) Find() (result []*model.OrderProduct, err error) {
	return result, a.tx.Find(&result)
}

func (a passengerOrderFormHasManyProductInfoTx) Append(values ...*model.OrderProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a passengerOrderFormHasManyProductInfoTx) Replace(values ...*model.OrderProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a passengerOrderFormHasManyProductInfoTx) Delete(values ...*model.OrderProduct) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a passengerOrderFormHasManyProductInfoTx) Clear() error {
	return a.tx.Clear()
}

func (a passengerOrderFormHasManyProductInfoTx) Count() int64 {
	return a.tx.Count()
}

type passengerOrderFormDo struct{ gen.DO }

type IPassengerOrderFormDo interface {
	gen.SubQuery
	Debug() IPassengerOrderFormDo
	WithContext(ctx context.Context) IPassengerOrderFormDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPassengerOrderFormDo
	WriteDB() IPassengerOrderFormDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPassengerOrderFormDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPassengerOrderFormDo
	Not(conds ...gen.Condition) IPassengerOrderFormDo
	Or(conds ...gen.Condition) IPassengerOrderFormDo
	Select(conds ...field.Expr) IPassengerOrderFormDo
	Where(conds ...gen.Condition) IPassengerOrderFormDo
	Order(conds ...field.Expr) IPassengerOrderFormDo
	Distinct(cols ...field.Expr) IPassengerOrderFormDo
	Omit(cols ...field.Expr) IPassengerOrderFormDo
	Join(table schema.Tabler, on ...field.Expr) IPassengerOrderFormDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPassengerOrderFormDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPassengerOrderFormDo
	Group(cols ...field.Expr) IPassengerOrderFormDo
	Having(conds ...gen.Condition) IPassengerOrderFormDo
	Limit(limit int) IPassengerOrderFormDo
	Offset(offset int) IPassengerOrderFormDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPassengerOrderFormDo
	Unscoped() IPassengerOrderFormDo
	Create(values ...*model.PassengerOrderForm) error
	CreateInBatches(values []*model.PassengerOrderForm, batchSize int) error
	Save(values ...*model.PassengerOrderForm) error
	First() (*model.PassengerOrderForm, error)
	Take() (*model.PassengerOrderForm, error)
	Last() (*model.PassengerOrderForm, error)
	Find() ([]*model.PassengerOrderForm, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PassengerOrderForm, err error)
	FindInBatches(result *[]*model.PassengerOrderForm, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PassengerOrderForm) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPassengerOrderFormDo
	Assign(attrs ...field.AssignExpr) IPassengerOrderFormDo
	Joins(fields ...field.RelationField) IPassengerOrderFormDo
	Preload(fields ...field.RelationField) IPassengerOrderFormDo
	FirstOrInit() (*model.PassengerOrderForm, error)
	FirstOrCreate() (*model.PassengerOrderForm, error)
	FindByPage(offset int, limit int) (result []*model.PassengerOrderForm, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPassengerOrderFormDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FGetAllOrderFormsFrom(id int64) (result []model.PassengerOrderForm, err error)
	FGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.PassengerOrderForm, err error)
	DGetAllOrderFormsFrom(id int64) (result []model.PassengerOrderForm, err error)
	DGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.PassengerOrderForm, err error)
	PGetAllOrderFormsFrom(id int64) (result []model.PassengerOrderForm, err error)
	PGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.PassengerOrderForm, err error)
}

// sql(SELECT * FROM @@table where @id =(SELECT factory_id from driver_order_forms where factory_id=@id))
func (p passengerOrderFormDo) FGetAllOrderFormsFrom(id int64) (result []model.PassengerOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM passenger_order_forms where ? =(SELECT factory_id from driver_order_forms where factory_id=?) ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(SELECT * FROM @@table where state=@mode and @id =(SELECT factory_id from driver_order_forms where factory_id=@id))
func (p passengerOrderFormDo) FGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.PassengerOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, mode)
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM passenger_order_forms where state=? and ? =(SELECT factory_id from driver_order_forms where factory_id=?) ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(SELECT * FROM @@table where @id =(SELECT id from drivers where id=@id))
func (p passengerOrderFormDo) DGetAllOrderFormsFrom(id int64) (result []model.PassengerOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM passenger_order_forms where ? =(SELECT id from drivers where id=?) ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(SELECT * FROM @@table where state=@mode and @id =(SELECT id from driver_order_forms where factory_id=@id))
func (p passengerOrderFormDo) DGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.PassengerOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, mode)
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM passenger_order_forms where state=? and ? =(SELECT id from driver_order_forms where factory_id=?) ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(SELECT * FROM @@table where @id =(SELECT id from passengers where id=@id))
func (p passengerOrderFormDo) PGetAllOrderFormsFrom(id int64) (result []model.PassengerOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM passenger_order_forms where ? =(SELECT id from passengers where id=?) ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(SELECT * FROM @@table where statestate=@mode and @id =(SELECT factory_id from driver_order_forms where factory_id=@id))
func (p passengerOrderFormDo) PGetByStateOrderForms(id int64, mode model.ORDERSTATE) (result []model.PassengerOrderForm, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, mode)
	params = append(params, id)
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM passenger_order_forms where statestate=? and ? =(SELECT factory_id from driver_order_forms where factory_id=?) ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p passengerOrderFormDo) Debug() IPassengerOrderFormDo {
	return p.withDO(p.DO.Debug())
}

func (p passengerOrderFormDo) WithContext(ctx context.Context) IPassengerOrderFormDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p passengerOrderFormDo) ReadDB() IPassengerOrderFormDo {
	return p.Clauses(dbresolver.Read)
}

func (p passengerOrderFormDo) WriteDB() IPassengerOrderFormDo {
	return p.Clauses(dbresolver.Write)
}

func (p passengerOrderFormDo) Session(config *gorm.Session) IPassengerOrderFormDo {
	return p.withDO(p.DO.Session(config))
}

func (p passengerOrderFormDo) Clauses(conds ...clause.Expression) IPassengerOrderFormDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p passengerOrderFormDo) Returning(value interface{}, columns ...string) IPassengerOrderFormDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p passengerOrderFormDo) Not(conds ...gen.Condition) IPassengerOrderFormDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p passengerOrderFormDo) Or(conds ...gen.Condition) IPassengerOrderFormDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p passengerOrderFormDo) Select(conds ...field.Expr) IPassengerOrderFormDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p passengerOrderFormDo) Where(conds ...gen.Condition) IPassengerOrderFormDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p passengerOrderFormDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPassengerOrderFormDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p passengerOrderFormDo) Order(conds ...field.Expr) IPassengerOrderFormDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p passengerOrderFormDo) Distinct(cols ...field.Expr) IPassengerOrderFormDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p passengerOrderFormDo) Omit(cols ...field.Expr) IPassengerOrderFormDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p passengerOrderFormDo) Join(table schema.Tabler, on ...field.Expr) IPassengerOrderFormDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p passengerOrderFormDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPassengerOrderFormDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p passengerOrderFormDo) RightJoin(table schema.Tabler, on ...field.Expr) IPassengerOrderFormDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p passengerOrderFormDo) Group(cols ...field.Expr) IPassengerOrderFormDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p passengerOrderFormDo) Having(conds ...gen.Condition) IPassengerOrderFormDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p passengerOrderFormDo) Limit(limit int) IPassengerOrderFormDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p passengerOrderFormDo) Offset(offset int) IPassengerOrderFormDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p passengerOrderFormDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPassengerOrderFormDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p passengerOrderFormDo) Unscoped() IPassengerOrderFormDo {
	return p.withDO(p.DO.Unscoped())
}

func (p passengerOrderFormDo) Create(values ...*model.PassengerOrderForm) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p passengerOrderFormDo) CreateInBatches(values []*model.PassengerOrderForm, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p passengerOrderFormDo) Save(values ...*model.PassengerOrderForm) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p passengerOrderFormDo) First() (*model.PassengerOrderForm, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassengerOrderForm), nil
	}
}

func (p passengerOrderFormDo) Take() (*model.PassengerOrderForm, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassengerOrderForm), nil
	}
}

func (p passengerOrderFormDo) Last() (*model.PassengerOrderForm, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassengerOrderForm), nil
	}
}

func (p passengerOrderFormDo) Find() ([]*model.PassengerOrderForm, error) {
	result, err := p.DO.Find()
	return result.([]*model.PassengerOrderForm), err
}

func (p passengerOrderFormDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PassengerOrderForm, err error) {
	buf := make([]*model.PassengerOrderForm, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p passengerOrderFormDo) FindInBatches(result *[]*model.PassengerOrderForm, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p passengerOrderFormDo) Attrs(attrs ...field.AssignExpr) IPassengerOrderFormDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p passengerOrderFormDo) Assign(attrs ...field.AssignExpr) IPassengerOrderFormDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p passengerOrderFormDo) Joins(fields ...field.RelationField) IPassengerOrderFormDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p passengerOrderFormDo) Preload(fields ...field.RelationField) IPassengerOrderFormDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p passengerOrderFormDo) FirstOrInit() (*model.PassengerOrderForm, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassengerOrderForm), nil
	}
}

func (p passengerOrderFormDo) FirstOrCreate() (*model.PassengerOrderForm, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PassengerOrderForm), nil
	}
}

func (p passengerOrderFormDo) FindByPage(offset int, limit int) (result []*model.PassengerOrderForm, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p passengerOrderFormDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p passengerOrderFormDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p passengerOrderFormDo) Delete(models ...*model.PassengerOrderForm) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *passengerOrderFormDo) withDO(do gen.Dao) *passengerOrderFormDo {
	p.DO = *do.(*gen.DO)
	return p
}
