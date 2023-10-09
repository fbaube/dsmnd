package dsmnd

import "fmt"

// ColumnSpec specifies a datum (i.e. a struct field and/or a
// DB column), including its generic/portable/DB-independent
// representation using the enumeration [dsmnd.Fundatype].
// It provides enough detail that DB operations can be based
// onn it. Some values for common DB columns are defined in
// file datadxnry.go . Field usage is as follows:
//   - Fundatype: e.g. [db.TEXT], [db.INTG], [db.PKEY], [db.FKEY].
//     TODO: Replace with semantic.
//   - StorName: the field name IN THE DB.
//     NOTE: For PKEY or FKEY, is this authoritative ? There
//     could be multiple unique indices, so maybe it should be.
//   - DispName: short description.
//     NOTE: Exception: for FKEY, the name of the ref'd DB table.
//   - Description: long description.
//     NOTE: For FKEY, describe the purpose+function of the ref.
//
// ColumnSpec's are useful in three situations:
//   - To generate CREATE TABLE statements
//   - To validate fields with semantic meaning (RSN)
//   - To document fields when displaying them in a UI
//
// .
type ColumnSpec Datum

func (p *ColumnSpec) String() string {
	// fields are Fundatype, StorName, DispName, Description
	if p == nil {
		return "ϕ"
	}
	return fmt.Sprintf("%s:%s:%s:\"%s\"", p.BaseDataType,
		p.StorName, p.DispName, p.Description)
}
