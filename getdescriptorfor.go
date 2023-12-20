package dsmnd

/*
Sadly, this does not coimpile. Not yet, anyways. 
func DescriptorFor[T ~string](s T) (TypeType, Datum) {
     // Sadly, neither of these works
     // switch z := s.(type) {
     // switch z := T.(type) {
     return TT_BASICK, Datum{}
}
*/

func DescriptorFor[T ~string](s T) (TypeType, Datum) {
     df, ok := sftMap[SemanticFieldType(s)]
     if !ok {
     	return TT_SEMFIELD, Datum(df)
	}
     dl, ok := sltMap[SemanticListType(s)]
     if !ok {
     	return TT_SEMLIST, Datum(dl)
	}
     dc, ok := sctMap[SemanticClxnType(s)]
     if !ok {
     	return TT_SEMCLXN, Datum(dc)
	}
     println("No descriptor found for:", s)
     return TT_NIL, Datum{}
}
