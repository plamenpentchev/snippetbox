



REM Will back up a volume(snippetbox_data, mounted at /var/lib/mysql) in a /Users/plamenpentchev/gomodules/volume-export.tgz file.
REM Will map the /tmp directory of the for-short-time-created container(based on the alpine image) to '/Users/plamenpentchev/gomodules' on the host
REM and will bundle the content of the /var/lib/mysql directoryand all its subdirectories into volume-export.tgz tar-file
docker run -it --rm -v snippetbox_data:/var/lib/mysql -v /tmp:/Users/plamenpentchev/gomodules alpine tar -cjf  /Users/plamenpentchev/gomodules/volume-export.tgz -C /var/lib/mysql ./

REM Will restore a volume from a /Users/plamenpentchev/gomodules/volume-export.tgz file.
REM into a snippetbox_data volume, For this purpose will create temporary container based on the alpine image
REM and will map this volume to its /var/lib/mysql directory
docker run -it --rm -v snippetbox_data:/var/lib/mysql -v /tmp:/Users/plamenpentchev/gomodules alpine sh -c "rm -rf /var/lib/mysql/* /varlib/mysql/..?* /varlib/mysql/.[!.]* ; tar -C /varlib/mysql/ -xjf /Benutzer/plamenpentchev/gomodules/volume-export.tgz"
