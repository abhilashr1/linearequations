package stringfilter

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// AddOneBefore adds the co-efficient value of 1 to expression strings
// which do not have any co-efficient values available
// For example: x-y =3 becomes 1x - 1y = 3
func AddOneBefore(s string, pos int) string {
	newstr := ""

	if len(s) < pos {
		return newstr
	}

	for i := 0; i < pos; i++ {
		newstr = newstr + string(s[i])
	}
	newstr += "1"
	for i := pos; i < len(s); i++ {
		newstr = newstr + string(s[i])
	}
	return newstr
}

// FilterUtil iterates through the expression and checks to see if there are no co-efficients found
func FilterUtil(text string) string {
	text = " " + text
	flag := 0
	for pos, char := range text {
		if unicode.IsLetter(char) {
			for pos >= 0 {
				ch := text[pos]
				if ch >= 48 && ch <= 57 {
					// A digit is found, so there is already a co-efficient
					break
				} else if ch == '+' || ch == '-' || pos == 0 {
					// Add one to the string since a coefficient value is absent
					text = AddOneBefore(text, pos+1)
					flag = 1
					break
				}
				pos--
			}
			if flag == 1 {
				break
			}
		}
	}
	if flag == 1 {
		text = FilterUtil(text)
	}
	return text
}

//GetVariables fetches the list of variables used in the expression
func GetVariables(text string) []string {
	text = strings.ToLower(strings.Replace(text, " ", "", -1))
	reText := regexp.MustCompile("[a-z]+")
	variablesText := reText.FindAllString(text, -1)
	return variablesText
}

//Filter is the main processing function which extracts the coefficients and variables,
//then converts them into float array form
func Filter(text1 string, text2 string) ([2][3]float32, error) {

	text1 = FilterUtil(strings.ToLower(strings.Replace(text1, " ", "", -1)))
	text2 = FilterUtil(strings.ToLower(strings.Replace(text2, " ", "", -1)))

	reText1 := regexp.MustCompile("[a-z]+")
	reText2 := regexp.MustCompile("[a-z]+")
	reNum1 := regexp.MustCompile("[+-]?[0-9]+([.][0-9]+)?")
	reNum2 := regexp.MustCompile("[+-]?[0-9]+([.][0-9]+)?")

	variablesText1 := reText1.FindAllString(text1, -1)
	variablesText2 := reText2.FindAllString(text2, -1)

	constantsText1 := reNum1.FindAllString(text1, -1)
	constantsText2 := reNum2.FindAllString(text2, -1)

	// Check if the variables are two in number
	if len(variablesText1) != 2 || len(variablesText2) != 2 || len(constantsText1) != 3 || len(constantsText2) != 3 {
		err := errors.New("Non-Uniform variables")
		return [2][3]float32{}, err
	}

	// check if both the variables are the same
	if variablesText1[0] == variablesText1[1] || variablesText2[0] == variablesText2[1] {
		err := errors.New("Repeated variables")
		return [2][3]float32{}, err
	}

	if variablesText1[0] != variablesText2[0] {
		// The order is switched. So, reverse it.
		temp := constantsText2[0]
		constantsText2[0] = constantsText2[1]
		constantsText2[1] = temp

		temp = variablesText2[0]
		variablesText2[0] = variablesText2[1]
		variablesText2[1] = temp
	}

	frfc, err := strconv.ParseFloat(constantsText1[0], 32)
	frsc, err := strconv.ParseFloat(constantsText1[1], 32)
	frtc, err := strconv.ParseFloat(constantsText1[2], 32)

	srfc, err := strconv.ParseFloat(constantsText2[0], 32)
	srsc, err := strconv.ParseFloat(constantsText2[1], 32)
	srtc, err := strconv.ParseFloat(constantsText2[2], 32)

	res := [2][3]float32{
		{float32(frfc), float32(frsc), float32(frtc)},
		{float32(srfc), float32(srsc), float32(srtc)}}

	return res, err

}
