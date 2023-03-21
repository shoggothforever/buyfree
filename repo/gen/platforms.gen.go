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

func newPlatform(db *gorm.DB, opts ...gen.DOOption) platform {
	_platform := platform{}

	_platform.platformDo.UseDB(db, opts...)
	_platform.platformDo.UseModel(&model.Platform{})

	tableName := _platform.platformDo.TableName()
	_platform.ALL = field.NewAsterisk(tableName)
	_platform.ID = field.NewInt64(tableName, "id")
	_platform.CreatedAt = field.NewTime(tableName, "created_at")
	_platform.UpdatedAt = field.NewTime(tableName, "updated_at")
	_platform.DeletedAt = field.NewField(tableName, "deleted_at")
	_platform.Balance = field.NewFloat64(tableName, "balance")
	_platform.Pic = field.NewString(tableName, "pic")
	_platform.Name = field.NewString(tableName, "name")
	_platform.Password = field.NewString(tableName, "password")
	_platform.Mobile = field.NewString(tableName, "mobile")
	_platform.IDCard = field.NewString(tableName, "id_card")
	_platform.Role = field.NewInt(tableName, "role")
	_platform.Level = field.NewInt(tableName, "level")
	_platform.PasswordSalt = field.NewString(tableName, "password_salt")
	_platform.AuthorizedDrivers = platformHasManyAuthorizedDrivers{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("AuthorizedDrivers", "model.Driver"),
		Cart: struct {
			field.RelationField
			Products struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("AuthorizedDrivers.Cart", "model.DriverCart"),
			Products: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("AuthorizedDrivers.Cart.Products", "model.OrderProduct"),
			},
		},
		Devices: struct {
			field.RelationField
			Products struct {
				field.RelationField
			}
			Advertisements struct {
				field.RelationField
				Devices struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("AuthorizedDrivers.Devices", "model.Device"),
			Products: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("AuthorizedDrivers.Devices.Products", "model.DeviceProduct"),
			},
			Advertisements: struct {
				field.RelationField
				Devices struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("AuthorizedDrivers.Devices.Advertisements", "model.Advertisement"),
				Devices: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("AuthorizedDrivers.Devices.Advertisements.Devices", "model.Device"),
				},
			},
		},
		DriverOrderForms: struct {
			field.RelationField
			ProductInfo struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("AuthorizedDrivers.DriverOrderForms", "model.DriverOrderForm"),
			ProductInfo: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("AuthorizedDrivers.DriverOrderForms.ProductInfo", "model.OrderProduct"),
			},
		},
	}

	_platform.Devices = platformHasManyDevices{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Devices", "model.Device"),
	}

	_platform.Advertisements = platformHasManyAdvertisements{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Advertisements", "model.Advertisement"),
	}

	_platform.fillFieldMap()

	return _platform
}

type platform struct {
	platformDo

	ALL               field.Asterisk
	ID                field.Int64
	CreatedAt         field.Time
	UpdatedAt         field.Time
	DeletedAt         field.Field
	Balance           field.Float64
	Pic               field.String
	Name              field.String
	Password          field.String
	Mobile            field.String
	IDCard            field.String
	Role              field.Int
	Level             field.Int
	PasswordSalt      field.String
	AuthorizedDrivers platformHasManyAuthorizedDrivers

	Devices platformHasManyDevices

	Advertisements platformHasManyAdvertisements

	fieldMap map[string]field.Expr
}

func (p platform) Table(newTableName string) *platform {
	p.platformDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p platform) As(alias string) *platform {
	p.platformDo.DO = *(p.platformDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *platform) updateTableName(table string) *platform {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Balance = field.NewFloat64(table, "balance")
	p.Pic = field.NewString(table, "pic")
	p.Name = field.NewString(table, "name")
	p.Password = field.NewString(table, "password")
	p.Mobile = field.NewString(table, "mobile")
	p.IDCard = field.NewString(table, "id_card")
	p.Role = field.NewInt(table, "role")
	p.Level = field.NewInt(table, "level")
	p.PasswordSalt = field.NewString(table, "password_salt")

	p.fillFieldMap()

	return p
}

func (p *platform) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *platform) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 16)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["balance"] = p.Balance
	p.fieldMap["pic"] = p.Pic
	p.fieldMap["name"] = p.Name
	p.fieldMap["password"] = p.Password
	p.fieldMap["mobile"] = p.Mobile
	p.fieldMap["id_card"] = p.IDCard
	p.fieldMap["role"] = p.Role
	p.fieldMap["level"] = p.Level
	p.fieldMap["password_salt"] = p.PasswordSalt

}

