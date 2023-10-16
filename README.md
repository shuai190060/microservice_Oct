## Brief:

This repo is to refactor code the “****Working with Microservices in Go (Golang)****” in Udemy, build, push to docker hub, and deploy them on kubernetes cluster provisioned by Kops in aws.

## Prerequisite

## Topology

## Port by service

| Service | Port | link |
| --- | --- | --- |
| auth-app | 80 | http://auth-app.app.svc.cluster.local/authenticate |
| broker-app | 80 | http://auth-app.app.svc.cluster.local/authenticate |
| postgresql  | 5432 | http://auth-app.app.svc.cluster.local/authenticate |
| mongodb | 20017 | http://auth-app.app.svc.cluster.local/authenticate |

## Ansible-playbook init

```jsx
// deploy tools, (nginx, )
ansible-playbook playbook.yaml -e "tools=true"

// provision the database
ansible-playbook playbook.yaml -e "postgres=true"

// provision the mongo database
ansible-playbook playbook.yaml -e "mongo=true"

// deploy broker
ansible-playbook playbook.yaml -e "app=true"

// deploy broker
ansible-playbook playbook.yaml -e "logger=true"


// update the broker link
ansible-playbook playbook.yaml -e "broker_link_update=true"







// clean up
ansible-playbook playbook.yaml -e "clean=true"
```
