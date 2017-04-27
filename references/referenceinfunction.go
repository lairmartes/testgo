package references

//ReferenceUpdater changes the reference values by that values sent in the parameters
func (t *ReferenceTestType) ReferenceUpdater(newName string, newAddress string) {
	t.Name = newName
	t.Address = newAddress
}
