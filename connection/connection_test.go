package connection

import (
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	db := GetConnection()
	if db == nil {
		t.Error("Failed to connect to the database")
	}
	defer db.Close()

	fmt.Println("Successfully connected to database")
}
