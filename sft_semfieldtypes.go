package dsmnd

import "errors"

// SemanticFieldType assigns the semantics of a simple field.
// Semantically-based field validation is desirable but TBD.
//  - Symbol names: "SFT_" + FIVE UPPER CASE letters
//  - Symbol values: five lower case alphanumeric 
//
// For more information about each field type,
// see [SemanticFieldDescriptors].
// .
type SemanticFieldType SemanticType

var sftMap map[SemanticFieldType]SemanticFieldDescriptor 

func (sft SemanticFieldType) S() string {
     return string(sft)
}

func (sft SemanticFieldType) DT() Datatype {
     return Datatype(sft)
}

func (sft SemanticFieldType) Descriptor() SemanticFieldDescriptor {
     if sft == "" {
     	return SemanticFieldDescriptors[0]
	}
     return sftMap[sft]
}

func (sft SemanticFieldType) BasicDatatype() BasicDatatype {
     if sft == "" {
     	return BDT_NIL
	}
     return BasicDatatype(sft.Descriptor().Datatype)
}

func SemanticFieldDescriptorByType(sft SemanticFieldType) (SemanticFieldDescriptor, error) { 
    v, ok := sftMap[sft]
    if !ok { return SemanticFieldDescriptors[0],
       errors.New("Bad SemFieldType: " + sft.S()) }
    return v, nil 
}

func init() {
     sftMap = make(map[SemanticFieldType]SemanticFieldDescriptor)
     for _, sftd := range SemanticFieldDescriptors {
     	 sftMap[SemanticFieldType(sftd.StorName)] = sftd
     }
}

const(
	SFT_NIL   = SemanticFieldType("nil")
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
	SFT_DATE_ = SemanticFieldType("date_")
	SFT_DDWWM = SemanticFieldType("ddwmm")
	SFT_TIME_ = SemanticFieldType("time_")
	SFT_SESON = SemanticFieldType("seson")
	SFT_DYPRT = SemanticFieldType("dyprt")
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
//  1. [BasicDatatype] (i.e. storage type) 
//  2. [StorName]: the internal string value of the field type:
//     FIVE UPPER CASE letters
//  3. [DispName]: the external (for users) string value of the
//     field type: a single word (or token), Capitalized
//  4. [Description]: Capitalized free text describing the field type
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
{BDT_NIL.DT(),  SFT_NIL.S(),   "nil", "NOT-FOUND"}, 
// ONE-OFFS (3) 
{BDT_FLOT.DT(), SFT_FLOAT.S(), "Float", "Generic FP number, size unspecified"},
{BDT_BLOB.DT(), SFT_BLOB_.S(), "Blob", "Binary large object (program / data"},
{BDT_NULL.DT(), SFT_NULL_.S(), "Null", "Generic Not-a-value"}, 
// INTEGERS (2)
{BDT_INTG.DT(), SFT_INTEG.S(), "Integer", "Generic integer, size unspecified"},
{BDT_INTG.DT(), SFT_BOOL_.S(), "Boolean", "Boolean (0|1)"},
// TEXTS (8)
{BDT_TEXT.DT(), SFT_STRNG.S(), "String", "Generic string, not text"},
{BDT_TEXT.DT(), SFT_TOKEN.S(), "Token", "Generic token (no spaces or punc.)"},
{BDT_TEXT.DT(), SFT_FTEXT.S(), "Free-text", "Generic free-flowing text"},
{BDT_TEXT.DT(), SFT_MTEXT.S(), "Markdown", "Markdown (or plain) text, incl LwDITA MDITA"},
{BDT_TEXT.DT(), SFT_JTEXT.S(), "JSON", "JSON content"},
{BDT_TEXT.DT(), SFT_XTEXT.S(), "XML-text", "XML text such as LwDITA XDITA"},
{BDT_TEXT.DT(), SFT_HTEXT.S(), "HTML5-text", "HTML[5!] text, incl LwDITA HDITA"},
{BDT_TEXT.DT(), SFT_MCFMT.S(), "Microformat", "Microformat record"},
// TEXT-BASED MISC. (5)
{BDT_TEXT.DT(), SFT_FONUM.S(), "Phone-nr.", "Telephone number"},
{BDT_TEXT.DT(), SFT_EMAIL.S(), "Email", "Email address"},
{BDT_TEXT.DT(), SFT_URLIN.S(), "URL/URI/URN", "Generic path ID (URL, URI, URN)"},
{BDT_TEXT.DT(), SFT_DATIM.S(), "Date/Time", "Date and/or time (ISO-8601/RFC-3339)"},
{BDT_TEXT.DT(), SFT_SEMVR.S(), "Sem.ver.nr.", "Semantic version number (x.y.z)"},
// KEYS (3)
{BDT_KEYY.DT(), SFT_PRKEY.S(), "Pri.key", "Primary table key (unique, non-NULL, int64"},
{BDT_KEYY.DT(), SFT_FRKEY.S(), "For.key", "Foreign table key (int64)"},
{BDT_KEYY.DT(), SFT_CXKEY.S(), "Complex for.key", "Complex for.key (e.g. 1/n to other table)"}, 
// DYTM/DATETIME (6)
{BDT_DYTM.DT(), SFT_DATE_.S(), "Date", "Date only, no time"},
{BDT_DYTM.DT(), SFT_DDWWM.S(), "TBD", "Day of week ? Week nr ? Month ?"},
{BDT_DYTM.DT(), SFT_DATIM.S(), "Date & time", "Timestamp (day+time)"},
{BDT_DYTM.DT(), SFT_TIME_.S(), "Time", "Time only, no date (recurring ?)"},
{BDT_DYTM.DT(), SFT_SESON.S(), "Season", "Season (W S M F)"},
{BDT_DYTM.DT(), SFT_DYPRT.S(), "Daypart", "Part ofo day (N M A E)"},

}
