class SolutionNaive:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i, val in enumerate(nums):
            for j, nval in enumerate(nums):
                if i != j and val + nval == target:
                    return [i, j]


class SolutionInternet:
    def twoSum(self, nums: List[int], target: int) -> List[int]:

        d_sum = {}
        for i, val in enumerate(nums):
            ntarget = target - val
            if ntarget in d_sum:
                return [d_sum[ntarget], i]
            d_sum[val] = i
