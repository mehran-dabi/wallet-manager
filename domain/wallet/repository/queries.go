package repository

const (
	walletTableName = "wallets"
)

const (
	insertWalletQuery = `INSERT INTO ` + walletTableName + ` SET user_id = ?, created_at = NOW(), updated_at = NOW()`

	addBalanceQuery = `UPDATE ` + walletTableName + ` SET balance = balance + ?, updated_at = NOW() WHERE id = ?`

	subtractBalanceQuery = `UPDATE ` + walletTableName + ` SET balance = balance - ?, updated_at = NOW() WHERE id = ?`

	getWalletByIDQuery = `SELECT id, balance, user_id, created_at, updated_at FROM ` + walletTableName + ` WHERE id = ?`

	getWalletByUserIDQuery = `SELECT id, balance, user_id, created_at, updated_at FROM ` + walletTableName + ` WHERE user_id = ?`
)
