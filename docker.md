set manager
docker swarm init --advertise-addr 212.85.27.3

join worker
docker swarm join --token SWMTKN-1-3as4i1umcszrj4uz3gn2q5m9swzp8yabrwb6b2eim1x9rgbirb-ex9lj7bg51t6o9gz7xctd6l4n 212.85.27.3:2377

show node
docker node ls

update label
docker node update --label-add name=worker-01 qeuhg4lmyu7zp0hfakrqwiec3

show detail node
docker node inspect qeuhg4lmyu7zp0hfakrqwiec3 --pretty

deploy stack / service
docker stack deploy -c traefik.yml traefik

force update
docker service update --force traefik_traefik

cek logs
docker service logs traefik_traefik -f

show all service
 docker service ls

ssl register with cloudflare
sudo certbot certonly   --dns-cloudflare   --dns-cloudflare-credentials /root/.secrets/certbot/cloudflare.ini   -d astrovia.id   -d '*.astrovia.id'

add auto renew ssl
crontab -l

add to cronjob
0 3 * * * certbot renew --dns-cloudflare --dns-cloudflare-credentials ~/.secrets/certbot/cloudflare.ini --quiet
0 3 * * * certbot renew --quiet --post-hook "docker exec jmdn-reverse-proxy nginx -s reload"


build image
cd /var/www/astrovia/astrovia-landing-page
docker build -t astrovia/landing:latest .

cd /var/www/astrovia/astrovia-app
docker build -t astrovia/app:latest .

docker login

push image
docker push astrovia/landing:latest
docker push astrovia/app:latest

harus sama dengan login username docker hub
# Landing
cd astrovia-landing-page
docker build -t fendriknj/astrovia-landing:latest .

# App
cd astrovia-app
docker build -t fendriknj/astrovia-app:latest .

docker push fendriknj/astrovia-landing:latest
docker push fendriknj/astrovia-app:latest

Cek service
docker service ls
docker service ps astrovia_landing
docker service ps astrovia_app

pull di worker
docker pull fendriknj/astrovia-landing-page:latest
docker pull fendriknj/astrovia-app:latest

