# At this stage we will create the redis configuration file from the content of
# the environment variables file. It might not be necessary, but it's very
# helpful
FROM redis:alpine AS builder

ARG REDIS_STORE_PASSWORD

RUN echo "requirepass \"${REDIS_STORE_PASSWORD}\"" >> /redis.conf

# After the building of the configuration file is done, we can copy the file and
# the rest of the configuration is the default one
FROM builder AS deployment

COPY --from=builder /redis.conf /usr/local/etc/redis/redis.conf

CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]
