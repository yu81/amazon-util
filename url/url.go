package url

import "regexp"

var ASINRegex = regexp.MustCompile("(B[A-Z0-9]{9})|([0-9]{10})")

func IsASIN(id string) bool {
	r := ASINRegex.Copy()
	return r.MatchString(id) && len(id) == 10
}

func ExtractASIN(u string) string {
	r := ASINRegex.Copy()
	matched := r.FindAllStringSubmatch(u, 1)
	if len(matched) > 0 && len(matched[0]) > 0 && IsASIN(matched[0][0]) {
		return matched[0][0]
	}
	return ""
}

func CreateAmazonJpURLFromASIN(id string) string {
	return "https://www.amazon.co.jp/dp/" + id
}

func CreateAmazonJpURLFromASINWithAffiliate(id, tag string) string {
	return CreateAmazonJpURLFromASIN(id) + "/?tag=" + tag
}

func ToSimpleAmazonLink(link string) string {
	return CreateAmazonJpURLFromASIN(ExtractASIN(link))
}

func ToSimpleAmazonLinkWithAffiliate(link, tag string) string {
	return CreateAmazonJpURLFromASINWithAffiliate(ExtractASIN(link), tag)
}
