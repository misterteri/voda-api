# Voda Scheduler API

Run the command for the API to listen and serve
```
go run main.go
```

Send a YAML file through `curl`
```
curl --form file="@example.yaml" http://localhost:8080/post
```