package rigidFramework

// Execute is not aware of any database
func Execute(callback func() error) error {
	return callback()
}
