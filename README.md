HOWTO

1. `docker-compose up`
2. Wait for a ledger with 3 transactions, example:
   ```
   2022-09-24T21:22:38.841 GD5KD [Ledger INFO] Got consensus: [seq=23, prev=76a607, txs=3, ops=3, sv: [ SIGNED@GD5KD txH: d58dbe, ct: 1664054558, upgrades: [ ] ]]
   ```
3. `docker cp captive:/stream .` will copy meta stream to current dir.
4. Inspect the stream. I prepared a small helper using Go: `go run lcm.go`