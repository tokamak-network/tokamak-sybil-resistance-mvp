# Build the project and generate the necessary artifacts
forge build

# Compile the Solidity contracts and generate binary files (bin)
forge compile --extra-output-files bin

# Extract the ABI from the JSON file and save it as a .abi file
jq .abi out/mvp/Sybil.sol/Sybil.json > out/mvp/Sybil.sol/Sybil.abi

# Use abigen to generate a Go package from the ABI and binary, and save it in the abi folder
abigen --abi=out/mvp/Sybil.sol/Sybil.abi --bin=out/mvp/Sybil.sol/Sybil.bin --pkg=sybil --out=../sequencer/eth/contracts/sybil.go
