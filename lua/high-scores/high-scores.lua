return function(scores)
	local T = {}
	T.score = scores

	function T:scores()
		return self.score
	end

	function T:latest()
		return self.score[#self.score]
	end

	function T:personal_best()
		local highest = 0
		for i = 1, #self.score, 1 do
			if self.score[i] > highest then
				highest = self.score[i]
			end
		end

		return highest
	end

	function T:personal_top_three()
		local sortedTbl = { table.unpack(T.score) }

		table.sort(sortedTbl, function(a, b)
			return a > b
		end)

		while #sortedTbl > 3 do
			table.remove(sortedTbl)
		end

		return sortedTbl
	end

	return T
end
