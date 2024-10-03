FROM google/cloud-sdk:495.0.0-debian_component_based
RUN apt-get install -y mtr
WORKDIR /app

COPY mtr-script.sh .
RUN chmod +x ./mtr-script.sh

CMD [ "./mtr-script.sh" ]