func (p platform) clone(db *gorm.DB) platform {
	p.platformDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p platform) replaceDB(db *gorm.DB) platform {
	p.platformDo.ReplaceDB(db)
	return p
}

type platformHasManyAuthorizedDrivers struct {
	db *gorm.DB

	field.RelationField

	Cart struct {
		field.RelationField
		Products struct {
			field.RelationField
		}
	}
	Devices struct {
		field.RelationField
		Products struct {
			field.RelationField
		}
		Advertisements struct {
			field.RelationField
			Devices struct {
				field.RelationField
			}
		}
	}
	DriverOrderForms struct {
		field.RelationField
		ProductInfo struct {
			field.RelationField
		}
	}
}

func (a platformHasManyAuthorizedDrivers) Where(conds ...field.Expr) *platformHasManyAuthorizedDrivers {
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

func (a platformHasManyAuthorizedDrivers) WithContext(ctx context.Context) *platformHasManyAuthorizedDrivers {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a platformHasManyAuthorizedDrivers) Model(m *model.Platform) *platformHasManyAuthorizedDriversTx {
	return &platformHasManyAuthorizedDriversTx{a.db.Model(m).Association(a.Name())}
}

type platformHasManyAuthorizedDriversTx struct{ tx *gorm.Association }

func (a platformHasManyAuthorizedDriversTx) Find() (result []*model.Driver, err error) {
	return result, a.tx.Find(&result)
}

func (a platformHasManyAuthorizedDriversTx) Append(values ...*model.Driver) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a platformHasManyAuthorizedDriversTx) Replace(values ...*model.Driver) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a platformHasManyAuthorizedDriversTx) Delete(values ...*model.Driver) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a platformHasManyAuthorizedDriversTx) Clear() error {
	return a.tx.Clear()
}

func (a platformHasManyAuthorizedDriversTx) Count() int64 {
	return a.tx.Count()
}

type platformHasManyDevices struct {
	db *gorm.DB

	field.RelationField
}

func (a platformHasManyDevices) Where(conds ...field.Expr) *platformHasManyDevices {
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

func (a platformHasManyDevices) WithContext(ctx context.Context) *platformHasManyDevices {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a platformHasManyDevices) Model(m *model.Platform) *platformHasManyDevicesTx {
	return &platformHasManyDevicesTx{a.db.Model(m).Association(a.Name())}
}

type platformHasManyDevicesTx struct{ tx *gorm.Association }

func (a platformHasManyDevicesTx) Find() (result []*model.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a platformHasManyDevicesTx) Append(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a platformHasManyDevicesTx) Replace(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a platformHasManyDevicesTx) Delete(values ...*model.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a platformHasManyDevicesTx) Clear() error {
	return a.tx.Clear()
}

func (a platformHasManyDevicesTx) Count() int64 {
	return a.tx.Count()
}

type platformHasManyAdvertisements struct {
	db *gorm.DB

	field.RelationField
}

func (a platformHasManyAdvertisements) Where(conds ...field.Expr) *platformHasManyAdvertisements {
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

func (a platformHasManyAdvertisements) WithContext(ctx context.Context) *platformHasManyAdvertisements {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a platformHasManyAdvertisements) Model(m *model.Platform) *platformHasManyAdvertisementsTx {
	return &platformHasManyAdvertisementsTx{a.db.Model(m).Association(a.Name())}
}

type platformHasManyAdvertisementsTx struct{ tx *gorm.Association }

func (a platformHasManyAdvertisementsTx) Find() (result []*model.Advertisement, err error) {
	return result, a.tx.Find(&result)
}

func (a platformHasManyAdvertisementsTx) Append(values ...*model.Advertisement) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a platformHasManyAdvertisementsTx) Replace(values ...*model.Advertisement) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a platformHasManyAdvertisementsTx) Delete(values ...*model.Advertisement) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a platformHasManyAdvertisementsTx) Clear() error {
	return a.tx.Clear()
}

func (a platformHasManyAdvertisementsTx) Count() int64 {
	return a.tx.Count()
}

type platformDo struct{ gen.DO }

