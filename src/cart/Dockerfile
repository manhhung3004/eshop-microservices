# Runtime Stage
FROM public.ecr.aws/amazonlinux/amazonlinux:2023

# Cài đặt Java 17
RUN dnf --setopt=install_weak_deps=False install -q -y \
    java-17-amazon-corretto-headless \
    && dnf clean all

# Thiết lập người dùng và thư mục làm việc
ENV APPUSER=appuser
ENV APPUID=1000
ENV APPGID=1000

# RUN useradd --home "/app" --create-home --user-group --uid "$APPUID" "$APPUSER"

# Thiết lập OpenTelemetry Agent
ARG aws_opentelemetry_agent_version=1.24.0
RUN curl -L https://github.com/aws-observability/aws-otel-java-instrumentation/releases/download/v${aws_opentelemetry_agent_version}/aws-opentelemetry-agent.jar -o /opt/aws-opentelemetry-agent.jar

ENV JAVA_TOOL_OPTIONS=-javaagent:/opt/aws-opentelemetry-agent.jar
ENV OTEL_JAVAAGENT_ENABLED=false

WORKDIR /app
# USER appuser

# Copy file JAR từ build artifact
# COPY --chown=appuser:appuser ./target/ui-0.0.1-SNAPSHOT.jar /app/app.jar

COPY ./target/carts-0.0.1-SNAPSHOT.jar /app/app.jar

EXPOSE 8080

# Khởi chạy ứng dụng
ENTRYPOINT ["sh", "-c", "java $JAVA_OPTS -jar /app/app.jar"]


