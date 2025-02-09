#!/bin/zsh

echo "Removing existing files..."
rm go.mod
rm go.sum
sudo rm -rf $GOPATH/pkg/mod/*

echo "Initializing new go.mod..."
go mod init github.com/sean-david-welch/farmec-v2/server

echo "Getting latest dependencies..."
dependencies=(
    "firebase.google.com/go"
    "github.com/aws/aws-sdk-go-v2"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/aws/smithy-go"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/joho/godotenv"
    "github.com/mattn/go-sqlite3"
    "github.com/sendgrid/sendgrid-go"
    "github.com/signintech/gopdf"
    "github.com/stripe/stripe-go/v76"
    "google.golang.org/api"
)

for dep in $dependencies; do
    echo "Getting $dep..."
    go get "$dep@latest"
done

echo "Running go mod tidy..."
go mod tidy

echo "Done!"