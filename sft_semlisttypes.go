package dsmnd

import "errors"

// SemanticListType assigns the semantics of a simple list. 
//  - Symbol names: "SFT_" + FIVE UPPER CASE letters
//  - Symbol values: five lower case letters
//
// For more information about each field type,
// see [SemanticListDescriptors].
// .
type SemanticListType SemanticType

var sltMap map[SemanticListType]SemanticListDescriptor 

func (slt SemanticListType) S() string {
     return string(slt)
}

func (slt SemanticListType) DT() Datatype {
     return Datatype(slt)
}

func SemanticListDescriptorByType(slt SemanticListType) (SemanticListDescriptor, error) { 
    v, ok := sltMap[slt]
    if !ok { return SemanticListDescriptors[0],
       errors.New("Bad SemListType: " + slt.S()) }
    return v, nil 
}

func init() {
     sltMap = make(map[SemanticListType]SemanticListDescriptor)
     for _, sltd := range SemanticListDescriptors {
     	 sltMap[SemanticListType(sltd.StorName)] = sltd
     }
}

const(	
	SLT_NIL   = SemanticListType("nil")
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
//  1. [Datatype] of each type of list is [BDT_LIST] 
//  2. [StorName]: the internal string value of the list type:
//     FIVE UPPER CASE letters
//  3. [DispName]: the external (for users) string value of the
//     list type, Capitalized
//  4. [Description]: Capitalized free text describing the list type
//
// ENUME is included because it (a) uses a Datum and 
// (b) also can implement the hierarchical naming 
// scheme for faceted metadata and other enumerations. 
// .
var SemanticListDescriptors = []SemanticListDescriptor{
{BDT_NIL.DT(),  SLT_NIL.S(),   "nil", "NOT FOUND"},
{BDT_LIST.DT(), SLT_OLIST.S(), "Ord'd list", "Generic ordered list"},
{BDT_LIST.DT(), SLT_ULIST.S(), "Unord list", "Generic unordered list"},
{BDT_LIST.DT(), SLT_RLIST.S(), "Rankd list", "List ordered by ranking"},
{BDT_LIST.DT(), SLT_SLIST.S(), "Seql. list", "List ordered as a sequence"},
{BDT_LIST.DT(), SLT_ELIST.S(), "Enum. list", "List of enumerated elements"},
{"enume",       "ENUME", "Enum. item", "Element of enumeration"},
{BDT_LIST.DT(), SLT_XLIST.S(), "Excl. list", "List (select one only)"}, // rbn
{BDT_LIST.DT(), SLT_MLIST.S(), "Mult. list", "List (select multiple)"}, // cbx
}
