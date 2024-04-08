package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/toxyl/flo"
	"github.com/toxyl/flo/utils"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Writes a YAML file to /tmp/test.yaml")
		fmt.Println("Usage: " + filepath.Base(os.Args[0]))
		return
	}
	testData := struct {
		Root struct {
			Test string  `mapstructure:"test"`
			Pi   float64 `mapstructure:"pi"`
			Phi  float64 `mapstructure:"phi"`
			Tau  float64 `mapstructure:"tau"`
		} `mapstructure:"root"`
	}{
		Root: struct {
			Test string  "mapstructure:\"test\""
			Pi   float64 "mapstructure:\"pi\""
			Phi  float64 "mapstructure:\"phi\""
			Tau  float64 "mapstructure:\"tau\""
		}{
			Test: "This is some content",
			Pi:   math.Pi,
			Phi:  math.Phi,
			Tau:  math.Pi * 2,
		},
	}

	st := testData
	s := fmt.Sprint(st)
	b := []byte(fmt.Sprint(st))
	y := testData
	j := testData
	g := testData
	gz := testData
	b64 := s
	url := s

	fnRW := func(ext string, in, out any, fnTest func(f *flo.FileObj, in, out any)) {
		st.Root.Pi *= 2
		ext = strings.TrimSpace(ext)
		f := flo.File("/tmp/test." + ext)
		fnTest(f, in, out)
		fmt.Printf("%s:\t%v\n", ext, utils.Dereference(out))
	}

	fmt.Println("--- WRITE, then READ ---")
	fnRW(" string", fmt.Sprint(st), &s, func(f *flo.FileObj, in, out any) { f.WriteString(in.(string)).ReadString(out.(*string)) })
	fnRW("stringz", fmt.Sprint(st), &s, func(f *flo.FileObj, in, out any) { f.WriteStringGZ(in.(string)).ReadStringGZ(out.(*string)) })
	fnRW("  bytes", []byte(fmt.Sprint(st)), &b, func(f *flo.FileObj, in, out any) { f.WriteBytes(in.([]byte)).ReadBytes(out.(*[]byte)) })
	fnRW(" bytesz", []byte(fmt.Sprint(st)), &b, func(f *flo.FileObj, in, out any) { f.WriteBytesGZ(in.([]byte)).ReadBytesGZ(out.(*[]byte)) })
	fnRW("   yaml", st, &y, func(f *flo.FileObj, in, out any) { f.WriteYAML(in).ReadYAML(out) })
	fnRW("   json", st, &j, func(f *flo.FileObj, in, out any) { f.WriteJSON(in).ReadJSON(out) })
	fnRW("    gob", st, &g, func(f *flo.FileObj, in, out any) { f.WriteGob(in).ReadGob(out) })
	fnRW("   gobz", st, &gz, func(f *flo.FileObj, in, out any) { f.WriteGobGZ(in).ReadGobGZ(out) })
	fnRW("b64_url", st, &b64, func(f *flo.FileObj, in, out any) { f.WriteBase64URL(in).ReadBase64URL(out) })
	fnRW("b64_std", st, &b64, func(f *flo.FileObj, in, out any) { f.WriteBase64Std(in).ReadBase64Std(out) })
	fnRW("    url", st, &url, func(f *flo.FileObj, in, out any) { f.WriteURL(in).ReadURL(out) })

	srcFile := flo.File("/tmp/test.json")
	altSrcFile := flo.File("/tmp/test.yaml")
	cmpFile := flo.File("/tmp/test.json.2")
	if time.Now().Nanosecond()%2 == 0 {
		_ = cmpFile.CopyFrom(srcFile)
	} else {
		_ = cmpFile.CopyFrom(altSrcFile)
	}

	fmt.Println("")
	fmt.Println("--- Hashes ---")
	fmt.Printf("MD5:   \t%v\n", srcFile.MD5())
	fmt.Printf("SHA1:  \t%v\n", srcFile.SHA1())
	fmt.Printf("SHA256:\t%v\n", srcFile.SHA256())
	fmt.Printf("SHA512:\t%v\n", srcFile.SHA512())
	fmt.Printf("CRC32: \t%v\n", srcFile.CRC32())
	fmt.Printf("CRC64: \t%v\n", srcFile.CRC64())

	fmt.Println("")
	fmt.Println("--- Match ---")
	fmt.Printf("File A == B: \t%v\n", srcFile.SameAs(cmpFile))

	fmt.Println("")
	fmt.Println("--- Encode As ---")
	fmt.Printf("(string):    \n\t\t%v vs \n\t\t%v\n", srcFile.AsString(), cmpFile.AsString())
	fmt.Printf("(stringz):   \n\t\t%v vs \n\t\t%v\n", srcFile.AsStringGZ(), cmpFile.AsStringGZ())
	fmt.Printf("(bytes):     \n\t\t%v vs \n\t\t%v\n", srcFile.AsBytes(), cmpFile.AsBytes())
	fmt.Printf("(bytesz):    \n\t\t%v vs \n\t\t%v\n", srcFile.AsBytesGZ(), cmpFile.AsBytesGZ())
	fmt.Printf("(B64_URL):   \n\t\t%v vs \n\t\t%v\n", srcFile.AsBase64URL(), cmpFile.AsBase64URL())
	fmt.Printf("(B64_STD):   \n\t\t%v vs \n\t\t%v\n", srcFile.AsBase64Std(), cmpFile.AsBase64Std())
	fmt.Printf("(URL):       \n\t\t%v vs \n\t\t%v\n", srcFile.AsURL(), cmpFile.AsURL())

	fmt.Println("")
	fmt.Println("--- Compare ---")
	fmt.Printf("MD5:         %v\n\t\t%v vs \n\t\t%v\n", srcFile.SameMD5As(cmpFile), srcFile.MD5(), cmpFile.MD5())
	fmt.Printf("SHA1:        %v\n\t\t%v vs \n\t\t%v\n", srcFile.SameSHA1As(cmpFile), srcFile.SHA1(), cmpFile.SHA1())
	fmt.Printf("SHA256:      %v\n\t\t%v vs \n\t\t%v\n", srcFile.SameSHA256As(cmpFile), srcFile.SHA256(), cmpFile.SHA256())
	fmt.Printf("SHA512:      %v\n\t\t%v vs \n\t\t%v\n", srcFile.SameSHA512As(cmpFile), srcFile.SHA512(), cmpFile.SHA512())
	fmt.Printf("CRC32:       %v\n\t\t%v vs \n\t\t%v\n", srcFile.SameCRC32As(cmpFile), srcFile.CRC32(), cmpFile.CRC32())
	fmt.Printf("CRC64:       %v\n\t\t%v vs \n\t\t%v\n", srcFile.SameCRC64As(cmpFile), srcFile.CRC64(), cmpFile.CRC64())

	fmt.Println("")
	fmt.Println("--- Checksum ---")
	fmt.Printf("File A == B:   \t%v\n\t\t%v vs \n\t\t%v\n", cmpFile.Checksum().Matches(srcFile.Checksum()), srcFile.Checksum().Get(), cmpFile.Checksum().Get())
	fmt.Printf("File B == srd: \t%v\n", cmpFile.Checksum().MatchesBytes(srcFile.AsBytes()))
	fmt.Printf("File B == alt: \t%v\n", cmpFile.Checksum().MatchesBytes(altSrcFile.AsBytes()))
	fmt.Printf("File B == alt: \t%v\n", cmpFile.Checksum().Matches(altSrcFile.Checksum()))
}
