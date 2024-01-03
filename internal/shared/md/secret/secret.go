package secret

import (
	"strings"

	"github.com/evgenivanovi/gophkeeper/internal/shared/md/common"
	"github.com/evgenivanovi/gpl/std"
)

/* __________________________________________________ */

type DecodedSecretDataModel struct {
	Name    string
	Type    string
	Content SecretContentModel
}

type OwnedDecodedSecretDataModel struct {
	UserID  int64
	Name    string
	Type    string
	Content SecretContentModel
}

type EncodedSecretDataModel struct {
	Name    string
	Type    string
	Content []byte
}

type OwnedEncodedSecretDataModel struct {
	UserID  int64
	Name    string
	Type    string
	Content []byte
}

/* __________________________________________________ */

type DecodedSecretModel struct {
	ID       int64
	Data     DecodedSecretDataModel
	Metadata common.MetadataModel
}

type OwnedDecodedSecretModel struct {
	ID       int64
	Data     OwnedDecodedSecretDataModel
	Metadata common.MetadataModel
}

type EncodedSecretModel struct {
	ID       int64
	Data     EncodedSecretDataModel
	Metadata common.MetadataModel
}

type OwnedEncodedSecretModel struct {
	ID       int64
	Data     OwnedEncodedSecretDataModel
	Metadata common.MetadataModel
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type SecretContentModel interface {
	SecretType() string
	String() string
}

//goland:noinspection GoNameStartsWithPackageName
type TextSecretContentModel struct {
	Text string
}

func (s *TextSecretContentModel) SecretType() string {
	return "TEXT"
}

func (s *TextSecretContentModel) String() string {
	return s.Text
}

//goland:noinspection GoNameStartsWithPackageName
type BinarySecretContentModel struct {
	Bytes []byte
}

func (s *BinarySecretContentModel) SecretType() string {
	return "BINARY"
}

func (s *BinarySecretContentModel) String() string {
	return string(s.Bytes)
}

//goland:noinspection GoNameStartsWithPackageName
type CredentialsSecretContentModel struct {
	Username string
	Password string
}

func (s *CredentialsSecretContentModel) SecretType() string {
	return "CREDENTIALS"
}

func (s *CredentialsSecretContentModel) String() string {
	var output strings.Builder
	output.WriteString("Username: " + s.Username)
	output.WriteString(std.NL)
	output.WriteString("Password: " + s.Password)
	return output.String()
}

//goland:noinspection GoNameStartsWithPackageName
type CardSecretContentModel struct {
	Num string
	CVV string
	Due string
}

func (s *CardSecretContentModel) SecretType() string {
	return "CARD"
}

func (s *CardSecretContentModel) String() string {
	var output strings.Builder
	output.WriteString("Num: " + s.Num)
	output.WriteString(std.NL)
	output.WriteString("CVV: " + s.CVV)
	output.WriteString(std.NL)
	output.WriteString("Due: " + s.Due)
	return output.String()
}

/* __________________________________________________ */
