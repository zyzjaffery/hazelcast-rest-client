FROM hazelcast/hazelcast:3.6.1

# Add your custom hazelcast.xml
ADD hazelcast.xml $HZ_HOME

# Run hazelcast
CMD ./server.sh
