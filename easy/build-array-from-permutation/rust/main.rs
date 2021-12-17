fn main() -> () {}

impl Solution {
    pub fn build_array(nums: Vec<i32>) -> Vec<i32> {
        let mut ans = vec![];

        for n in nums.clone() {
            ans.push(nums[n as usize]);
        }

        ans
    }
}
