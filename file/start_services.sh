#!/bin/bash

# Start your services
service ebb_ricardBorneServ start
service ebb_ricardgoBack start

# Restart Apache
service apache2 restart

# Keep the container running by tailing the Apache logs
tail -f /var/log/apache2/access.log /var/log/apache2/error.log
