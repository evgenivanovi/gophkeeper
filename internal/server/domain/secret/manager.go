package secret

import (
	"context"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	"github.com/evgenivanovi/gpl/search"
)

/* __________________________________________________ */

type Manager interface {
	CreateEncoded(
		ctx context.Context, data secret.OwnedEncodedSecretData,
	) (secret.OwnedEncodedSecret, error)

	CreateDecoded(
		ctx context.Context, data secret.OwnedDecodedSecretData,
	) (secret.OwnedEncodedSecret, error)

	UpdateEncoded(
		ctx context.Context, entity secret.OwnedEncodedSecret,
	) (secret.OwnedEncodedSecret, error)

	UpdateDecoded(
		ctx context.Context, entity secret.OwnedDecodedSecret,
	) (secret.OwnedEncodedSecret, error)

	RemoveByID(
		ctx context.Context, id secret.SecretID,
	) (secret.OwnedEncodedSecret, error)

	RemoveByName(
		ctx context.Context, name string,
	) (secret.OwnedEncodedSecret, error)
}

type ManagerService struct {
	repo Repository
	enc  secret.SecretContentEncoder
	dec  secret.SecretContentDecoder
}

func ProvideManagerService(
	repo Repository,
	enc secret.SecretContentEncoder,
	dec secret.SecretContentDecoder,
) *ManagerService {
	return &ManagerService{
		repo: repo,
		enc:  enc,
		dec:  dec,
	}
}

func (m *ManagerService) CreateEncoded(
	ctx context.Context, data secret.OwnedEncodedSecretData,
) (secret.OwnedEncodedSecret, error) {
	sec, err := m.repo.AutoSave(ctx, data, *core.NewNowMetadata())
	if err != nil {
		return secret.NewEmptyOwnedEncodedSecret(), err
	}
	return *sec, nil
}

func (m *ManagerService) CreateDecoded(
	ctx context.Context, data secret.OwnedDecodedSecretData,
) (secret.OwnedEncodedSecret, error) {
	enc, err := m.enc.Encode(data.Content)
	if err != nil {
		return secret.NewEmptyOwnedEncodedSecret(), err
	}
	return m.CreateEncoded(ctx, data.ToOwnedEncodedSecretData(enc))
}

func (m *ManagerService) UpdateEncoded(
	ctx context.Context, entity secret.OwnedEncodedSecret,
) (secret.OwnedEncodedSecret, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ManagerService) UpdateDecoded(
	ctx context.Context, entity secret.OwnedDecodedSecret,
) (secret.OwnedEncodedSecret, error) {
	enc, err := m.encodeSecret(entity)
	if err != nil {
		return secret.NewEmptyOwnedEncodedSecret(), err
	}
	return m.UpdateEncoded(ctx, enc)
}

func (m *ManagerService) RemoveByID(
	ctx context.Context, id secret.SecretID,
) (secret.OwnedEncodedSecret, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ManagerService) RemoveByName(
	ctx context.Context, name string,
) (secret.OwnedEncodedSecret, error) {
	//TODO implement me
	panic("implement me")
}

func (m *ManagerService) encodeSecretData(
	data secret.OwnedDecodedSecretData,
) (secret.OwnedEncodedSecretData, error) {
	enc, err := m.enc.Encode(data.Content)
	if err != nil {
		return secret.OwnedEncodedSecretData{}, err
	}
	return data.ToOwnedEncodedSecretData(enc), nil
}

func (m *ManagerService) encodeSecret(
	entity secret.OwnedDecodedSecret,
) (secret.OwnedEncodedSecret, error) {
	enc, err := m.encodeSecretData(entity.Data())
	if err != nil {
		return secret.NewEmptyOwnedEncodedSecret(), err
	}
	return entity.ToOwnedEncodedSecret(enc), nil
}

func (m *ManagerService) searchSessionByID(
	ctx context.Context,
	id secret.SecretID,
) (*secret.OwnedEncodedSecret, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(IdentityCondition(id))
	return m.repo.FindOneBySpec(ctx, spec)
}

func (m *ManagerService) searchSessionByName(
	ctx context.Context,
	name string,
) (*secret.OwnedEncodedSecret, error) {
	spec := search.
		NewSpecificationTemplate().
		WithSearch(NameCondition(name))
	return m.repo.FindOneBySpec(ctx, spec)
}

/* __________________________________________________ */
