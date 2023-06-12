package dsmnd

// TableSummary sums up a DB table (but not its columns).
// Field usage is as follows:
//   - Fundatype: TABL
//   - StorName: the name of the table IN THE DB - a "long name"
//   - DispName: a short name (3 ltrs!) for use in building up
//     field names (mainly: primary & foreign keys)
//   - Description: long description
type TableSummary Datum
