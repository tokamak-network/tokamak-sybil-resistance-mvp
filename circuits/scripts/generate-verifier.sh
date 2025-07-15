#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}=================================${NC}"
echo -e "${BLUE}  Circomkit Verifier Generator   ${NC}"
echo -e "${BLUE}=================================${NC}"
echo ""

# Record script start time
SCRIPT_START_TIME=$(date +%s)

# Get script directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="$( cd "$SCRIPT_DIR/.." && pwd )"

# Change to project directory
cd "$PROJECT_DIR" || { echo -e "${RED}Failed to change to project directory${NC}"; exit 1; }

# Check if we're in the right directory
if [ ! -f "package.json" ]; then
    echo -e "${RED}❌ Error: package.json not found. Please run from circuits directory${NC}"
    exit 1
fi

# Check if circomkit is installed
if ! npm list circomkit >/dev/null 2>&1; then
    echo -e "${YELLOW}Installing circomkit...${NC}"
    npm install circomkit || { echo -e "${RED}Failed to install circomkit${NC}"; exit 1; }
fi

# Parameter input with validation
echo -e "${BLUE}Circuit Parameters:${NC}"
read -p "Number of transactions (nTx) [16]: " NTX
read -p "Tree depth (nLevels) [20]: " NLEVELS
NTX=${NTX:-16}
NLEVELS=${NLEVELS:-20}

# Validate parameters
if ! [[ "$NTX" =~ ^[0-9]+$ ]] || [ "$NTX" -lt 1 ]; then
    echo -e "${RED}❌ Invalid nTx: must be a positive integer${NC}"
    exit 1
fi

if ! [[ "$NLEVELS" =~ ^[0-9]+$ ]] || [ "$NLEVELS" -lt 1 ]; then
    echo -e "${RED}❌ Invalid nLevels: must be a positive integer${NC}"
    exit 1
fi

echo ""
echo -e "${YELLOW}Configuration:${NC}"
echo -e "  nTx: ${GREEN}$NTX${NC}"
echo -e "  nLevels: ${GREEN}$NLEVELS${NC}"
echo ""

# Backup existing config files if they exist
if [ -f "circomkit.json" ]; then
    cp circomkit.json circomkit.json.backup
    echo -e "${YELLOW}Backed up existing circomkit.json${NC}"
fi

if [ -f "circuits.json" ]; then
    cp circuits.json circuits.json.backup
    echo -e "${YELLOW}Backed up existing circuits.json${NC}"
fi

# Update circuits.json with new parameters
echo -e "${YELLOW}Updating circuits.json...${NC}"
cat > circuits.json << EOF
{
  "batch_main": {
    "file": "syb_rollup/batch-main",
    "template": "BatchMain",
    "params": [$NTX, $NLEVELS]
  }
}
EOF

echo -e "${GREEN}✅ Configuration updated${NC}"
echo ""

# Create necessary directories
mkdir -p build ptau

# Function to run command with better error handling
run_command() {
    local step=$1
    local desc=$2
    local cmd=$3
    
    echo -e "${YELLOW}Step $step: $desc...${NC}"
    echo -e "${BLUE}Executing: $cmd${NC}"
    
    # Record start time
    local start_time=$(date +%s)
    
    # Run command and capture output
    if output=$(eval "$cmd" 2>&1); then
        local end_time=$(date +%s)
        local duration=$((end_time - start_time))
        echo "$output"
        echo -e "${GREEN}✅ $desc completed in ${duration} seconds${NC}"
        echo ""
    else
        local end_time=$(date +%s)
        local duration=$((end_time - start_time))
        echo "$output"
        echo -e "${RED}❌ $desc failed after ${duration} seconds${NC}"
        echo -e "${RED}Command: $cmd${NC}"
        echo -e "${RED}Exit code: $?${NC}"
        exit 1
    fi
}

# Execute circomkit commands
run_command "1" "Compiling circuit" "npx circomkit compile batch_main"

# Check if r1cs file was generated
if [ ! -f "build/batch_main/batch_main.r1cs" ]; then
    echo -e "${RED}❌ Error: r1cs file not generated${NC}"
    exit 1
fi

run_command "2" "Circuit info" "npx circomkit info batch_main"

# Calculate required ptau power
echo -e "${YELLOW}Calculating required ptau power...${NC}"
CONSTRAINTS=$(npx circomkit info batch_main 2>/dev/null | grep -i "constraints" | grep -o '[0-9]\+' | tail -1)
if [ -z "$CONSTRAINTS" ]; then
    echo -e "${YELLOW}Could not determine constraints, using default ptau power 21${NC}"
    PTAU_POWER=21
