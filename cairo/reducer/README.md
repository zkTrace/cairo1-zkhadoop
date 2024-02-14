# Cairo Matrix Operations
This project provides implementations of matrix multiplication and matrix-vector multiplication in the Cairo language, designed to work within zero-knowledge proof and distributed system.

# Overview
The repository includes two main components:

Matrix Multiplication: A Cairo program that performs multiplication of two matrices using a map reduce approach.
Matrix-Vector Multiplication: A Cairo program that multiplies a matrix by a vector, both of which are crucial in various computational and cryptographic contexts, especially within zk-STARKs.
These implementations are designed to demonstrate how complex mathematical operations can be executed within the Cairo language, maintaining the integrity and privacy guarantees provided by zk proofs.

# Prerequisites
Before you begin, ensure you have the environment installed:

## Install asdf
### For Linux
```rust
sudo apt install curl git
git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.13.1

//add this to the ~/.bashsrc file
. "$HOME/.asdf/asdf.sh"
. "$HOME/.asdf/completions/asdf.bash"

code ~/.bashrc
source ~/.bashrc
```

### For Mac

```rust
brew install asdf
```
If using macOS Catalina or newer, the default shell has changed to ZSH. Unless changing back to Bash, follow the ZSH instructions.

Add asdf.sh to your ~/.bash_profile with:

```rust
echo -e "\n. \"$(brew --prefix asdf)/libexec/asdf.sh\"" >> ~/.bash_profile
```

Completions will need to be configured as per Homebrew's instructions or with the following:

```rust
echo -e "\n. \"$(brew --prefix asdf)/etc/bash_completion.d/asdf.bash\"" >> ~/.bash_profile
```

### Install Scarb
```
asdf plugin add scarb
asdf install scarb latest
asdf global scarb latest

scarb --version
```

### To run the code
``` rust
scarb cairo-run --available-gas=200000000  
```

### To run the tests
``` rust
scarb cairo-test -f tests 
```