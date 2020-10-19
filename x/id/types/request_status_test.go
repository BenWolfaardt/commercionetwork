package types_test

import (
	"testing"

	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stretchr/testify/require"

	"github.com/commercionetwork/commercionetwork/x/id/types"
)

func TestRequestStatus_Validate(t *testing.T) {
	tests := []struct {
		name    string
		rq      types.RequestStatus
		wantErr error
	}{
		{
			"invalid status type",
			types.NewRequestStatus("invalid", "message"),
			sdkErr.Wrap(sdkErr.ErrUnknownRequest, "Invalid status type: invalid"),
		},
		{
			"\"rejected\" type",
			types.NewRequestStatus("rejected", ""),
			nil,
		},
		{
			"\"canceled\" type",
			types.NewRequestStatus("canceled", ""),
			nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr != nil {
				require.EqualError(t, tt.rq.Validate(), tt.wantErr.Error())
			} else {
				require.NoError(t, tt.rq.Validate())
			}
		})
	}
}
