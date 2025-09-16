local component = require("component")
local internet = require("internet")
local os = require("os")

-- Target URL
local url = "http://192.168.1.83:8081/putTemp?temp=2134.98"

while true do
  -- Perform GET request
  local handle, result = internet.request(url)

  if not handle then
    io.stderr:write("Request failed: " .. tostring(result) .. "\n")
    return
  end

  -- Read response
  local response = ""
  for chunk in handle do
    response = response .. chunk
  end

  print("Response from server:")
  print(response)
  os.sleep(5)
end