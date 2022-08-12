package handler

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// base short url
var BASE_SHORT_URL string = "https://r.banksalad.com/"

const (
	// Similar to Base64, but modified to avoid both non-alphanumeric characters (+ and /)
	// and letters that might look ambiguous when printed
	// (0 – zero, I – capital i, O – capital o and l – lower-case L)
	// bitcoin address style
	BASE58_ALPHA_NUM = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	BASE58_LEN       = int64(len(BASE58_ALPHA_NUM))
)

// convert int64 to base58
func convertInt64ToBase58(num int64) string {
	if num == 0 {
		return string(BASE58_ALPHA_NUM[0])
	}

	var base58Str string = ""
	for ; num > 0; num = num / BASE58_LEN {
		base58Str = string(BASE58_ALPHA_NUM[num%BASE58_LEN]) + base58Str
	}

	return base58Str
}

// convert base58 to int64
func convertBase58ToInt64(base58Str string) (int64, error) {
	var num int64 = 0
	for _, c := range []byte(base58Str) {
		i := strings.IndexByte(BASE58_ALPHA_NUM, c)
		if i < 0 {
			return 0, fmt.Errorf("unexpected character %c in base58 literal", c)
		}
		num = BASE58_LEN*num + int64(i)
	}

	return num, nil
}

func urlEscape(rawUrl string) string {
	u, err := url.Parse(rawUrl)
	if err != nil {
		fmt.Println("error UrlEscape")
	}
	escUrl := u.String()

	return escUrl
}

func urlUnescape(escUrl string) string {
	uescUrl, err := url.QueryUnescape(escUrl)
	if err != nil {
		fmt.Println("error UrlUnescape")
	}

	return uescUrl
}

func createShortUrl(orgUrl string) string {
	// original url -> sha256 hash -> folding int64 -> encode base58 string

	// url escaping 하고 시작
	escOrgUrl := urlEscape(orgUrl)
	byteData := []byte(escOrgUrl)
	sha256Hash := sha256.Sum256(byteData)
	sha256HashStr := fmt.Sprintf("%x", sha256Hash)

	// sha256 -> folding
	// splitLen 만큼씩 나눠서 16진수->10진수 변환해서 더하기
	var splitLen int = 8
	var sha256Len int = len(sha256HashStr)
	var splitCnt int = int(math.Ceil(float64(sha256Len) / float64(splitLen)))
	var hashNum int64 = 0
	for i := 0; i < splitCnt; i++ {
		start := i * splitLen
		end := start + splitLen
		if end > sha256Len {
			end = sha256Len
		}

		partNum, err := strconv.ParseInt(sha256HashStr[start:end], 16, 64)
		if err != nil {
			fmt.Println("error ParseInt")
		}

		hashNum += partNum
	}

	shortUrlKey := convertInt64ToBase58(hashNum)

	// shortUrlKey로 redis에 저장
	// key : shortUrlKey
	// value : escOrgUrl

	shortUrl := fmt.Sprintf("%s%s", BASE_SHORT_URL, shortUrlKey)

	return shortUrl
}

// get original url
func getOriginalUrl(shortUrl string) string {

	// short url -> parse short url -> get short url key -> decode base58 -> int64 -> original url
	shortUrlKey := parseShortUrlKey(shortUrl)

	hashNum, err := convertBase58ToInt64(shortUrlKey)
	if err != nil {
		fmt.Println("error ConvertBase58ToInt64")
	}

	fmt.Println("GetOrgUrl shortUrlKey=", shortUrlKey)
	fmt.Println("GetOrgUrl hashNum=", hashNum)
	orgUrl := "dev ing"

	// org url은 redis에서 얻어와야한다
	// key : hashNum
	// value : original url string(url encoded)

	return orgUrl
}

func parseShortUrlKey(shortUrl string) string {
	u, err := url.Parse(shortUrl)
	if err != nil {
		fmt.Println("error parse short url key")
	}

	fmt.Println("input shortUrl=", shortUrl)
	fmt.Println("u.Scheme=", u.Scheme)
	fmt.Println("u.Host=", u.Host)
	fmt.Println("u.Path=", u.Path)
	fmt.Println("u.RawQuery=", u.RawQuery)

	// invalid check
	if u.Scheme != "https" {
		fmt.Println("error not https request")
	}

	if strings.Contains(u.Host, "banksalad") != true {
		fmt.Println("error invalid hostname")
	}

	base58Str := u.Path
	if u.Path[0] == '/' {
		base58Str = u.Path[1:]
	}

	return base58Str
}
