package redis

import "errors"

var redis = map[string]string{}

func Set(key, value string) {
	redis[key] = value
}
func Get(key string) (string, error) {
	if v, ok := redis[key]; ok {
		return v, nil
	}
	return "", errors.New("key not found")
}
