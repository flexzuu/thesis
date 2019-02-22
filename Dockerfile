FROM golang:1.11
ENV GO111MODULE on
WORKDIR /go/src/app
# <- COPY go.mod and go.sum files to the workspace
COPY go.mod . 
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .



CMD ["bash"]