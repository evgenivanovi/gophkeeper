package config

import "os"

/* __________________________________________________ */

type Reader interface {
	Read(path string) (ConfigObject, error)
}

type Writer interface {
	Write(target ConfigObject, path string) error
}

/* __________________________________________________ */

type ReaderService struct {
	parser Parser
}

func ProvideReaderService(parser Parser) *ReaderService {
	return &ReaderService{
		parser: parser,
	}
}

func (r *ReaderService) Read(path string) (ConfigObject, error) {

	in, err := os.ReadFile(path)
	if err != nil {
		return ConfigObject{}, err
	}

	origin, err := r.parser.Unmarshal(in)
	if err != nil {
		return ConfigObject{}, err
	}

	return origin, nil

}

/* __________________________________________________ */

type WriterService struct {
	parser Parser
}

func ProvideWriterService(parser Parser) *WriterService {
	return &WriterService{
		parser: parser,
	}
}

func (w WriterService) Write(target ConfigObject, path string) error {

	out, err := w.parser.Marshal(target)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, out, os.ModePerm)
	if err != nil {
		return err
	}

	return nil

}

/* __________________________________________________ */
