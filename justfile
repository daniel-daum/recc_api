build:
    docker build -t malicious_pickle .

run:
    docker run -d --rm -p 8000:8000 --name malicious_pickle malicious_pickle

clear:
    docker system prune -f

start:
    just build
    just run
    just clear

stop:
    docker stop malicious_pickle
