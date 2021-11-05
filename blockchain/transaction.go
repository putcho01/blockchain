package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

const reward = 100

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

//Important to note that each output is Indivisible.
//You cannot "make change" with any output.
//If the Value is 10, in order to give someone 5, we need to make two five coin outputs.

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]

}

//ノード上の誰かがブロックの "マイニング "に成功したときに実行される関数
func CoinbaseTx(toAddress, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", toAddress)
	}
	//Since this is the "first" transaction of the block, it has no previous output to reference.
	//This means that we initialize it with no ID, and it's OutputIndex is -1
	txIn := TxInput{[]byte{}, -1, data}
	//txOut will represent the amount of tokens(reward) given to the person(toAddress) that executed CoinbaseTx
	txOut := TxOutput{reward, toAddress} // You can see it follows {value, PubKey}

	tx := Transaction{nil, []TxInput{txIn}, []TxOutput{txOut}}

	return &tx

}
func (tx *Transaction) IsCoinbase() bool {
	//トランザクションをチェックし、新しく鋳造された "コイン "である場合にのみtrueを返す。
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: Not enough funds!")
	}
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})

	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
