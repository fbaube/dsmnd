package dsmnd

import "errors"

// PandocInlineType takes values from
// https://hackage.haskell.org/package/pandoc-types-1.23.1/docs/Text-Pandoc-Definition.html
//
// For more information about each field type,
// see [PandocInlineDescriptors].
//
// These values should be compared to the "Mark" values for the Tiptap
// editor, documented at https://tiptap.dev/docs/editor/extensions/marks
// .
type PandocInlineType PandocType

var pitMap map[PandocInlineType]PandocInlineDescriptor 

func (pit PandocInlineType) S() string {
     return string(pit)
}

func (pit PandocInlineType) DT() Datatype {
     return Datatype(pit)
}

func PandocInlineDescriptorByType(pit PandocInlineType) (PandocInlineDescriptor, error) { 
    v, ok := pitMap[pit]
    if !ok { return PandocInlineDescriptors[0],
       errors.New("Bad SemClxnType: " + pit.S()) }
    return v, nil 
}

func init() {
     pitMap = make(map[PandocInlineType]PandocInlineDescriptor)
     for _, pitd := range PandocInlineDescriptors {
     	 pitMap[PandocInlineType(pitd.StorName)] = pitd
     }
}

// PandocInlineDescriptor is TBS.
type PandocInlineDescriptor SemanticDescriptor

// Tiptap https://tiptap.dev/docs/editor/extensions/nodes
// Bold Code Highlight Italic Link Strike Subscript
// Superscript TextStyle Underline

// TextStyle "This mark renders a <span> HTML tag and enables you to 
// add a list of styling related attributes, for example font-family,
// font-size, or color. The extension doesn’t add any styling attri-
// bute by default, but other extensions use it as the foundation,
// for example FontFamily or Color.

const(
    PIT_NIL  = PandocInlineType("nil")
    PIT_TEXT = PandocInlineType("text")
    PIT_PARA = PandocInlineType("para")
    PIT_CODE = PandocInlineType("code")
    PIT_QUOT = PandocInlineType("quot")
    PIT_PRE = PandocInlineType("pre") // Line Block 
    PIT_RAW = PandocInlineType("raw")
    PIT_OL  = PandocInlineType("ol")
    PIT_UL  = PandocInlineType("ul")
    PIT_DL  = PandocInlineType("dl")
    PIT_HDR = PandocInlineType("hdr")
    PIT_HR  = PandocInlineType("hr")
    PIT_TBL = PandocInlineType("tbl")
    PIT_FIG = PandocInlineType("fig")
    PIT_DIV = PandocInlineType("div")
)

// {BDT_TEXT.DT(), SFT_TOKEN.S(), "Token", "Generic token or datum tag (no spaces or punc.)"},

// PandocInlineDescriptors describe values of [PandocInlineType].
// .
var PandocInlineDescriptors = []PandocInlineDescriptor{
{BDT_NIL.DT(),  PIT_NIL.S(),   "nil", "NOT FOUND"},
{BDT_TEXT.DT(), PIT_TEXT.S(), "Text", "Plain text [Inline], not a paragraph"},
{BDT_TEXT.DT(), PIT_PARA.S(), "Para", "Paragraph [Inline]"},
{BDT_TEXT.DT(), PIT_CODE.S(), "Code", "Code block (literal), with attributes"},
{BDT_TEXT.DT(), PIT_QUOT.S(), "Quot", "Block quote (list of blocks)"},
{BDT_TEXT.DT(), PIT_PRE.S(), "Pre", "Pre, line block, mult. non-breaking lines"},
{BDT_TEXT.DT(), PIT_RAW.S(), "Raw", "Raw block, with format"},
{BDT_TEXT.DT(), PIT_OL.S(),  "OLst", "Ordered list, with attributes and a list of items, each a list of blocks)"},
{BDT_TEXT.DT(), PIT_UL.S(),  "ULst", "Unordered list (bullets), with a list of items, each a list of blocks)"},
{BDT_TEXT.DT(), PIT_DL.S(),  "DLst", "Definition list: pairs: term (a list of inlines) + 1+ definitions (each a list of blocks)"},
{BDT_TEXT.DT(), PIT_HDR.S(), "Hdr", "Header, with level (integer) and text (inlines)"},
{BDT_TEXT.DT(), PIT_HR.S(), "HorzR", "Horizontal rule"},
{BDT_TEXT.DT(), PIT_TBL.S(), "Tbl", "Table Attr Caption [ColSpec] <br/> TableHead [TableBody] TableFoot (a) Table, with attributes, caption, optional short caption, column alignments and <br/> widths (required), table head, table bodies, and table foot <br/> (b) Table w [Inline] [Alignment] [Double] [TableCell] [[TableCell]] <br/> (c) Table, with caption, column alignments, relative column widths (0 = default), <br/> column headers (each a list of blocks), and rows (each a list of lists of blocks)"},
{BDT_TEXT.DT(), PIT_FIG.S(), "Fig", "Figure, with attributes, caption, and content (list of blocks)"},
{BDT_TEXT.DT(), PIT_DIV.S(), "Div", "Generic block(s) container with attributes"},
}

