package dsmnd

import "fmt"

// When used as an enum node (ELIST) or value (ENUME), 
// will need addition of namespace and isSet/hasKids. 

// Datum is a datum descriptor that contains 
// a storage name (for persistence), a short
// display name (for users), and a description.
//
// It can be used to describe any of the following:
//   - A facet or enumeration
//   - One value in a facet or enumeration
//   - An enumeration (set), or a higher node in a
//     hierarchical structure of facets/enumerations
//   - A table spec (schema) for a database
//   - A table spec (schema) as present in a database
//   - A column spec for a DB table
//   - A field value in a Go struct
//
// Some use cases (TODO: this needs revising): 
//
// An element of an enumeration (example: Language)
//   - BasicDataType: TEXT ("text")
//   - StorName: "EN", "FR"
//   - DispName: "English", "French"
//   - Description: "English (default USA)", "French (default France)"
//
// A database table:
//   - e.g. TableSpec{TABL, "inbatch", "INB", "Batch import of files"}
//   - BasicDataType: TABL ("tabl")
//   - StorName: "inbatch" (in SQL "CREATE TABLE")
//   - DispName: "inb" (when used as a prefix, e.g. "inb_idx"))
//   - Description: "Input batch of imported files"
//
// A database column:
//   - e.g. ColumnSpec{TEXT, "relfp", "Rel. path", "Rel.FP (from CLI)"}
//   - BasicDataType: PKEY ("pkey") OOPS, PATH?
//   - StorName: "relfp"
//   - DispName: "Rel. path"
//   - Description: "Rel.FP (from CLI)"
//
// The BasicDataType can be considered optional, but not other fields.
//
// The following discussion (ref: https://www.sqlite.org/datatype3.html)
// is mostly obsolete now, because we have expanded the set of permissible
// values for BasicDataType...
//
// The SQLite concepts "storage class" and "affinity" are quite similar.
// The values for these are "INTEGER", "REAL", "NUMERIC, "TEXT", "BLOB".
// (Note that primary and foreign keys are overloaded with "INTEGER").
// It turns out that all of them are both classes/affinities AND ("mostly")
// specific data types, so they are simple to use if we can ignore details.
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
//  2. declare every column to be either "INTG" or "TEXT" (in BasicDataType),
//  3. but use "INTEGER" (i.e. "KEYY") for a key (primary or foreign), and
//  4. not worry about "mistakes" like assigning a BLOB (like an image file)
//     to a TEXT column cos SQLite has got us covered.
// .
type Datum struct {
	// BasicDatatype is an SQLite fundamental datatype,
	// enhanced with Primary Key and Foreign Key. 
	// TODO: SemanticType (FUTURE).
	BasicDatatype
	// StorName is a short unique string token - no 
	// spaces or punctuation. For robustness we use 
	// string codes not iota-based integer values, 
	// because values based on iota could change.
	//  1) When a Datum describes a DB item (table or 
	//     column or row's column value), StorName is 
	//     the actual name of the DB field or table, 
	//     and should be all lower case
	//  2) When a Datum describes an enumeration (set) 
	//     or facet, StorName should be all upper case, 
	//     and all values of a StorNames in a particular 
	//     enumeration "should" be of the same length.
	StorName string
	// DispName is for common use by users when brief 
	// is better, such as for column headers.
	DispName string
	// Description is a long-form description.
	Description string
}

func (d Datum) String() string {
	return fmt.Sprintf("\"%s\",\"%s\",\"%s\",\"%s\"",
		d.BasicDatatype, d.StorName, d.DispName, d.Description)
}
