package secret

import (
	"bytes"
	"encoding/gob"
)

/* __________________________________________________ */

type GobSecretContentEncoder struct{}

func ProvideGobSecretContentEncoder() GobSecretContentEncoder {
	return GobSecretContentEncoder{}
}

func (svc GobSecretContentEncoder) Encode(secret SecretContent) ([]byte, error) {
	buf := bytes.Buffer{}

	err := gob.NewEncoder(&buf).Encode(secret)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

/* __________________________________________________ */

type GobSecretContentDecoder struct {
	textDec        GobTextSecretContentDecoder
	binaryDec      GobBinarySecretContentDecoder
	credentialsDec GobCredentialsSecretContentDecoder
	cardDec        GobCardSecretContentDecoder
}

func ProvideGobSecretContentDecoder() *GobSecretContentDecoder {
	return &GobSecretContentDecoder{
		textDec:        ProvideGobTextSecretContentDecoder(),
		binaryDec:      ProvideGobBinarySecretContentDecoder(),
		credentialsDec: ProvideGobCredentialsSecretContentDecoder(),
		cardDec:        ProvideGobCardSecretContentDecoder(),
	}
}

func (svc *GobSecretContentDecoder) Decode(
	data []byte, kind SecretType,
) (SecretContent, error) {

	switch kind.String() {
	case TextName:
		{
			content, err := svc.textDec.Decode(data)
			if err != nil {
				return nil, err
			}
			return &content, nil
		}
	case BinaryName:
		{
			content, err := svc.binaryDec.Decode(data)
			if err != nil {
				return nil, err
			}
			return &content, nil
		}
	case CredentialsName:
		{
			content, err := svc.credentialsDec.Decode(data)
			if err != nil {
				return nil, err
			}
			return &content, nil
		}
	case CardName:
		{
			content, err := svc.cardDec.Decode(data)
			if err != nil {
				return nil, err
			}
			return &content, nil
		}
	}

	panic("unknown secret type")

}

/* __________________________________________________ */

type GobTextSecretContentDecoder struct{}

func ProvideGobTextSecretContentDecoder() GobTextSecretContentDecoder {
	return GobTextSecretContentDecoder{}
}

func (svc GobTextSecretContentDecoder) Decode(
	data []byte,
) (TextSecretContent, error) {
	res := TextSecretContent{}

	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&res)
	if err != nil {
		return TextSecretContent{}, err
	}

	return res, nil
}

/* __________________________________________________ */

type GobBinarySecretContentDecoder struct{}

func ProvideGobBinarySecretContentDecoder() GobBinarySecretContentDecoder {
	return GobBinarySecretContentDecoder{}
}

func (svc GobBinarySecretContentDecoder) Decode(
	data []byte,
) (BinarySecretContent, error) {
	res := BinarySecretContent{}

	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&res)
	if err != nil {
		return BinarySecretContent{}, err
	}

	return res, nil
}

/* __________________________________________________ */

type GobCredentialsSecretContentDecoder struct{}

func ProvideGobCredentialsSecretContentDecoder() GobCredentialsSecretContentDecoder {
	return GobCredentialsSecretContentDecoder{}
}

func (svc GobCredentialsSecretContentDecoder) Decode(
	data []byte,
) (CredentialsSecretContent, error) {
	res := CredentialsSecretContent{}

	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&res)
	if err != nil {
		return CredentialsSecretContent{}, err
	}

	return res, nil
}

/* __________________________________________________ */

type GobCardSecretContentDecoder struct{}

func ProvideGobCardSecretContentDecoder() GobCardSecretContentDecoder {
	return GobCardSecretContentDecoder{}
}

func (svc GobCardSecretContentDecoder) Decode(
	data []byte,
) (CardSecretContent, error) {
	res := CardSecretContent{}

	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&res)
	if err != nil {
		return CardSecretContent{}, err
	}

	return res, nil
}

/* __________________________________________________ */
