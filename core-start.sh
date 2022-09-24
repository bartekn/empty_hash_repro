#!/usr/bin/env bash
set -e

source /etc/profile
# work within the current docker working dir
if [ ! -f "./stellar-core.cfg" ]; then
   cp /stellar-core.cfg ./
fi

# initialize new db
stellar-core new-db

# initialize for new history archive path, remove any pre-existing on same path from base image
rm -rf ./history
stellar-core new-hist vs

# serve history archives to horizon on port 1570
pushd ./history/vs/
python3 -m http.server 1570 &
popd

exec stellar-core run --console &

sleep 10

# Upgrade to Soroban
curl "http://localhost:11626/upgrades?mode=set&upgradetime=1970-01-01T00:00:00Z&protocolversion=20&basereserve=5000000"

sleep infinity