#include <DHT.h>
#include <Wire.h>
#include <XBee.h>
#include <DS1337.h>
#include <DS1621.h>

#define SENSOR_VOLTAGE_MAX 3.0
#define INPUT_PIN_MAX 1023
#define SAMPLE_FREQUENCY 10

// PINS
#define DHTPIN 2
#define DHTTYPE DHT22 
#define SOIL_MOISTURE_PIN 0
#define SOIL_TEMPERATURE_PIN 1
byte ds1621_addr = (0x90 >> 1) | 0;  // replace the 0 with the value you set on pins A2, A1 and A0

DHT dht(DHTPIN, DHTTYPE);
DS1621 enclosure_temp_sensor = DS1621(ds1621_addr);
XBee xbee = XBee();
uint8_t seconds_counter;

void setup() {
  // Two-wire Init
  Wire.begin();

  // Serial port
  Serial.begin(9600);

  // XBee Init
  xbee.setSerial(Serial);

  // Timer config
  ds1337_set_alarm_1_every_second();
  ds1337_clear_status();
  ds1337_set_control(DS1337_A1_INT);
  seconds_counter = 0;

  // external temp sensor
  dht.begin();

  // Temperature sensor config
  enclosure_temp_sensor.startConversion(false);                       // stop if presently set to continuous
  enclosure_temp_sensor.setConfig(DS1621::POL | DS1621::ONE_SHOT);    // Tout = active high; 1-shot mode
}

void loop() {
  uint8_t status;
  ds1337_get_status(&status);
  if (status & DS1337_A1_FLAG) {
    ds1337_clear_status();
    if (++seconds_counter == SAMPLE_FREQUENCY) {
      sendSensorReadings();
      seconds_counter = 0;
    }
  }
}

/**
   Take sensor readings and send to relay
*/
void sendSensorReadings() {
  unsigned char buffer[50];
  float soil_moisture_1_f = readSoilMoistureValue(SOIL_MOISTURE_PIN);
  float soil_temp_1_f = readSoilTemperatureValue(SOIL_TEMPERATURE_PIN);
  float enclosure_temp_f = readEnclosureTemp();
  float external_humidity = dht.readHumidity();
  float external_temp = dht.readTemperature();

  char soil_moisture_1_str[8];
  char soil_temp_1_str[8];
  char enclosure_temp_str[8];
  char ext_temp_str[8];
  char ext_humidity_str[8];
  dtostrf(soil_moisture_1_f, 7, 3, soil_moisture_1_str);
  dtostrf(soil_temp_1_f, 7, 3, soil_temp_1_str);
  dtostrf(enclosure_temp_f, 7, 3, enclosure_temp_str);
  dtostrf(external_temp, 7, 3, ext_temp_str);
  dtostrf(external_humidity, 7, 3, ext_humidity_str);

  sprintf((char*)buffer, "%s,%s,%s,%s,%s", enclosure_temp_str, soil_moisture_1_str, soil_temp_1_str, ext_temp_str, ext_humidity_str);

  // send data to the XBee coordinator
  XBeeAddress64 addr64 = XBeeAddress64(0x00000000, 0x0000FFFF);
  ZBTxRequest zbTx = ZBTxRequest(addr64, buffer, strlen((const char*)buffer));
  xbee.send(zbTx);
}

/**
   Read temperature of enclosure in whole degrees C
*/
float readEnclosureTemp() {
  int tC = enclosure_temp_sensor.getHrTemp();
  return (float) tC / 100;
}

/**
   Read soil temperature from the given pin
*/
float readSoilTemperatureValue(int pin) {
  float pinVoltage = analogRead(pin) * (SENSOR_VOLTAGE_MAX / INPUT_PIN_MAX);
  return (pinVoltage * 41.67) - 40;
}

/**
   Read soil moisture from the given pin
*/
float readSoilMoistureValue(int pin) {
  float pinVoltage = analogRead(pin) * (SENSOR_VOLTAGE_MAX / INPUT_PIN_MAX);
  if ((pinVoltage >= 0.0) && (pinVoltage <= 1.1)) {
    return (10 * pinVoltage) - 1;
  } else if ((pinVoltage >= 1.1) && (pinVoltage <= 1.3)) {
    return (25 * pinVoltage) - 17.5;
  } else if ((pinVoltage >= 1.3) && (pinVoltage <= 1.82)) {
    return (48.08 * pinVoltage) - 47.5;
  } else if ((pinVoltage >= 1.82) && (pinVoltage <= 2.2)) {
    return (26.32 * pinVoltage) - 7.89;
  } else {
    return 50;
  }
}

