# rlyCTF
rlyCTF (relay CTF) is a simple capture the flag challenge to emulate real-world SSRF attacks. 

# installation:
```
git clone https://github.com/C0RB3N/rlyCTF
cd rlyCTF & vim rlyserver.go
```
change the line 
`path := "/root/rlyCTF"`

to 
`path := "/<yourpath>/rlyCTF"`


```
cd install && chmod +x install.sh
./install.sh
service rlyserver start
```

