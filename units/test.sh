go test $(go list ./... | grep -v "vendor" | grep -v "docs" ) 
