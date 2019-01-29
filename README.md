# Remove Services
    $docker-compose build taskmgr usermgr email api
    
# Start All Micro Services
    $docker-compose up -d datastore consul broker
    $sleep 5s
    $docker-compose up -d taskmgr usermgr email api
    
# Stop Micro Services
    $docker-compose stop taskmgr usermgr email api
    
# Remove Micro Services
    $docker-compose rm taskmgr usermgr email api

# Run test
    $docker-compose up -d taskmgr-client
    $docker-compose stop taskmgr-client
    $docker-compose rm taskmgr-client
    
