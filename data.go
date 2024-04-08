package flo

import (
	"bytes"

	c "github.com/toxyl/flo/codec"
	"github.com/toxyl/flo/log"
	"github.com/toxyl/flo/utils"
)

func (f *FileObj) read(codec *c.Codec, target any) error {
	return codec.Decode(f.Open(), target)
}

func (f *FileObj) readStr(codec *c.Codec) string {
	s := ""
	f.read(codec, &s)
	return s
}

func (f *FileObj) readBytes(codec *c.Codec) []byte {
	b := []byte{}
	f.read(codec, &b)
	return b
}

func (f *FileObj) mustRead(codec *c.Codec, target any) *FileObj {
	if err := f.read(codec, target); err != nil {
		log.Panic("failed to read %s: %v", f.path, err)
	}
	return f
}

func (f *FileObj) encodeStr(codec *c.Codec) string {
	var buf bytes.Buffer
	codec.Encode(f.AsString(), &buf)
	return buf.String()
}

func (f *FileObj) encodeBytes(codec *c.Codec) []byte {
	var buf bytes.Buffer
	codec.Encode(f.AsBytes(), &buf)
	return buf.Bytes()
}

func (f *FileObj) AsBytes() []byte     { return f.readBytes(c.BYTES) }
func (f *FileObj) AsBytesGZ() []byte   { return f.encodeBytes(c.BYTESGZ) }
func (f *FileObj) AsString() string    { return f.readStr(c.STRING) }
func (f *FileObj) AsStringGZ() string  { return f.encodeStr(c.STRINGGZ) }
func (f *FileObj) AsBase64URL() string { return f.encodeStr(c.BASE64_URL) }
func (f *FileObj) AsBase64Std() string { return f.encodeStr(c.BASE64_STD) }
func (f *FileObj) AsURL() string       { return f.encodeStr(c.URL) }

func (f *FileObj) LoadBytes(target *[]byte) error    { return f.read(c.BYTES, target) }
func (f *FileObj) LoadBytesGZ(target *[]byte) error  { return f.read(c.BYTESGZ, target) }
func (f *FileObj) LoadString(target *string) error   { return f.read(c.STRING, target) }
func (f *FileObj) LoadStringGZ(target *string) error { return f.read(c.STRINGGZ, target) }
func (f *FileObj) LoadGob(target any) error          { return f.read(c.GOB, target) }
func (f *FileObj) LoadGobGZ(target any) error        { return f.read(c.GOBGZ, target) }
func (f *FileObj) LoadYAML(target any) error         { return f.read(c.YAML, target) }
func (f *FileObj) LoadJSON(target any) error         { return f.read(c.JSON, target) }
func (f *FileObj) LoadBase64URL(target any) error    { return f.read(c.BASE64_URL, target) }
func (f *FileObj) LoadBase64Std(target any) error    { return f.read(c.BASE64_STD, target) }
func (f *FileObj) LoadURL(target any) error          { return f.read(c.URL, target) }

func (f *FileObj) ReadBytes(target *[]byte) *FileObj    { return f.mustRead(c.BYTES, target) }
func (f *FileObj) ReadBytesGZ(target *[]byte) *FileObj  { return f.mustRead(c.BYTESGZ, target) }
func (f *FileObj) ReadString(target *string) *FileObj   { return f.mustRead(c.STRING, target) }
func (f *FileObj) ReadStringGZ(target *string) *FileObj { return f.mustRead(c.STRINGGZ, target) }
func (f *FileObj) ReadGob(target any) *FileObj          { return f.mustRead(c.GOB, target) }
func (f *FileObj) ReadGobGZ(target any) *FileObj        { return f.mustRead(c.GOBGZ, target) }
func (f *FileObj) ReadYAML(target any) *FileObj         { return f.mustRead(c.YAML, target) }
func (f *FileObj) ReadJSON(target any) *FileObj         { return f.mustRead(c.JSON, target) }
func (f *FileObj) ReadBase64URL(target any) *FileObj    { return f.mustRead(c.BASE64_URL, target) }
func (f *FileObj) ReadBase64Std(target any) *FileObj    { return f.mustRead(c.BASE64_STD, target) }
func (f *FileObj) ReadURL(target any) *FileObj          { return f.mustRead(c.URL, target) }

func (f *FileObj) write(codec *c.Codec, data any) error {
	defer f.updateInfo()
	file, err := utils.Mkfile(f.Path(), f.Parent().info.Mode, f.info.Mode)
	if err != nil {
		return err
	}
	return codec.Encode(data, file)
}

func (f *FileObj) mustWrite(codec *c.Codec, data any) *FileObj {
	if err := f.write(codec, data); err != nil {
		log.Panic("failed to write %s: %v", f.path, err)
	}
	return f
}

func (f *FileObj) StoreBytes(data []byte) error    { return f.write(c.BYTES, data) }
func (f *FileObj) StoreBytesGZ(data []byte) error  { return f.write(c.BYTESGZ, data) }
func (f *FileObj) StoreString(data string) error   { return f.write(c.STRING, data) }
func (f *FileObj) StoreStringGZ(data string) error { return f.write(c.STRINGGZ, data) }
func (f *FileObj) StoreGob(data any) error         { return f.write(c.GOB, data) }
func (f *FileObj) StoreGobGZ(data any) error       { return f.write(c.GOBGZ, data) }
func (f *FileObj) StoreYAML(data any) error        { return f.write(c.YAML, data) }
func (f *FileObj) StoreJSON(data any) error        { return f.write(c.JSON, data) }
func (f *FileObj) StoreBase64URL(data any) error   { return f.write(c.BASE64_URL, data) }
func (f *FileObj) StoreBase64Std(data any) error   { return f.write(c.BASE64_STD, data) }
func (f *FileObj) StoreURL(data any) error         { return f.write(c.URL, data) }

func (f *FileObj) WriteBytes(data []byte) *FileObj    { return f.mustWrite(c.BYTES, data) }
func (f *FileObj) WriteBytesGZ(data []byte) *FileObj  { return f.mustWrite(c.BYTESGZ, data) }
func (f *FileObj) WriteString(data string) *FileObj   { return f.mustWrite(c.STRING, data) }
func (f *FileObj) WriteStringGZ(data string) *FileObj { return f.mustWrite(c.STRINGGZ, data) }
func (f *FileObj) WriteGob(data any) *FileObj         { return f.mustWrite(c.GOB, data) }
func (f *FileObj) WriteGobGZ(data any) *FileObj       { return f.mustWrite(c.GOBGZ, data) }
func (f *FileObj) WriteYAML(data any) *FileObj        { return f.mustWrite(c.YAML, data) }
func (f *FileObj) WriteJSON(data any) *FileObj        { return f.mustWrite(c.JSON, data) }
func (f *FileObj) WriteBase64URL(data any) *FileObj   { return f.mustWrite(c.BASE64_URL, data) }
func (f *FileObj) WriteBase64Std(data any) *FileObj   { return f.mustWrite(c.BASE64_STD, data) }
func (f *FileObj) WriteURL(data any) *FileObj         { return f.mustWrite(c.URL, data) }
