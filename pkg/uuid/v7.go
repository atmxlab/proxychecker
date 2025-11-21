package uuid

import (
	"github.com/atmxlab/proxychecker/pkg/errors"
	"github.com/google/uuid"
)

type UUID = uuid.UUID

func MustV7() uuid.UUID {
	id, err := uuid.NewV7()
	if err != nil {
		panic(errors.Wrap(err, "uuid.NewV7: failed to generate uuid v7"))
	}

	return id
}
