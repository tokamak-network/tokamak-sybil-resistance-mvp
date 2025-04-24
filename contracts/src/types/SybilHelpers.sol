// SPDX-License-Identifier: AGPL-3.0
pragma solidity 0.8.23;

error InvalidPoseidonAddress(string elementType);

/**
 * @dev Interface poseidon hash function 2 elements
 */
interface PoseidonUnit2 {
    function poseidon(uint256[2] memory) external pure returns (uint256);
}

/**
 * @dev Interface poseidon hash function 3 elements
 */
interface PoseidonUnit3 {
    function poseidon(uint256[3] memory) external pure returns (uint256);
}

/**
 * @dev Sybil helper functions
 */
contract MVPSybilHelpers {
    PoseidonUnit2 _insPoseidonUnit2;
    PoseidonUnit3 _insPoseidonUnit3;

    /**
     * @dev Load poseidon smart contract

     */
    function _initializeHelpers(
        address _poseidon2Elements,
        address _poseidon3Elements
    ) internal {
        if (_poseidon2Elements == address(0)) {
            revert InvalidPoseidonAddress("poseidon2Elements");
        }
        if (_poseidon3Elements == address(0)) {
            revert InvalidPoseidonAddress("poseidon3Elements");
        }

        _insPoseidonUnit2 = PoseidonUnit2(_poseidon2Elements);
        _insPoseidonUnit3 = PoseidonUnit3(_poseidon3Elements);
    }

    // /**
    //  * @dev Builds the state for the Merkle tree.
    //  *
    //  * @param amount The amount to be included in the state.
    //  * @param user The address of the user associated with the state.
    //  *
    //  * @return A uint256 array representing the state for the Merkle tree.
    // */
    function _buildTreeState(
        uint192 amount,
        address user
    ) internal pure returns (uint256[2] memory) {
        uint256[2] memory state;
        state[0] = amount;
        state[1] = uint256(uint160(user)); // Convert address to uint256
        return state;
    }

    /**
     * @dev Hash poseidon for 2 elements
     * @param inputs Poseidon input array of 2 elements
     * @return Poseidon hash
     */
    function _hash2Elements(
        uint256[2] memory inputs
    ) internal view returns (uint256) {
        return _insPoseidonUnit2.poseidon(inputs);
    }

    /**
     * @dev Hash poseidon for 3 elements
     * @param inputs Poseidon input array of 3 elements
     * @return Poseidon hash
     */
    function _hash3Elements(
        uint256[3] memory inputs
    ) internal view returns (uint256) {
        return _insPoseidonUnit3.poseidon(inputs);
    }

    /**
     * @dev Hash poseidon for sparse merkle tree final nodes
     * @param key Input element array
     * @param value Input element array
     * @return Poseidon hash
     */
    function _hashFinalNode(
        uint256 key,
        uint256 value
    ) public view returns (uint256) {
        uint256[3] memory inputs;
        inputs[0] = key;
        inputs[1] = value;
        inputs[2] = 1;
        return _hash3Elements(inputs);
    }

    /**
     * @dev Verify sparse merkle tree proof
     * @param scoreRoot Root to verify
     * @param siblings Siblings necessary to compute the merkle proof
     * @param idx Key to verify
     * @param stateHash Value to verify
     * @return True if verification is correct, false otherwise
     */
    function _smtVerifier(
        uint256 scoreRoot,
        uint256[] calldata siblings,
        uint256 idx,
        uint256 stateHash
    ) internal view returns (bool) {
        // Step 2: Calcuate root
        uint256 nextHash = _hashFinalNode(idx, stateHash);
        uint256 siblingTmp;
        for (int256 i = int256(siblings.length) - 1; i >= 0; i--) {
            siblingTmp = siblings[uint256(i)];
            bool leftRight = (uint8(idx >> uint256(i)) & 0x01) == 1;
            nextHash = leftRight
                ? _hashNode(siblingTmp, nextHash)
                : _hashNode(nextHash, siblingTmp);
        }

        // Step 3: Check root
        return scoreRoot == nextHash;
    }

    /**
     * @dev Hash poseidon for sparse merkle tree nodes
     * @param left Input element array
     * @param right Input element array
     * @return Poseidon hash
     */
    function _hashNode(
        uint256 left,
        uint256 right
    ) public view returns (uint256) {
        uint256[2] memory inputs;
        inputs[0] = left;
        inputs[1] = right;
        return _hash2Elements(inputs);
    }
}
