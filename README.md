# Notes

- Docker Hub:
- Run Hello World

```
docker run hello-world
```

- alpine

- run with shell

```
docker run -it alpine sh
```

```
docker push abarakatm/rafee-hello-world:tagname

docker run -it --platform=linux/amd64 abarakatm/rafee-hello-world:0.0.1 sh

docker run --platform=linux/amd64 abarakatm/rafee-hello-world:0.0.1 /newapp/hello
```

- create employee
```
curl -X POST http://localhost:3000/api/employees \
-H "Content-Type: application/json" \
-d '{
"firstName": "John",
"lastName": "Doe",
"email": "john.doe@example.com",
"position": "Software Engineer",
"salary": 75000
}'

```
- get all employees

```
curl http://localhost:3000/api/employees
```

- get an employee
```
curl http://localhost:3000/api/employees/{employee-id}
```

- delete an employee
```
curl -X DELETE http://localhost:3000/api/employees/{employee-id}

```
--
AWS:
ECS: Elastic Container Services
EKS: Elastic Kubernetes Services
