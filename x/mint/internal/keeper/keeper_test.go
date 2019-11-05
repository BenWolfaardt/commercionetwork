package keeper

import (
	"fmt"
	"testing"

	"github.com/commercionetwork/commercionetwork/x/mint/internal/types"
	"github.com/commercionetwork/commercionetwork/x/pricefeed"
	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// --------------
// --- Credits
// --------------

func TestKeeper_SetCreditsDenom(t *testing.T) {
	ctx, _, _, k := SetupTestInput()

	denom := "test"
	k.SetCreditsDenom(ctx, denom)

	store := ctx.KVStore(k.storeKey)
	denomBz := store.Get([]byte(types.CreditsDenomStoreKey))
	assert.Equal(t, denom, string(denomBz))
}

func TestKeeper_GetCreditsDenom(t *testing.T) {
	ctx, _, _, k := SetupTestInput()
	denom := "test"
	k.SetCreditsDenom(ctx, denom)
	actual := k.GetCreditsDenom(ctx)
	assert.Equal(t, denom, actual)
}

// --------------
// --- CDPs
// --------------

func TestKeeper_AddCdp(t *testing.T) {
	testData := []struct {
		name             string
		cdps             types.Cdps
		newCdp           types.Cdp
		shouldBeInserted bool
	}{
		{
			name:             "Existing CDP is not inserted",
			cdps:             types.Cdps{testCdp},
			newCdp:           testCdp,
			shouldBeInserted: false,
		},
		{
			name:             "New CDP is inserted properly",
			cdps:             types.Cdps{},
			newCdp:           testCdp,
			shouldBeInserted: true,
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, _, _, k := SetupTestInput()

			for _, cdp := range test.cdps {
				k.AddCdp(ctx, cdp)
			}

			k.AddCdp(ctx, test.newCdp)

			result := k.GetCdps(ctx)

			if test.shouldBeInserted {
				assert.Len(t, result, len(test.cdps)+1)
			} else {
				assert.Len(t, result, len(test.cdps))
			}
		})
	}
}

func TestKeeper_OpenCdp(t *testing.T) {
	testData := []struct {
		name            string
		owner           sdk.AccAddress
		amount          sdk.Coins
		tokenPrice      pricefeed.Price
		userFunds       sdk.Coins
		error           sdk.Error
		returnedCredits sdk.Coins
	}{
		{
			name:       "Invalid deposited amount",
			owner:      testCdp.Owner,
			amount:     sdk.NewCoins(sdk.NewInt64Coin("testcoin", 0)),
			tokenPrice: pricefeed.EmptyPrice(),
			error: sdk.ErrInvalidCoins(fmt.Sprintf(
				"Invalid deposit amount: %s",
				sdk.NewCoins(sdk.NewInt64Coin("testcoin", 0)),
			)),
		},
		{
			name:       "Token price not found",
			owner:      testCdp.Owner,
			amount:     testCdp.DepositedAmount,
			tokenPrice: pricefeed.EmptyPrice(),
			error:      sdk.ErrUnknownRequest(fmt.Sprintf("No current price for given token: %s", testCdp.DepositedAmount[0].Denom)),
		},
		{
			name:       "Not enough funds inside user wallet",
			amount:     testCdp.DepositedAmount,
			owner:      testCdp.Owner,
			tokenPrice: pricefeed.NewPrice("ucommercio", sdk.NewDec(10), sdk.NewInt(1000)),
			error: sdk.ErrInsufficientCoins(fmt.Sprintf(
				"insufficient account funds; %s < %s",
				sdk.Coins{},
				sdk.NewCoins(sdk.NewInt64Coin("ucommercio", 100)),
			)),
		},
		{
			name:            "Successful opening",
			amount:          testCdp.DepositedAmount,
			owner:           testCdp.Owner,
			tokenPrice:      pricefeed.NewPrice(testLiquidityDenom, sdk.NewDec(10), sdk.NewInt(1000)),
			userFunds:       testCdp.DepositedAmount,
			returnedCredits: sdk.NewCoins(sdk.NewInt64Coin(testCreditsDenom, 10*50)),
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, bk, pfk, k := SetupTestInput()

			// Setup
			if !test.userFunds.Empty() {
				_ = bk.SetCoins(ctx, test.owner, test.userFunds)
			}
			if !test.tokenPrice.Equals(pricefeed.EmptyPrice()) {
				pfk.SetCurrentPrice(ctx, test.tokenPrice)
			}

			err := k.OpenCdp(ctx, test.owner, test.amount)
			assert.Equal(t, test.error, err, "Invalid error: %s, %s", test.error, err)

			if !test.returnedCredits.IsEqual(sdk.Coins{}) {
				actual := bk.GetCoins(ctx, test.owner)
				assert.Equal(t, test.returnedCredits, actual)
			}
		})
	}

}

