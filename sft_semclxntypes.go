package dsmnd

// SemanticClxnType characterizes the semantics of data structures 
// more complex than simple lists. ENUME is included because it
// (a) uses a Datum and (b) also can implement the hierarchical
// naming scheme for faceted metadata and other enumerations.
//  - Symbol names: "SCT_" + FIVE UPPER CASE letters
//  - Symbol values: five lower case letters
//
// For more information about each field type,
// see [SemanticClxnDescriptors].
// .
type SemanticClxnType SemanticType

func (sft SemanticClxnType) S() string {
     return string(sft)
}

// SemanticColxnDescriptor is TBS.
type SemanticColxnDescriptor SemanticDescriptor

const(
        SCT_TABLE = SemanticClxnType("TABLE")
        SCT_UTREE = SemanticClxnType("UTREE")
        SCT_OTREE = SemanticClxnType("OTREE")
)

// SemanticClxnDescriptors describe data structures that are 
// more complex than simple lists. ENUME is included because it
// (a) uses a Datum and (b) also can implement the hierarchical
// naming scheme for faceted metadata and other enumerations. 
// .
var SemanticColxnDescriptors = []SemanticColxnDescriptor{
{BDT_CLXN, "TABLE", "Tableframe", "Table / dataframe"},
{BDT_CLXN, "UTREE", "Unord tree", "Tree of unordered children (e.g. XML data)"},
{BDT_CLXN, "OTREE", "Ord'd tree", "Tree of ordered children (XML mixed content)"},
}

