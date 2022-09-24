package main

import (
	"fmt"
	"os"

	xdr3 "github.com/stellar/go-xdr/xdr3"
	"github.com/stellar/go/xdr"
)

func main() {
	file, err := os.Open("./stream")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := xdr3.NewDecoder(file)

	for {
		_, err := xdr.ReadFrameLength(decoder)
		if err != nil {
			panic(err)
		}

		var xlcm xdr.LedgerCloseMeta
		_, err = xlcm.DecodeFrom(decoder)
		if err != nil {
			panic(err)
		}

		if xlcm.V == 2 {
			// Ledger with invalid data has 3 txs
			if len(xlcm.V2.TxProcessing) == 3 {
				fmt.Printf("%+v\n", xlcm.V2)
				// Nothing interesting later on...
				break
			}
		}
	}
}
