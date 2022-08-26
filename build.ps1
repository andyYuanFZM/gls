echo "building gls.exe gls-cli.exe"
go env -w CGO_ENABLED=1
go build  -v -o gls.exe github.com/andyYuanFZM/gls
go build  -v -o gls-cli.exe github.com/andyYuanFZM/gls/cli

echo "build end"
