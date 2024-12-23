pragma circom 2.0.0;

include "../../node_modules/circomlib/circuits/smt/smtprocessor.circom";
include "../../node_modules/circomlib/circuits/eddsaposeidon.circom";
include "../../node_modules/circomlib/circuits/gates.circom";
include "../../node_modules/circomlib/circuits/mux1.circom";

include "./balance-updater.circom";
include "./batch-tx-states.circom";
include "./lib/hash-state.circom";
include "./lib/decode-float.circom";

template BatchTx(nLevels) {

}