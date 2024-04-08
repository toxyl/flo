package codec

import (
	"io"

	"github.com/toxyl/flo/errors"
	"github.com/toxyl/flo/utils"
)

type Codec struct {
	Name         string
	Decode       func(input io.Reader, output any) error
	DecodeString func(input string, output any) error
	Encode       func(input any, output io.Writer) error
	EncodeString func(input any) string
}

func NewCodec(
	name string,
	fnEncode func(input any, output io.Writer) error,
	fnDecode func(input io.Reader, output any) error,
) *Codec {
	c := &Codec{
		Name: name,
		Encode: func(source any, target io.Writer) error {
			return errors.ErrNoEncoderImplemented
		},
		EncodeString: func(input any) string {
			return "no encoder implemented for " + name
		},
		Decode: func(source io.Reader, target any) error {
			return errors.ErrNoDecoderImplemented
		},
		DecodeString: func(source string, target any) error {
			return errors.ErrNoDecoderImplemented
		},
	}
	if fnEncode != nil {
		c.Encode = func(source any, target io.Writer) error {
			switch t := target.(type) {
			case io.WriteCloser:
				defer t.Close()
			}
			return fnEncode(utils.Dereference(source), target)
		}
		c.EncodeString = func(input any) string {
			res := utils.NewStringIO("")
			if err := c.Encode(input, res); err != nil {
				return ""
			}
			return res.String()
		}
	}

	if fnDecode != nil {
		c.Decode = func(source io.Reader, target any) error {
			switch t := source.(type) {
			case io.ReadCloser:
				defer t.Close()
			}
			if !utils.IsPointer(target) {
				return errors.ErrMustBePointer(target)
			}
			return fnDecode(source, target)
		}
		c.DecodeString = func(source string, target any) error {
			if !utils.IsPointer(target) {
				return errors.ErrMustBePointer(target)
			}
			in := utils.NewStringIO(source)
			if err := c.Decode(in, target); err != nil {
				return err
			}
			return nil
		}
	}

	return c
}
