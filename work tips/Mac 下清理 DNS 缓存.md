```bash
# Mojave (version 10.14)
# High Sierra (version 10.13)
# Sierra (version 10.12) 
# Mountain Lion (version 10.8) 
# Lion (version 10.7)
sudo killall -HUP mDNSResponder

# El Capitan (version 10.11) 
# Mavericks (version 10.9)
sudo dscacheutil -flushcache
sudo killall -HUP mDNSResponder

# Yosemite (version 10.10)
sudo discoveryutil mdnsflushcache

# Snow Leopard (version 10.6) 
# Leopard (version 10.5)
sudo dscacheutil -flushcache

# Tiger (version 10.4)
lookupd -flushcache
```

