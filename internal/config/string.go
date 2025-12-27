package config

import "encoding/json"

func (c Config) String() string {
	buffer, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	return string(buffer)
}
