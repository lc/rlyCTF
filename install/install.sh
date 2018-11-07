sudo apt-get update -y && sudo apt-get upgrade -y
sudo apt-get install golang
cd ../
go build rlyserver.go
sudo cp rlyserver /usr/bin
cp install/rlyserver.service /lib/systemd/system/rlyserver.service
sudo apt-get install postfix -y
sudo echo 'header_checks = regexp:/etc/postfix/ctf_flag' >> /etc/postfix/main.cf
sudo echo '/^Subject:(.*)/ REPLACE Subject: flag{oMg_suCh_h4x_sSrf}' > /etc/postfix/ctf_flag
service postfix reload
ufw allow 22
ufw allow 80
ufw allow from 127.0.0.1 to 127.0.0.1 port 25 proto tcp
