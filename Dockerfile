FROM alpine

WORKDIR /ProductParser

COPY gopars /usr/bin/ 
COPY files .