package task05

import (
	"io"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(data []byte) error {
	// TODO: implement.
	return nil
}

type Config struct {
	Address string   `json:"address"`
	Workers int      `json:"workers"`
	Timeout Duration `json:"timeout"`
}

func LoadConfig(r io.Reader) (Config, error) {
	// TODO: implement.
	return Config{}, nil
}
