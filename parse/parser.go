package parse

import (
	"github.com/RichardKyire/ormazing/dialect"
	"go/ast"
	"reflect"
)

type DbField struct {
	Name string
	Type string
	Tag  string
}

type DbTable struct {
	Model      interface{}
	Name       string
	Fields     []*DbField
	FieldNames []string
	fieldMap   map[string]*DbField
}

func (table *DbTable) GetField(name string) *DbField {
	return table.fieldMap[name]
}

func ParseToTable(dest interface{}, d dialect.Dialect) *DbTable {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	table := &DbTable{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*DbField),
	}
	for i := 0; i < modelType.NumField(); i++ {
		gField := modelType.Field(i)
		if !gField.Anonymous && ast.IsExported(gField.Name) {
			dbField := &DbField{
				Name: gField.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(gField.Type))),
			}
			if v, ok := gField.Tag.Lookup("ormazing"); ok {
				dbField.Tag = v
			}
			table.Fields = append(table.Fields, dbField)
			table.FieldNames = append(table.FieldNames, dbField.Name)
			table.fieldMap[dbField.Name] = dbField
		}
	}
	return table
}
