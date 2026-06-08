// Shared logic for PostgreSQL connection, migrations etc
package db

import "fmt"

func GetDBConnectionURL(username, password, host, port, database string) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)
}
