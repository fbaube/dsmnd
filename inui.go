package dsmnd

// InUI is meant to be embedded in any struct
// that uses a Datum to store an entry in a UI.
type InUI struct {
	Visible, Enabled, Selected bool
	Position                   int
}
