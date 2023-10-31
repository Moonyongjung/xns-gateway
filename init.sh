rm -rf ~/.xns
make install

# Intialize XNS gateway 
xns i

# Key recover
xns k recover

# Store all XNS contracts
xns e c s

# Instantiate all XNS contracts
xns e c i

# Connect all XNS contracts
xns e c c

# Start XNS gateway
xns start