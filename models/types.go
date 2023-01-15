package models

import (
    "database/sql/driver"
    "github.com/google/uuid"
)

type DBUUID uuid.UUID

// StringToDBUUID -> parse string to DBUUID
func StringToDBUUID(s string) (DBUUID, error) {
    id, err := uuid.Parse(s)
    return DBUUID(id), err
}

//String -> String Representation of Binary16
func (my DBUUID) String() string {
    return uuid.UUID(my).String()
}

//GormDataType -> sets type to binary(16)
func (my DBUUID) GormDataType() string {
	return "binary(16)"
}

func (my DBUUID) MarshalJSON() ([]byte, error) {
    s := uuid.UUID(my)
    str := "\"" + s.String() + "\""
    return []byte(str), nil
}

func (my *DBUUID) UnmarshalJSON(by []byte) error {
    s, err := uuid.ParseBytes(by)
    *my = DBUUID(s)
    return err
}

// Scan --> From DB
func (my *DBUUID) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = DBUUID(parseByte)
	return err
}

// Value -> TO DB
func (my DBUUID) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}
