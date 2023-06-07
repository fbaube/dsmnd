package dsmnd

// ColumnSpec specifies a datum (i.e. a struct field and/or a
// DB column), including its generic/portable/DB-independent
// representation using the enumeration Datum.Fundatype).
// Some values for common DB columns are defined in file
// datadxnry.go . Field usage is as follows:
//   - Fundatype: TEXT or INTG // TODO Replace w semantic
//   - StorName: the field name IN THE DB
//   - DispName: short description (FIXME: for FKEY, the name
//     of the referenced DB table; also maybe the name of the
//     referenced DB field, cos it might not be "ID")
//   - Description: long description
//
// ColumnSpec's are useful in three situations:
//   - To generate CREATE TABLE statements
//   - To validate fields with semantic meaning (RSN)
//   - To document fields when displaying them in a UI
//
// .
type ColumnSpec Datum // DbDescr
