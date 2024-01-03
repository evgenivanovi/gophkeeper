package secret

import (
	"errors"
	"strings"

	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/common"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type SecretID struct {
	id int64
}

func NewSecretID(id int64) SecretID {
	return SecretID{
		id: id,
	}
}

func (s SecretID) ID() int64 {
	return s.id
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type DecodedSecret struct {
	id       SecretID
	data     DecodedSecretData
	metadata core.Metadata
}

func NewDecodedSecret(
	id SecretID,
	data DecodedSecretData,
	metadata core.Metadata,
) *DecodedSecret {
	return &DecodedSecret{
		id:       id,
		data:     data,
		metadata: metadata,
	}
}

func NewEmptyDecodedSecret() DecodedSecret {
	return DecodedSecret{}
}

func (e *DecodedSecret) Identity() SecretID {
	return e.id
}

func (e *DecodedSecret) Data() DecodedSecretData {
	return e.data
}

func (e *DecodedSecret) Metadata() core.Metadata {
	return e.metadata
}

func (e *DecodedSecret) ToEncodedSecret(data EncodedSecretData) EncodedSecret {
	return *NewEncodedSecret(e.id, data, e.metadata)
}

func ToDecodedSecretPointers(entities []DecodedSecret) []*DecodedSecret {
	result := make([]*DecodedSecret, 0)
	for _, entity := range entities {
		result = append(result, &entity)
	}
	return result
}

func ToDecodedSecretValues(entities []*DecodedSecret) []DecodedSecret {
	result := make([]DecodedSecret, 0)
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type OwnedDecodedSecret struct {
	id       SecretID
	data     OwnedDecodedSecretData
	metadata core.Metadata
}

func NewOwnedDecodedSecret(
	id SecretID,
	data OwnedDecodedSecretData,
	metadata core.Metadata,
) *OwnedDecodedSecret {
	return &OwnedDecodedSecret{
		id:       id,
		data:     data,
		metadata: metadata,
	}
}

func NewEmptyOwnedDecodedSecret() OwnedDecodedSecret {
	return OwnedDecodedSecret{}
}

func (e *OwnedDecodedSecret) Identity() SecretID {
	return e.id
}

func (e *OwnedDecodedSecret) Data() OwnedDecodedSecretData {
	return e.data
}

func (e *OwnedDecodedSecret) Metadata() core.Metadata {
	return e.metadata
}

func (e *OwnedDecodedSecret) ToDecodedSecret() DecodedSecret {
	return *NewDecodedSecret(e.id, e.data.ToDecodedSecretData(), e.metadata)
}

func (e *OwnedDecodedSecret) ToEncodedSecret(data EncodedSecretData) EncodedSecret {
	return *NewEncodedSecret(e.id, data, e.metadata)
}

func (e *OwnedDecodedSecret) ToOwnedEncodedSecret(data OwnedEncodedSecretData) OwnedEncodedSecret {
	return *NewOwnedEncodedSecret(e.id, data, e.metadata)
}

func ToOwnedDecodedSecretPointers(entities []OwnedDecodedSecret) []*OwnedDecodedSecret {
	result := make([]*OwnedDecodedSecret, 0)
	for _, entity := range entities {
		result = append(result, &entity)
	}
	return result
}

func ToOwnedDecodedSecretValues(entities []*OwnedDecodedSecret) []OwnedDecodedSecret {
	result := make([]OwnedDecodedSecret, 0)
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type EncodedSecret struct {
	id       SecretID
	data     EncodedSecretData
	metadata core.Metadata
}

func NewEncodedSecret(
	id SecretID,
	data EncodedSecretData,
	metadata core.Metadata,
) *EncodedSecret {
	return &EncodedSecret{
		id:       id,
		data:     data,
		metadata: metadata,
	}
}

func NewEmptyEncodedSecret() EncodedSecret {
	return EncodedSecret{}
}

func (e *EncodedSecret) Identity() SecretID {
	return e.id
}

func (e *EncodedSecret) Data() EncodedSecretData {
	return e.data
}

func (e *EncodedSecret) Metadata() core.Metadata {
	return e.metadata
}

func (e *EncodedSecret) ToDecodedSecret(data DecodedSecretData) DecodedSecret {
	return *NewDecodedSecret(e.id, data, e.metadata)
}

func ToEncodedSecretPointers(entities []EncodedSecret) []*EncodedSecret {
	result := make([]*EncodedSecret, 0)
	for _, entity := range entities {
		result = append(result, &entity)
	}
	return result
}

func ToEncodedSecretValues(entities []*EncodedSecret) []EncodedSecret {
	result := make([]EncodedSecret, 0)
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type OwnedEncodedSecret struct {
	id       SecretID
	data     OwnedEncodedSecretData
	metadata core.Metadata
}

func NewOwnedEncodedSecret(
	id SecretID,
	data OwnedEncodedSecretData,
	metadata core.Metadata,
) *OwnedEncodedSecret {
	return &OwnedEncodedSecret{
		id:       id,
		data:     data,
		metadata: metadata,
	}
}

func NewEmptyOwnedEncodedSecret() OwnedEncodedSecret {
	return OwnedEncodedSecret{}
}

func (e *OwnedEncodedSecret) Identity() SecretID {
	return e.id
}

func (e *OwnedEncodedSecret) Data() OwnedEncodedSecretData {
	return e.data
}

func (e *OwnedEncodedSecret) Metadata() core.Metadata {
	return e.metadata
}

func (e *OwnedEncodedSecret) ToEncodedSecret() EncodedSecret {
	return *NewEncodedSecret(e.id, e.data.ToEncodedSecretData(), e.metadata)
}

func (e *OwnedEncodedSecret) ToDecodedSecret(data DecodedSecretData) DecodedSecret {
	return *NewDecodedSecret(e.id, data, e.metadata)
}

func (e *OwnedEncodedSecret) ToOwnedDecodedSecret(data OwnedDecodedSecretData) OwnedDecodedSecret {
	return *NewOwnedDecodedSecret(e.id, data, e.metadata)
}

func ToOwnedEncodedSecretPointers(entities []OwnedEncodedSecret) []*OwnedEncodedSecret {
	result := make([]*OwnedEncodedSecret, 0)
	for _, entity := range entities {
		result = append(result, &entity)
	}
	return result
}

func ToOwnedEncodedSecretValues(entities []*OwnedEncodedSecret) []OwnedEncodedSecret {
	result := make([]OwnedEncodedSecret, 0)
	for _, entity := range entities {
		result = append(result, *entity)
	}
	return result
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type DecodedSecretData struct {
	Name    string
	Secret  SecretType
	Content SecretContent
}

func NewDecodedSecretData(
	name string,
	secret SecretType,
	content SecretContent,
) *DecodedSecretData {
	return &DecodedSecretData{
		Name:    name,
		Secret:  secret,
		Content: content,
	}
}

func (d *DecodedSecretData) ToEncodedSecretData(
	content []byte,
) EncodedSecretData {
	return *NewEncodedSecretData(
		d.Name,
		d.Secret,
		content,
	)
}

//goland:noinspection GoNameStartsWithPackageName
type OwnedDecodedSecretData struct {
	UserID common.UserID

	Name    string
	Secret  SecretType
	Content SecretContent
}

func NewOwnedDecodedSecretData(
	userID common.UserID,
	name string,
	secret SecretType,
	content SecretContent,
) *OwnedDecodedSecretData {
	return &OwnedDecodedSecretData{
		UserID:  userID,
		Name:    name,
		Secret:  secret,
		Content: content,
	}
}

func (d *OwnedDecodedSecretData) ToDecodedSecretData() DecodedSecretData {
	return *NewDecodedSecretData(
		d.Name,
		d.Secret,
		d.Content,
	)
}

func (d *OwnedDecodedSecretData) ToEncodedSecretData(
	content []byte,
) EncodedSecretData {
	return *NewEncodedSecretData(
		d.Name,
		d.Secret,
		content,
	)
}

func (d *OwnedDecodedSecretData) ToOwnedEncodedSecretData(
	content []byte,
) OwnedEncodedSecretData {
	return *NewOwnedEncodedSecretData(
		d.UserID,
		d.Name,
		d.Secret,
		content,
	)
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type EncodedSecretData struct {
	Name    string
	Secret  SecretType
	Content []byte
}

func NewEncodedSecretData(
	name string,
	secret SecretType,
	content []byte,
) *EncodedSecretData {
	return &EncodedSecretData{
		Name:    name,
		Secret:  secret,
		Content: content,
	}
}

func (d *EncodedSecretData) ToDecodedSecretData(
	content SecretContent,
) DecodedSecretData {
	return *NewDecodedSecretData(
		d.Name,
		d.Secret,
		content,
	)
}

//goland:noinspection GoNameStartsWithPackageName
type OwnedEncodedSecretData struct {
	UserID common.UserID

	Name    string
	Secret  SecretType
	Content []byte
}

func NewOwnedEncodedSecretData(
	userID common.UserID,
	name string,
	secret SecretType,
	content []byte,
) *OwnedEncodedSecretData {
	return &OwnedEncodedSecretData{
		UserID:  userID,
		Name:    name,
		Secret:  secret,
		Content: content,
	}
}

func (d *OwnedEncodedSecretData) ToEncodedSecretData() EncodedSecretData {
	return *NewEncodedSecretData(
		d.Name,
		d.Secret,
		d.Content,
	)
}

func (d *OwnedEncodedSecretData) ToDecodedSecretData(
	content SecretContent,
) DecodedSecretData {
	return *NewDecodedSecretData(
		d.Name,
		d.Secret,
		content,
	)
}

func (d *OwnedEncodedSecretData) ToOwnedDecodedSecretData(
	content SecretContent,
) OwnedDecodedSecretData {
	return *NewOwnedDecodedSecretData(
		d.UserID,
		d.Name,
		d.Secret,
		content,
	)
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type SecretContent interface {
	SecretType() SecretType
}

type TextSecretContent struct {
	Text string
}

func NewTextSecretContent(text string) *TextSecretContent {
	return &TextSecretContent{
		Text: text,
	}
}

func (s *TextSecretContent) SecretType() SecretType {
	return Text
}

type BinarySecretContent struct {
	Bytes []byte
}

func NewBinarySecretContent(bytes []byte) *BinarySecretContent {
	return &BinarySecretContent{
		Bytes: bytes,
	}
}

func (s *BinarySecretContent) SecretType() SecretType {
	return Binary
}

type CredentialsSecretContent struct {
	Username string
	Password string
}

func NewCredentialsSecretContent(username string, password string) *CredentialsSecretContent {
	return &CredentialsSecretContent{
		Username: username,
		Password: password,
	}
}

func (s *CredentialsSecretContent) SecretType() SecretType {
	return Credentials
}

type CardSecretContent struct {
	Num string
	CVV string
	Due string
}

func NewCardSecretContent(num string, cvv string, due string) *CardSecretContent {
	return &CardSecretContent{
		Num: num,
		CVV: cvv,
		Due: due,
	}
}

func (s *CardSecretContent) SecretType() SecretType {
	return Card
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type SecretType string

const (
	TextName        string = "TEXT"
	BinaryName      string = "BINARY"
	CredentialsName string = "CREDENTIALS"
	CardName        string = "CARD"
)

const (
	Text        = SecretType(TextName)
	Binary      = SecretType(BinaryName)
	Credentials = SecretType(CredentialsName)
	Card        = SecretType(CardName)
)

func (el SecretType) String() string {
	return string(el)
}

func (el SecretType) MarshalText() ([]byte, error) {
	return []byte(el.String()), nil
}

func TypeFromString(value string) (SecretType, error) {
	switch strings.ToUpper(value) {
	default:
		return "", errors.New("invalid secret type")
	case TextName:
		return Text, nil
	case BinaryName:
		return Binary, nil
	case CredentialsName:
		return Credentials, nil
	case CardName:
		return Card, nil
	}
}

func MustTypeFromString(value string) SecretType {
	kind, err := TypeFromString(value)
	if err != nil {
		panic(err)
	}
	return kind
}

/* __________________________________________________ */
