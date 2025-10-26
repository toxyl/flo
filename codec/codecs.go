package codec

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"io"
	"net/url"
	"strconv"

	"gopkg.in/yaml.v3"
)

var (
	GOB = NewCodec("gob",
		func(source any, target io.Writer) error {
			return gob.NewEncoder(target).Encode(source)
		},
		func(source io.Reader, target any) error {
			return gob.NewDecoder(source).Decode(target)
		},
	)
	GOBGZ = NewCodec("gobz",
		func(source any, target io.Writer) error {
			gzw := gzip.NewWriter(target)
			defer gzw.Close()
			return gob.NewEncoder(gzw).Encode(source)
		},
		func(source io.Reader, target any) error {
			gzr, err := gzip.NewReader(source)
			if err != nil {
				return err
			}
			defer gzr.Close()
			return gob.NewDecoder(gzr).Decode(target)
		},
	)
	STRINGGZ = NewCodec("stringsz",
		func(source any, target io.Writer) error {
			gzw := gzip.NewWriter(target)
			defer gzw.Close()
			_, err := gzw.Write([]byte(source.(string)))
			return err
		},
		func(source io.Reader, target any) error {
			gzr, err := gzip.NewReader(source)
			if err != nil {
				return err
			}
			defer gzr.Close()

			buf, err := io.ReadAll(gzr)
			if err != nil {
				return err
			}
			*(target.(*string)) = string(buf)
			return err
		},
	)
	BYTESGZ = NewCodec("bytesz",
		func(source any, target io.Writer) error {
			gzw := gzip.NewWriter(target)
			defer gzw.Close()
			_, err := gzw.Write(source.([]byte))
			return err
		},
		func(source io.Reader, target any) error {
			gzr, err := gzip.NewReader(source)
			if err != nil {
				return err
			}
			defer gzr.Close()

			buf, err := io.ReadAll(gzr)
			if err != nil {
				return err
			}
			*(target.(*[]byte)) = buf
			return err
		},
	)
	YAML = NewCodec("yaml",
		func(source any, target io.Writer) error {
			return yaml.NewEncoder(target).Encode(source)
		},
		func(source io.Reader, target any) error {
			return yaml.NewDecoder(source).Decode(target)
		},
	)
	JSON = NewCodec("json",
		func(source any, target io.Writer) error {
			return json.NewEncoder(target).Encode(source)
		},
		func(source io.Reader, target any) error {
			return json.NewDecoder(source).Decode(target)
		},
	)
	STRING = NewCodec("string",
		func(source any, target io.Writer) error {
			_, err := target.Write([]byte(source.(string)))
			return err
		},
		func(source io.Reader, target any) error {
			var buf bytes.Buffer
			_, err := io.Copy(&buf, source)
			if err != nil {
				return err
			}
			*(target.(*string)) = buf.String()
			return nil
		},
	)
	BYTES = NewCodec("bytes",
		func(source any, target io.Writer) error {
			_, err := target.Write(source.([]byte))
			return err
		},
		func(source io.Reader, target any) error {
			var buf bytes.Buffer
			_, err := io.Copy(&buf, source)
			if err != nil {
				return err
			}
			*(target.(*[]byte)) = buf.Bytes()
			return nil
		},
	)
	BASE64_URL = NewCodec("b64_url",
		func(source any, target io.Writer) error {
			encoder := base64.NewEncoder(base64.RawURLEncoding, target)
			defer encoder.Close()
			_, err := encoder.Write([]byte(fmt.Sprint(source)))
			return err
		},
		func(source io.Reader, target any) error {
			decoder := base64.NewDecoder(base64.RawURLEncoding, source)
			decoded, err := io.ReadAll(decoder)
			if err != nil {
				return err
			}
			*(target.(*string)) = string(decoded)
			return nil
		},
	)
	BASE64_STD = NewCodec("b64_std",
		func(source any, target io.Writer) error {
			encoder := base64.NewEncoder(base64.RawStdEncoding, target)
			defer encoder.Close()
			_, err := encoder.Write([]byte(fmt.Sprint(source)))
			return err
		},
		func(source io.Reader, target any) error {
			decoder := base64.NewDecoder(base64.RawStdEncoding, source)
			decoded, err := io.ReadAll(decoder)
			if err != nil {
				return err
			}
			*(target.(*string)) = string(decoded)
			return nil
		},
	)
	URL = NewCodec("url",
		func(source any, target io.Writer) error {
			encoded := url.QueryEscape(fmt.Sprintf("%v", source))
			_, err := target.Write([]byte(encoded))
			return err
		},
		func(source io.Reader, target any) error {
			decoded, err := io.ReadAll(source)
			if err != nil {
				return err
			}
			decodedStr, err := url.QueryUnescape(string(decoded))
			if err != nil {
				return err
			}
			*(target.(*string)) = decodedStr
			return nil
		},
	)
	SHA1 = NewCodec("sha1",
		func(input any, output io.Writer) error {
			buf := input.([]byte)
			br := bytes.NewReader(buf)
			hashed := sha1.New()
			if _, err := io.Copy(hashed, br); err != nil {
				return err
			}
			_, err := output.Write([]byte(hex.EncodeToString(hashed.Sum(nil))))
			return err
		},
		func(source io.Reader, target any) error {
			hashed := sha1.New()
			if _, err := io.Copy(hashed, source); err != nil {
				return err
			}
			*(target.(*string)) = hex.EncodeToString(hashed.Sum(nil))
			return nil
		},
	)
	SHA256 = NewCodec("sha256",
		func(input any, output io.Writer) error {
			buf := input.([]byte)
			br := bytes.NewReader(buf)
			hashed := sha256.New()
			if _, err := io.Copy(hashed, br); err != nil {
				return err
			}
			_, err := output.Write([]byte(hex.EncodeToString(hashed.Sum(nil))))
			return err
		},
		func(source io.Reader, target any) error {
			hashed := sha256.New()
			if _, err := io.Copy(hashed, source); err != nil {
				return err
			}
			*(target.(*string)) = hex.EncodeToString(hashed.Sum(nil))
			return nil
		},
	)
	SHA512 = NewCodec("sha512",
		func(input any, output io.Writer) error {
			buf := input.([]byte)
			br := bytes.NewReader(buf)
			hashed := sha512.New()
			if _, err := io.Copy(hashed, br); err != nil {
				return err
			}
			_, err := output.Write([]byte(hex.EncodeToString(hashed.Sum(nil))))
			return err
		},
		func(source io.Reader, target any) error {
			hashed := sha512.New()
			if _, err := io.Copy(hashed, source); err != nil {
				return err
			}
			*(target.(*string)) = hex.EncodeToString(hashed.Sum(nil))
			return nil
		},
	)
	MD5 = NewCodec("md5",
		func(input any, output io.Writer) error {
			buf := input.([]byte)
			br := bytes.NewReader(buf)
			hashed := md5.New()
			if _, err := io.Copy(hashed, br); err != nil {
				return err
			}
			_, err := output.Write([]byte(hex.EncodeToString(hashed.Sum(nil))))
			return err
		},
		func(source io.Reader, target any) error {
			hashed := md5.New()
			if _, err := io.Copy(hashed, source); err != nil {
				return err
			}
			*(target.(*string)) = hex.EncodeToString(hashed.Sum(nil))
			return nil
		},
	)
	CRC32 = NewCodec("crc32",
		func(input any, output io.Writer) error {
			buf := input.([]byte)
			br := bytes.NewReader(buf)
			table := crc32.MakeTable(crc32.IEEE)
			hashed := crc32.New(table)
			if _, err := io.Copy(hashed, br); err != nil {
				return err
			}
			_, err := output.Write(hashed.Sum(nil))
			return err
		},
		func(source io.Reader, target any) error {
			content, err := io.ReadAll(source)
			if err != nil {
				return err
			}
			checksum := crc32.ChecksumIEEE(content)
			*(target.(*string)) = strconv.FormatUint(uint64(checksum), 10)
			return nil
		},
	)
	CRC64 = NewCodec("crc64",
		func(input any, output io.Writer) error {
			buf := input.([]byte)
			table := crc64.MakeTable(crc64.ISO)
			checksum := crc64.Checksum(buf, table)
			str := strconv.FormatUint(checksum, 10)
			_, err := output.Write([]byte(str))
			return err
		},
		func(source io.Reader, target any) error {
			buf, err := io.ReadAll(source)
			if err != nil {
				return err
			}
			table := crc64.MakeTable(crc64.ISO)
			checksum := crc64.Checksum(buf, table)
			str := strconv.FormatUint(checksum, 10)
			*(target.(*string)) = str
			return nil
		},
	)
)
