package keeper

import (
	"fmt"
	"testing"

	sdkErr "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stretchr/testify/require"

	"github.com/commercionetwork/commercionetwork/x/commerciomint/types"
	pricefeedTypes "github.com/commercionetwork/commercionetwork/x/pricefeed/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeper_StoreCdp(t *testing.T) {
	ctx, bk, _, _, _, k := SetupTestInput()
	// handler := NewHandler(k)

	_, _ = bk.AddCoins(ctx, k.supplyKeeper.GetModuleAddress(types.ModuleName), testCdp.Collateral)
	_ = bk.SetCoins(ctx, testCdp.Owner, sdk.NewCoins(testCdp.Credits))
	require.Equal(t, 0, len(k.GetAllPositions(ctx)))
	k.SetPosition(ctx, testCdp)
	require.Equal(t, 1, len(k.GetAllPositions(ctx)))
	cdp, found := k.GetPosition(ctx, testCdp.Owner, testCdp.CreatedAt)
	require.True(t, found)
	require.Equal(t, testCdp.Owner, cdp.Owner)
	require.Equal(t, testCdp.CreatedAt, cdp.CreatedAt)
}

// --------------
// --- Credits
// --------------

func TestKeeper_SetCreditsDenom(t *testing.T) {
	ctx, _, _, _, _, k := SetupTestInput()
	denom := "test"
	k.SetCreditsDenom(ctx, denom)

	store := ctx.KVStore(k.storeKey)
	denomBz := store.Get([]byte(types.CreditsDenomStoreKey))
	require.Equal(t, denom, string(denomBz))
}

func TestKeeper_GetCreditsDenom(t *testing.T) {
	ctx, _, _, _, _, k := SetupTestInput()
	denom := "test"
	k.SetCreditsDenom(ctx, denom)
	actual := k.GetCreditsDenom(ctx)
	require.Equal(t, denom, actual)
}

// --------------
// --- CDPs
// --------------

func TestKeeper_StoreCdpBasic(t *testing.T) {
	testData := []struct {
		name             string
		cdps             []types.Position
		newCdp           types.Position
		shouldBeInserted bool
	}{
		{
			name:             "Existing CDP is not inserted",
			cdps:             []types.Position{testCdp},
			newCdp:           testCdp,
			shouldBeInserted: false,
		},
		{
			name:             "New CDP is inserted properly",
			cdps:             nil,
			newCdp:           testCdp,
			shouldBeInserted: true,
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, _, _, _, _, k := SetupTestInput()
			for _, cdp := range test.cdps {
				k.SetPosition(ctx, cdp)
			}

			if test.shouldBeInserted {
				require.NotPanics(t, func() { k.SetPosition(ctx, test.newCdp) })
				require.Len(t, k.GetAllPositions(ctx), len(test.cdps)+1)
				return
			}

			if !test.shouldBeInserted {
				require.Panics(t, func() { k.SetPosition(ctx, test.newCdp) })
				require.Len(t, k.GetAllPositions(ctx), len(test.cdps))
				return
			}
		})
	}
}

func TestKeeper_OpenCdp(t *testing.T) {
	testData := []struct {
		name            string
		owner           sdk.AccAddress
		amount          sdk.Coins
		tokenPrice      pricefeedTypes.Price
		userFunds       sdk.Coins
		error           error
		returnedCredits sdk.Coins
	}{
		{
			name:       "invalid deposited amount",
			owner:      testCdp.Owner,
			amount:     sdk.NewCoins(sdk.NewInt64Coin("ucommercio", 0)),
			tokenPrice: pricefeedTypes.NewPrice("ucommercio", sdk.NewDec(10), sdk.NewInt(1000)),
			error:      fmt.Errorf("invalid position: invalid deposit amount: "),
		},
		{
			name:       "Token price not found",
			owner:      testCdp.Owner,
			amount:     testCdp.Collateral,
			tokenPrice: pricefeedTypes.EmptyPrice(),
			error:      sdkErr.Wrap(sdkErr.ErrUnknownRequest, fmt.Sprintf("no current price for given denom: %s", testCdp.Collateral[0].Denom)),
		},
		{
			name:       "Not enough funds inside user wallet",
			amount:     testCdp.Collateral,
			owner:      testCdp.Owner,
			tokenPrice: pricefeedTypes.NewPrice("ucommercio", sdk.NewDec(10), sdk.NewInt(1000)),
			error: sdkErr.Wrap(sdkErr.ErrInsufficientFunds, fmt.Sprintf(
				"insufficient account funds; %s < %s",
				sdk.NewCoins(sdk.NewInt64Coin("stake", 500)),
				sdk.NewCoins(sdk.NewInt64Coin("ucommercio", 100)),
			)),
		},
		{
			name:            "Successful opening",
			amount:          testCdp.Collateral,
			owner:           testCdp.Owner,
			tokenPrice:      pricefeedTypes.NewPrice(testLiquidityDenom, sdk.NewDec(10), sdk.NewInt(1000)),
			userFunds:       testCdp.Collateral,
			returnedCredits: sdk.NewCoins(sdk.NewInt64Coin(testCreditsDenom, 10*50)),
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, bk, pfk, _, _, k := SetupTestInput()
			ctx = ctx.WithBlockHeight(10)

			// Setup
			if !test.userFunds.Empty() {
				_ = bk.SetCoins(ctx, test.owner, test.userFunds)
			}
			if !test.tokenPrice.Equals(pricefeedTypes.EmptyPrice()) {
				pfk.SetCurrentPrice(ctx, test.tokenPrice)
			}

			err := k.NewPosition(ctx, test.owner, test.amount)
			if test.error != nil {
				require.Error(t, err)
				require.Equal(t, test.error.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}

			if !test.returnedCredits.IsEqual(sdk.Coins{}) {
				actual := bk.GetCoins(ctx, test.owner)
				require.Equal(t, test.returnedCredits, actual)
			}
		})
	}

}

