@echo off


"C:\Program Files\Git\bin\bash.exe" -c " \"/c/Program Files/Docker Toolbox/start.sh\" \"%*\"" 
echo "Starting snippetbox_Db"
docker start snippetbox_db