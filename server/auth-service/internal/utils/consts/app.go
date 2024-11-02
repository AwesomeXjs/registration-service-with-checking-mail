package consts

const (
	EnvPath = ".env" // EnvPath - contains path to .env file

	// PgDsn is the environment variable key for the PostgreSQL Data Source Name (DSN).
	// This should be set in the .env file to configure the database connection.
	PgDsn = "PG_DSN"

	// GrpcHost is the environment variable key for the gRPC server hostname.
	// It specifies where the gRPC server can be reached, set in the .env file.
	GrpcHost = "GRPC_HOST"

	// GrpcPort is the environment variable key for the gRPC server port number.
	// It indicates the port on which the gRPC server listens for connections, also set in the .env file.
	GrpcPort = "GRPC_PORT"
)
