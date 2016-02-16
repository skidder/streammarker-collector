Feature: Sensor Readings Collector API
  As a StreamMarker relay
  I want to be able to record sensor readings

  @happy
  Scenario: Record a single sensor reading
    Given I have the Sensor request "valid_single_sensor_reading"
    When I call POST to "/api/v1/sensor_readings"
    Then the result should be a 201
    And a message should be sent by relay id "53644F1C-2480-4F9B-9CBA-26D66139D221" for sensor "D195E2B3-A1F6-407C-83F5-334FA46B5F6C"

  @sad
  Scenario: Attempt to record sensor reading with invalid API Key
    Given I have the Sensor request "valid_single_sensor_reading"
    When I call POST with API key "badkey" to "/api/v1/sensor_readings"
    Then the result should be a 401

  @sad
  Scenario: Attempt to record sensor reading without an API Key
    Given I have the Sensor request "valid_single_sensor_reading"
    When I call POST without API key to "/api/v1/sensor_readings"
    Then the result should be a 401

  @happy
  Scenario: Record a multiple sensor readings
    Given I have the Sensor request "valid_multiple_sensor_readings"
    When I call POST to "/api/v1/sensor_readings"
    Then the result should be a 201
    And "3" messages should be sent by relay id "53644F1C-2480-4F9B-9CBA-26D66139D221" for sensor "D195E2B3-A1F6-407C-83F5-334FA46B5F6C"
