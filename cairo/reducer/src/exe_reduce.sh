source ~/.bash_profile
$1

scarb cairo-run "[$1]" 
# scarb cairo-run --available-gas=200000000 "[$1]" 

# Todo: Create a trace of our computation
# when i import the json need to create a new cair file thne
# cargo run ../cairo/reducer/src/lib.cairo --trace_file ../cairo/map/src/reducer.trace.bin --memory_file ../cairo/map/src/reducer.memory.bin --layout starknet_with_keccak --proof_mode --args '1'                  