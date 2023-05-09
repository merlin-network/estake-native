#!/bin/bash
<<<<<<< HEAD
interchain-queries keys restore --chain $(jq -r ".chains[0].name" /configs/keys.json) --home /icq perKey
=======
interchain-queries keys restore --chain $(jq -r ".chains[0].name" /configs/keys.json) --home /icq elyKey
>>>>>>> 4b25098 (::)
