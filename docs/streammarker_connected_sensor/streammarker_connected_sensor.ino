#include <Wire.h>
#include <Event.h>
#include <DS1337.h>


#define SAMPLE_DELAY 5000
#define SENSOR_VOLTAGE_MAX 3.0
#define INPUT_PIN_MAX 614

// PINS
#define SOIL_MOISTURE_PIN 0
#define SOIL_TEMPERATURE_PIN 1

// Vegetronix VH400 piecewise linear curve
// Returns the temperature in degrees Celsius.
float readSoilTemperatureValue(int pin) {
  float pinVoltage = analogRead(pin) * ((float)SENSOR_VOLTAGE_MAX / (float)INPUT_PIN_MAX);
  return (pinVoltage * 41.67) - 40;
}

// Vegetronix THERM200
// Voltage ranges from 0-3 Volts, which the analog sensors reports as being between 0-614.
// The Arduino analog input ranges from 0-5 Volts, reported as ranging from 0-1023.
// So, a voltage ranging from 0-3 Volts will have a max reported as (0.6 * 1023), or 614.
float readSoilMoistureValue(int pin) {
  float pinVoltage = analogRead(pin) * ((float)SENSOR_VOLTAGE_MAX / (float)INPUT_PIN_MAX);
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

void setup() {

  // Initialize the I2C bus
  Wire.begin();
  
  // Initialize the serial interface
  Serial.begin(9600);
  while(!Serial);
  
  // Once serial has been inialized, print a welcome message
  Serial.println("Maxim DS1337 RTC Example"); 

  ds1337_set_alarm_1_every_second();
  ds1337_clear_status();
  ds1337_set_control(DS1337_A1_INT);
}

void loop() {
  uint8_t status;
  ds1337_get_status(&status);
  if (status & DS1337_A1_FLAG) {
    sendSensorReadings();
    ds1337_clear_status();
  }
}

void sendSensorReadings() {
  char buffer[50];
  float a = readSoilMoistureValue(SOIL_MOISTURE_PIN);
  float b = readSoilTemperatureValue(SOIL_TEMPERATURE_PIN);
  char a_str[15];
  char b_str[15];
  dtostrf(a, 10, 5, a_str);
  dtostrf(b, 10, 5, b_str);
  sprintf(buffer, "%s,%s", a_str, b_str);

  Serial.println(buffer);
}
