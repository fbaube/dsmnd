// Package dsmnd is short for Data Structure Semantics,
// Metadata, and Descriptors. Nicknamed "Desmond". It
// contains constants and structs and other materials 
// for describing data, data structures, database
// schemata, and a sort of taxonomy of UI elements. 
//
// It has no dependencies aside from the stdlib. 
//
// A digression about naming, vis-à-vis SQLite 
// 
// In principle, ALL names and ALL characters are allowed
// except beginning with "sqlite_". You can use keywords
// ("TABLE"), special characters (line noise), and even
// the empty string ("").
//
// If you use brackets or quotes you can use any name
// and there is no restriction:
//   create table [--This is a_valid.table+name!?] (x int);
//
// But table names that don't have surrounding brackets 
// should be an alphanumeric combination that does not
// start with a digit and does not contain any spaces.
// In such case you can use underline and $ but you can 
// not use symbols like: + - ? ! * @ % ^ & # = / \ : " '
//
// SQLite's documentation on identifiers says that square
// brackets are allowed for compatibility with MS Access
// and SQL Server, but double quotes are standard and
// appear to be equivalent.
//
// If you want to use a keyword as a name, quote it.
// There are four ways of quoting keywords in SQLite:
//  - 'keyword'       A keyword in single quotes is a string literal.
//  - "keyword"       A keyword in double-quotes is an identifier.
//  - [keyword]       A keyword in square brackets is an identifier.
//    (This is non-standard SQL, for MS Access and SQL Server compatibility.)
//  - `keyword`       A keyword enclosed in grave accents (ASCII code 96)
//    is an identifier. (This is non-standard SQL, for MySQL compatibility.)
//
// So, if you double quote the table name you can use any characters, 
// and [tablename] can be used but not a standard SQL.
// .
package dsmnd
