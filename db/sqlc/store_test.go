package sqlcDB

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	// Printing the account  balance before transaction
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	// run n concurrent transfer transaction
	n := 2
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		txName := fmt.Sprintf("tx %d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.TransferTX(ctx, TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// checking result
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		assert.NoError(t, err)

		result := <-results
		assert.NotEmpty(t, result)

		//checking Transfer
		transfer := result.Transfer
		assert.NotEmpty(t, transfer)
		assert.Equal(t, account1.ID, transfer.FromAccountID)
		assert.Equal(t, account2.ID, transfer.ToAccountID)
		assert.Equal(t, amount, transfer.Amount)
		assert.NotZero(t, transfer.ID)
		assert.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		assert.NoError(t, err)

		// checking enrty
		fromEntry := result.FromEntry
		assert.NotEmpty(t, fromEntry)
		assert.Equal(t, account1.ID, fromEntry.AccountID)
		assert.Equal(t, -amount, fromEntry.Amount)
		assert.NotZero(t, fromEntry.ID)
		assert.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		assert.NoError(t, err)

		toEntry := result.ToEntry
		assert.NotEmpty(t, toEntry)
		assert.Equal(t, account2.ID, toEntry.AccountID)
		assert.Equal(t, amount, toEntry.Amount)
		assert.NotZero(t, toEntry.ID)
		assert.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		assert.NoError(t, err)

		// USING TEST DRIVEN
		// checking Account
		fromAccount := result.FromAccount
		assert.NotEmpty(t, fromAccount)
		assert.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		assert.NotEmpty(t, toAccount)
		assert.Equal(t, account2.ID, toAccount.ID)

		// Printing the account balance after each transaction
		fmt.Println(">> Tx:", fromAccount.Balance, toAccount.Balance)

		// check accounts' balance
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		assert.Equal(t, diff1, diff2)
		assert.True(t, diff1 > 0)
		assert.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		assert.True(t, k >= 1 && k <= n)
		assert.NotContains(t, existed, k)
		existed[k] = true
	}

	//  checking the final updated balance
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	assert.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	assert.NoError(t, err)

	// Printing the account  balance after transaction
	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)

	assert.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	assert.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}
