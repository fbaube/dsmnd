package dsmnd

import "fmt"

// Datum is a datum descriptor that contains a storage name (for
// persistence), a display name (for users), and a description.
//
// It can be used to describe any of the following:
//   - A facet or enumeration
//   - One value in a facet or enumeration
//   - A table spec in a database
//   - A column spec in a DB table
//   - A field value in a Go struct
//
// Some use cases:
//
// An element of an enumeration (example: Language)
//   - Fundatype: TEXT ("text")
//   - StorName: "EN", "FR"
//   - DispName: "English", "French"
//   - Description: "English (default USA)", "French (default France)"
//
// A database table:
//   - e.g. TableSpec{TABL, "inbatch", "INB", "Batch import of files"}
//   - Fundatype: TABL ("tabl")
//   - StorName: "inbatch" (in SQL "CREATE TABLE")
//   - DispName: "inb" (when used as a prefix, e.g. "inb_idx"))
//   - Description: "Input batch of imported files"
//
// A database column:
//   - e.g. ColumnSpec{TEXT, "relfp", "Rel. path", "Rel.FP (from CLI)"}
//   - Fundatype: PKEY ("pkey") OOPS, PATH?
//   - StorName: "relfp"
//   - DispName: "Rel. path"
//   - Description: "Rel.FP (from CLI)"
//
// The Fundatype can be considered optional, but not other fields.
//
// The following discussion (ref: https://www.sqlite.org/datatype3.html)
// is mostly obsolete now, because we have expanded the set of permissible
// values for Fundatype...
//
// The SQLite concepts "storage class" and "affinity" are quite similar.
// The values for these are "INTEGER", "REAL", "NUMERIC, "TEXT", "BLOB".
// It turns out that all of them are both classes/affinities AND specific
// data types, so they are simple to use if we don't care about details.
//
// For purposes of comparison, note that the dataframe library "gota"
// uses "string", "int", "float", "bool". We will not use a BOOL data
// type anywhere, and anyways a BOOL can be stored in SQLite as an
// INTEGER limited to (0,1).
//
// Furthermore, in SQLite it seems that "BLOBs are always stored as
// BLOBs regardless of column affinity".
//
// This all means that we can basically
//  1. completely ignore SQLite NUMERIC, SQLite BLOB, and gota BOOL,
//  2. declare every column to be either "INT" or "TEXT" (in Fundatype),
//  3. be sure to use "INTEGER" when specifying a key (primary or foreign), and
//  4. not worry about "mistakes" like assigning a BLOB (like an image file)
//     to a TEXT column cos SQLite has got us covered.
//
// .
type Datum struct {
	// Fundatype can be an SQLite fundamental datatype,
	// but note that we have enhanced the list with some
	// semantics, such as (for example) Primary Key.
	Fundatype
	// StorName is a short unique string token - no spaces
	// or punctuation. For robustness we use string codes
	// not iota-based integer values, because values based
	// on iota could change.
	//  1) When a Datum describes a DB item (table or column
	//     or row's column value), StorName is the actual
	//     name of the DB field or table, and should be all
	//     lower case, and ideally all of the same length.
	//  2) When a Datum describes an enumeration, StorName
	//     should be all upper case, and all values of a
	//     StorNames in a particular enumeration "should"
	//     be of the same length.
	StorName string
	// DispName is for common use by users when brief is
	// better, such as for column headers.
	DispName string
	// Description is a long-form description.
	Description string
}

func (d Datum) String() string {
	return fmt.Sprintf("\"%s\",\"%s\",\"%s\",\"%s\"",
		d.Fundatype, d.StorName, d.DispName, d.Description)
}
