#!/usr/bin/env bash

do_issuedenom()
{
	echo "y" | rizond tx nft issue nft-001 --from=my_key --name=game --symbol="https://gblobscdn.gitbook.com/spaces%2F-MaSOXNsWsuvtyz9G2Xu%2Favatar-1623141242009.png\?alt\=media" --mint-restricted=true --update-restricted=true --schema="https://lh3.googleusercontent.com/8yUPxXmB49iPZDrnv2tuclkh_Ea19wULbtCuTQvY0Jg4kXiyDa2xeh7RscmC42nFeirqpzL4fxeJQ7N_bk8qd5s2cpa7isErKl5bVw\=w600" --chain-id=my_testnet --fees=10stake --keyring-backend test
}

do_qtx()
{
	rizond q tx $1
}

do_test()
{
	cmd=$1
	echo "cmd : $cmd"
	if [ $cmd == "issuedenom" ]; then
		do_issuedenom
	elif [ $cmd == "qtx" ]; then
		do_qtx $2 
	fi
}

echo $cmd
if [ $# -le 1 ]; then
	echo "Usage: $0 issuedenom"
	echo "Usage: $0 qtx [txhash]"
	exit -1
else
	do_test $1 $2
fi

