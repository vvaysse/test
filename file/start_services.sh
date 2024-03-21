#!/bin/bash

# Start your services
systemctl enable ebb_ricardBorneServ 
systemctl enable ebb_ricardgoBack 
systemctl start ebb_ricardBorneServ 
systemctl start ebb_ricardgoBack 

# Restart Apache
service apache2 restart

# Keep the container running by tailing the Apache logs
tail -f /var/log/apache2/access.log /var/log/apache2/error.log
