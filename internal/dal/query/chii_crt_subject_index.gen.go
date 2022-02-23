// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/bangumi/server/internal/dal/dao"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newCharacterSubjects(db *gorm.DB) characterSubjects {
	_characterSubjects := characterSubjects{}

	_characterSubjects.characterSubjectsDo.UseDB(db)
	_characterSubjects.characterSubjectsDo.UseModel(&dao.CharacterSubjects{})

	tableName := _characterSubjects.characterSubjectsDo.TableName()
	_characterSubjects.ALL = field.NewField(tableName, "*")
	_characterSubjects.CrtID = field.NewUint32(tableName, "crt_id")
	_characterSubjects.SubjectID = field.NewUint32(tableName, "subject_id")
	_characterSubjects.SubjectTypeID = field.NewUint8(tableName, "subject_type_id")
	_characterSubjects.CrtType = field.NewUint8(tableName, "crt_type")
	_characterSubjects.CtrAppearEps = field.NewString(tableName, "ctr_appear_eps")
	_characterSubjects.CrtOrder = field.NewUint8(tableName, "crt_order")
	_characterSubjects.Character = characterSubjectsHasOneCharacter{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Character", "dao.Character"),
	}

	_characterSubjects.fillFieldMap()

	return _characterSubjects
}

type characterSubjects struct {
	characterSubjectsDo characterSubjectsDo

	ALL           field.Field
	CrtID         field.Uint32
	SubjectID     field.Uint32
	SubjectTypeID field.Uint8
	CrtType       field.Uint8
	CtrAppearEps  field.String
	CrtOrder      field.Uint8
	Character     characterSubjectsHasOneCharacter

	fieldMap map[string]field.Expr
}

func (c characterSubjects) Table(newTableName string) *characterSubjects {
	c.characterSubjectsDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c characterSubjects) As(alias string) *characterSubjects {
	c.characterSubjectsDo.DO = *(c.characterSubjectsDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *characterSubjects) updateTableName(table string) *characterSubjects {
	c.ALL = field.NewField(table, "*")
	c.CrtID = field.NewUint32(table, "crt_id")
	c.SubjectID = field.NewUint32(table, "subject_id")
	c.SubjectTypeID = field.NewUint8(table, "subject_type_id")
	c.CrtType = field.NewUint8(table, "crt_type")
	c.CtrAppearEps = field.NewString(table, "ctr_appear_eps")
	c.CrtOrder = field.NewUint8(table, "crt_order")

	c.fillFieldMap()

	return c
}

func (c *characterSubjects) WithContext(ctx context.Context) *characterSubjectsDo {
	return c.characterSubjectsDo.WithContext(ctx)
}

func (c characterSubjects) TableName() string { return c.characterSubjectsDo.TableName() }

func (c *characterSubjects) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *characterSubjects) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["crt_id"] = c.CrtID
	c.fieldMap["subject_id"] = c.SubjectID
	c.fieldMap["subject_type_id"] = c.SubjectTypeID
	c.fieldMap["crt_type"] = c.CrtType
	c.fieldMap["ctr_appear_eps"] = c.CtrAppearEps
	c.fieldMap["crt_order"] = c.CrtOrder

}

func (c characterSubjects) clone(db *gorm.DB) characterSubjects {
	c.characterSubjectsDo.ReplaceDB(db)
	return c
}

type characterSubjectsHasOneCharacter struct {
	db *gorm.DB

	field.RelationField
}

func (a characterSubjectsHasOneCharacter) Where(conds ...field.Expr) *characterSubjectsHasOneCharacter {
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

func (a characterSubjectsHasOneCharacter) WithContext(ctx context.Context) *characterSubjectsHasOneCharacter {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a characterSubjectsHasOneCharacter) Model(m *dao.CharacterSubjects) *characterSubjectsHasOneCharacterTx {
	return &characterSubjectsHasOneCharacterTx{a.db.Model(m).Association(a.Name())}
}

type characterSubjectsHasOneCharacterTx struct{ tx *gorm.Association }

func (a characterSubjectsHasOneCharacterTx) Find() (result *dao.Character, err error) {
	return result, a.tx.Find(&result)
}

func (a characterSubjectsHasOneCharacterTx) Append(values ...*dao.Character) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a characterSubjectsHasOneCharacterTx) Replace(values ...*dao.Character) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a characterSubjectsHasOneCharacterTx) Delete(values ...*dao.Character) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a characterSubjectsHasOneCharacterTx) Clear() error {
	return a.tx.Clear()
}

func (a characterSubjectsHasOneCharacterTx) Count() int64 {
	return a.tx.Count()
}

type characterSubjectsDo struct{ gen.DO }

