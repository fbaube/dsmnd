package dsmnd

// BlockSemType is TBS.
type BlockSemType Datum

/* ULIST of SQLite (FieldSemTypes &) BlockSemTypes

An app that uses this package should know how
to handle each of the following output block
types, so that the app can in each case

1. create the appropriate SQLite field, and
2. create a translation layer btwn the app and SQLite, and
3. create the proper rendering in HTML (or plain text), and
4. maybe do basic validation

N.B. An "enume" is: one value in an enumeration
   (can include StorName, DispName, Description)

| Date & Time | How | Generic "when" - _gets messy_ |
|--|:-:|--|
| DATE_ | picker  | Date only, no time |
| DDWWM | picker  | Day of week ? Week nr ? Month ? |
| DATIM | picker  | Day & time |
| TIME_ | picker  | Time only, no date (recurring ?) |
| SESON | enum ol | Season  &nbsp; W S M F |
| DYPRT | enum ol | Daypart &nbsp; N M A E |
*/

// BlockSemTypes are data structures related to simple
// semantically-meaningful block types, each FIVE characters.
// .
var BlockSemTypes = []BlockSemType{
	// LISTS (8)
	{LIST, "OLIST", "Ord'd list", "Generic ordered list"},
	{LIST, "ULIST", "Unord list", "Generic unordered list"},
	{LIST, "RLIST", "Rankd list", "List ordered by ranking"},
	{LIST, "SLIST", "Seql. list", "List ordered as a sequence"},
	{LIST, "ELIST", "Enum. list", "List of enumerated elements"},
	{LIST, "ENUME", "Enum. item", "Element of enumeration"},
	{LIST, "XLIST", "Excl. list", "List (select one only)"}, // rbn
	{LIST, "MLIST", "Mult. list", "List (select multiple)"}, // cbx
	// MISC. STRUX (3)
	{OTHR, "TABLE", "Tableframe", "Table / dataframe"},
	{OTHR, "UTREE", "Unord tree", "Tree of unordered children (e.g. XML data)"},
	{OTHR, "OTREE", "Ord'd tree", "Tree of ordered children (e.g. XML mixed content)"},
}