type IPlatformDo interface {
	gen.SubQuery
	Debug() IPlatformDo
	WithContext(ctx context.Context) IPlatformDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPlatformDo
	WriteDB() IPlatformDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPlatformDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPlatformDo
	Not(conds ...gen.Condition) IPlatformDo
	Or(conds ...gen.Condition) IPlatformDo
	Select(conds ...field.Expr) IPlatformDo
	Where(conds ...gen.Condition) IPlatformDo
	Order(conds ...field.Expr) IPlatformDo
	Distinct(cols ...field.Expr) IPlatformDo
	Omit(cols ...field.Expr) IPlatformDo
	Join(table schema.Tabler, on ...field.Expr) IPlatformDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPlatformDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPlatformDo
	Group(cols ...field.Expr) IPlatformDo
	Having(conds ...gen.Condition) IPlatformDo
	Limit(limit int) IPlatformDo
	Offset(offset int) IPlatformDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPlatformDo
	Unscoped() IPlatformDo
	Create(values ...*model.Platform) error
	CreateInBatches(values []*model.Platform, batchSize int) error
	Save(values ...*model.Platform) error
	First() (*model.Platform, error)
	Take() (*model.Platform, error)
	Last() (*model.Platform, error)
	Find() ([]*model.Platform, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Platform, err error)
	FindInBatches(result *[]*model.Platform, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Platform) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPlatformDo
	Assign(attrs ...field.AssignExpr) IPlatformDo
	Joins(fields ...field.RelationField) IPlatformDo
	Preload(fields ...field.RelationField) IPlatformDo
	FirstOrInit() (*model.Platform, error)
	FirstOrCreate() (*model.Platform, error)
	FindByPage(offset int, limit int) (result []*model.Platform, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPlatformDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id int64) (result model.Platform, err error)
	GetByName(id int64) (result model.Platform, err error)
}

// SELECT * FROM @@table WHERE id=@id
func (p platformDo) GetByID(id int64) (result model.Platform, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM platforms WHERE id=? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table WHERE name=@id
func (p platformDo) GetByName(id int64) (result model.Platform, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM platforms WHERE name=? ")

	var executeSQL *gorm.DB
	executeSQL = p.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (p platformDo) Debug() IPlatformDo {
	return p.withDO(p.DO.Debug())
}

func (p platformDo) WithContext(ctx context.Context) IPlatformDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p platformDo) ReadDB() IPlatformDo {
	return p.Clauses(dbresolver.Read)
}

func (p platformDo) WriteDB() IPlatformDo {
	return p.Clauses(dbresolver.Write)
}

func (p platformDo) Session(config *gorm.Session) IPlatformDo {
	return p.withDO(p.DO.Session(config))
}

func (p platformDo) Clauses(conds ...clause.Expression) IPlatformDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p platformDo) Returning(value interface{}, columns ...string) IPlatformDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p platformDo) Not(conds ...gen.Condition) IPlatformDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p platformDo) Or(conds ...gen.Condition) IPlatformDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p platformDo) Select(conds ...field.Expr) IPlatformDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p platformDo) Where(conds ...gen.Condition) IPlatformDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p platformDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPlatformDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p platformDo) Order(conds ...field.Expr) IPlatformDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p platformDo) Distinct(cols ...field.Expr) IPlatformDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p platformDo) Omit(cols ...field.Expr) IPlatformDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p platformDo) Join(table schema.Tabler, on ...field.Expr) IPlatformDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p platformDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPlatformDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p platformDo) RightJoin(table schema.Tabler, on ...field.Expr) IPlatformDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p platformDo) Group(cols ...field.Expr) IPlatformDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p platformDo) Having(conds ...gen.Condition) IPlatformDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p platformDo) Limit(limit int) IPlatformDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p platformDo) Offset(offset int) IPlatformDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p platformDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPlatformDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p platformDo) Unscoped() IPlatformDo {
	return p.withDO(p.DO.Unscoped())
}

func (p platformDo) Create(values ...*model.Platform) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p platformDo) CreateInBatches(values []*model.Platform, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p platformDo) Save(values ...*model.Platform) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p platformDo) First() (*model.Platform, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Platform), nil
	}
}

func (p platformDo) Take() (*model.Platform, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Platform), nil
	}
}

func (p platformDo) Last() (*model.Platform, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Platform), nil
	}
}

func (p platformDo) Find() ([]*model.Platform, error) {
	result, err := p.DO.Find()
	return result.([]*model.Platform), err
}

func (p platformDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Platform, err error) {
	buf := make([]*model.Platform, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p platformDo) FindInBatches(result *[]*model.Platform, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p platformDo) Attrs(attrs ...field.AssignExpr) IPlatformDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p platformDo) Assign(attrs ...field.AssignExpr) IPlatformDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p platformDo) Joins(fields ...field.RelationField) IPlatformDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p platformDo) Preload(fields ...field.RelationField) IPlatformDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p platformDo) FirstOrInit() (*model.Platform, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Platform), nil
	}
}

func (p platformDo) FirstOrCreate() (*model.Platform, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Platform), nil
	}
}

func (p platformDo) FindByPage(offset int, limit int) (result []*model.Platform, count int64, err error) {
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

func (p platformDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p platformDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p platformDo) Delete(models ...*model.Platform) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *platformDo) withDO(do gen.Dao) *platformDo {
	p.DO = *do.(*gen.DO)
	return p
}
