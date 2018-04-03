# Zocker

---

Build Docker with go, just study and reasearch.

```bash
wget http://dl-cdn.alpinelinux.org/alpine/v3.6/releases/x86_64/alpine-minirootfs-3.6.2-x86_64.tar.gz
mkdir /var/lib/alpine
tar -xzvf alpine-minirootfs-3.6.2-x86_64.tar.gz -C /var/lib/alpine
go build
sudo ./zocker run /bin/bash
```