# Sequencer for SYB

SYB sequencer is a zk-rollup sequencer designed to compute uniqueness score for each user account based on the vouches made between accounts. This score will reflect the interactions and endorsements that one account provides to another. Acting as a backend middleware, sequencer facilitates communication between the smart contract and the Circuit, which is responsible for calculating the uniqueness score. 

By utilizing a zk-rollup architecture, we aim to significantly reduce the computational costs associated with calculating scores directly through smart contracts, which would otherwise be prohibitively high. This approach ensures efficiency and scalability while maintaining the integrity of our computations.

## Setup
```bash
cp .env.example .env
brew install go-task # if you are on different OS: https://taskfile.dev/installation/
brew install golangci-lint

# from the root directory
cp githooks/pre-commit .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
```

If you are using VSCode, install the [go extension](https://marketplace.visualstudio.com/items?itemName=golang.go) and add the following to your user settings.json file to lint `go` files on save (note that if you push unlinted go files, triggered actions will fail):
```json
"go.lintTool": "golangci-lint",
"go.lintFlags": [
    "--fast"
],
"go.lintOnSave": "file",
"go.formatTool": "gofmt",
"go.useLanguageServer": true,
"[go]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
        "source.organizeImports": "always"
    }
}
```

## Run Sequencer
```bash
task run-seq
```

## Update keystore address
```bash
mkdir -p var/tokamak/ethkeystore  # create this directory inside sequencer

# Once file is created,New address will be given in terminal, Please update this address as forger address in cfg.toml file

cfg.Coordinator.ForgerAddress  # Update this address
```

## Run Tests
```bash
task test-<name> # for example: task test-historydb
```

## Architecture

### E2E flow
<img src="../doc/images/sequencer_e2e_flow.png" />

### Sync flow ([interactive link](https://viewer.diagrams.net/?tags=%7B%7D&lightbox=1&highlight=0000ff&edit=_blank&layers=1&nav=1#G10tKc2c3VyREzzdtekl2dcNMwI4HWOSfr#%7B%22pageId%22%3A%22mWZ3KBgQXANqmTgwyxpi%22%7D))
<img src="../doc/images/sequencer_sync_flow.png" />

## CMD to run sequencer in sync mode

```
go run main.go run --mode sync --cfg cfg.toml
```


### Coord flow ([interactive link](https://viewer.diagrams.net/?tags=%7B%7D&lightbox=1&highlight=0000ff&edit=_blank&layers=1&nav=1#G10tKc2c3VyREzzdtekl2dcNMwI4HWOSfr#%7B%22pageId%22%3A%22MOlNjBzEnPvgUMVi-x9F%22%7D))
<img src="../doc/images/sequencer_coord_flow.png" />

## CMD to run sequencer in coord mode

```
go run main.go run --mode coord --cfg cfg.toml
```
# Generate .go file from the Sybil contract
Navigate to the contracts folder inside the tokamak-sybil-resistance-mvp directory.

Run the followig command:

```sh sybil_go.sh```

## Testing Smart Contract Events with the Sequencer

### Step 1: Deploy and Trigger Events
First, we need to publish events on the smart contract using a test script:

1. Set the Sybil contract address in `contracts/.env`:
```bash
SYBIL_CONTRACT_ADDRESS="0x14F39A3380100f724c075814d4540a3a34f93C53"
```

2. Run the test script to create a account and deposit event:
```bash
forge script script/TestSyncEvents.s.sol:TestSyncEvents --rpc-url <rpc> --private-key <private-key> --broadcast
```

### Step 2: Configure Sequencer
After the transaction is confirmed:

1. Find the transaction hash in the block explorer
   - Example hash: `0x19517b80de62a1ee22212ff5a3b26cd216249e3b4daa88be13a8f07dd7a51694` (Hash which is used with the deployed contract to test the events)

2. Note the block number for this transaction
3. Update `sequencer/.env` to start syncing from slightly before this block: (Keeping the block number slightly before the transaction block to monitor the sync functionality easily)
```bash
ROLLUP_START_BLOCK_NUM=<transaction_block_number - 1>
```

### Step 3: Run and Verify
1. Start the sequencer in sync mode:
```bash
task run-seq
```

2. Verify the event processing:
   - Check the `tx` table in the database
   - You should see a new entry with type `CreateAccountDeposit`
