package epochid

import (
	"fmt"
	"strconv"
	"strings"

	goerrors "github.com/TudorHulban/go-errors"
)

type paramsValidationString struct {
	Caller string
	Name   string
	Value  string
}

func validationString(params *paramsValidationString, result *EpochID) error {
	if len(params.Value) != 19 {
		return goerrors.ErrValidation{
			Caller: params.Caller,

			Issue: goerrors.ErrInvalidInput{
				InputName: params.Name,
			},
		}
	}

	numericValue, errConvToNumeric := strconv.Atoi(
		strings.TrimSpace(params.Value),
	)
	if errConvToNumeric != nil {
		return goerrors.ErrValidation{
			Caller: params.Caller,

			Issue: fmt.Errorf(
				"%s: %w",
				params.Name,
				errConvToNumeric,
			),
		}
	}

	if numericValue < 0 {
		return goerrors.ErrValidation{
			Caller: params.Caller,

			Issue: goerrors.ErrNegativeInput{
				InputName: params.Name,
			},
		}
	}

	if numericValue == 0 {
		return goerrors.ErrValidation{
			Caller: params.Caller,

			Issue: goerrors.ErrZeroInput{
				InputName: params.Name,
			},
		}
	}

	*result = EpochID(numericValue)

	return nil
}
