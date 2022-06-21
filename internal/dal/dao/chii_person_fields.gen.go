// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"github.com/bangumi/server/internal/model"
)

const TableNamePersonField = "chii_person_fields"

// PersonField mapped from table <chii_person_fields>
type PersonField struct {
	OwnerType string         `gorm:"column:prsn_cat;type:enum('prsn','crt');primaryKey"`
	OwnerID   model.PersonID `gorm:"column:prsn_id;type:int(8) unsigned;primaryKey;index:prsn_id,priority:1"`
	Gender    uint8          `gorm:"column:gender;type:tinyint(4) unsigned;not null"`
	Bloodtype uint8          `gorm:"column:bloodtype;type:tinyint(4) unsigned;not null"`
	BirthYear uint16         `gorm:"column:birth_year;type:year(4);not null"`
	BirthMon  uint8          `gorm:"column:birth_mon;type:tinyint(2) unsigned;not null"`
	BirthDay  uint8          `gorm:"column:birth_day;type:tinyint(2) unsigned;not null"`
}

// TableName PersonField's table name
func (*PersonField) TableName() string {
	return TableNamePersonField
}
