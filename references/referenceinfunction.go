package references

//TestReference tests the reference right after func declaration
func (t *ReferenceTestType) TestReference(newName string, newAddress string) {
	t.Name = newName
	t.Address = newAddress
}
