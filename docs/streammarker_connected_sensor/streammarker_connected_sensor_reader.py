#!/usr/bin/python

import os
import httplib
import serial
import re
import json

ser = serial.Serial('/dev/ttyAMA0', 9600)

headers = {"Content-type": "application/json",
           "Accept": "application/json"}
if 'STREAMMARKER_API_KEY' in os.environ:
  headers["X-API-KEY"] = os.environ['STREAMMARKER_API_KEY']

r_unwanted = re.compile("[\n\r]")
while True:
    try:
        msg = ser.readline()
        msg = r_unwanted.sub("", msg)
        readings = msg.split(',')
        body = json.dumps({"timestamp": 0,"relay_id": "3AC91DEC-B0DD-4DB5-B56D-7682B7C9B28C","status": "ok","sensors": [{"id":"31E541A7-815A-4527-88FC-CEA39808FCAC","readings": [{"timestamp": 0,"measurements": [{"name":"soil_moisture_1","value": float(readings[0]),"unit":"VWC"}, {"name":"soil_temperature_1","value": float(readings[1]),"unit":"Celsius"}]}]}]})
        print body

    except KeyboardInterrupt:
        break
    except: # catch-all
        e = sys.exc_info()[0]
        print(e)

ser.close()