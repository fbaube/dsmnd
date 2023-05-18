package dsmnd

// See https://pkg.go.dev/github.com/mattn/go-sqlite3:
//
// Currently, go-sqlite3 supports the following data types:
//
//   +------------------------------+
//   |go        | sqlite3           |
//   |----------|-------------------|
//   |nil       | null              |
//   |int       | integer           |
//   |int64     | integer           |
//   |bool      | integer           |
//   |float64   | float             |
//   |[]byte    | blob              |
//   |string    | text              |
//   |time.Time | timestamp/datetime|
//   +------------------------------+
// .
