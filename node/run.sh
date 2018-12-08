#!/bin/bash

# This is public address of operator at test geth node
#ADDRESS="0xfFb85D98FC040123756ABCA2a7489A2E2716189b"
#
## remove old node at demon
#sudo docker rm -f geth
#
## run test geth node at demon (development config at truffle)
#sudo docker run -it -d -p 8545:8545 --name geth -p 8546:8546 -p 30303:30303 erage/gethplasma
#
#cd ../contracts
#
#sleep 3
## give plasma owner 100 ether
#( echo "eth.sendTransaction({from: eth.accounts[0],to:'${ADDRESS}',value: web3.toWei(100, 'ether')})" ) | geth attach http://127.0.0.1:8545
#

# deploy bankex plasma contract to ethereum
truffle migrate --reset
#truffle compile
#node migrations/deploy.js



cd ../node/operator/

# run operator
GIN_MODE=release go run operator.go &

cd ../verifier/

# clear db before run - NOT FOR PRODUCTION
# sudo rm -rf database

# run verifier
GIN_MODE=release go run verifier.go
