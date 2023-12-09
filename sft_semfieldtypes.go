package dsmnd

// SemanticFieldType assigns the semantics of a simple field.
// Semantically-based field validation is desirable but TBD.
//  - Symbol names: "SFT_" + FIVE UPPER CASE letters
//  - Symbol values: five lower case letters
//
// For more information about each field type,
// see [SemanticFieldDescriptors].
// .
type SemanticFieldType SemanticType

func (sft SemanticFieldType) S() string {
     return string(sft)
}

const(
        SFT_FLOAT = SemanticFieldType("float")
        SFT_BLOB_ = SemanticFieldType("blob_")
        SFT_NULL_ = SemanticFieldType("null_")
        // INTEGERS (2) (also needs BYTE_, LONG_ ?)
        SFT_INTEG = SemanticFieldType("integ")
        SFT_BOOL_ = SemanticFieldType("bool_")
        // TEXTS (8)
        SFT_STRNG = SemanticFieldType("strng")
        SFT_TOKEN = SemanticFieldType("token")
        SFT_FTEXT = SemanticFieldType("ftext")
        SFT_MTEXT = SemanticFieldType("mtext")
        SFT_JTEXT = SemanticFieldType("jtext")
        SFT_XTEXT = SemanticFieldType("xtext")
        SFT_HTEXT = SemanticFieldType("htext")
        SFT_MCFMT = SemanticFieldType("mcfmt")
	// TEXT-BASED MISC. (5)
        SFT_FONUM = SemanticFieldType("fonum")
        SFT_EMAIL = SemanticFieldType("email")
        SFT_URLIN = SemanticFieldType("urlin")
        SFT_DATIM = SemanticFieldType("datim") // ISO-8601
        SFT_SEMVR = SemanticFieldType("semvr")
        // KEYS (3)
        SFT_PRKEY = SemanticFieldType("prkey")
        SFT_FRKEY = SemanticFieldType("frkey")
        SFT_CXKEY = SemanticFieldType("cxkey")
)

// SemanticFieldDescriptor details the semantics of a simple field.
// An app that uses SemanticFieldDescriptor's should know how to
// handle each field type, so that the app can in each case:
//  1. create the appropriate SQLite field, and
//  2. create a translation layer btwn the app and SQLite, and
//  3. create the proper rendering in HTML (or plain text), and
//  4. maybe do basic validation
//
// N.B. An "enume" is: one value in an enumeration.
// .
type SemanticFieldDescriptor SemanticDescriptor

// SemanticFieldDescriptors are semantically-meaningful descriptors
// of simple fields. Descriptor fields:
//  1. [BasicDatatype]
//  2. StorName: the internal string value of the field type:
//     FIVE UPPER CASE letters
//  3. DispName: the external (for users) string value of the
//     field type: a single word (or token), Capitalized
//  4. Description: Capitalized free text describing the field type
// 
// One might expect that this would be the place where we would
// associate an [SqliteDatatype] with each SemanticFieldDescriptor.
// However we do not do it here, because
//  1. maybe the user wants to override it, and
//  2. maybe we will want to set this up for a different
//     DB with a different set of field types, and
//  3. FIXME: ?? anyways the values are dead simple:
//     mostly just Text (or Integer or Key) 
// .
var SemanticFieldDescriptors = []SemanticFieldDescriptor{
// ONE-OFFS (3) 
{BDT_FLOT, "FLOAT", "Float", "Generic FP number, size unspecified"},
{BDT_BLOB, "BLOB_", "Blob", "Binary large object (program / data"},
{BDT_NULL, "NULL_", "Null", "Generic Not-a-value"}, 
// INTEGERS (2)
{BDT_INTG, "INTEG", "Integer", "Generic integer, size unspecified"},
{BDT_INTG, "BOOL_", "Boolean", "Boolean (0|1)"},
// TEXTS (8)
{BDT_TEXT, "STRNG", "String", "Generic string, not text"},
{BDT_TEXT, "TOKEN", "Token", "Generic token (no spaces or punc.)"},
{BDT_TEXT, "FTEXT", "Free-text", "Generic free-flowing text"},
{BDT_TEXT, "MTEXT", "Markdown", "Markdown (or plain) text, incl LwDITA MDITA"},
{BDT_TEXT, "JTEXT", "JSON", "JSON content"},
{BDT_TEXT, "XTEXT", "XML-text", "XML text such as LwDITA XDITA"},
{BDT_TEXT, "HTEXT", "HTML5-text", "HTML[5!] text, incl LwDITA HDITA"},
{BDT_TEXT, "MCFMT", "Microformat", "Microformat record"},
// TEXT-BASED MISC. (5)
{BDT_TEXT, "FONUM", "Phone-nr.", "Telephone number"},
{BDT_TEXT, "EMAIL", "Email", "Email address"},
{BDT_TEXT, "URLIN", "URL/URI/URN", "Generic path ID (URL, URI, URN)"},
{BDT_TEXT, "DATIM", "Date/Time", "Date and/or time (ISO-8601/RFC-3339)"},
{BDT_TEXT, "SEMVR", "Sem.ver.nr.", "Semantic version number (x.y.z)"},
// KEYS (3)
{BDT_KEYY, "PRKEY", "Pri.key", "Primary table key (unique, non-NULL, int64"},
{BDT_KEYY, "FRKEY", "For.key", "Foreign table key (int64)"},
{BDT_KEYY, "CXKEY", "Complex for.key", "Complex for.key (e.g. 1/n to other table)"}, 
// DYTM/DATETIME (6)
{BDT_DYTM, "DATE_", "Date", "Date only, no time"},
{BDT_DYTM, "DDWWM", "TBD", "Day of week ? Week nr ? Month ?"},
{BDT_DYTM, "DATIM", "Date & time", "Timestamp (day+time)"},
{BDT_DYTM, "TIME_", "Time", "Time only, no date (recurring ?)"},
{BDT_DYTM, "SESON", "Season", "Season (W S M F)"},
{BDT_DYTM, "DYPRT", "Daypart", "Part ofo day (N M A E)"},

}
