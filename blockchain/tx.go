package blockchain

//TxOutput represents a transaction in the blockchain
//For Example, I sent you 5 coins. Value would == 5, and it would have my unique PubKey
type TxOutput struct {
	Value int
	//取引のコインの量
	PubKey string
	//Pubkeyは、アウトプット内のコインをアンロックするために必要。送信者であることを示す。
}

//TxInput:以前のTxOutputへの参照.
type TxInput struct {
	ID []byte
	//特定のoutputが内部にあるトランザクションを見つける.
	Out int
	Sig string
	//アウトプットのPubKeyにデータを追加するスクリプト
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