else
    # Calculate required power (log2 of constraints, rounded up)
    # For 1818151 constraints, we need at least power 21
    if [ "$CONSTRAINTS" -lt 1024 ]; then
        PTAU_POWER=10
    elif [ "$CONSTRAINTS" -lt 2048 ]; then
        PTAU_POWER=11
    elif [ "$CONSTRAINTS" -lt 4096 ]; then
        PTAU_POWER=12
    elif [ "$CONSTRAINTS" -lt 8192 ]; then
        PTAU_POWER=13
    elif [ "$CONSTRAINTS" -lt 16384 ]; then
        PTAU_POWER=14
    elif [ "$CONSTRAINTS" -lt 32768 ]; then
        PTAU_POWER=15
    elif [ "$CONSTRAINTS" -lt 65536 ]; then
        PTAU_POWER=16
    elif [ "$CONSTRAINTS" -lt 131072 ]; then
        PTAU_POWER=17
    elif [ "$CONSTRAINTS" -lt 262144 ]; then
        PTAU_POWER=18
    elif [ "$CONSTRAINTS" -lt 524288 ]; then
        PTAU_POWER=19
    elif [ "$CONSTRAINTS" -lt 1048576 ]; then
        PTAU_POWER=20
    elif [ "$CONSTRAINTS" -lt 2097152 ]; then
        PTAU_POWER=21
    elif [ "$CONSTRAINTS" -lt 4194304 ]; then
        PTAU_POWER=22
    elif [ "$CONSTRAINTS" -lt 8388608 ]; then
        PTAU_POWER=23
    elif [ "$CONSTRAINTS" -lt 16777216 ]; then
        PTAU_POWER=24
    elif [ "$CONSTRAINTS" -lt 33554432 ]; then
        PTAU_POWER=25
    elif [ "$CONSTRAINTS" -lt 67108864 ]; then
        PTAU_POWER=26
    elif [ "$CONSTRAINTS" -lt 134217728 ]; then
        PTAU_POWER=27
    else
        PTAU_POWER=28
    fi
    echo -e "${BLUE}Constraints: $CONSTRAINTS, using ptau power: $PTAU_POWER${NC}"
fi

# Step 3: Check PTAU file
echo -e "${YELLOW}Step 3: Checking PTAU file...${NC}"
if [ "$PTAU_POWER" -eq 28 ]; then
    PTAU_FILE="./ptau/powersOfTau28_hez_final.ptau"
else
    PTAU_FILE="./ptau/powersOfTau28_hez_final_${PTAU_POWER}.ptau"
fi

if [ -f "$PTAU_FILE" ]; then
    echo -e "${GREEN}✅ Default PTAU file found: $PTAU_FILE${NC}"
else
    echo -e "${YELLOW}⚠️  Default PTAU file not found: $PTAU_FILE${NC}"
fi

echo ""

# Step 4: Circuit setup (Key generation)
echo -e "${YELLOW}Step 4: Circuit setup (Key generation)...${NC}"
echo -e "${BLUE}Using PTAU: $PTAU_FILE${NC}"
echo -e "${BLUE}Constraints: $CONSTRAINTS${NC}"
echo ""

# Check if keys already exist for current parameters
PKEY_FILE="build/batch_main/${NTX}_${NLEVELS}_groth16_pkey.zkey"
VKEY_FILE="build/batch_main/${NTX}_${NLEVELS}_groth16_vkey.json"
VERIFIER_FILE="build/batch_main/${NTX}_${NLEVELS}_groth16_verifier.sol"

