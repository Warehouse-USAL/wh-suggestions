FROM golang:1.25 AS build

WORKDIR /src

# Download dependencies first (cached layer — only re-runs if go.mod/go.sum change)
COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/wh-suggestions ./cmd/wh-suggestions

# --- Runtime image ---
FROM gcr.io/distroless/static-debian12:nonroot AS runtime

COPY --from=build /out/wh-suggestions /wh-suggestions
USER nonroot

EXPOSE 8080

ENTRYPOINT ["/wh-suggestions"]
