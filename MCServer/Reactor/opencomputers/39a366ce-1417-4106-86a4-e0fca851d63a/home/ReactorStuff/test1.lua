local component = require "component"
local event = require "event"
local term = require "term"
local computer = require "computer"
reactor = component.nc_fission_reactor

computer.beep()

print("Initialising...")

repeat
 local power = reactor.getEnergyChange()
 term.clear()
  if (power < 0) then
print("Reactor ONLINE")
  else
print("Reactor OFFLINE")
  end
  currentHeat = reactor.getHeatLevel()
  maxHeat = 2400000
  print("")
  print("Heat")
  print(currentHeat .. "/" .. maxHeat)
  cells = reactor.getNumberOfCells()
  print("")
  print("Cells Remaining")
  print(cells)
 
  if (currentHeat > 0) then
computer.beep()
  end
  if (currentHeat > (maxHeat/2)) then
reactor.deactivate()    
  end
until event.pull(1) == "interrupted"
