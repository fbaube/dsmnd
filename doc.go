// Package dsmnd is short for Data semantics and metadata.
// Nickname "Desmond". It contains constants and structs
// and other stuff for describing data, data structures,
// and database schemata.
//
// It should have no dependencies, aside from the stdlib
// (and maybe some 3rd-party DB stuff).
//
// About naming
// 
// All names are allowed, except beginning with "sqlite_".
// CREATE TABLE "TABLE"("#!@""'☺\", "");
// You can use keywords ("TABLE"), special characters
// (""#!@""'☺\"), and even the empty string ("").
//
// If you use brackets or quotes you can use any name
// and there is no restriction :
// create table [--This is a_valid.table+name!?] (x int);
//
// But table names that don't have brackets around them
// should be any alphanumeric combination that doesn't
// start with a digit and does not contain any spaces.
// 
// You can use underline and $ but you can not use
// symbols like: + - ? ! * @ % ^ & # = / \ : " '
//
// SQLite's documentation on identifiers says that square
// brackets are allowed for compatibility with MS Access
// and SQL Server but double quotes are standard and
// appear to be equivalent.
//
// If you want to use a keyword as a name, quote it.
// There are four ways of quoting keywords in SQLite:
//
// 'keyword'       A keyword in single quotes is a string literal.
// "keyword"       A keyword in double-quotes is an identifier.
// [keyword]       A keyword in square brackets is an identifier.
//  This is non-standard SQL, for MS Access and SQL Server compatibility.
// `keyword`       A keyword enclosed in grave accents (ASCII code 96)
//  is an identifier. This is non-standard SQL, for MySQL compatibility.
//
// So, double quoting the table name and you can use any chars.
// [tablename] can be used but not a standard SQL.
// .
package dsmnd
