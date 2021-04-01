FROM maven:3.5-jdk-8 AS build
COPY ms /usr/src/app/ms
RUN mvn -f /usr/src/app/ms/pom.xml clean package

#using Google's distroless base image, which strives to provide just enough run-time for a java app.
FROM gcr.io/distroless/java 
COPY --from=build /usr/src/app/app.jar /usr/app/app.jar  
EXPOSE 9002
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/usr/app/app.jar"]