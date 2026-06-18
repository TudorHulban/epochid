package epochid

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tudorhulban/hxerrors"
)

type paramsValidationString struct {
	Caller string
	Name   string
	Value  string
}

func validationString(params *paramsValidationString, result *EpochID) error {
	if len(params.Value) != 19 {
		return hxerrors.ErrValidation{
			Caller: params.Caller,

			Issue: hxerrors.ErrInvalidInput{
				InputName: params.Name,
			},
		}
	}

	numericValue, errConvToNumeric := strconv.Atoi(
		strings.TrimSpace(params.Value),
	)
	if errConvToNumeric != nil {
		return hxerrors.ErrValidation{
			Caller: params.Caller,

			Issue: fmt.Errorf(
				"%s: %w",
				params.Name,
				errConvToNumeric,
			),
		}
	}

	if numericValue < 0 {
		return hxerrors.ErrValidation{
			Caller: params.Caller,

			Issue: hxerrors.ErrNegativeInput{
				InputName: params.Name,
			},
		}
	}

	if numericValue == 0 {
		return hxerrors.ErrValidation{
			Caller: params.Caller,

			Issue: hxerrors.ErrZeroInput{
				InputName: params.Name,
			},
		}
	}

	*result = EpochID(numericValue)

	return nil
}
