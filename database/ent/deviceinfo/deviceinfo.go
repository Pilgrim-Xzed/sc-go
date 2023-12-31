// Code generated by ent, DO NOT EDIT.

package deviceinfo

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the deviceinfo type in the database.
	Label = "device_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldOs holds the string denoting the os field in the database.
	FieldOs = "os"
	// FieldBrowser holds the string denoting the browser field in the database.
	FieldBrowser = "browser"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeGenerations holds the string denoting the generations edge name in mutations.
	EdgeGenerations = "generations"
	// EdgeUpscales holds the string denoting the upscales edge name in mutations.
	EdgeUpscales = "upscales"
	// EdgeVoiceovers holds the string denoting the voiceovers edge name in mutations.
	EdgeVoiceovers = "voiceovers"
	// Table holds the table name of the deviceinfo in the database.
	Table = "device_info"
	// GenerationsTable is the table that holds the generations relation/edge.
	GenerationsTable = "generations"
	// GenerationsInverseTable is the table name for the Generation entity.
	// It exists in this package in order to avoid circular dependency with the "generation" package.
	GenerationsInverseTable = "generations"
	// GenerationsColumn is the table column denoting the generations relation/edge.
	GenerationsColumn = "device_info_id"
	// UpscalesTable is the table that holds the upscales relation/edge.
	UpscalesTable = "upscales"
	// UpscalesInverseTable is the table name for the Upscale entity.
	// It exists in this package in order to avoid circular dependency with the "upscale" package.
	UpscalesInverseTable = "upscales"
	// UpscalesColumn is the table column denoting the upscales relation/edge.
	UpscalesColumn = "device_info_id"
	// VoiceoversTable is the table that holds the voiceovers relation/edge.
	VoiceoversTable = "voiceovers"
	// VoiceoversInverseTable is the table name for the Voiceover entity.
	// It exists in this package in order to avoid circular dependency with the "voiceover" package.
	VoiceoversInverseTable = "voiceovers"
	// VoiceoversColumn is the table column denoting the voiceovers relation/edge.
	VoiceoversColumn = "device_info_id"
)

// Columns holds all SQL columns for deviceinfo fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldOs,
	FieldBrowser,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
