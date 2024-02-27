#!/bin/bash


# generating trace (in docker)
(cd /app/cairo-vm && ./cairo1-run /app/cairo/map/src/agg.cairo \
    --trace_file /app/cairo/map/src/mapper.trace.bin --memory_file \
    /app/cairo/map/src/mapper.memory.bin --layout starknet_with_keccak --proof_mode \
    > /dev/null) || (echo "Failed to generate trace" && exit 1)


# generating proof
platinum-prover prove mapper.trace.bin mapper.memory.bin mapper.proof

cp mapper.proof /app/server/data/mr-tmp
