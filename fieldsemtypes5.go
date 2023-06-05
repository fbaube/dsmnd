package dsmnd

// FieldSemType describes a simple field that has semantics.
// Field validation is highly desirable but TBD.
type FieldSemType Datum

/* ULIST of SQLite FieldSemTypes (& BlockSemTypes)

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

// FieldSemTypes are simple semantically-meaningful simple fields,
// each FIVE characters. Note that we would expect that this is
// the place where we would associate an SQLite field type with
// each FieldSemType. However we do not do it here, because
//  1. maybe the user wants to override it, and
//  2. maybe we will want to set this up for a different
//     DB with a different set of field types, and
//  3. FIXME: ?? anyways the values are dead simple: only INTEGER or TEXT.
//
// .
var FieldSemTypes = []FieldSemType{
	// INTEGERS (5)
	{INTG, "INTEG", "Integer", "Generic integer, size unspecified"},
	{INTG, "BOOL_", "Boolean", "Boolean (0|1)"},
	{INTG, "PRKEY", "Pri. key", "Primary table key (unique, non-NULL"},
	{INTG, "FRKEY", "For. key", "Foreign table key"},
	{FLOT, "FLOAT", "Float", "Generic non-integer, size unspecified"},
	// TEXTS (8)
	{TEXT, "STRNG", "String", "Generic string, not text"},
	{TEXT, "TOKEN", "Token", "Generic token (no spaces, punc.)"},
	{TEXT, "FTEXT", "Free text", "Generic free-flowing text"},
	{TEXT, "MTEXT", "Markdown", "Markdown (or plain) text"},
	{TEXT, "JTEXT", "JSON", "JSON content"},
	{TEXT, "XTEXT", "XML text", "XML text such as (Lw)DITA"},
	{TEXT, "HTEXT", "HTML5 text", "HTML5 (or previous) text"},
	{TEXT, "MCFMT", "Microformat", "Microformat record"},
	// TEXT-BASED MISC. (5)
	{TEXT, "FONUM", "Phone nr.", "Telephone number"},
	{TEXT, "EMAIL", "Email", "Email address"},
	{TEXT, "URLIN", "URL/URI/URN", "Generic path ID (URL, URI, URN)"},
	{TEXT, "DATIM", "Date / Time", "Date and/or time (ISO-8601/RFC-3339)"},
	{TEXT, "SEMVR", "Sem. ver. nr.", "Semantic version number (x.y.z)"},
}
