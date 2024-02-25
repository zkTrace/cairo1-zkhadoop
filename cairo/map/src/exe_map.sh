#!/bin/bash

# Append asdf.sh to your .bash_profile
# echo -e "\n. \"$(brew --prefix asdf)/libexec/asdf.sh\"" >> ~/.bash_profile

# Append asdf.bash completion to your .bash_profile
# echo -e "\n. \"$(brew --prefix asdf)/etc/bash_completion.d/asdf.bash\"" >> ~/.bash_profile

# Source your .bash_profile to apply changes
source ~/.bash_profile

# Run your scarb cairo-run command
scarb cairo-run --available-gas=200000000 

# Todo: Create a trace of our computation
# when i import the json need to create a new cair file thne
# cargo run ../cairo/map/src/lib.cairo --trace_file ../cairo/map/src/mapper.trace.bin --memory_file ../cairo/map/src/mapper.memory.bin --layout starknet_with_keccak --proof_mode

