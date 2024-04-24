#!/bin/bash

docker buildx build --push \
--platform linux/amd64,linux/arm64 -f Dockerfile \
--tag  tribehealth/supabase-backend:v0.0.1 .
