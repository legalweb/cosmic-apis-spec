# Cosmic Public API specification

## Grabbing go-swagger (v0.22)

```sh
wget -O swagger-22 https://github.com/go-swagger/go-swagger/releases/download/v0.22.0/swagger_linux_amd64
chmod +x ./swagger-22
mv ./swagger-22 "${HOME}/bin"
```

## Rebuiling Go vendor

```sh
bazel run //:gazelle -- update-repos -from_file=./go.mod -to_macro "go_vendor.bzl%go_repositories"
```
