local ArmstrongNumbers = {}

function ArmstrongNumbers.is_armstrong_number(number)
	if number == 0 then
		return true
	end

	Sum = 0
	local original_number = number
	local len_number = #tostring(original_number)

	while number > 0 do
		local digit = number % 10
		Sum = Sum + (digit ^ len_number)
		number = math.floor(number / 10)
	end

	return Sum == original_number
end

return ArmstrongNumbers