if [ -f "$PKEY_FILE" ] && [ -f "$VKEY_FILE" ]; then
    echo -e "${GREEN}✅ Proving and verification keys already exist for parameters (nTx=${NTX}, nLevels=${NLEVELS})!${NC}"
    echo -e "${BLUE}Key files found:${NC}"
    echo -e "  - Proving key: $PKEY_FILE"
    echo -e "  - Verification key: $VKEY_FILE"
    echo ""
    
    echo -e "${YELLOW}Choose an option:${NC}"
    echo -e "  ${GREEN}1${NC}) Use existing keys (skip key generation)"
    echo -e "  ${YELLOW}2${NC}) Regenerate keys (DEMO mode - not for production)"
    echo -e "  ${RED}3${NC}) Exit"
    echo ""
    read -p "Select option [1]: " SETUP_CHOICE
    SETUP_CHOICE=${SETUP_CHOICE:-1}
    
    case $SETUP_CHOICE in
        1)
            echo -e "${GREEN}✅ Using existing keys${NC}"
            echo ""
            ;;
        2)
            echo -e "${RED}⚠️  DEMO MODE - NOT FOR PRODUCTION USE ⚠️${NC}"
            echo -e "${RED}This will regenerate keys with unsafe randomness!${NC}"
            echo ""
            read -p "Are you sure you want to continue? (y/N): " DEMO_CONFIRM
            if [[ "$DEMO_CONFIRM" =~ ^[Yy]$ ]]; then
                if [ -f "$PTAU_FILE" ]; then
                    echo -e "${YELLOW}Removing existing keys...${NC}"
                    rm -f "$PKEY_FILE"
                    rm -f "$VKEY_FILE"
                    rm -f "$VERIFIER_FILE"
                    
                    SELECTED_PTAU_FILE="$PTAU_FILE"
                    USE_EXISTING_ZKEY=false
                    DEMO_MODE=true
                    
                    echo -e "${YELLOW}Generating new keys in DEMO mode...${NC}"
                    # Continue to key generation below
                else
                    echo -e "${RED}❌ Auto-detected PTAU file not found: $PTAU_FILE${NC}"
                    echo -e "${YELLOW}Please download the PTAU file first${NC}"
                    echo ""
                fi
            else
                echo -e "${BLUE}Operation cancelled${NC}"
                echo ""
            fi
            ;;
        3)
            echo -e "${BLUE}Exiting...${NC}"
            exit 0
            ;;
        *)
            echo -e "${RED}Invalid option. Using existing keys.${NC}"
            ;;
    esac
else
    echo -e "${YELLOW}No existing keys found for parameters (nTx=${NTX}, nLevels=${NLEVELS}).${NC}"
    echo ""
    
    while true; do
        echo -e "${YELLOW}Choose key generation method:${NC}"
        echo -e "  ${GREEN}1${NC}) Use existing final zkey file (recommended for production)"
        echo -e "  ${YELLOW}2${NC}) Generate keys with DEMO mode (development only)"
        echo -e "  ${RED}3${NC}) Exit"
        echo ""
        read -p "Select option [1]: " NEW_KEY_CHOICE
        NEW_KEY_CHOICE=${NEW_KEY_CHOICE:-1}
        
        case $NEW_KEY_CHOICE in
            1)
                echo ""
                echo -e "${YELLOW}Enter path to final zkey file (Phase 2 complete, tab completion enabled):${NC}"
                echo -e "${BLUE}Note: This should be a .zkey file, not a .ptau file${NC}"
                read -e -p "Path: " FINAL_ZKEY_PATH
                if [ -f "$FINAL_ZKEY_PATH" ]; then
                    # Check if it's actually a zkey file
                    if [[ ! "$FINAL_ZKEY_PATH" =~ \.zkey$ ]]; then
                        echo -e "${RED}❌ Error: File doesn't appear to be a zkey file (expected .zkey extension)${NC}"
                        echo -e "${RED}   You provided: $FINAL_ZKEY_PATH${NC}"
                        echo -e "${YELLOW}   Please provide a zkey file generated from Phase 2 of the trusted setup${NC}"
                        echo ""
                        continue
                    fi
                    
                    echo -e "${GREEN}✅ Final zkey file found: $FINAL_ZKEY_PATH${NC}"
                    
                    # Skip verification for now - just use the file
                    SELECTED_ZKEY_FILE="$FINAL_ZKEY_PATH"
                    USE_EXISTING_ZKEY=true
                    break
                else
                    echo -e "${RED}❌ Zkey file not found: $FINAL_ZKEY_PATH${NC}"
                    echo ""
                fi
                ;;
            2)
                echo -e "${RED}⚠️  DEMO MODE - NOT FOR PRODUCTION USE ⚠️${NC}"
                echo -e "${RED}This will generate keys with unsafe randomness!${NC}"
                echo -e "${RED}Keys generated this way MUST NOT be used in production!${NC}"
                echo ""
                read -p "Are you sure you want to continue? (y/N): " DEMO_CONFIRM
                if [[ "$DEMO_CONFIRM" =~ ^[Yy]$ ]]; then
                    if [ -f "$PTAU_FILE" ]; then
                        SELECTED_PTAU_FILE="$PTAU_FILE"
                        USE_EXISTING_ZKEY=false
                        DEMO_MODE=true
                        break
                    else
                        echo -e "${RED}❌ Auto-detected PTAU file not found: $PTAU_FILE${NC}"
                        echo ""
                    fi
                else
                    echo -e "${BLUE}Operation cancelled${NC}"
                    echo ""
                fi
                ;;
            3)
                echo -e "${BLUE}Exiting...${NC}"
                exit 0
                ;;
            *)
                echo -e "${RED}Invalid option. Please try again.${NC}"
                echo ""
                ;;
        esac
    done
    echo ""
