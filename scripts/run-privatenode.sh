#!/usr/bin/env bash

rizonhome=~/.rizon

if [ -d $rizonhome ];then
 	rm -rf $rizonhome
else
	echo "$rizonhome doesn't exist"
fi

rizond init my_node --chain-id my_testnet
rizond keys add my_key --keyring-backend test
rizond add-genesis-account $(rizond keys show my_key -a --keyring-backend test) 10000000000stake
rizond gentx my_key 1000000000stake --chain-id my_testnet --keyring-backend test
rizond collect-gentxs
rizond validate-genesis
rizond start
