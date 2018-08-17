FROM golang:1.10

RUN mkdir -p /go/src/github.com/apexlearning/fake-secretsmanager
RUN mkdir -p /opt/fake-secretsmanager/data

# You really ought to link an actual "secrets" file into this docker container
# rather than just use the default empty secrets file

COPY ./minimal.json /opt/fake-secretsmanager/data/secrets.json
COPY . /go/src/github.com/apexlearning/fake-secretsmanager

WORKDIR /go/src/github.com/apexlearning/fake-secretsmanager
RUN go get -v -d && \
  go install github.com/apexlearning/fake-secretsmanager

EXPOSE 7887

# This should be more better (and more flexible), but it will do for the moment.
# Starting off we'll assume that any secrets file is being linked in from
# outside to secrets.json and nowhere else.

ENTRYPOINT [ "/go/bin/fake-secretsmanager" ]
CMD [ "-f", "/opt/fake-secretsmanager/data/secrets.json" ]
