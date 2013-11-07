package helper

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

var ErrHelperInvalidPathLength = errors.New("Invalid Path Length")

func ParseUrl(url *url.URL) ([]string, int64, float32, error) {
	path := strings.Split(url.Path, "/")
	l := len(path)

	if l < 3 {
		return nil, 0, 0, ErrHelperInvalidPathLength
	}

	t, err := strconv.ParseInt(path[l-2], 10, 64)
	if err != nil {
		return nil, 0, 0, err
	}

	//TODO: validate time

	v, err2 := strconv.ParseFloat(path[l-1], 32)
	if err2 != nil {
		return nil, 0, 0, err
	}

	return path[:l-3], t, float32(v), nil
}
