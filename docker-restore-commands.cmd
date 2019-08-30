
REM ____ RESTORE AND BACKUP OF DOCKER CONTAINERS ON DIFFERENT HOSTS_________________________________________
REM 1). - ON SOURCE HOST- Create an image(tagged snippetbox_db_project) from container(snippetbox_db, stop it first, if it is already running)
REM (src: https://www.scalyr.com/blog/create-docker-image/)
docker commit snippetbox_db snippetbox_db_project

REM 2). - ON SOURCE HOST-Saves the image(tagged snippetbox_db_project) to the '/Users/plamenpentchev/gomodules/snippetbox_db.tgz' tar-file.
REM (src:https://stackoverflow.com/questions/23935141/how-to-copy-docker-images-from-one-host-to-another-without-using-a-repository)
docker save -o /Users/plamenpentchev/gomodules/snippetbox_db.tgz snippetbox_db_project

REM ----- COPY tar-file from source to target hosts-------------------------------

REM 3). - ON TARGET HOST- Loads the image(tagged snippetbox_db_project) to the '/Users/plamenpentchev/gomodules/snippetbox_db.tgz' tar-file.
REM (src:https://stackoverflow.com/questions/23935141/how-to-copy-docker-images-from-one-host-to-another-without-using-a-repository)
docker load -i <path to image tar file>







REM ____ RESTORE AND BACKUP OF DOCKER NAMED VOLUMES_________________________________________
REM (src. https://loomchild.net/2017/03/26/backup-restore-docker-named-volumes/)

REM Will create a container from the slim image busybox and will destroy it(rm) after it has suceeded.
REM will bind(-v) this container to the snippet_databox volume.
REM Volume is snippetbox_data
REM Mount point in the newly created container is /var/lib/mysql
REM will start the tar command in busybox and will create a tar file with the content of the /var/lib/mysql directory
REM Will subsequently send the tar-file to stdout.
$ docker run --rm -v snippetbox_data:/var/lib/mysql busybox sh -c 'tar -cOzf - /var/lib/mysql' > volume-export.tgz

REM --- ANOTHER WAY

REM Will back up a volume(snippetbox_data, mounted at /var/lib/mysql) in a /Users/plamenpentchev/gomodules/volume-export.tgz file.
REM Will map the /tmp directory of the for-short-time-created container(based on the alpine image) to '/Users/plamenpentchev/gomodules' on the host
REM and will bundle the content of the /var/lib/mysql directoryand all its subdirectories into volume-export.tgz tar-file
docker run -it --rm -v snippetbox_data:/var/lib/mysql -v /tmp:/Users/plamenpentchev/gomodules alpine tar -cjf  /Users/plamenpentchev/gomodules/volume-export.tgz -C /var/lib/mysql ./

REM Will restore a volume from a /Users/plamenpentchev/gomodules/volume-export.tgz file.
REM into a snippetbox_data volume, For this purpose will create temporary container based on the alpine image
REM and will map this volume to its /var/lib/mysql directory
docker run -it --rm -v snippetbox_data:/var/lib/mysql -v /tmp:/Users/plamenpentchev/gomodules alpine sh -c "rm -rf /var/lib/mysql/* /varlib/mysql/..?* /varlib/mysql/.[!.]* ; tar -C /varlib/mysql/ -xjf /Benutzer/plamenpentchev/gomodules/volume-export.tgz"
