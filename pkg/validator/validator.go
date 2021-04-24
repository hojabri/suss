package validator

import (
	"errors"
	"fmt"
	"regexp"
)



// Pattern validates data with a regex pattern
func Pattern(data, pattern string) error {
	re, err := compileRegexp(pattern)
	if err != nil {
		return errors.New(fmt.Sprintf("pattern is %s, but is invalid: %s", pattern, err.Error()))
	}
	if !re.MatchString(data) {
		return errors.New(fmt.Sprintf("data is not match with pattern. %s, %s", pattern, data))
	}
	return nil
}

func compileRegexp(pattern string) (*regexp.Regexp, error) {

	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return r, nil
}
