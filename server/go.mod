module server

go 1.12

require (
	github.com/bnkamalesh/webgo v2.5.1+incompatible
	github.com/jinzhu/gorm v1.9.10 // indirect
	github.com/revel/modules v0.21.0
	gopkg.in/gorp.v2 v2.0.0 // indirect
	homesync.com/foldermonitor v0.0.0
)

replace homesync.com/foldermonitor => ../foldermonitor/
