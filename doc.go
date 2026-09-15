// Package dsmnd is short for Data Structure Semantics,
// Metadata, and Descriptors. Nicknamed "Desmond". It
// contains constants and structs and other materials 
// for describing a variety of items:
//  - data types
//  - data enumeration types and values 
//  - data structures
//  - database schemata
//  - a (sort of) taxonomy of UI elements 
//
// It has no dependencies aside from the stdlib. 
//
// (The package description ends here.)
// 
// A digression about naming, vis-à-vis SQLite 
//
// First of all, pathological possibilities are brackets
// for MS Access and SQL Server compatibility) and grave
// accents (ASCII code 96) for MySQL compatibility.)
// AVOID. They are ignored in the following discussion.
// 
// In principle, ALL names and ALL characters are allowed
// except beginning with "sqlite_". You can use keywords
// ("TABLE"), special characters (line noise), and even
// the empty string ("").
//
// Table names should be an alphanumeric combination 
// that does not start with a digit and has no spaces.
// In such case you can use underline and $ but you can 
// not use symbols like: + - ? ! * @ % ^ & # = / \ : " '
//
// If you want to use a keyword as a name, quote it.
// There are two ways to quote keywords in SQLite:
//  - 'keyword'       A keyword in single quotes is a string literal.
//  - "keyword"       A keyword in double-quotes is an identifier.
//
// So, if you double quote the table name you can use any characters.
// .
package dsmnd
