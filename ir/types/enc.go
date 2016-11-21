package types

// global returns the global identifier corresponding to the given ID.
func global(id string) string {
	return "@" + enc(id)
}

// local returns the local identifier corresponding to the given ID.
func local(id string) string {
	return "%" + enc(id)
}

// enc encodes special characters in the given string.
func enc(s string) string {
	// TODO: Encode id if containing special characters.
	return s
}
