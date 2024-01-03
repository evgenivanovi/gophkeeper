package secret

import (
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	secretshareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	commonsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/common"
)

/* __________________________________________________ */

func ToDecodedSecret(
	model DecodedSecretModel,
) secretshareddm.DecodedSecret {
	id := secretshareddm.NewSecretID(model.ID)
	data := ToDecodedSecretData(model.Data)
	metadata := commonsharedmd.ToMetadata(model.Metadata)
	return *secretshareddm.NewDecodedSecret(id, data, metadata)
}

func FromDecodedSecret(
	entity secretshareddm.DecodedSecret,
) DecodedSecretModel {
	return DecodedSecretModel{
		ID:       entity.Identity().ID(),
		Data:     FromDecodedSecretData(entity.Data()),
		Metadata: commonsharedmd.FromMetadata(entity.Metadata()),
	}
}

func ToOwnedDecodedSecret(
	model OwnedDecodedSecretModel,
) secretshareddm.OwnedDecodedSecret {
	id := secretshareddm.NewSecretID(model.ID)
	data := ToOwnedDecodedSecretData(model.Data)
	metadata := commonsharedmd.ToMetadata(model.Metadata)
	return *secretshareddm.NewOwnedDecodedSecret(id, data, metadata)
}

func FromOwnedDecodedSecret(
	entity secretshareddm.OwnedDecodedSecret,
) OwnedDecodedSecretModel {
	return OwnedDecodedSecretModel{
		ID:       entity.Identity().ID(),
		Data:     FromOwnedDecodedSecretData(entity.Data()),
		Metadata: commonsharedmd.FromMetadata(entity.Metadata()),
	}
}

/* __________________________________________________ */

func ToEncodedSecret(
	model EncodedSecretModel,
) secretshareddm.EncodedSecret {
	id := secretshareddm.NewSecretID(model.ID)
	data := ToEncodedSecretData(model.Data)
	metadata := commonsharedmd.ToMetadata(model.Metadata)
	return *secretshareddm.NewEncodedSecret(id, data, metadata)
}

func FromEncodedSecret(
	entity secretshareddm.EncodedSecret,
) EncodedSecretModel {
	return EncodedSecretModel{
		ID:       entity.Identity().ID(),
		Data:     FromEncodedSecretData(entity.Data()),
		Metadata: commonsharedmd.FromMetadata(entity.Metadata()),
	}
}

func ToOwnedEncodedSecret(
	model OwnedEncodedSecretModel,
) secretshareddm.OwnedEncodedSecret {
	id := secretshareddm.NewSecretID(model.ID)
	data := ToOwnedEncodedSecretData(model.Data)
	metadata := commonsharedmd.ToMetadata(model.Metadata)
	return *secretshareddm.NewOwnedEncodedSecret(id, data, metadata)
}

func FromOwnedEncodedSecret(
	entity secretshareddm.OwnedEncodedSecret,
) OwnedEncodedSecretModel {
	return OwnedEncodedSecretModel{
		ID:       entity.Identity().ID(),
		Data:     FromOwnedEncodedSecretData(entity.Data()),
		Metadata: commonsharedmd.FromMetadata(entity.Metadata()),
	}
}

/* __________________________________________________ */

func ToDecodedSecretData(
	model DecodedSecretDataModel,
) secretshareddm.DecodedSecretData {
	return secretshareddm.DecodedSecretData{
		Name:    model.Name,
		Secret:  secretshareddm.MustTypeFromString(model.Type),
		Content: ToSecretContent(model.Content),
	}
}

func FromDecodedSecretData(
	value secretshareddm.DecodedSecretData,
) DecodedSecretDataModel {
	return DecodedSecretDataModel{
		Name:    value.Name,
		Type:    value.Secret.String(),
		Content: FromSecretContent(value.Content),
	}
}

func ToOwnedDecodedSecretData(
	model OwnedDecodedSecretDataModel,
) secretshareddm.OwnedDecodedSecretData {
	return secretshareddm.OwnedDecodedSecretData{
		UserID:  common.NewUserID(model.UserID),
		Name:    model.Name,
		Secret:  secretshareddm.MustTypeFromString(model.Type),
		Content: ToSecretContent(model.Content),
	}
}

func FromOwnedDecodedSecretData(
	value secretshareddm.OwnedDecodedSecretData,
) OwnedDecodedSecretDataModel {
	return OwnedDecodedSecretDataModel{
		UserID:  value.UserID.ID(),
		Name:    value.Name,
		Type:    value.Secret.String(),
		Content: FromSecretContent(value.Content),
	}
}

/* __________________________________________________ */

func ToEncodedSecretData(
	model EncodedSecretDataModel,
) secretshareddm.EncodedSecretData {
	return secretshareddm.EncodedSecretData{
		Name:    model.Name,
		Secret:  secretshareddm.MustTypeFromString(model.Type),
		Content: model.Content,
	}
}

func FromEncodedSecretData(
	value secretshareddm.EncodedSecretData,
) EncodedSecretDataModel {
	return EncodedSecretDataModel{
		Name:    value.Name,
		Type:    value.Secret.String(),
		Content: value.Content,
	}
}

func ToOwnedEncodedSecretData(
	model OwnedEncodedSecretDataModel,
) secretshareddm.OwnedEncodedSecretData {
	return secretshareddm.OwnedEncodedSecretData{
		UserID:  common.NewUserID(model.UserID),
		Name:    model.Name,
		Secret:  secretshareddm.MustTypeFromString(model.Type),
		Content: model.Content,
	}
}

func FromOwnedEncodedSecretData(
	entity secretshareddm.OwnedEncodedSecretData,
) OwnedEncodedSecretDataModel {
	return OwnedEncodedSecretDataModel{
		UserID:  entity.UserID.ID(),
		Name:    entity.Name,
		Type:    entity.Secret.String(),
		Content: entity.Content,
	}
}

/* __________________________________________________ */

func ToSecretContent(
	model SecretContentModel,
) secretshareddm.SecretContent {

	switch kind := model.(type) {
	case *TextSecretContentModel:
		return secretshareddm.NewTextSecretContent(kind.Text)
	case *BinarySecretContentModel:
		return secretshareddm.NewBinarySecretContent(kind.Bytes)
	case *CredentialsSecretContentModel:
		return secretshareddm.NewCredentialsSecretContent(kind.Username, kind.Password)
	case *CardSecretContentModel:
		return secretshareddm.NewCardSecretContent(kind.Num, kind.CVV, kind.Due)
	}

	panic("invalid secret content type")

}

func FromSecretContent(
	value secretshareddm.SecretContent,
) SecretContentModel {

	switch kind := value.(type) {
	case *secretshareddm.TextSecretContent:
		return &TextSecretContentModel{
			Text: kind.Text,
		}
	case *secretshareddm.BinarySecretContent:
		return &BinarySecretContentModel{
			Bytes: kind.Bytes,
		}
	case *secretshareddm.CredentialsSecretContent:
		return &CredentialsSecretContentModel{
			Username: kind.Username,
			Password: kind.Password,
		}
	case *secretshareddm.CardSecretContent:
		return &CardSecretContentModel{
			Num: kind.Num,
			CVV: kind.CVV,
			Due: kind.Due,
		}
	}

	panic("invalid secret content type")

}

/* __________________________________________________ */
