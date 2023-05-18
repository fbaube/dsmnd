package dsmnd

// FLDTP describes a simple field that has semantics.
// Field validation is highly desirable but TBD.
type FLDTP Datum

// BLKTP is TBS.
type BLKTP Datum

/* ULIST of SQLite BLKTPs

An app that uses this package should know how
to handle each of the following output block
types, so that the app can in each case

1. create the appropriate SQLite field, and
2. create a translation layer btwn the app and SQLite, and
3. create the proper rendering in HTML (or plain text)

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

// FLDTPs are semantic simple fields. Note that we would expect that
// this is the place where we would associate an SQLite field type
// with each FLDTP. However we do not do it here, because
//  1. maybe the user wants to override it, and
//  2. maybe we will want to set this up for a different
//     DB with a different set of field types, and
//  3. FIXME: ?? anyways the values are dead simple: only INTEGER or TEXT.
//
// .
var FLDTPs = []FLDTP{
	// INTEGERS (5)
	{D_INTG, "INTEG", "Integer", "Generic integer, size unspecified"},
	{D_INTG, "BOOL_", "Boolean", "Boolean (0|1)"},
	{D_INTG, "PRKEY", "Pri. key", "Primary table key (unique, non-NULL"},
	{D_INTG, "FRKEY", "For. key", "Foreign table key"},
	{D_FLOT, "FLOAT", "Float", "Generic non-integer, size unspecified"},
	// TEXTS (8)
	{D_TEXT, "STRNG", "String", "Generic string, not text"},
	{D_TEXT, "TOKEN", "Token", "Generic token (no spaces, punc.)"},
	{D_TEXT, "FTEXT", "Free text", "Generic free-flowing text"},
	{D_TEXT, "MTEXT", "Markdown", "Markdown (or plain) text"},
	{D_TEXT, "JTEXT", "JSON", "JSON content"},
	{D_TEXT, "XTEXT", "XML text", "XML text such as (Lw)DITA"},
	{D_TEXT, "HTEXT", "HTML5 text", "HTML5 (or previous) text"},
	{D_TEXT, "MCFMT", "Microformat", "Microformat record"},
	// TEXT-BASED MISC. (5)
	{D_TEXT, "FONUM", "Phone nr.", "Telephone number"},
	{D_TEXT, "EMAIL", "Email", "Email address"},
	{D_TEXT, "URLIN", "URL/URI/URN", "Generic path ID (URL, URI, URN)"},
	{D_TEXT, "DATIM", "Date / Time", "Date and/or time (ISO-8601/RFC-3339)"},
	{D_TEXT, "SEMVR", "Sem. ver. nr.", "Semantic version number (x.y.z)"},
}

// BLKTPs are data structures related to semantic field types.
var BLKTPs = []BLKTP{
	// LISTS (8)
	{D_LIST, "OLIST", "Ord'd list", "Generic ordered list"},
	{D_LIST, "ULIST", "Unord list", "Generic unordered list"},
	{D_LIST, "RLIST", "Rankd list", "List ordered by ranking"},
	{D_LIST, "SLIST", "Seql. list", "List ordered as a sequence"},
	{D_LIST, "ELIST", "Enum. list", "List of enumerated elements"},
	{D_LIST, "ENUME", "Enum. item", "Element of enumeration"},
	{D_LIST, "XLIST", "Excl. list", "List (select one only)"}, // rbn
	{D_LIST, "MLIST", "Mult. list", "List (select multiple)"}, // cbx
	// MISC. STRUX (3)
	{D_OTHR, "TABLE", "Tableframe", "Table / dataframe"},
	{D_OTHR, "UTREE", "Unord tree", "Tree of unordered children (e.g. XML data)"},
	{D_OTHR, "OTREE", "Ord'd tree", "Tree of ordered children (e.g. XML mixed content)"},
}
