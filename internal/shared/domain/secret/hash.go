package secret

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type OwnedSecretEncoderDecoder interface {
	Decode(sec OwnedEncodedSecret) (OwnedDecodedSecret, error)
	Encode(sec OwnedDecodedSecret) (OwnedEncodedSecret, error)
}

type OwnedSecretEncoderDecoderService struct {
	enc SecretContentEncoder
	dec SecretContentDecoder
}

func ProvideOwnedSecretEncoderDecoderService(
	enc SecretContentEncoder,
	dec SecretContentDecoder,
) *OwnedSecretEncoderDecoderService {
	return &OwnedSecretEncoderDecoderService{
		enc: enc,
		dec: dec,
	}
}

func (svc *OwnedSecretEncoderDecoderService) Decode(
	sec OwnedEncodedSecret,
) (OwnedDecodedSecret, error) {

	bytes := sec.Data().Content
	content, err := svc.dec.Decode(bytes, sec.Data().Secret)
	if err != nil {
		return NewEmptyOwnedDecodedSecret(), err
	}

	data := sec.Data()
	decoded := sec.ToOwnedDecodedSecret(data.ToOwnedDecodedSecretData(content))
	return decoded, nil

}

func (svc *OwnedSecretEncoderDecoderService) Encode(
	sec OwnedDecodedSecret,
) (OwnedEncodedSecret, error) {

	content, err := svc.enc.Encode(sec.Data().Content)
	if err != nil {
		return NewEmptyOwnedEncodedSecret(), err
	}

	data := sec.Data()
	encoded := sec.ToOwnedEncodedSecret(data.ToOwnedEncodedSecretData(content))
	return encoded, nil

}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type SecretEncoderDecoder interface {
	Decode(sec EncodedSecret) (DecodedSecret, error)
	Encode(sec DecodedSecret) (EncodedSecret, error)
}

//goland:noinspection GoNameStartsWithPackageName
type SecretEncoderDecoderService struct {
	enc SecretContentEncoder
	dec SecretContentDecoder
}

func ProvideSecretEncoderDecoderService(
	enc SecretContentEncoder,
	dec SecretContentDecoder,
) *SecretEncoderDecoderService {
	return &SecretEncoderDecoderService{
		enc: enc,
		dec: dec,
	}
}

func (svc *SecretEncoderDecoderService) Decode(
	sec EncodedSecret,
) (DecodedSecret, error) {

	bytes := sec.Data().Content
	content, err := svc.dec.Decode(bytes, sec.Data().Secret)
	if err != nil {
		return NewEmptyDecodedSecret(), err
	}

	data := sec.Data()
	decoded := sec.ToDecodedSecret(data.ToDecodedSecretData(content))
	return decoded, nil

}

func (svc *SecretEncoderDecoderService) Encode(
	sec DecodedSecret,
) (EncodedSecret, error) {

	content, err := svc.enc.Encode(sec.Data().Content)
	if err != nil {
		return NewEmptyEncodedSecret(), err
	}

	data := sec.Data()
	encoded := sec.ToEncodedSecret(data.ToEncodedSecretData(content))
	return encoded, nil

}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type SecretContentEncoder interface {
	Encode(secret SecretContent) ([]byte, error)
}

//goland:noinspection GoNameStartsWithPackageName
type SecretContentDecoder interface {
	Decode(data []byte, kind SecretType) (SecretContent, error)
}

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
type TextSecretContentDecoder interface {
	Decode(data []byte) (TextSecretContent, error)
}

//goland:noinspection GoNameStartsWithPackageName
type BinarySecretContentDecoder interface {
	Decode(data []byte) (BinarySecretContent, error)
}

//goland:noinspection GoNameStartsWithPackageName
type CredentialsSecretContentDecoder interface {
	Decode(data []byte) (CredentialsSecretContent, error)
}

//goland:noinspection GoNameStartsWithPackageName
type CardSecretContentDecoder interface {
	Decode(data []byte) (CardSecretContent, error)
}

/* __________________________________________________ */
