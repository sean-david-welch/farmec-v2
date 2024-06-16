package stores

import (
	"database/sql"
	"fmt"
	"reflect"
)

func ScanStruct(row interface{}, dest interface{}) error {
	// Validate the type of dest
	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr || destVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dest must be a pointer to a struct, got %T", dest)
	}

	// Ensure row is of a valid type
	var scanner interface {
		Scan(dest ...interface{}) error
	}
	switch value := row.(type) {
	case *sql.Row:
		scanner = value
	case *sql.Rows:
		scanner = value
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}

	// Prepare the scan targets
	destElem := destVal.Elem()
	numFields := destElem.NumField()
	scanTargets := make([]interface{}, numFields)
	for i := 0; i < numFields; i++ {
		field := destElem.Field(i)
		scanTargets[i] = field.Addr().Interface()
	}

	// Perform the scan
	if err := scanner.Scan(scanTargets...); err != nil {
		return err
	}

	return nil
}
