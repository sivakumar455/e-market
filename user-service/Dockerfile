FROM siva443/java:17-corr-alpine AS builder

WORKDIR /app

COPY src /app/user-service/src
COPY pom.xml /app/user-service/pom.xml

RUN apk add --no-cache maven

COPY settings.xml /root/.m2/settings.xml

WORKDIR /app/user-service

RUN mvn clean package

FROM siva443/java:17-corr-alpine 

COPY --from=builder /app/user-service/target/user-service-0.0.1-SNAPSHOT.jar /app/user-service-0.0.1-SNAPSHOT.jar
WORKDIR /app
ENTRYPOINT ["java","-jar","./user-service-0.0.1-SNAPSHOT.jar"]
CMD ["--server.port=8081"]

# docker build -t user-service-ms .