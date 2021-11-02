module github.com/el10savio/distMergeSort

go 1.17

require (
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/daixiang0/gci v0.2.9 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/el10savio/distMergeSort/handlers => ../handlers

replace github.com/el10savio/distMergeSort/sort => ../sort
