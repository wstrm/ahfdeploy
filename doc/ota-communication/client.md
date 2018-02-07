Over-the-air communication as of 76e15e2
========================================

 * Go into `V3.0_docker/dev/examples/hello-ahf-java-pure`
 * Change the IP (`130.240.152.197`) to a core services instance, and run:
 ```
 docker run --rm \
           --name hello-ahf-java-pure \
           --hostname hello.docker.ahf \
           --net-alias hello.docker.ahf \
           -p 8888:8888 \
           --volume core_tls:/tls \
           --network core_ahf \
           --add-host="simpleservicediscovery.docker.ahf:130.240.152.197" \
           --add-host="glassfish.docker.ahf:130.240.152.197" \
           hello-ahf-java-pure
 ```
 * 