func TestKeeper_GetCdpsByOwner(t *testing.T) {
	t.Run("Empty list is returned properly", func(t *testing.T) {
		ctx, _, _, k := SetupTestInput()
		assert.Empty(t, k.GetCdpsByOwner(ctx, testCdpOwner))
	})

	t.Run("Existing list is returned properly", func(t *testing.T) {
		ctx, _, _, k := SetupTestInput()

		k.AddCdp(ctx, testCdp)

		store := ctx.KVStore(k.storeKey)
		var cdps types.Cdps
		k.cdc.MustUnmarshalBinaryBare(store.Get(k.getCdpKey(testCdp.Owner)), &cdps)
		assert.Equal(t, types.Cdps{testCdp}, k.GetCdpsByOwner(ctx, testCdpOwner))
	})
}

func TestKeeper_CloseCdp(t *testing.T) {

	t.Run("Non existing CDP returns error", func(t *testing.T) {
		ctx, _, _, k := SetupTestInput()

		err := k.CloseCdp(ctx, testCdp.Owner, testCdp.Timestamp)
		errMsg := fmt.Sprintf("CDP for user with address %s and timestamp %d does not exist", testCdpOwner, testCdp.Timestamp)
		assert.Equal(t, sdk.ErrUnknownRequest(errMsg), err)
	})

	t.Run("Existing CDP is closed properly", func(t *testing.T) {
		ctx, bk, _, k := SetupTestInput()

		k.AddCdp(ctx, testCdp)
		_ = k.supplyKeeper.MintCoins(ctx, types.ModuleName, testLiquidityPool)
		_, _ = bk.AddCoins(ctx, testCdpOwner, testCdp.CreditsAmount)

		assert.NoError(t, k.CloseCdp(ctx, testCdpOwner, testCdp.Timestamp))
		assert.Equal(t, testCdp.DepositedAmount, bk.GetCoins(ctx, testCdpOwner))
	})

}

func TestKeeper_DeleteCdp(t *testing.T) {
	testData := []struct {
		name            string
		existingCdps    types.Cdps
		deletedCdp      types.Cdp
		shouldBeDeleted bool
	}{
		{
			name:            "Existing CDP is deleted",
			existingCdps:    types.Cdps{testCdp},
			deletedCdp:      testCdp,
			shouldBeDeleted: true,
		},
		{
			name:         "Non existent CDP is not deleted",
			existingCdps: types.Cdps{testCdp},
			deletedCdp: types.Cdp{
				Owner:           testCdp.Owner,
				DepositedAmount: testCdp.DepositedAmount,
				CreditsAmount:   testCdp.CreditsAmount,
				Timestamp:       testCdp.Timestamp + 1,
			},
			shouldBeDeleted: false,
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, _, _, k := SetupTestInput()

			for _, cdp := range test.existingCdps {
				k.AddCdp(ctx, cdp)
			}

			k.deleteCdp(ctx, test.deletedCdp)

			result := k.GetCdps(ctx)
			if test.shouldBeDeleted {
				assert.Len(t, result, len(test.existingCdps)-1)
			} else {
				assert.Len(t, result, len(test.existingCdps))
			}
		})
	}
}