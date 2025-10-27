package script

import (
	"fmt"
	"net/url"
	"regexp"
	"scripter/pkg/api"
	"strings"
)

func CheckPackage(packageName string) (string, error) {
	if strings.Contains(packageName, "github") {
		return packageName, nil
	}

	baseUrl := "https://api.github.com/search/repositories"
	params := url.Values{}
	params.Add("q", fmt.Sprintf("%s in:name language:go", packageName))
	params.Add("sort", "stars")
	params.Add("order", "desc")
	params.Add("per_page", "2")

	reqUrl := baseUrl + "?" + params.Encode()

	name, err := api.AskGithub(reqUrl)
	if err != nil {
		return "", err
	}

	return "github.com/" + name, nil
}

func GetPackageNameFromErrOutput(input string, re *regexp.Regexp) string {
    matches := re.FindStringSubmatch(input)
    if len(matches) > 1 {
        return matches[1]
    }
    return ""
}
