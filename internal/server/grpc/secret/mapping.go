package secret

import (
	secretpb "github.com/evgenivanovi/gophkeeper/api/pb/secret"
	"github.com/evgenivanovi/gophkeeper/internal/server/grpc/common"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
)

/* __________________________________________________ */

func ToDecodedSecretModel(
	pb *secretpb.DecodedSecret,
) secretsharedmd.DecodedSecretModel {
	return secretsharedmd.DecodedSecretModel{
		ID:       pb.GetId(),
		Data:     ToDecodedSecretDataModel(pb.GetData()),
		Metadata: common.ToMetadataModel(pb.GetMetadata()),
	}
}

func FromDecodedSecretModel(
	model secretsharedmd.DecodedSecretModel,
) *secretpb.DecodedSecret {
	return &secretpb.DecodedSecret{
		Id:       model.ID,
		Data:     FromDecodedSecretDataModel(model.Data),
		Metadata: common.FromMetadataModel(model.Metadata),
	}
}

func ToEncodedSecretModel(
	pb *secretpb.EncodedSecret,
) secretsharedmd.EncodedSecretModel {
	return secretsharedmd.EncodedSecretModel{
		ID:       pb.GetId(),
		Data:     ToEncodedSecretDataModel(pb.GetData()),
		Metadata: common.ToMetadataModel(pb.GetMetadata()),
	}
}

func FromEncodedSecretModel(
	model secretsharedmd.EncodedSecretModel,
) *secretpb.EncodedSecret {
	return &secretpb.EncodedSecret{
		Id:       model.ID,
		Data:     FromEncodedSecretDataModel(model.Data),
		Metadata: common.FromMetadataModel(model.Metadata),
	}
}

func ToOwnedDecodedSecretModel(
	pb *secretpb.DecodedSecret,
	userID int64,
) secretsharedmd.OwnedDecodedSecretModel {
	return secretsharedmd.OwnedDecodedSecretModel{
		ID:       pb.GetId(),
		Data:     ToOwnedDecodedSecretDataModel(pb.GetData(), userID),
		Metadata: common.ToMetadataModel(pb.GetMetadata()),
	}
}

func FromOwnedDecodedSecretModel(
	model secretsharedmd.OwnedDecodedSecretModel,
) *secretpb.DecodedSecret {
	return &secretpb.DecodedSecret{
		Id:       model.ID,
		Data:     FromOwnedDecodedSecretDataModel(model.Data),
		Metadata: common.FromMetadataModel(model.Metadata),
	}
}

func ToOwnedEncodedSecretModel(
	pb *secretpb.EncodedSecret,
	userID int64,
) secretsharedmd.OwnedEncodedSecretModel {
	return secretsharedmd.OwnedEncodedSecretModel{
		ID:       pb.GetId(),
		Data:     ToOwnedEncodedSecretDataModel(pb.GetData(), userID),
		Metadata: common.ToMetadataModel(pb.GetMetadata()),
	}
}

func FromOwnedEncodedSecretModel(
	model secretsharedmd.OwnedEncodedSecretModel,
) *secretpb.EncodedSecret {
	return &secretpb.EncodedSecret{
		Id:       model.ID,
		Data:     FromOwnedEncodedSecretDataModel(model.Data),
		Metadata: common.FromMetadataModel(model.Metadata),
	}
}

/* __________________________________________________ */

func ToDecodedSecretDataModel(
	pb *secretpb.DecodedSecretData,
) secretsharedmd.DecodedSecretDataModel {
	return secretsharedmd.DecodedSecretDataModel{
		Name:    pb.GetName(),
		Type:    pb.GetType(),
		Content: ToSecretContentModel(pb.GetContent()),
	}
}

func FromDecodedSecretDataModel(
	model secretsharedmd.DecodedSecretDataModel,
) *secretpb.DecodedSecretData {
	return &secretpb.DecodedSecretData{
		Name:    model.Name,
		Type:    model.Type,
		Content: FromSecretContentModel(model.Content),
	}
}

func ToEncodedSecretDataModel(
	pb *secretpb.EncodedSecretData,
) secretsharedmd.EncodedSecretDataModel {
	return secretsharedmd.EncodedSecretDataModel{
		Name:    pb.GetName(),
		Type:    pb.GetType(),
		Content: pb.GetContent(),
	}
}

