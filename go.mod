module homesync

go 1.12

require homesync/client v0.0.0

replace homesync/client => ./client

require homesync/server v0.0.0

replace homesync/server => ./server

require homesync/foldermonitor v0.0.0

replace homesync/foldermonitor => ./foldermonitor