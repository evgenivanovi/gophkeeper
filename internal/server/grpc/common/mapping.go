package common

import (
	commonpb "github.com/evgenivanovi/gophkeeper/api/pb/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/md/common"
	"github.com/evgenivanovi/gophkeeper/pkg/proto"
)

/* __________________________________________________ */

func ToMetadataModel(pb *commonpb.Metadata) common.MetadataModel {
	return common.MetadataModel{
		CreatedAt: *proto.NewTimeFromTimestamp(pb.GetCreatedAt()),
		UpdatedAt: proto.NewTimeFromOptionalTimestamp(pb.GetUpdatedAt()),
		DeletedAt: proto.NewTimeFromOptionalTimestamp(pb.GetDeletedAt()),
	}
}

func FromMetadataModel(model common.MetadataModel) *commonpb.Metadata {
	return &commonpb.Metadata{
		CreatedAt: proto.NewTimestampFromTime(&model.CreatedAt),
		UpdatedAt: proto.NewOptionalTimestampFromTime(model.UpdatedAt),
		DeletedAt: proto.NewOptionalTimestampFromTime(model.DeletedAt),
	}
}

/* __________________________________________________ */