func FromEncodedSecretDataModel(
	model secretsharedmd.EncodedSecretDataModel,
) *secretpb.EncodedSecretData {
	return &secretpb.EncodedSecretData{
		Name:    model.Name,
		Type:    model.Type,
		Content: model.Content,
	}
}

func ToOwnedDecodedSecretDataModel(
	pb *secretpb.DecodedSecretData,
	userID int64,
) secretsharedmd.OwnedDecodedSecretDataModel {
	return secretsharedmd.OwnedDecodedSecretDataModel{
		UserID:  userID,
		Name:    pb.GetName(),
		Type:    pb.GetType(),
		Content: ToSecretContentModel(pb.GetContent()),
	}
}

func FromOwnedDecodedSecretDataModel(
	model secretsharedmd.OwnedDecodedSecretDataModel,
) *secretpb.DecodedSecretData {
	return &secretpb.DecodedSecretData{
		Name:    model.Name,
		Type:    model.Type,
		Content: FromSecretContentModel(model.Content),
	}
}

func ToOwnedEncodedSecretDataModel(
	pb *secretpb.EncodedSecretData,
	userID int64,
) secretsharedmd.OwnedEncodedSecretDataModel {
	return secretsharedmd.OwnedEncodedSecretDataModel{
		UserID:  userID,
		Name:    pb.GetName(),
		Type:    pb.GetType(),
		Content: pb.GetContent(),
	}
}

func FromOwnedEncodedSecretDataModel(
	model secretsharedmd.OwnedEncodedSecretDataModel,
) *secretpb.EncodedSecretData {
	return &secretpb.EncodedSecretData{
		Name:    model.Name,
		Type:    model.Type,
		Content: model.Content,
	}
}

/* __________________________________________________ */

func ToSecretContentModel(
	pb *secretpb.SecretContent,
) secretsharedmd.SecretContentModel {

	switch kind := pb.GetKind().(type) {
	case *secretpb.SecretContent_Text:
		return &secretsharedmd.TextSecretContentModel{
			Text: kind.Text.GetText(),
		}
	case *secretpb.SecretContent_Binary:
		return &secretsharedmd.BinarySecretContentModel{
			Bytes: kind.Binary.GetBytes(),
		}
	case *secretpb.SecretContent_Credentials:
		return &secretsharedmd.CredentialsSecretContentModel{
			Username: kind.Credentials.GetUsername(),
			Password: kind.Credentials.GetPassword(),
		}
	case *secretpb.SecretContent_Card:
		return &secretsharedmd.CardSecretContentModel{
			Num: kind.Card.GetNum(),
			CVV: kind.Card.GetCvv(),
			Due: kind.Card.GetDue(),
		}
	}

	panic("invalid secret content type")

}

func FromSecretContentModel(
	model secretsharedmd.SecretContentModel,
) *secretpb.SecretContent {

	switch kind := model.(type) {
	case *secretsharedmd.TextSecretContentModel:
		return &secretpb.SecretContent{
			Kind: &secretpb.SecretContent_Text{
				Text: &secretpb.TextSecretContent{
					Text: kind.Text,
				},
			},
		}
	case *secretsharedmd.BinarySecretContentModel:
		return &secretpb.SecretContent{
			Kind: &secretpb.SecretContent_Binary{
				Binary: &secretpb.BinarySecretContent{
					Bytes: kind.Bytes,
				},
			},
		}
	case *secretsharedmd.CredentialsSecretContentModel:
		return &secretpb.SecretContent{
			Kind: &secretpb.SecretContent_Credentials{
				Credentials: &secretpb.CredentialsSecretContent{
					Username: kind.Username,
					Password: kind.Password,
				},
			},
		}

	case *secretsharedmd.CardSecretContentModel:
		return &secretpb.SecretContent{
			Kind: &secretpb.SecretContent_Card{
				Card: &secretpb.CardSecretContent{
					Num: kind.Num,
					Cvv: kind.CVV,
					Due: kind.Due,
				},
			},
		}
	}

	panic("invalid secret content type")

}

/* __________________________________________________ */
