cd $1;
echo "RUNNING GO MOD INIT";
git init;
go mod init;
go mod download;
go mod tidy;
git add .;
git commit -m "first commit";
