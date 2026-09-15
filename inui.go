package dsmnd

// InUI is meant to be embedded in any struct
// that uses a Datum to store an entry in a UI.
// Does not include Focused or Expanded/Collapsed.
type InUI struct {
	Visible, Enabled, Selected bool
	Position                   int
}
