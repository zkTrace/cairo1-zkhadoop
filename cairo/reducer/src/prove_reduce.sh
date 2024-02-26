#!/bin/bash


# generating trace (in docker)
(cd /app/cairo-vm && ./cairo1-run /app/cairo/reducer/src/agg.cairo \
    --trace_file /app/cairo/reducer/src/reducer.trace.bin --memory_file \
    /app/cairo/reducer/src/reducer.memory.bin --layout starknet_with_keccak --proof_mode \
    --args '0'
    > /dev/null) || (echo "Failed to generate trace" && exit 1)


# generating proof
platinum-prover prove reducer.trace.bin reducer.memory.bin reducer.proof

cp reducer.proof /app/server/data/mr-tmp
