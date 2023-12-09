package dsmnd

// SemanticListType assigns the semantics of a simple list. 
//  - Symbol names: "SFT_" + FIVE UPPER CASE letters
//  - Symbol values: five lower case letters
//
// For more information about each field type,
// see [SemanticListDescriptors].
// .
type SemanticListType SemanticType

func (sft SemanticListType) S() string {
     return string(sft)
}

const(	
        SLT_OLIST = SemanticListType("olist")
        SLT_ULIST = SemanticListType("ulist")
        SLT_RLIST = SemanticListType("rlist")
        SLT_SLIST = SemanticListType("slist")
        SLT_ELIST = SemanticListType("elist")
        SLT_ENUME = SemanticListType("enume")
        SLT_XLIST = SemanticListType("xlist")
        SLT_MLIST = SemanticListType("mlist")
)

// SemanticListDescriptor details the semantics of a simple list. 
// An app that uses SemanticListDescriptor's should know how to
// handle each field type, so that the app can in each case:
//  1. create the appropriate SQLite field (single/multi value), and
//  2. create a translation layer btwn the app and SQLite, and
//  3. create the proper HTML widget (or plain text UI)
//
// N.B. An "enume" is: one value in an enumeration.
// .
type SemanticListDescriptor SemanticDescriptor


// SemanticListDescriptors are semantically-meaningful descriptors
// of simple lists. Descriptor fields:
//  1. [BasicDatatype] of each list item 
//  2. StorName: the internal string value of the list type:
//     FIVE UPPER CASE letters
//  3. DispName: the external (for users) string value of the
//     list type, Capitalized
//  4. Description: Capitalized free text describing the list type
//
// ENUME is included because it (a) uses a Datum and 
// (b) also can implement the hierarchical naming 
// scheme for faceted metadata and other enumerations. 
// .
var SemanticListDescriptors = []SemanticListDescriptor{
{BDT_LIST, "OLIST", "Ord'd list", "Generic ordered list"},
{BDT_LIST, "ULIST", "Unord list", "Generic unordered list"},
{BDT_LIST, "RLIST", "Rankd list", "List ordered by ranking"},
{BDT_LIST, "SLIST", "Seql. list", "List ordered as a sequence"},
{BDT_LIST, "ELIST", "Enum. list", "List of enumerated elements"},
{BDT_LIST, "ENUME", "Enum. item", "Element of enumeration"},
{BDT_LIST, "XLIST", "Excl. list", "List (select one only)"}, // rbn
{BDT_LIST, "MLIST", "Mult. list", "List (select multiple)"}, // cbx
}
