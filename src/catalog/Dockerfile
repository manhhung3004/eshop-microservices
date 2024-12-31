# Production Stage
FROM alpine:3.19.1

ENV APPUSER=appuser
ENV APPUID=1000
ENV APPGID=1000

RUN apk add --no-cache shadow

RUN addgroup -g $APPGID $APPUSER

RUN adduser -D -u $APPUID -G $APPUSER -s /bin/sh $APPUSER

WORKDIR /app
USER $APPUSER

COPY --chown=appuser:appuser /main /app/
COPY --chown=appuser:appuser db/ /app/db/

ENV GIN_MODE=release    

ENTRYPOINT ["/app/main"]