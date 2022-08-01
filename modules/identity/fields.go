package identity

type UploadedByField struct {
	UploadedByID uint64 ``
	UploadedBy   *User  ``
}

type CreatedByField struct {
	CreatedByID uint64 ``
	CreatedBy   *User  ``
}

type UpdatedByField struct {
	UpdatedByID *uint64 ``
	UpdatedBy   *User   ``
}
