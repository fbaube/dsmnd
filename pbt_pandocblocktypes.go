package dsmnd

import "errors"

// PandocBlockType takes values from
// https://hackage.haskell.org/package/pandoc-types-1.23.1/docs/Text-Pandoc-Definition.html
//
// For more information about each field type,
// see [PandocBlockDescriptors].
// .
type PandocBlockType PandocType

var pbtMap map[PandocBlockType]PandocBlockDescriptor 

func (pbt PandocBlockType) S() string {
     return string(pbt)
}

func (pbt PandocBlockType) DT() Datatype {
     return Datatype(pbt)
}

func PandocBlockDescriptorByType(pbt PandocBlockType) (PandocBlockDescriptor, error) { 
    v, ok := pbtMap[pbt]
    if !ok { return PandocBlockDescriptors[0],
       errors.New("Bad SemClxnType: " + pbt.S()) }
    return v, nil 
}

func init() {
     pbtMap = make(map[PandocBlockType]PandocBlockDescriptor)
     for _, pbtd := range PandocBlockDescriptors {
     	 pbtMap[PandocBlockType(pbtd.StorName)] = pbtd
     }
}

// PandocBlockDescriptor is TBS.
type PandocBlockDescriptor SemanticDescriptor

const(
    PBT_NIL  = PandocBlockType("nil")
    PBT_TEXT = PandocBlockType("text")
    PBT_PARA = PandocBlockType("para")
    PBT_CODE = PandocBlockType("code")
    PBT_QUOT = PandocBlockType("quot")
    PBT_PRE = PandocBlockType("pre") // Line Block 
    PBT_RAW = PandocBlockType("raw")
    PBT_OL  = PandocBlockType("ol")
    PBT_UL  = PandocBlockType("ul")
    PBT_DL  = PandocBlockType("dl")
    PBT_HDR = PandocBlockType("hdr")
    PBT_HR  = PandocBlockType("hr")
    PBT_TBL = PandocBlockType("tbl")
    PBT_FIG = PandocBlockType("fig")
    PBT_DIV = PandocBlockType("div")
)

// {BDT_TEXT.DT(), SFT_TOKEN.S(), "Token", "Generic token or datum tag (no spaces or punc.)"},

// PandocBlockDescriptors describe values of [PandocBlockType].
// .
var PandocBlockDescriptors = []PandocBlockDescriptor{
{BDT_NIL.DT(),  PBT_NIL.S(),   "nil", "NOT FOUND"},
{BDT_TEXT.DT(), PBT_TEXT.S(), "Text", "Plain text [Inline], not a paragraph"},
{BDT_TEXT.DT(), PBT_PARA.S(), "Para", "Paragraph [Inline]"},
{BDT_TEXT.DT(), PBT_CODE.S(), "Code", "Code block (literal), with attributes"},
{BDT_TEXT.DT(), PBT_QUOT.S(), "Quot", "Block quote (list of blocks)"},
{BDT_TEXT.DT(), PBT_PRE.S(), "Pre", "Pre, line block, mult. non-breaking lines"},
{BDT_TEXT.DT(), PBT_RAW.S(), "Raw", "Raw block, with format"},
{BDT_TEXT.DT(), PBT_OL.S(),  "OLst", "Ordered list, with attributes and a list of items, each a list of blocks)"},
{BDT_TEXT.DT(), PBT_UL.S(),  "ULst", "Unordered list (bullets), with a list of items, each a list of blocks)"},
{BDT_TEXT.DT(), PBT_DL.S(),  "DLst", "Definition list: pairs: term (a list of inlines) + 1+ definitions (each a list of blocks)"},
{BDT_TEXT.DT(), PBT_HDR.S(), "Hdr", "Header, with level (integer) and text (inlines)"},
{BDT_TEXT.DT(), PBT_HR.S(), "HorzR", "Horizontal rule"},
{BDT_TEXT.DT(), PBT_TBL.S(), "Tbl", "Table Attr Caption [ColSpec] <br/> TableHead [TableBody] TableFoot (a) Table, with attributes, caption, optional short caption, column alignments and <br/> widths (required), table head, table bodies, and table foot <br/> (b) Table w [Inline] [Alignment] [Double] [TableCell] [[TableCell]] <br/> (c) Table, with caption, column alignments, relative column widths (0 = default), <br/> column headers (each a list of blocks), and rows (each a list of lists of blocks)"},
{BDT_TEXT.DT(), PBT_FIG.S(), "Fig", "Figure, with attributes, caption, and content (list of blocks)"},
{BDT_TEXT.DT(), PBT_DIV.S(), "Div", "Generic block(s) container with attributes"},
}

type PdcTblCell struct {
     // Cell Attr Alignment RowSpan ColSpan [Block]
     PdcTblCellAlnmt
     rowSpan, colSpan int
     }

// https://hackage.haskell.org/package/pandoc-types-1.10/docs/src/Text-Pandoc-Definition.html

// Alignment of a table column 
type PdcTblCellAlnmt string
var  PdcTblCellAlnmts = []PdcTblCellAlnmt {
	"AlignLeft", "AlignRight", "AlignCenter", "AlignDefault", 
	}

// Table cells [contents] are list of Blocks 
// type TableCell = [Block] 

// Formats for raw blocks 
// type Format = String 

