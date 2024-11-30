FROM siva443/java:17-corr-alpine AS builder

WORKDIR /app

COPY src /app/order-service/src
COPY pom.xml /app/order-service/pom.xml

RUN apk add --no-cache maven

COPY settings.xml /root/.m2/settings.xml

WORKDIR /app/order-service

RUN mvn clean package

FROM siva443/java:17-corr-alpine 

COPY --from=builder /app/order-service/target/order-service-0.0.1-SNAPSHOT.jar /app/order-service-0.0.1-SNAPSHOT.jar
WORKDIR /app
#COPY ./target/order-service-0.0.1-SNAPSHOT.jar .
ENTRYPOINT ["java","-jar","./order-service-0.0.1-SNAPSHOT.jar"]
CMD ["--server.port=8081"]
#  docker build -t order-service-ms .