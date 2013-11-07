package helper

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

import "log"

var ErrHelperInvalidPathLength = errors.New("Invalid Path Length")

func ParseUrl(url *url.URL) ([]string, int64, float32, error) {
	path := strings.Split(url.Path, "/")
	l := len(path)

	if l < 4 {
		return nil, 0, 0, ErrHelperInvalidPathLength
	}

	//XXX: ugly
	if path[0] == "" { // won't it always? eg: /a/path => ["" "a" "path"]
		path = path[1:]
		l = len(path)
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

	return path[0 : l-2], t, float32(v), nil
}