/*

// https://hackage.haskell.org/package/pandoc-types-1.23.1/docs/Text-Pandoc-Definition.html 

  <tr>
    <td> Inl_TEXT </td> 
    <td> Str Text </td>
    <td> Text (string) </td>
  </tr>
  <tr>
    <td> Inl_EMPH </td>
    <td> Emph [Inline] </td>
    <td> Emphasized text (list of inlines) </td>
  </tr>
  <tr> 
    <td> Inl_UNDRLN </td>
    <td> Underline [Inline] </td>
    <td> Underlined text (list of inlines) </td>
  </tr>
  <tr>
    <td> Inl_STRONG </td>
    <td> Strong [Inline] </td>
    <td> Strongly emphasized text (list of inlines) </td>
  </tr>
  <tr>
    <td> Inl_STRIKE </td>
    <td> Strikeout [Inline] </td>
    <td> Strikeout text (list of inlines) </td>
  </tr>
  <tr>
    <td> Inl_SUPER </td>
    <td> Superscript [Inline] </td>
    <td> Superscripted text (list of inlines) </td>
  </tr>
  <tr>
    <td> Inl_SUB </td>
    <td> Subscript [Inline] </td>
    <td> Subscripted text (list of inlines) </td>
  </tr>
  <tr>
    <td> Inl_SMLCAP </td>
    <td> SmallCaps [Inline] </td>
    <td> Small caps text (list of inlines) </td>
  </tr>
  <tr>
    <td> Inl_QUOTE </td>
    <td> Quoted QuoteType [Inline] </td>
    <td> Quoted text (list of inlines) </td>
  </tr>
  <tr>
    <td> Inl_CITE </td> 
    <td> Cite [Citation] [Inline] </td>
    <td> Citation (list of inlines) </td>
  </tr>
  <tr>
    <td> Inl_CODE </td>
    <td> Code Attr Text </td>
    <td> Inline code (literal) </td>
  </tr>
  <tr>
    <td> Inl_SPACE </td>
    <td> Space </td>
    <td> Inter-word space </td>
  </tr>
  <tr>
    <td> Inl_SFTBRK </td>
    <td> SoftBreak </td>
    <td> Soft line break </td>
  </tr>
  <tr>
    <td> Inl_LINBRK </td>
    <td> LineBreak </td>
    <td> Hard line break </td>
  </tr>
  <tr>
    <td> Inl_MATH </td>
    <td> Math MathType Text </td>
    <td> TeX math (literal) </td>
  </tr>
  <tr>
    <td> Inl_RAW </td>
    <td> RawInline Format Text </td>
    <td> Raw inline </td>
  </tr>
  <tr>
    <td> Inl_LINK </td>
    <td> Link Attr [Inline] Target </td>
    <td> Hyperlink: alt text (list of inlines), target </td>
  </tr>
  <tr>
    <td> Inl_IMAGE </td>
    <td> Image Attr [Inline] Target </td>
    <td> Image: alt text (list of inlines), target </td>
  </tr>
  <tr>
    <td> Inl_NOTE </td>
    <td> Note [Block] </td>
    <td> Footnote or endnote </td>
  </tr>
  <tr>
    <td> Inl_SPAN </td>
    <td> Span Attr [Inline] </td>
    <td> Generic inline container with attributes </td>
  </tr>
</table>

<p> type ListAttributes = (Int, ListNumberStyle, ListNumberDelim) <br/>
List attributes. The Int element is the start number of the list. </p>

<p> data ListNumberStyle <br/>
Style of list numbers. <br/>
DefaultStyle, Example, Decimal, <br/>
LowerRoman, UpperRoman, LowerAlpha, UpperAlpha </p>

<p> data ListNumberDelim <br/>
Delimiter of list numbers. <br/>
DefaultDelim, Period, OneParen, TwoParens </p> 

<p> data Alignment <br/>
Alignment of a table column. <br/>
AlignLeft, AlignRight, AlignCenter, AlignDefault </p>

<p> data ColWidth <br/>
The width of a table column, as a percentage of the text width. <br/>
ColWidth (Double), ColWidthDefault </p>

<p> type ColSpec = (Alignment, ColWidth) <br/>
The specification for a single table column. </p>

<p> data QuoteType <br/>
Type of quotation marks to use in Quoted inline. <br/>
SingleQuote, DoubleQuote </p>

<p> Link target (URL, title). <br/>
type Target = (String, String) </p>

<p> Type of math element (display or inline). <br/>
data MathType = DisplayMath | InlineMath </p>

https://hackage.haskell.org/package/pandoc-types-1.10/docs/src/Text-Pandoc-Definition.html

*/