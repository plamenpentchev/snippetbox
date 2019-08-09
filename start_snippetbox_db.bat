# create volume to persist data bases
# docker volume create snippetbox_data
# create container (based on the mysql image) and map the docker volume
docker run --name snippetbox_db -v snippetbox_data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=web -e MYSQL_PASSWORD=sn1pp3tb0x -e MYSQL_DATABASE=snippetbox -p 3306:3306 -d mysql