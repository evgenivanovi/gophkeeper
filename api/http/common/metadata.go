package common

import "time"

/* __________________________________________________ */

type MetadataModel struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

/* __________________________________________________ */
