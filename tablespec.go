package dsmnd

// TableSpec specifies a DB table (but not its columns).
// Field usage is as follows:
//   - Fundatype: TABL
//   - StorName: a short name (3 ltrs!) for use in building
//     up field names (mainly: primary & foreign keys)
//   - DispName: the name of the table IN THE DB - a "long name"
//   - Description: long description
type TableSpec Datum // DbDescr
