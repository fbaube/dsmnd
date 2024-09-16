package dsmnd

// NOTE: THIS _IS_ USED

// TableSummary sums up a DB table (but not its columns).
// Field usage is as follows:
//   - [Datatype]: SCT_TABL
//   - [Datum.StorName]: the name of the table IN THE DB - a "long name"
//   - [Datum.DispName]: a short name (3 ltrs!) for use in building up
//     field names (mainly: primary & foreign keys)
//   - [Datum.Description]: long description
// .
type TableSummary Datum