//GetBySubject ...
//
//where(subject_id=@id)
func (c characterSubjectsDo) GetBySubject(id uint32) (result []*dao.CharacterSubjects, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("subject_id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = c.UnderlyingDB().Where(generateSQL.String(), params).Find(&result)
	} else {
		executeSQL = c.UnderlyingDB().Where(generateSQL.String()).Find(&result)
	}
	err = executeSQL.Error
	return
}

//GetByCharacter ...
//
//where(crt_id=@id)
func (c characterSubjectsDo) GetByCharacter(id uint32) (result []*dao.CharacterSubjects, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["id"] = id
	generateSQL.WriteString("crt_id=@id ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = c.UnderlyingDB().Where(generateSQL.String(), params).Find(&result)
	} else {
		executeSQL = c.UnderlyingDB().Where(generateSQL.String()).Find(&result)
	}
	err = executeSQL.Error
	return
}

func (c characterSubjectsDo) Debug() *characterSubjectsDo {
	return c.withDO(c.DO.Debug())
}

func (c characterSubjectsDo) WithContext(ctx context.Context) *characterSubjectsDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c characterSubjectsDo) Clauses(conds ...clause.Expression) *characterSubjectsDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c characterSubjectsDo) Not(conds ...gen.Condition) *characterSubjectsDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c characterSubjectsDo) Or(conds ...gen.Condition) *characterSubjectsDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c characterSubjectsDo) Select(conds ...field.Expr) *characterSubjectsDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c characterSubjectsDo) Where(conds ...gen.Condition) *characterSubjectsDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c characterSubjectsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *characterSubjectsDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c characterSubjectsDo) Order(conds ...field.Expr) *characterSubjectsDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c characterSubjectsDo) Distinct(cols ...field.Expr) *characterSubjectsDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c characterSubjectsDo) Omit(cols ...field.Expr) *characterSubjectsDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c characterSubjectsDo) Join(table schema.Tabler, on ...field.Expr) *characterSubjectsDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c characterSubjectsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *characterSubjectsDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c characterSubjectsDo) RightJoin(table schema.Tabler, on ...field.Expr) *characterSubjectsDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c characterSubjectsDo) Group(cols ...field.Expr) *characterSubjectsDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c characterSubjectsDo) Having(conds ...gen.Condition) *characterSubjectsDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c characterSubjectsDo) Limit(limit int) *characterSubjectsDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c characterSubjectsDo) Offset(offset int) *characterSubjectsDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c characterSubjectsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *characterSubjectsDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c characterSubjectsDo) Unscoped() *characterSubjectsDo {
	return c.withDO(c.DO.Unscoped())
}

func (c characterSubjectsDo) Create(values ...*dao.CharacterSubjects) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c characterSubjectsDo) CreateInBatches(values []*dao.CharacterSubjects, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c characterSubjectsDo) Save(values ...*dao.CharacterSubjects) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c characterSubjectsDo) First() (*dao.CharacterSubjects, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.CharacterSubjects), nil
	}
}

func (c characterSubjectsDo) Take() (*dao.CharacterSubjects, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.CharacterSubjects), nil
	}
}

func (c characterSubjectsDo) Last() (*dao.CharacterSubjects, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.CharacterSubjects), nil
	}
}

func (c characterSubjectsDo) Find() ([]*dao.CharacterSubjects, error) {
	result, err := c.DO.Find()
	return result.([]*dao.CharacterSubjects), err
}

func (c characterSubjectsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.CharacterSubjects, err error) {
	buf := make([]*dao.CharacterSubjects, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c characterSubjectsDo) FindInBatches(result *[]*dao.CharacterSubjects, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c characterSubjectsDo) Attrs(attrs ...field.AssignExpr) *characterSubjectsDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c characterSubjectsDo) Assign(attrs ...field.AssignExpr) *characterSubjectsDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c characterSubjectsDo) Joins(field field.RelationField) *characterSubjectsDo {
	return c.withDO(c.DO.Joins(field))
}

func (c characterSubjectsDo) Preload(field field.RelationField) *characterSubjectsDo {
	return c.withDO(c.DO.Preload(field))
}

func (c characterSubjectsDo) FirstOrInit() (*dao.CharacterSubjects, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.CharacterSubjects), nil
	}
}

func (c characterSubjectsDo) FirstOrCreate() (*dao.CharacterSubjects, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.CharacterSubjects), nil
	}
}

func (c characterSubjectsDo) FindByPage(offset int, limit int) (result []*dao.CharacterSubjects, count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	if limit <= 0 {
		return
	}

	result, err = c.Offset(offset).Limit(limit).Find()
	return
}

func (c characterSubjectsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c *characterSubjectsDo) withDO(do gen.Dao) *characterSubjectsDo {
	c.DO = *do.(*gen.DO)
	return c
}
