Generate swagger docs
swag init -g main.go

Start gateway
./serve gateway

Start service user
./serve user


git init
git pull git@github.com:fendriknuruljadid/astrovia-landing-page.git main
git remote add origin git@github.com:fendriknuruljadid/astrovia-landing-page.git
git fetch origin

git remote -v
git pull origin main
git ls-remote --heads origin

Ganti branch lokal ke main
git fetch origin
git checkout -b main origin/main

Set upstream supaya pull otomatis pakai main:
git branch --set-upstream-to=origin/main main

git pull