func TestKeeper_GetAllPositionsOwnedBy(t *testing.T) {
	t.Run("Empty list is returned properly", func(t *testing.T) {
		ctx, _, _, _, _, k := SetupTestInput()
		require.Empty(t, k.GetAllPositionsOwnedBy(ctx, testCdpOwner))
	})
	t.Run("Existing list is returned properly", func(t *testing.T) {
		ctx, _, _, _, _, k := SetupTestInput()
		k.SetPosition(ctx, testCdp)
		require.Equal(t, []types.Position{testCdp}, k.GetAllPositionsOwnedBy(ctx, testCdpOwner))
	})
}

func TestKeeper_CloseCdp(t *testing.T) {
	t.Run("Non existing CDP returns error", func(t *testing.T) {
		ctx, _, _, _, _, k := SetupTestInput()

		err := k.BurnCCC(ctx, testCdp.Owner, testCdp.CreatedAt)
		errMsg := fmt.Sprintf("position for user with address %s and timestamp %d does not exist", testCdpOwner, testCdp.CreatedAt)
		require.Equal(t, sdkErr.Wrap(sdkErr.ErrUnknownRequest, errMsg).Error(), err.Error())
	})

	t.Run("Existing CDP is closed properly", func(t *testing.T) {
		ctx, bk, _, _, _, k := SetupTestInput()

		k.SetPosition(ctx, testCdp)
		_ = k.supplyKeeper.MintCoins(ctx, types.ModuleName, testLiquidityPool)
		_, _ = bk.AddCoins(ctx, testCdpOwner, sdk.NewCoins(testCdp.Credits))

		require.NoError(t, k.BurnCCC(ctx, testCdpOwner, testCdp.CreatedAt))
		require.Equal(t, testCdp.Collateral, bk.GetCoins(ctx, testCdpOwner))
	})

}

func TestKeeper_DeleteCdp(t *testing.T) {
	testData := []struct {
		name            string
		existingCdps    []types.Position
		deletedCdp      types.Position
		shouldBeDeleted bool
	}{
		{
			name:            "Existing CDP is deleted",
			existingCdps:    []types.Position{testCdp},
			deletedCdp:      testCdp,
			shouldBeDeleted: true,
		},
		{
			name:         "Non existent CDP is not deleted",
			existingCdps: []types.Position{testCdp},
			deletedCdp: types.Position{
				Owner:      testCdp.Owner,
				Collateral: testCdp.Collateral,
				Credits:    testCdp.Credits,
				CreatedAt:  testCdp.CreatedAt + 1,
			},
			shouldBeDeleted: false,
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, _, _, _, _, k := SetupTestInput()

			for _, cdp := range test.existingCdps {
				k.SetPosition(ctx, cdp)
			}

			if test.shouldBeDeleted {
				require.NotPanics(t, func() { k.deletePosition(ctx, test.deletedCdp) })
			} else {
				require.Panics(t, func() { k.deletePosition(ctx, test.deletedCdp) })
			}

			result := k.GetAllPositions(ctx)
			if test.shouldBeDeleted {
				require.Len(t, result, len(test.existingCdps)-1)
			} else {
				require.Len(t, result, len(test.existingCdps))
			}
		})
	}
}

// --------------
// --- CdpCollateralRate
// --------------

func TestKeeper_SetCdpCollateralRate(t *testing.T) {
	ctx, _, _, _, _, k := SetupTestInput()
	require.Error(t, k.SetConversionRate(ctx, sdk.NewInt(0).ToDec()))
	require.Error(t, k.SetConversionRate(ctx, sdk.NewInt(-1).ToDec()))
	require.NoError(t, k.SetConversionRate(ctx, sdk.NewInt(2).ToDec()))
	rate := sdk.NewDec(3)
	require.NoError(t, k.SetConversionRate(ctx, rate))

	var got sdk.Dec
	k.cdc.MustUnmarshalBinaryBare(ctx.KVStore(k.storeKey).Get([]byte(types.CollateralRateKey)), &got)
	require.True(t, rate.Equal(got), got.String())
}

func TestKeeper_GetCdpCollateralRate(t *testing.T) {
	ctx, _, _, _, _, k := SetupTestInput()
	rate := sdk.NewDec(3)
	require.NoError(t, k.SetConversionRate(ctx, rate))
	require.Equal(t, rate, k.GetConversionRate(ctx))
}
