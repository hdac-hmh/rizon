#!/usr/bin/env bash

do_issuedenom()
{
  denom_id=$1
  update_block=$2
	echo "y" | rizond tx nft issue $denom_id --from=my_key --name=game --symbol="https://gblobscdn.gitbook.com/spaces%2F-MaSOXNsWsuvtyz9G2Xu%2Favatar-1623141242009.png\?alt\=media" --mint-restricted=true --update-restricted=$update_block --schema="https://lh3.googleusercontent.com/8yUPxXmB49iPZDrnv2tuclkh_Ea19wULbtCuTQvY0Jg4kXiyDa2xeh7RscmC42nFeirqpzL4fxeJQ7N_bk8qd5s2cpa7isErKl5bVw\=w600" --chain-id=my_testnet --fees=10stake --keyring-backend test
}

do_qtx()
{
	rizond q tx $1
}

do_qdenom()
{
  denom_id=$1
  if [ $denom_id == "all" ]; then
    rizond q nft denoms
  else
    rizond q nft denom $denom_id
  fi
}

do_mint()
{
  denom_id=$1
  nft_id=$2
  nft_name="test-nft"
  nft_data="https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
  uri="https://gblobscdn.gitbook.com/spaces%2F-MaSOXNsWsuvtyz9G2Xu%2Favatar-1623141242009.png\?alt\=media"
  echo "y" | rizond tx nft mint $denom_id $nft_id --uri=$uri --name=$nft_name --data=$nft_data --recipient=$(rizond keys show my_key -a --keyring-backend test) --from=my_key --chain-id=my_testnet --fees=10stake --keyring-backend test
}

do_qtoken()
{
  denom_id=$1
  nft_id=$2
  rizond q nft token $denom_id $nft_id
}

do_edit()
{
  denom_id=$1
  nft_id=$2
  uri=$3
  echo "y" | rizond tx nft edit $denom_id $nft_id --uri=$uri --from=my_key --chain-id=my_testnet --fees=10stake --keyring-backend test
}

do_test()
{
	cmd=$1
	if [ $cmd == "issuedenom" ]; then
		do_issuedenom $2 $3
	elif [ $cmd == "qtx" ]; then
		do_qtx $2
  elif [ $cmd == "qdenom" ]; then
    do_qdenom $2
  elif [ $cmd == "mint" ]; then
    do_mint $2 $3
  elif [ $cmd == "qtoken" ]; then
    do_qtoken $2 $3
  elif [ $cmd == "edit" ]; then
    do_edit $2 $3 $4
  else
    echo "Unknown Command $cmd"
	fi
}

if [ $# -le 1 ]; then
	echo "Usage: $0 issuedenom [denom-id] [updateblock=true or false]"
	echo "Usage: $0 qtx [txhash]"
	echo "Usage: $0 qdenom [denom-id | all]"
	echo "Usage: $0 mint [denom-id] [nft-id]"
	echo "Usage: $0 qtoken [denom-id] [nft-id]"
	echo "Usage: $0 edit [denom-id] [nft-id] [uri]"
	exit -1
elif [ $# -eq 2 ]; then
  do_test $1 $2
elif [ $# -eq 3 ]; then
  do_test $1 $2 $3
elif [ $# -eq 4 ]; then
  do_test $1 $2 $3 $4
fi
