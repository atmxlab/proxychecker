package config

import (
	"github.com/atmxlab/atmc"
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/samber/lo"
)

func LoadAndScan[T any](path string) (T, error) {
	scanner, err := atmc.New().Load(path)
	if err != nil {
		return lo.Empty[T](), errors.Wrap(err, "atmc.New.Load")
	}

	var cfg T
	if err = scanner.Scan(&cfg); err != nil {
		return lo.Empty[T](), errors.Wrap(err, "scanner.Scan")
	}

	return cfg, nil
}
