package identity

type UserRefField struct {
	CreatedByID uint64 ``
	CreatedBy   *User  ``
}

type CreatedByField struct {
	UserRefField
}

type UpdatedByField struct {
	UserRefField
}
