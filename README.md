This hazelcast-rest-client repo demonstrates how GO interacts with Hazelcast through its REST interface.

This project doesn't produce any executable, instead, provides a SessionManager interface defining CRUD operations
of (key, value) pair, also it provides one implementation communicating with Hazelcast via REST.

#Set up a Hazelcast cluster locally

The tests in this repo require a running Hazelcast node/cluster. We are going to bring up a cluster of 2 nodes using 
Docker.

##Prerequisite

Docker is installed locally, and the docker deamon is running.

##Build a customized Hazelcast Docker Image

Hazelcast officially publishes its Docker image in [Docker Hub] (https://hub.docker.com/r/hazelcast/hazelcast/).

In order to provide our own configuration, this repo provides a Dockerfile and Hazelcast configuration file (hazelcast.xml)
in the root directory. The Dockerfile extends the official Hazelcast OSS 3.6.1 docker image, and wires in our own hazelcast.xml.
Users can edit the hazelcast.xml so that the configuration can be wired in the image to be built.

As we don't have private Docker repo yet, users need to build the image from the Dockerfile themselves by:

```
docker build -t {imagename:tag} .
```

##Bring up a Hazelcast Cluster

Once you have built the image, it's straight forward to bring up the Hazelcast cluster. 

The following two lines bring up two Hazelcast nodes, they will auto-detect each other and form a cluster. 
And the host port 5701 and 5702 are bind to the two hazelcast 5701 ports in container.

```
docker run --rm -it -p 5701:5701 {imagename:tag}
docker run --rm -it -p 5702:5701 {imagename:tag}
```

If everything is fine, user will see similar message as below in the second node console:

>Members [2] {

>	Member [172.17.0.3]:5701

>	Member [172.17.0.2]:5701 this

>}

##How to access the hazelcast node(s)

{docker_deamon_ip}:{host_bind_port}

**docker_deamon_ip** can be found by ```docker-machine ip```

**host_bind_port** will be 5701 and 5702 for node1 and node2 respectively.

*Run tests

Once you have a Hazelcast node or cluster running locally, you can run tests by

```
cd rest
go test
go test -bench .
```

The tests themselves are self-explanatory. 