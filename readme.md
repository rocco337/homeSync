#TODO

##Client
* Test that creates random files in folders
* Scan folder every X seconds
    * Inputs: interval, path to monitor
    * Output: file structure, with filename/hash/modified date    
* Detect new/deleted/modified files
    * Input: local structure, remote structure
    * Output: files to upload/remove
* Send new files and command to delete remote files
    * Detect servers in local network or use preconfigured dns name
    * Use computer name and username as unique username on server

##Server
* Receive instructions to create delete files
* Default folder is username from request
* Send server discovery signal