package blkparser

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestParseTx(t *testing.T) {
	// Arbitrarily chosen MAZA transaction.
	rawTx, err := hex.DecodeString("010000000188592292671feedac902dc146017cfc500743f90159928771bc9c87553560160000000006a47304402200aa84b1ade118ecf874452a9320d05f989f5e840b64a91b7dc353342ac3f6e46022018b9dcf2ff67ea2705574133c905c58dd8e1353014387a35b63ea6aa530b4ca4012103c05cedc14205dd668786238dfa19eff1782d35f6ae12fbbe0c3ce89e5a266123ffffffff02002d3101000000001976a914ba22650461f865965ce648040049bedc1127450388aca14d9b1c000000001976a914f124cfab9186731879ec4a4d362a1f6b087b708988ac00000000")
	if err != nil {
		t.Error(err)
	}
	tx, _ := NewTx(rawTx)

	if tx.Version != 1 {
		t.Error("For Tx version, expected 1, got", tx.Version)
	}

	// Test Tx input.
	if tx.TxInCnt != 1 {
		t.Error("For TxIn count, expected 1, got", tx.TxInCnt)
	}

	txIn := tx.TxIns[0]
	if txIn.InputHash != "6001565375c8c91b77289915903f7400c5cf176014dc02c9daee1f6792225988" {
		t.Error("For tx input 0 hash, expected 6001565375c8c91b77289915903f7400c5cf176014dc02c9daee1f6792225988, got", txIn.InputHash)
	}
	if txIn.InputVout != 0 {
		t.Error("For tx input 0 index, expected 0, got", txIn.InputVout)
	}
	actualScriptSig, _ := hex.DecodeString("47304402200aa84b1ade118ecf874452a9320d05f989f5e840b64a91b7dc353342ac3f6e46022018b9dcf2ff67ea2705574133c905c58dd8e1353014387a35b63ea6aa530b4ca4012103c05cedc14205dd668786238dfa19eff1782d35f6ae12fbbe0c3ce89e5a266123")
	if bytes.Equal(txIn.ScriptSig, actualScriptSig) != true {
		t.Errorf("For tx input 0 script, expected 47304402200aa84b1ade118ecf874452a9320d05f989f5e840b64a91b7dc353342ac3f6e46022018b9dcf2ff67ea2705574133c905c58dd8e1353014387a35b63ea6aa530b4ca4012103c05cedc14205dd668786238dfa19eff1782d35f6ae12fbbe0c3ce89e5a266123, got %x", txIn.ScriptSig)
	}
	if txIn.Sequence != 0xffffffff {
		t.Error("For tx input 0 sequence, expected 0xffffffff, got", txIn.Sequence)
	}

	// Test Tx output.
	if tx.TxOutCnt != 2 {
		t.Error("For TxOut count, expected 2, got", tx.TxOutCnt)
	}

	txOut := tx.TxOuts[0]
	if txOut.Addr != "MQsM7tXb2rSDzymQzMxXAMeDrRbMmFbpuw" {
		t.Error("For tx output 0 address, expected MQsM7tXb2rSDzymQzMxXAMeDrRbMmFbpuw, got", txOut.Addr)
	}
	if txOut.Value != 20000000 {
		t.Error("For tx output 0 value, expected 20000000, got", txOut.Value)
	}
	actualOutputScript, _ := hex.DecodeString("76a914ba22650461f865965ce648040049bedc1127450388ac")
	if !bytes.Equal(actualOutputScript, txOut.Pkscript) {
		t.Errorf("For tx output 0 script, expected 76a914ba22650461f865965ce648040049bedc1127450388ac, got %x", txOut.Pkscript)
	}

	txOut = tx.TxOuts[1]
	if txOut.Addr != "MVtDBMCnBRPnALQPAbmzC7X5yNKmTuusWt" {
		t.Error("For tx output 1 address, expected MVtDBMCnBRPnALQPAbmzC7X5yNKmTuusWt, got", txOut.Addr)
	}
	if txOut.Value != 479940001 {
		t.Error("For tx output 1 value, expected 479940001, got", txOut.Value)
	}
	actualOutputScript, _ = hex.DecodeString("76a914f124cfab9186731879ec4a4d362a1f6b087b708988ac")
	if !bytes.Equal(actualOutputScript, txOut.Pkscript) {
		t.Errorf("For tx output 1 script, expected 76a914f124cfab9186731879ec4a4d362a1f6b087b708988ac, got %x", txOut.Pkscript)
	}

	if tx.LockTime != 0 {
		t.Error("For Tx locktime, expected 0, got", tx.LockTime)
	}

}
