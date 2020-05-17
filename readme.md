## Command line interface to upload and download single file or multiple files in parallel from s3

### Uses 
- First set environmental variable to your bucket
`export my_bucket=<your backet name>`
- Make sure aws credential exists in your home directory
- To download one file use (if **filepath** is empty then file will be stored in current directory)
`go run main.go download test.txt -filepath ~/Documents/test_folder/test.txt`
 - To upload on file use (if key is empty it will use filename as key)
 `go run main.go upload ~/Documents/test_folder/test.txt -key test.txt`
- For multiple upload 
`go run main.go upload ~/Documents/test_folder/test4.txt ~/Documents/test_folder/test3.txt`
- For multiple download 
`go run main.go download tes2.txt test.txt`