fi

# Generate or use existing keys
if [ ! -f "$PKEY_FILE" ] || [ ! -f "$VKEY_FILE" ]; then
    START_TIME=$(date +%s)
    
    if [ "$USE_EXISTING_ZKEY" = true ]; then
        # Use existing final zkey file
        echo -e "${YELLOW}Using existing final zkey file...${NC}"
        
        # Copy the zkey file to our parameter-specific location
        cp "$SELECTED_ZKEY_FILE" "$PKEY_FILE"
        echo -e "${GREEN}✅ Proving key copied: $PKEY_FILE${NC}"
        
        # Export verification key
        echo -e "${BLUE}Exporting verification key...${NC}"
        npx snarkjs zkey export verificationkey "$PKEY_FILE" "$VKEY_FILE"
        echo -e "${GREEN}✅ Verification key exported: $VKEY_FILE${NC}"
        
        END_TIME=$(date +%s)
        DURATION=$((END_TIME - START_TIME))
        echo -e "${GREEN}✅ Key setup completed in ${DURATION} seconds${NC}"
        echo ""
        
    else
        # Perform circuit-specific setup (Phase 2 only - Phase 1 is the PTAU file)
        if [ "$DEMO_MODE" = true ]; then
            echo -e "${RED}⚠️  PERFORMING DEMO SETUP - NOT FOR PRODUCTION ⚠️${NC}"
        fi
        
        echo -e "${BLUE}Creating initial zkey from PTAU (circuit-specific setup)...${NC}"
        echo -e "${YELLOW}This may take a few minutes. Please wait...${NC}"
        npx snarkjs groth16 setup "build/batch_main/batch_main.r1cs" "$SELECTED_PTAU_FILE" "build/batch_main/circuit_0000.zkey"
        
        if [ "$DEMO_MODE" = true ]; then
            echo -e "${BLUE}Adding DEMO contribution (unsafe)...${NC}"
            # Use --entropy flag for non-interactive mode
            npx snarkjs zkey contribute "build/batch_main/circuit_0000.zkey" "build/batch_main/circuit_0001.zkey" --name="Demo Contributor" -e="demo_random_$(date +%s)"
            
            echo -e "${BLUE}Applying random beacon phase...${NC}"
            npx snarkjs zkey beacon "build/batch_main/circuit_0001.zkey" "$PKEY_FILE" 0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f 10 -n="Final Beacon"
        fi
        
        echo -e "${BLUE}Verifying final zkey...${NC}"
        npx snarkjs zkey verify "build/batch_main/batch_main.r1cs" "$SELECTED_PTAU_FILE" "$PKEY_FILE"
        
        echo -e "${BLUE}Exporting verification key...${NC}"
        npx snarkjs zkey export verificationkey "$PKEY_FILE" "$VKEY_FILE"
        
        # Clean up intermediate files
        rm -f build/batch_main/circuit_000*.zkey
        
        END_TIME=$(date +%s)
        DURATION=$((END_TIME - START_TIME))
        
        if [ "$DEMO_MODE" = true ]; then
            echo -e "${GREEN}✅ DEMO key generation completed in ${DURATION} seconds${NC}"
            echo -e "${RED}⚠️  WARNING: These keys are NOT safe for production use!${NC}"
        else
            echo -e "${GREEN}✅ Trusted setup completed in ${DURATION} seconds${NC}"
            echo -e "${GREEN}✅ Keys are suitable for production use${NC}"
        fi
        echo ""
    fi
fi

# Add a check before contract generation
echo -e "${YELLOW}Checking generated files before contract export...${NC}"
if [ -f "$VKEY_FILE" ]; then
    echo -e "${GREEN}✅ Verification key found: $VKEY_FILE${NC}"
    
    # Use the parameter-specific vkey for contract generation
    if [ ! -f "build/batch_main/groth16_vkey.json" ]; then
        cp "$VKEY_FILE" "build/batch_main/groth16_vkey.json"
        echo -e "${BLUE}Copied verification key for contract generation${NC}"
    fi
else
    echo -e "${RED}❌ Verification key not found: $VKEY_FILE${NC}"
    exit 1
fi

# Step 5: Generate Solidity verifier contract
echo -e "${YELLOW}Step 5: Exporting Solidity verifier contract...${NC}"

