package dsmnd

// SQLite fundamental datatypes:
// see https://sqlite.org/c3ref/c_blob.html
// There's also complicating factors like
// NUMERIC (and REAL).
// . 
type SqliteFundamentalDatatype int

const(
	SQLITE_INTEGER = SqliteFundamentalDatatype(1) // 64-bit signed integer
	SQLITE_FLOAT   = SqliteFundamentalDatatype(2) // 64-bit IEEE FP number
	SQLITE_TEXT    = SqliteFundamentalDatatype(3) // string 
	SQLITE_BLOB    = SqliteFundamentalDatatype(4)
	SQLITE_NULL    = SqliteFundamentalDatatype(5)
	SQLYT_DATETIME = SqliteFundamentalDatatype(6) // dusnt exist, but should 
)

// BaseDataType is a string that expands upon the SQLite 
// fundamental data types, mainly by adding "key" and "list"
// and "clxn" (collection). Each symbol is four upper case 
// characters; each symbol value is four lower-case letters.
// Collections (lists and sets/enumerations) are included. 
// But in any case see https://sqlite.org/c3ref/c_blob.html
// 
type BaseDataType string

// NOTE: These BaseDataType's _ARE_ USED

const (
	INTG = BaseDataType("1234") // "INT"   // SQLITE_INTEGER 1
	FLOT = BaseDataType("1.0f") // "FLOAT" // SQLITE_FLOAT   2
	TEXT = BaseDataType("AaZz") // "TEXT"  // SQLITE_TEXT    3
	BLOB = BaseDataType("blOb") // "BLOB"  // SQLITE_BLOB    4
	NULL = BaseDataType("null") // "NULL"  // SQLITE_NULL    5
	KEYY = BaseDataType("keyy") // PRIMARY/FOREIGN/OTHER KEY (SQLite "INTEGER")
	LIST = BaseDataType("list") // List (a simple one-dimensional list) 
	CLXN = BaseDataType("clxn") // Collection (a more-complicated data structure)
	OTHR = BaseDataType("othr") // reserved: expansion 
)

type SemanticType Datum

// SemanticFieldType describes the semantics of a simple field.
// Semantically-based field validation is desirable but TBD.
// 
// An app that uses SemanticFieldType's should know how to handle
// each of the output types, so that the app can in each case:
//  1. create the appropriate SQLite field, and
//  2. create a translation layer btwn the app and SQLite, and
//  3. create the proper rendering in HTML (or plain text), and
//  4. maybe do basic validation
//
// N.B. An "enume" is: one value in an enumeration.
// . 
type SemanticFieldType SemanticType

// SemanticFieldTypes are semantically-meaningful simple fields, and
// each name is FIVE characters. Note that we would expect that this 
// is the place where we would associate an [SqliteFundamentalDatatype]
// with each SemanticFieldType. However we do not do it here, because
//  1. maybe the user wants to override it, and
//  2. maybe we will want to set this up for a different
//     DB with a different set of field types, and
//  3. FIXME: ?? anyways the values are dead simple:
//     only INTEGER or TEXT (or KEYY) 
// .
var SemanticFieldTypes = []SemanticFieldType{
	{FLOT, "FLOAT", "Float", "Generic FP number, size unspecified"},
	{BLOB, "BLOB_", "Blob", "Binary large object (program / data"},
	{NULL, "NULL_", "Null", "Generic Not-a-value"}, 
	// INTEGERS (2)
	{INTG, "INTEG", "Integer", "Generic integer, size unspecified"},
	{INTG, "BOOL_", "Boolean", "Boolean (0|1)"},
	// TEXTS (8)
	{TEXT, "STRNG", "String", "Generic string, not text"},
	{TEXT, "TOKEN", "Token", "Generic token (no spaces or punc.)"},
	{TEXT, "FTEXT", "Free-text", "Generic free-flowing text"},
	{TEXT, "MTEXT", "Markdown", "Markdown (or plain) text, incl LwDITA MDITA"},
	{TEXT, "JTEXT", "JSON", "JSON content"},
	{TEXT, "XTEXT", "XML-text", "XML text such as LwDITA XDITA"},
	{TEXT, "HTEXT", "HTML5-text", "HTML[5!] text, incl LwDITA HDITA"},
	{TEXT, "MCFMT", "Microformat", "Microformat record"},
	// TEXT-BASED MISC. (5)
	{TEXT, "FONUM", "Phone-nr.", "Telephone number"},
	{TEXT, "EMAIL", "Email", "Email address"},
	{TEXT, "URLIN", "URL/URI/URN", "Generic path ID (URL, URI, URN)"},
	{TEXT, "DATIM", "Date/Time", "Date and/or time (ISO-8601/RFC-3339)"},
	{TEXT, "SEMVR", "Sem.ver.nr.", "Semantic version number (x.y.z)"},
	// KEYS (3)
	{KEYY, "PRKEY", "Pri.key", "Primary table key (unique, non-NULL, int64"},
	{KEYY, "FRKEY", "For.key", "Foreign table key (int64)"},
	{KEYY, "CXKEY", "Complex for.key", "Complex for.key (e.g. 1/n to other table)"}, 
}

// SemanticColxnType is TBS.
type SemanticColxnType SemanticType

// SemanticColxnTypes are data structures, each FIVE characters.
// ENUME is included because it (a) uses a Datum and (b) also
// implements the hierarchical naming scheme for faceted metadata
// and other enumerations. 
// .
var SemanticColxnTypes = []SemanticColxnType{
	// LISTS (8)
	{LIST, "OLIST", "Ord'd list", "Generic ordered list"},
	{LIST, "ULIST", "Unord list", "Generic unordered list"},
	{LIST, "RLIST", "Rankd list", "List ordered by ranking"},
	{LIST, "SLIST", "Seql. list", "List ordered as a sequence"},
	{LIST, "ELIST", "Enum. list", "List of enumerated elements"},
	{LIST, "ENUME", "Enum. item", "Element of enumeration"},
	{LIST, "XLIST", "Excl. list", "List (select one only)"}, // rbn
	{LIST, "MLIST", "Mult. list", "List (select multiple)"}, // cbx
	// MORE-COMPLEX DATA STRUCTURES (3)
	{CLXN, "TABLE", "Tableframe", "Table / dataframe"},
	{CLXN, "UTREE", "Unord tree", "Tree of unordered children (e.g. XML data)"},
	{CLXN, "OTREE", "Ord'd tree", "Tree of ordered children (XML mixed content)"},

	// Additional types for DATIM
	/* | Date & Time | How | Generic "when" - _gets messy_ |
	   |--|:-:|--|
	   | DATE_	| picker  | Date only, no time |
	   | DDWWM | picker  | Day of week ? Week nr ? Month ? |
	   | DATIM | picker  | Day & time |  
	   | TIME_ | picker  | Time only, no date (recurring ?) |
	   | SESON | enum ol | Season  &nbsp; W S M F |	     
	   | DYPRT | enum ol | Daypart &nbsp; N M A E |
	*/
}

