package consts

const (
	// PgDsn is the environment variable key for the PostgreSQL Data Source Name (DSN).
	// It should be used to fetch the DSN from the environment, typically specified in the .env file.
	PgDsn = "PG_DSN"

	// TableName specifies the name of the database table for user records.
	TableName = "users"

	// IDColumn defines the column name for the user ID in the database.
	IDColumn = "id"

	// EmailColumn defines the column name for storing user email addresses.
	EmailColumn = "email"

	// HashPasswordColumn specifies the column for storing hashed user passwords.
	HashPasswordColumn = "hash_password"

	// RoleColumn denotes the column used to store the user's role (e.g., "admin" or "user").
	RoleColumn = "role"

	// CreatedAtColumn defines the column that records when a user record was created.
	CreatedAtColumn = "created_at"

	// UpdatedAtColumn defines the column that records when a user record was last updated.
	UpdatedAtColumn = "updated_at"

	// ReturningID is a SQL clause used to return the ID of a newly inserted record.
	ReturningID = "RETURNING id"
)
