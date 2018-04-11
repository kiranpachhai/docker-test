What I have deployed:
1) Two directories ansible and sandbox
2) Installed ansible locally and added localhost to the host list. The file at ./ansible/prerequisites/ansible.hosts is the same as my local /etc/ansible/hosts
3) Wrote an ansible-playbook called ./ansible/deploy_server.yaml that puts the nginx configuration, self-signed certificate and docker-compose.yml to localhost to the appropriate place
4) Whenever the nginx config file is changed or docker-compose.yml file is updated, it'll trigger a restart of docker-compose so the services that host the actual web server running gets restarted with appropriate changes being made

Contents of directory "sandbox"
1) I wrote a simple go program that starts an http server at port 6000 that just prints "CHALLENGE" and then packaged it to be a docker image
2) The contents in ./sandbox/nginx/ is automatically managed by ansible playbook above. The docker container that runs nginx just uses these configs for its own deployment
3) ./sandbox/docker-compose.yml is the main file. It contains two docker containers(our website docker image and nginx image) that are linked together
4) The actual nginx config located at ./sandbox/nginx/webserver.conf is the one that handles all the routing. Anytime the user tries to access the website at a particular URL(in our case, it's http://localhost:8000), it just redirects to https). If, instead the user directly goes to https://localhost:9000, it works as expected


Pipeline:
- Run ansible playbook using "ansible-playbook deploy_server.yml"
- This will start docker containers and start running our website via a docker container

Assumptions:
- localhost already has docker installed and ansible installed
