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
