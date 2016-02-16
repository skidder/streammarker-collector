Given(/^I have the Sensor request "(.*)"$/) do |name|
  @request = get_request(name)
end

When(/^I call POST to "(.*)"$/) do | path |
  @response = HTTParty.post(APPLICATION_ENDPOINT + "#{path}",
                            body: @request,
                            headers: { 'Content-Type' => 'application/json',
                                        'X-API-KEY' => ENV['STREAMMARKER_COLLECTOR_API_TOKENS'] })
end

When(/^I call POST with API key "(.*)" to "(.*)"$/) do | api_key, path |
  @response = HTTParty.post(APPLICATION_ENDPOINT + "#{path}",
                            body: @request,
                            headers: { 'Content-Type' => 'application/json',
                                        'X-API-KEY' => api_key })
end

When(/^I call POST without API key to "(.*)"$/) do | path |
  @response = HTTParty.post(APPLICATION_ENDPOINT + "#{path}",
                            body: @request,
                            headers: { 'Content-Type' => 'application/json' })
end

Given(/^I have the Asset update request "(.*)"$/) do |name|
  @request = get_request(name)
end

And(/^the asset ID should not have changed for the asset "(.*)"$/) do |name|
  json = JSON.parse(@response.body)
  @assets[name]['id'].should eq(json['id'])
end

# Hack hack hack! This is to allow us to stash the asset in our internal model
# when we are doing a single "POST" test
Then(/^I save the response as the asset "(.*?)"$/) do |name|
  @assets[name] = JSON.parse(@response.body)
end

Then(/^a message should be sent by relay id "(.*?)" for sensor "(.*?)"$/) do |relay_device, sensor_id|
  m = wait_for_sqs_message(ENV['STREAMMARKER_SQS_QUEUE_URL'], 5)
  json = JSON.parse(m[:body])
  json['relay_id'].should eq (relay_device)
  json['sensor_id'].should eq (sensor_id)
end

Then(/^"(.*?)" messages should be sent by relay id "(.*?)" for sensor "(.*?)"$/) do |message_count, relay_device, sensor_id|
  m = wait_for_sqs_message(ENV['STREAMMARKER_SQS_QUEUE_URL'], 5, message_count.to_i)
  json = JSON.parse(m[:body])
  json['relay_id'].should eq (relay_device)
  json['sensor_id'].should eq (sensor_id)
end

def wait_for_sqs_message(queue, timeout, message_count=1)
  puts "Waiting for message in queue : " + queue + " timeout: " + timeout.to_s

  sqs = AWS::SQS.new(:access_key_id   => 'x',
                   :secret_access_key => 'y',
                   :use_ssl           => false,
                   :sqs_endpoint      => FAKESQS_HOST,
                   :sqs_port          => FAKESQS_PORT.to_i
                   )
  resp = {}

  (1..timeout).each do
    resp = sqs.client.receive_message(queue_url: queue, visibility_timeout: 10, max_number_of_messages: 10)
    break unless resp.data[:messages].empty?
    sleep 1
  end

  if resp.data[:messages].empty?
    raise "No messages arrived after #{timeout} seconds"
  end

  if resp.data[:messages].length > message_count
    raise "Got more than #{message_count} message(s)"
  end

  msg = resp.data[:messages].first
  sqs.client.delete_message(queue_url: queue, receipt_handle: msg[:receipt_handle])

  msg
end
