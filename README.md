# Introduction to Docker for Developers

Talk given at Full Stack Hants on 12 March 2018.

## Basics of Docker

Introducing namespaces, control groups, the union file system and many of the basic commands.

``` bash
uname -r
cat /etc/*release
whoami
ps aux
ls /home
hostname
docker version
docker search ubuntu
docker pull ubuntu
docker images ubuntu
docker run ubuntu echo "Hello World"
docker ps
docker ps -a
docker rm <ID>
docker run -it --rm ubuntu
uname -r
cat /etc/*release
whoami
ls /home
ps aux
hostname
sleep 321
Ctrl-p-q
ps aux | grep 321
docker attach <ID>
Ctrl-c
touch hello.txt
Ctrl-p-q
sudo find /var/lib -name "hello.txt"
sudo ls /var/lib/..../diff/....
sudo ls /var/lib/..../merged/...
docker logs <ID>
docker diff <ID>
docker commit <ID> touched
docker stop <ID>
docker images
docker history touched
docker run -it --rm touched
ls
Ctrl-c
docker run -d --name stress progrium/stress --cpu 2
docker stats
docker kill stress
docker rm stress
docker run -d --name stress progrium/stress --cpu 2
docker stats
docker rm -f stress
```

## Working with Docker images

```bash
cd flask
view Dockerfile
docker build -t flask .
docker run -d -p 5000:5000 --name flask flask
docker logs flask
open http://localhost:5000
docker rm -f flask
vi demo.py
[Edit string]
docker build -t flask .
docker run -d -p 5000:5000 --name flask flask
docker logs flask
[Refresh browser]
docker rm -f flask
docker run -d -v $(pwd):/app -p 5000:5000 --name flask flask
vi demo.py
[Edit string]
[Refresh browser]
docker rm -f flask
docker tag flask dcurrie/flask
docker push dcurrie/flask
docker rmi flask dcurrie/flask
docker run -d -p 5000:5000 --name flask dcurrie/flask
docker rm -f flask
```

## Docker Compose

```bash
cd ../flask-mongo
view docker-compose.yml
docker-compose up
```

## Further topics

* Multi-step builds
* Docker Swarm vs Kubernetes
* Windows Containers
