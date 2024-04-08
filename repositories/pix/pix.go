package pix

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"unicode"

	"github.com/EnubeRepos/ElevenST_BFF/model/pix"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	pixLib "github.com/fonini/go-pix/pix"
	"github.com/r10r/crc16"
)

type intMap map[int]interface{}

func GeneratePixCopyPast(amt float64, name, descpt, pixKey string) (string, error) {
	options := pixLib.Options{
		Name:          name,
		Key:           pixKey,
		City:          "Sao Paulo",
		Amount:        amt,          // optional
		Description:   "Invoice #4", // optional
		TransactionID: "0001",       // optional
	}

	copyPaste, err := pixLib.Pix(options)

	if err != nil {
		panic(err)
	}

	return copyPaste, err
}

func GeneratePaste(amt float64, name, descpt, pixKey string) (string, error) {
	opts := pix.PixOpts{
		Name:        name,
		Description: descpt,
		Amount:      amt,
		Key:         pixKey,
		City:        "São Paulo",
	}

	paste, err := Pix(opts)
	if err != nil {
		return paste, err
	}
	return paste, nil
}

// Pix generates a Copy and Paste Pix code
func Pix(options pix.PixOpts) (string, error) {
	if err := validateData(options); err != nil {
		return "", err
	}
	data := buildDataMap(options)
	str := parseData(data)
	// Add the CRC at the end
	str += "6304"
	crc, err := calculateCRC16(str)
	if err != nil {
		return "", err
	}
	str += crc
	return str, nil
}

func validateData(options pix.PixOpts) error {
	if options.Key == "" {
		return errors.New("key must not be empty")
	}
	if options.Name == "" {
		return errors.New("name must not be empty")
	}
	// if utf8.RuneCountInString(options.Name) > 25 {
	// 	return errors.New("name must be at least 25 characters long")
	// }
	return nil
}

func buildDataMap(options pix.PixOpts) intMap {
	data := make(intMap)
	// Payload Format Indicator
	data[0] = "01"
	// Merchant Account Information
	data[26] = intMap{0: "BR.GOV.BCB.PIX", 1: options.Key, 2: options.Description}
	// Merchant Category Code
	data[52] = "0000"
	// Transaction Currency - Brazilian Real - ISO4217
	data[53] = "986"
	// Transaction Amount
	data[54] = options.Amount
	// Country Code - ISO3166-1 alpha 2
	data[58] = "BR"
	// Merchant Name. 25 characters maximum
	data[59] = options.Name
	// Merchant City. 15 characters maximum
	// Transaction ID
	data[62] = intMap{5: "***", 50: intMap{0: "BR.GOV.BCB.BRCODE", 1: "1.0.0"}}
	return data
}

func parseData(data intMap) string {
	var str string

	keys := sortKeys(data)
	for _, k := range keys {
		v := reflect.ValueOf(data[k])
		switch v.Kind() {
		case reflect.String:
			value := data[k].(string)
			str += fmt.Sprintf("%02d%02d%s", k, len(value), value)
		case reflect.Float64:
			value := strconv.FormatFloat(v.Float(), 'f', 2, 64)
			str += fmt.Sprintf("%02d%02d%s", k, len(value), value)
		case reflect.Map:
			// If the element is another map, do a recursive call
			content := parseData(data[k].(intMap))
			str += fmt.Sprintf("%02d%02d%s", k, len(content), content)
		}
	}
	return str
}

func sortKeys(data intMap) []int {
	keys := make([]int, len(data))
	i := 0

	for k := range data {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}

func calculateCRC16(str string) (string, error) {
	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
	h := crc16.New(table)
	_, err := h.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%04X", h.Sum16()), nil
}
func NormalizeInput(input string) (string, error) {

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	res, _, err := transform.String(t, input)

	if err != nil {

		fmt.Println(err.Error())

		return input, err

	}

	fmt.Println("input normalize to " + res)

	return res, nil

}
