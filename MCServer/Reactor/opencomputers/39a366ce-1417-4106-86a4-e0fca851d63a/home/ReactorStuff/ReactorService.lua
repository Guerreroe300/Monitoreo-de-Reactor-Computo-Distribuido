local component = require("component")
local internet = require("internet")
local os = require("os")
local serialization = require("serialization")

reactor = component.nc_fission_reactor

-- /putTemp?temp=2134.98
-- /getCmd
-- Target URL
local url = "http://192.168.1.83:"
local tempPort = "8081"
local cmdPort = "8082"

print("Starting Reactor Service")

while true do
  -- Get The current Reactor temp
  local power = reactor.getHeatLevel()

  -- Send out Temperature
  local handle, result = internet.request(url .. tempPort .. "/putTemp?temp=" .. power)

  if not handle then
    io.stderr:write("Request to temp failed: " .. tostring(result) .. "\n")
    return
  end

  -- For reading response, ion need to but needs to close and it happens here somehow
  local response = ""
  for chunk in handle do
    response = response .. chunk
  end

  -- Get commands to runn
  local handlec, resultc = internet.request(url .. cmdPort .. "/getCmd")
  if not handlec then
  	io.stderr:write("Request to cmd failed: " .. tostring(resultc) .. "\n")
	return
  end

  local code = handlec.response()

  if code ~= 404 then
    local responsec = ""
    for chunk in handlec do
      responsec = responsec .. chunk
    end
    print("Response body:", responsec)
    local cmd = serialization.unserialize(responsec)
    if cmd == "Shutdown" then
      print("Shutting down by Command")
      break
    end
  end

  os.sleep(5)
end
