package ir

// global returns the global identifier corresponding to the given ID.
func global(id string) string {
	// TODO: Encode id if containing special characters.
	return "@" + id
}

// local returns the local identifier corresponding to the given ID.
func local(id string) string {
	// TODO: Encode id if containing special characters.
	return "%" + id
}
