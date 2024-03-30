### 1. Pull the docker image for proxy-server
```bash
docker pull fazlulkarim105925/codeserver:v1.0
```
### 2. Pull the docker images for codeservers :node
```bash
docker pull poridhi/codeserver-node:v1.1
```
### 3. Pull the docker images for codeservers :python
```bash
docker pull poridhi/codeserver-python:v1.2
```
### 4.Run the docker images for codeservers: node
```bash
docker run -it -p 5001:8080 poridhi/codeserver-node:v1.1
```
### 5.Run the docker images for codeservers: python
```bash
docker run -p 5002:8080 poridhi/codeserver-python:v1.2
```
### 6.Run the docker image of proxy-server
```bash
docker run -p 8000:8000 fazlulkarim105925/codeserver:v1.0
```
### To Test the code server use
```bash
http://localhost:8000/server1/
http://localhost:8000/server2/
```
