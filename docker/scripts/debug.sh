#!/bin/sh
(cd /app && dlv --listen=:40000 --headless=true --api-version=2 --accept-multiclient exec ./gorecon "$TARGET")