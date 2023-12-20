package dsmnd

var DD_RelFP = ColumnSpec{BDT_TEXT.DT(), "relfp", "Rel. path", "Rel.FP (from CLI)"}
var DD_AbsFP = ColumnSpec{BDT_TEXT.DT(), "absfp", "Abs. path", "Absolute filepath"}
var DD_T_Cre = ColumnSpec{BDT_TEXT.DT(), "t_cre", "Cre. time", "Creation date+time"}
var DD_T_Imp = ColumnSpec{BDT_TEXT.DT(), "t_imp", "Imp. time", "DB import date+time"}
var DD_T_Edt = ColumnSpec{BDT_TEXT.DT(), "t_edt", "Edit time", "Last edit date+time"}
