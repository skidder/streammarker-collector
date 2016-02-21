def get_request(name)
  get_json_from_fixture_file_as_hash('requests.json', name).to_json
end

def get_response(name)
  get_json_from_fixture_file_as_hash('responses.json', name).to_json
end

def get_json_from_fixture_file_as_hash(file, name)
  request = get_fixture_file_as_string(file)
  json = JSON.parse(request)[name]
  raise "Unable to find key '#{name}' in '#{file}'" if json.nil?
  json
end

def get_fixture_file_as_string(filename)
  File.read(File.join(CUCUMBER_BASE, 'fixtures', filename))
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
