source ~/.bash_profile
$1
scarb cairo-run --available-gas=200000000 "[$1]" | tee "../callCairo/files/output_files/reduce_res_$1.txt"
