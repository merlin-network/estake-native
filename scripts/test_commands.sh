#!/bin/sh

estaked config chain-id test
estaked config keyring-backend test
estaked config
estaked tx gov submit-proposal change-cosmos-validator-weights scripts/data/cosmos_validator_set_proposal.json --from test --gas auto -y -b block
estaked tx gov vote 1 yes --from test -y -b block
estaked tx gov submit-proposal  change-orchestrator-validator-weights scripts/data/oracle_validator_set_proposal.json --from test --gas auto -y -b block
estaked tx gov vote 2 yes --from test  -y -b block
estaked tx bank send did:fury:e1hcqg5wj9t42zawqkqucs7la85ffyv08ljhhesu did:fury:e1lcck2cxh7dzgkrfk53kysg9ktdrsjj6jzkd4ea 10000000stake --chain-id test --from test --keyring-backend=test -y -b block
estaked tx bank send did:fury:e1lcck2cxh7dzgkrfk53kysg9ktdrsjj6jzkd4ea did:fury:e1hcqg5wj9t42zawqkqucs7la85ffyv08ljhhesu 10stake --chain-id test --from test1 --keyring-backend=test -y -b block
sleep 25
estaked q gov proposals
estaked tx cosmos set-orchestrator-address did:fury:evaloper1hcqg5wj9t42zawqkqucs7la85ffyv08lmnhye9 did:fury:e1lcck2cxh7dzgkrfk53kysg9ktdrsjj6jzkd4ea --from test -y -b block --gas 400000
estaked tx gov submit-proposal enable-module scripts/data/module_enable_proposal.json --from test -y -b block
estaked tx gov vote 3 yes --from test -y -b block


sleep 25
estaked tx cosmos incoming did:fury:e10khgeppewe4rgfrcy809r9h00aquwxxxrk6glr did:fury:e1lcck2cxh7dzgkrfk53kysg9ktdrsjj6jzkd4ea 10000000stake cosmoshub-4 AE9ADDF593D45DDB09C8371F534AA773EB8CF288F63B09C160110338D362177B 100000 --from test1 -y -b block --gas 400000
estaked q bank balances did:fury:e10khgeppewe4rgfrcy809r9h00aquwxxxrk6glr
