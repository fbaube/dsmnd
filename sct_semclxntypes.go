package dsmnd

import "errors"

// SemanticClxnType characterizes the semantics of data structures 
// that are more complex than simple lists. ENUME is included because 
// it (a) uses a Datum and (b) also can implement the hierarchical
// naming scheme for faceted metadata and other enumerations.
//  - Symbol names are `"SCT_"` + FIVE UPPER CASE letters
//  - Symbol values are five lower case letters
//
// For more information about each field type,
// see [SemanticClxnDescriptors].
// .
type SemanticClxnType SemanticType

var sctMap map[SemanticClxnType]SemanticClxnDescriptor 

func (sct SemanticClxnType) S() string {
     return string(sct)
}

func (sct SemanticClxnType) DT() Datatype {
     return Datatype(sct)
}

func SemanticClxnDescriptorByType(sct SemanticClxnType) (SemanticClxnDescriptor, error) { 
    v, ok := sctMap[sct]
    if !ok { return SemanticClxnDescriptors[0],
       errors.New("Bad SemClxnType: " + sct.S()) }
    return v, nil 
}

func init() {
     sctMap = make(map[SemanticClxnType]SemanticClxnDescriptor)
     for _, sctd := range SemanticClxnDescriptors {
     	 sctMap[SemanticClxnType(sctd.StorName)] = sctd
     }
}

// SemanticClxnDescriptor is TBS.
type SemanticClxnDescriptor SemanticDescriptor

const(
        SCT_NIL   = SemanticClxnType("nil")
        SCT_TABLE = SemanticClxnType("table") // 2-D, rectangular 
        SCT_UTREE = SemanticClxnType("utree") // e.g. XML for data 
        SCT_OTREE = SemanticClxnType("otree") // e.g. markup
)

// SemanticClxnDescriptors describe data structures that are 
// more complex than simple lists. ENUME is included because it
// (a) uses a Datum and (b) also can implement the hierarchical
// naming scheme for faceted metadata and other enumerations. 
// .
var SemanticClxnDescriptors = []SemanticClxnDescriptor{
{BDT_NIL.DT(),  SCT_NIL.S(),   "nil", "NOT FOUND"},
{BDT_CLXN.DT(), SCT_TABLE.S(), "Tableframe", "Table / dataframe"},
{BDT_CLXN.DT(), SCT_UTREE.S(), "Unord tree", "Tree of unordered children (e.g. XML data)"},
{BDT_CLXN.DT(), SCT_OTREE.S(), "Ord'd tree", "Tree of ordered children (XML mixed content)"},
}

