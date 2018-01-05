package main

// START OMIT
type ObjectA struct{} // Do we really need to have one interface?
type ObjectB struct{} // Or we can keep them separately
func (o ObjectA) Run(task int) bool {}
func (o ObjectB) Do(task int) bool {}
// END OMIT