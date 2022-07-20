// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/internal/dal/dao"
)

func newPersonField(db *gorm.DB) personField {
	_personField := personField{}

	_personField.personFieldDo.UseDB(db)
	_personField.personFieldDo.UseModel(&dao.PersonField{})

	tableName := _personField.personFieldDo.TableName()
	_personField.ALL = field.NewField(tableName, "*")
	_personField.OwnerType = field.NewString(tableName, "prsn_cat")
	_personField.OwnerID = field.NewField(tableName, "prsn_id")
	_personField.Gender = field.NewUint8(tableName, "gender")
	_personField.Bloodtype = field.NewUint8(tableName, "bloodtype")
	_personField.BirthYear = field.NewUint16(tableName, "birth_year")
	_personField.BirthMon = field.NewUint8(tableName, "birth_mon")
	_personField.BirthDay = field.NewUint8(tableName, "birth_day")

	_personField.fillFieldMap()

	return _personField
}

type personField struct {
	personFieldDo personFieldDo

	ALL       field.Field
	OwnerType field.String
	OwnerID   field.Field
	Gender    field.Uint8
	Bloodtype field.Uint8
	BirthYear field.Uint16
	BirthMon  field.Uint8
	BirthDay  field.Uint8

	fieldMap map[string]field.Expr
}

func (p personField) Table(newTableName string) *personField {
	p.personFieldDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p personField) As(alias string) *personField {
	p.personFieldDo.DO = *(p.personFieldDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *personField) updateTableName(table string) *personField {
	p.ALL = field.NewField(table, "*")
	p.OwnerType = field.NewString(table, "prsn_cat")
	p.OwnerID = field.NewField(table, "prsn_id")
	p.Gender = field.NewUint8(table, "gender")
	p.Bloodtype = field.NewUint8(table, "bloodtype")
	p.BirthYear = field.NewUint16(table, "birth_year")
	p.BirthMon = field.NewUint8(table, "birth_mon")
	p.BirthDay = field.NewUint8(table, "birth_day")

	p.fillFieldMap()

	return p
}

func (p *personField) WithContext(ctx context.Context) *personFieldDo {
	return p.personFieldDo.WithContext(ctx)
}

func (p personField) TableName() string { return p.personFieldDo.TableName() }

func (p personField) Alias() string { return p.personFieldDo.Alias() }

func (p *personField) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *personField) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["prsn_cat"] = p.OwnerType
	p.fieldMap["prsn_id"] = p.OwnerID
	p.fieldMap["gender"] = p.Gender
	p.fieldMap["bloodtype"] = p.Bloodtype
	p.fieldMap["birth_year"] = p.BirthYear
	p.fieldMap["birth_mon"] = p.BirthMon
	p.fieldMap["birth_day"] = p.BirthDay
}

func (p personField) clone(db *gorm.DB) personField {
	p.personFieldDo.ReplaceDB(db)
	return p
}

type personFieldDo struct{ gen.DO }

func (p personFieldDo) Debug() *personFieldDo {
	return p.withDO(p.DO.Debug())
}

func (p personFieldDo) WithContext(ctx context.Context) *personFieldDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personFieldDo) ReadDB() *personFieldDo {
	return p.Clauses(dbresolver.Read)
}

func (p personFieldDo) WriteDB() *personFieldDo {
	return p.Clauses(dbresolver.Write)
}

func (p personFieldDo) Clauses(conds ...clause.Expression) *personFieldDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personFieldDo) Returning(value interface{}, columns ...string) *personFieldDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personFieldDo) Not(conds ...gen.Condition) *personFieldDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personFieldDo) Or(conds ...gen.Condition) *personFieldDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personFieldDo) Select(conds ...field.Expr) *personFieldDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personFieldDo) Where(conds ...gen.Condition) *personFieldDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personFieldDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *personFieldDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p personFieldDo) Order(conds ...field.Expr) *personFieldDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personFieldDo) Distinct(cols ...field.Expr) *personFieldDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personFieldDo) Omit(cols ...field.Expr) *personFieldDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personFieldDo) Join(table schema.Tabler, on ...field.Expr) *personFieldDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personFieldDo) LeftJoin(table schema.Tabler, on ...field.Expr) *personFieldDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personFieldDo) RightJoin(table schema.Tabler, on ...field.Expr) *personFieldDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personFieldDo) Group(cols ...field.Expr) *personFieldDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personFieldDo) Having(conds ...gen.Condition) *personFieldDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personFieldDo) Limit(limit int) *personFieldDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personFieldDo) Offset(offset int) *personFieldDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personFieldDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *personFieldDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personFieldDo) Unscoped() *personFieldDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personFieldDo) Create(values ...*dao.PersonField) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personFieldDo) CreateInBatches(values []*dao.PersonField, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personFieldDo) Save(values ...*dao.PersonField) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personFieldDo) First() (*dao.PersonField, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonField), nil
	}
}

func (p personFieldDo) Take() (*dao.PersonField, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonField), nil
	}
}

func (p personFieldDo) Last() (*dao.PersonField, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonField), nil
	}
}

func (p personFieldDo) Find() ([]*dao.PersonField, error) {
	result, err := p.DO.Find()
	return result.([]*dao.PersonField), err
}

func (p personFieldDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.PersonField, err error) {
	buf := make([]*dao.PersonField, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personFieldDo) FindInBatches(result *[]*dao.PersonField, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personFieldDo) Attrs(attrs ...field.AssignExpr) *personFieldDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personFieldDo) Assign(attrs ...field.AssignExpr) *personFieldDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personFieldDo) Joins(fields ...field.RelationField) *personFieldDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personFieldDo) Preload(fields ...field.RelationField) *personFieldDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personFieldDo) FirstOrInit() (*dao.PersonField, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonField), nil
	}
}

func (p personFieldDo) FirstOrCreate() (*dao.PersonField, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.PersonField), nil
	}
}

func (p personFieldDo) FindByPage(offset int, limit int) (result []*dao.PersonField, count int64, err error) {
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

func (p personFieldDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p personFieldDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p personFieldDo) Delete(models ...*dao.PersonField) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *personFieldDo) withDO(do gen.Dao) *personFieldDo {
	p.DO = *do.(*gen.DO)
	return p
}
