# OpenVPN Rest Server

## Install OpenVPN server

See: https://github.com/angristan/openvpn-install

```bash
curl -O https://raw.githubusercontent.com/angristan/openvpn-install/master/openvpn-install.sh
chmod +x openvpn-install.sh
./openvpn-install.sh
```

## Editing the basic certificate template

```bash
sudo nano /etc/openvpn/client-template.txt
```

Example:

```ini
client
proto udp
explicit-exit-notify
remote {HOST_ADDRESS_OF_YOUR_OPENVPN_SERVER} 1194
dev tun
resolv-retry infinite
nobind
persist-key
persist-tun
remote-cert-tls server
verify-x509-name server_AOUH2vX4ukuu96nG name
auth SHA256
auth-nocache
cipher AES-128-GCM
tls-client
tls-version-min 1.2
tls-cipher TLS-ECDHE-ECDSA-WITH-AES-128-GCM-SHA256
ignore-unknown-option block-outside-dns
setenv opt block-outside-dns # Prevent Windows 10 DNS leak
verb 3
route-nopull # <!--- ALLOW NON-VPN TRAFFIC 
route 10.0.0.0 255.0.0.0 # <!--- IP RANGE YOU WANT
```
