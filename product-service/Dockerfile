FROM siva443/java:17-corr-alpine AS builder

WORKDIR /app

COPY src /app/product-service/src
COPY pom.xml /app/product-service/pom.xml

RUN apk add --no-cache maven

COPY settings.xml /root/.m2/settings.xml

WORKDIR /app/product-service

RUN mvn clean package

FROM siva443/java:17-corr-alpine 

COPY --from=builder /app/product-service/target/product-service-0.0.1-SNAPSHOT.jar /app/product-service-0.0.1-SNAPSHOT.jar
WORKDIR /app
ENTRYPOINT ["java","-jar","./product-service-0.0.1-SNAPSHOT.jar"]
CMD ["--server.port=8081"]

# docker build -t  product-service-ms .