package models

// UserGroup groups users admins, managers, support
type UserGroup string

const (
	// SuperAdmin ...
	SuperAdmin UserGroup = "Super Administrator"

	// AdminGroup ...
	AdminGroup UserGroup = "Administrator"

	// ManagerGroup ...
	ManagerGroup UserGroup = "Manager"

	// SupportGroup ...
	SupportGroup UserGroup = "Support"
)
