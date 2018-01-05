type A struct {
	B B
}
type B struct {
	C C
}
type C struct {
	D string
}

a := A{}
a.B.C.D // access D through chain

a.GetD() // NO helper methods // HL

a.C.D // NO aliases // HL