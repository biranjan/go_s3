## Command line cli to upload and download file from s3

### Uses 
- First set environmental variable to your bucket
`export my_bucket=<your backet name>`
- Make sure aws credential exists in your home directory
- To download use (if file path is empty then it will store in current directory)
`go run main.go download -filename test.txt -filepath ~/Documents/test_folder/test.txt`
 - To upload use (if key is empty it will use filename as key)
 `go run main.go upload -filename ~/Documents/test_folder/test.txt -key test.txt`
