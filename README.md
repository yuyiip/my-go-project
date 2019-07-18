# My Go Project

## Configuration
conf/config.prod.yaml is meant for production

conf/config.yaml is meant for development

## Development
run

```fresh main.go```
    
to enable hot reload

## Build for production
For amd64 (Linux instance), run ```make```

## Build Info
run

```./my-go-project -v```

## Docker Deployment
```docker build -t my-go-project:latest .```

```docker push my-go-project```