START_TIME=$(date +%s)

# Run contract command and monitor output
{
    # Use a temporary file to capture output
    TEMP_OUTPUT=$(mktemp)
    
    # Start the contract command in background
    npx circomkit contract batch_main > "$TEMP_OUTPUT" 2>&1 &
    CONTRACT_PID=$!
    
    # Monitor the output file
    tail -f "$TEMP_OUTPUT" &
    TAIL_PID=$!
    
    # Wait for completion signal
    while kill -0 $CONTRACT_PID 2>/dev/null; do
        if grep -q "Created at:.*groth16_verifier.sol" "$TEMP_OUTPUT" 2>/dev/null; then
            # Kill the contract process
            kill $CONTRACT_PID 2>/dev/null
            sleep 1
            kill -9 $CONTRACT_PID 2>/dev/null
            break
        fi
        sleep 1
    done
    
    # Stop tailing
    kill $TAIL_PID 2>/dev/null
    
    # Calculate duration
    END_TIME=$(date +%s)
    DURATION=$((END_TIME - START_TIME))
    
    # Check if we found the completion signal
    if grep -q "Created at:.*groth16_verifier.sol" "$TEMP_OUTPUT" 2>/dev/null; then
        echo -e "${GREEN}✅ Verifier contract generated in ${DURATION} seconds${NC}"
        
        # Copy verifier contract to parameter-specific name
        if [ -f "build/batch_main/groth16_verifier.sol" ]; then
            cp "build/batch_main/groth16_verifier.sol" "$VERIFIER_FILE"
            echo -e "${GREEN}✅ Verifier contract saved as: $VERIFIER_FILE${NC}"
        fi
        echo ""
    else
        echo -e "${RED}❌ Contract generation may have failed${NC}"
        rm -f "$TEMP_OUTPUT"
        exit 1
    fi
    
    # Clean up
    rm -f "$TEMP_OUTPUT"
}

# Verify generated files
echo -e "${YELLOW}Verifying generated files...${NC}"
MISSING_FILES=()

[ ! -f "$VERIFIER_FILE" ] && MISSING_FILES+=("${NTX}_${NLEVELS}_groth16_verifier.sol")
[ ! -f "$PKEY_FILE" ] && MISSING_FILES+=("${NTX}_${NLEVELS}_groth16_pkey.zkey")
[ ! -f "$VKEY_FILE" ] && MISSING_FILES+=("${NTX}_${NLEVELS}_groth16_vkey.json")

if [ ${#MISSING_FILES[@]} -eq 0 ]; then
    echo -e "${GREEN}✅ All files generated successfully${NC}"
    echo ""
    
    # Calculate total time
    SCRIPT_END_TIME=$(date +%s)
    TOTAL_TIME=$((SCRIPT_END_TIME - SCRIPT_START_TIME))
    MINUTES=$((TOTAL_TIME / 60))
    SECONDS=$((TOTAL_TIME % 60))
    
    echo -e "${GREEN}🎉 Verifier generation completed!${NC}"
    echo -e "${BLUE}Total time: ${MINUTES}m ${SECONDS}s${NC}"
    echo ""
    echo -e "${BLUE}Generated files (nTx=${NTX}, nLevels=${NLEVELS}):${NC}"
    echo -e "  📄 Verifier contract: ${GREEN}$VERIFIER_FILE${NC}"
    echo -e "  🔑 Proving key: ${GREEN}$PKEY_FILE${NC}"
    echo -e "  🔓 Verification key: ${GREEN}$VKEY_FILE${NC}"
    echo -e "  📊 Circuit info: ${GREEN}build/batch_main/batch_main_artifacts.json${NC}"
    echo ""
    
    # Show file sizes
    echo -e "${BLUE}File sizes:${NC}"
    ls -lh build/batch_main/${NTX}_${NLEVELS}_groth16_* | awk '{print "  " $9 ": " $5}'
    echo ""
    
    echo -e "${BLUE}Next steps:${NC}"
    echo -e "  1. Deploy the verifier contract to your blockchain"
    echo -e "  2. Use the proving key to generate proofs"
    echo -e "  3. Submit proofs to the verifier contract"
    echo ""
    echo -e "${YELLOW}Note:${NC} To copy the verifier contract elsewhere, use:"
    echo -e "  cp $VERIFIER_FILE <destination>"
else
    echo -e "${RED}❌ Missing files:${NC}"
    for file in "${MISSING_FILES[@]}"; do
        echo -e "  - $file"
    done
    exit 1
fi