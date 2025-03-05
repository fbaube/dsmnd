package dsmnd

import "fmt"

// ColumnSpec specifies a datum - a struct field and/or a
// DB column. This includes its generic/portable/DB-independent
// representation that uses the enumeration [SemanticFieldType]
// (or else [BasicDatatype]).
// It provides enough detail that DB operations can be based
// on it. Some values for common DB columns are defined in
// file `datadxnry.go` . Field usage is as follows:
//   - [Datatype]: a value from [SemanticFieldTypes] (via 
//     which one can get the value of the [BasicDatatype])
//   - [Datum.StorName]: the field name IN THE DB.
//     NOTE: For PKEY or FKEY, is this authoritative ? There
//     could be multiple unique indices, so maybe it should be.
//   - [Datum.DispName]: short description.
//     NOTE the exception that for FKEY, this is the name of the ref'd DB table.
//   - [Datum.Description]: long description.
//     NOTE that for FKEY, this describes the purpose+function of the ref.
//
// ColumnSpec's are useful in at least three situations:
//   - To generate CREATE TABLE statements
//   - To validate fields with semantic meaning (RSN)
//   - To document fields when displaying them in a UI
// .
type ColumnSpec Datum

func (p *ColumnSpec) String() string {
	// fields are Datatype, StorName, DispName, Description
	if p == nil {
		return "ϕ"
	}
	return fmt.Sprintf("<%s:%s(%s)\"%s\">", 
		p.StorName, p.DispName, p.Datatype, p.Description)
}
