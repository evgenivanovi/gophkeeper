package common

import "time"

/* __________________________________________________ */

type MetadataModel struct {
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

/* __________________________________________________ */
