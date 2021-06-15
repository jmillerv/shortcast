#!/bin/bash

# Add general warning about how this isn't secure to plug into a network
# TODO look up how to prevent internal network browsing in the hostapd.conf file

apt update
apt install hostapd

# we may need this
# wget http://d.dropbox.com/u/1663660/hostapd/hostapd.zip
# unzip hostapd.zip
# cp hosapd /usr/sbin/hostapd


# modify hosts file for wifi usage
echo "interface=wlan0" | tee -a /etc/hostapd/hostapd.conf > /dev/null
echo "driver=rt187xdrv" | tee -a /etc/hostapd/hostapd.conf > /dev/null # we may not need the driver line
# TODO the wifi should be a variable that seeks user input
echo "ssid=Ceasar Chavez Park Wifi" | tee -a /etc/hostapd/hostapd.conf > /dev/null
# TODO look up modes to understand this
echo "hw_mode=g" | tee -a /etc/hostapd/hostapd.conf > /dev/null
# TODO look up channels for optimal choice
echo "channel=11" | tee -a /etc/hostapd/hostapd.conf > /dev/null
echo "macaddr_acl=0" | tee -a /etc/hostapd/hostapd.conf > /dev/null

# always run at start up
sed -i '$d' /etc/rc.local # TODO this is a hack and there's probably a way to insert at the penultimate line
echo "hostapd -B /etc/hostapd/hostapd.conf" | tee -a /etc/rc.local > /dev/null
echo "named -u bind" | tee -a /etc/rc.local > /dev/null
echo "nodogsplash" | tee -a /etc/rc.local > /dev/null
echo "exit 0" | tee -a /etc/rc.local > /dev/null

# install DHCP server
apt install isc-dhcp-server

# Edit DHCP server configuration
AuthA="#authoritative"
AuthB="authoritative"
sed -i "s/$AuthA/$AuthB/g" /etc/dhcp/dhcpd.conf
echo "subnet 10.0.10.0 netmask 255.255.255.0 {
            range 10.0.10.2 10.0.10.254
            option domain-name neighborhoodhistory.com
            option domain-name-servers 10.0.10.1
            option routers 10.0.10.1
            interface wlan0
      }" | tee -a /etc/dhcp/dhcpd.conf > /dev/null

# TODO search for a specifc spot to place this
# /etc/network/interfaces
# change ifcace wlan0 inet manual to iface wlan0 inet static
# add following below
# address 10.0.10.1
# netmask 255.255.255.0

# Enable ip forwarding
# remove comment from /etc/sysctl.conf
NetA="#net.ipv4.ip_forward=1"
NetB="net.ipv4.ip_forward=1"
sed -i "s/$NetA/$NetB/g" /etc/systl.conf

# Network address translation
iptable -t nat -A POSTROUTING -o etho0 -j MASQUERADE

# persist our changes
apt install ipstables-persistent

# Configure as captive portal
wget https://github.com/nodogsplash/nodogsplash/archive/master.zip
unzip master.zip
rm master.zip
cd nodogsplash-master || exit
make
make install

# Configure nodogsplash
echo "GatewayInterface wlan0" | tee -a /etc/nodogsplash/nodogsplash.conf > /dev/null
echo "GatewayAddress 10.0.10.1" | tee -a /etc/nodogsplash/nodogsplash.conf > /dev/null
echo "MaxClients 150" | tee -a /etc/nodogsplash/nodogsplash.conf > /dev/null
echo "ClientIdleTimeout 480" | tee -a /etc/nodogsplash/nodogsplash.conf > /dev/null

# TODO  Edit webpage /etc/nodogsplash/htdocs/splash.html
# https://nodogsplashdocs.readthedocs.io/en/stable/splash.html
# If it's possible execute the Hugo binary that we created in place of a splash page

# add cronjob to monitor who connects
echo "*/5 * * * * ndsctl status >> /root/results5.txt" >> mycron
echo "*/15 * * * * ndsctl status >> /root/results15.txt" >> mycron
crontab mycron
rm mycron

# configure as captive portal
apt install bind9
echo "zone \".\" {
            type master;
            file \"/etc/bind/db.catchall\";
      };" | tee -a /etc/bind/named.conf.local > /dev/null

# edit /etc/bind/named.conf.options
#hopefully variables work in sed
STR1="dir=/etc/namedb"
STR2="/etc/bind"
sed -i "s/$STR1/$STR2/g" /etc/bind/named.conf.options
 touch /etc/bind/db.catchall
echo "$TTL    604800
      @       IN    SOA     .  root.localhost. (
                              1         ; Serial
                              604800    ; Refresh
                              86400     ; Retry
                              2419200   ; Expire
                              604800)    ; Negative Cache TTL
              IN     NS     .
      .       IN     A      10.0.10.1
      *.      IN     A      10.0.10.1" | tee -a /etc/bind/db.catchall > /dev/null


# Run the captive portal software
nodogsplash