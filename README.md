## Brief:

This repo is to refactor code the “****Working with Microservices in Go (Golang)****” in Udemy, build, push to docker hub, and deploy them on kubernetes cluster provisioned by Kops in aws.

## Prerequisite

## Topology

## Port by service

| Service | Port |
| --- | --- |
| auth-app | 8081 |
| broker-app | 80 |
| postgresql  | 5432 |

## Ansible-playbook init

```jsx
// deploy tools, (nginx, )
ansible-playbook playbook.yaml -e "tools=true"

// deploy broker
ansible-playbook playbook.yaml -e "app=true"

// update the broker link
ansible-playbook playbook.yaml -e "broker_link_update=true"

// provision the database
ansible-playbook playbook.yaml -e "postgres=true"
```