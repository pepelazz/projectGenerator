package types

import (
	"fmt"
	"github.com/pepelazz/projectGenerator/src/utils"
)

const (
	FldTypeString = "string"
	FldTypeText = "text"
	FldTypeInt = "int"
	FldTypeDouble = "double"
	FldTypeDate = "date"
	FldTypeDatetime = "datetime"
	FldVueTypeSelect = "select"
	FldVueTypeMultipleSelect = "multipleSelect"
)

type (
	FldType struct {
		Name   string
		NameRu string
		Type   string
		Vue    FldVue
		Sql    FldSql
	}

	FldVue struct {
		Name      string
		NameRu    string
		Type      string
		RowCol    [][]int
		Class     []string
		IsRequred bool
		Ext       map[string]string
		Options   []FldVueOptionsItem
	}

	FldSql struct {
		IsSearch   bool
		IsRequired bool
		Ref        string
		IsUniq     bool
		Size       int
		IsOptionFld bool
	}

	FldVueOptionsItem struct {
		Label string `json:"label"`
		Value interface{} `json:"value"`
	}
)

func (fld *FldType) PrintPgModel() string {
	typeStr := fmt.Sprintf(`type="%s"`, fld.Type)
	extStr := ""
	if fld.Type == "string" {
		if fld.Sql.Size > 0 {
			typeStr = fmt.Sprintf("type=\"char\",\tsize=%v", fld.Sql.Size)
		} else {
			typeStr = `type="text"`
		}
	}
	if utils.CheckContainsSliceStr(fld.Type, FldTypeDate, FldTypeDatetime) {
		typeStr = `type="timestamp"`
	}
	if fld.Sql.IsRequired {
		extStr = "not null"
	}
	// ext может быть пустой
	ext := ""
	if len(extStr) > 0 {
		ext = fmt.Sprintf(" \text=\"%s\",", extStr)
	}
	res := fmt.Sprintf("\t{name=\"%s\",\t\t\t\t\t%s,%s\t comment=\"%s\"}", fld.Name, typeStr, ext, fld.NameRu)

	return res
}

func (fld *FldType) PgInsertType() string {
	switch fld.Type {
	case FldTypeDouble:
		return "double precision"
	case FldTypeString:
		return "text"
	case FldTypeDate, FldTypeDatetime:
		return "timestamp"
	default:
		return fld.Type
	}
}

func (fld *FldType) PgUpdateType() string {
	switch fld.Type {
	case FldTypeInt, FldTypeDouble:
		return "number"
	case FldTypeString:
		return "text"
	case FldTypeDate, FldTypeDatetime:
		return "timestamp"
	default:
		return fld.Type
	}
}
