package dsmnd

var DD_RelFP = ColumnSpec{SFT_URLIN.DT(), "relfp", "Rel. path", "Rel.FP (from CLI)"}
var DD_AbsFP = ColumnSpec{SFT_URLIN.DT(), "absfp", "Abs. path", "Absolute filepath"}
var DD_T_Cre = ColumnSpec{SFT_DATIM.DT(), "t_cre", "Cre. time", "Creation date+time"}
var DD_T_Imp = ColumnSpec{SFT_DATIM.DT(), "t_imp", "Imp. time", "DB import date+time"}
var DD_T_Edt = ColumnSpec{SFT_DATIM.DT(), "t_edt", "Edit time", "Last edit date+time"}
