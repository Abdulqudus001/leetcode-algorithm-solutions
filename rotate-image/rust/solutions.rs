struct Solution {

}
impl Solution {
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        let n = matrix[0].len();
        let inner_count = n / 2;
        let total_count = inner_count + (n % 2);
        for i in 0..total_count {
            for j in 0..inner_count {
                let new_value = matrix[n - 1 - j][i];
                matrix[n - 1 - j][i] = matrix[n - 1 - i][n - j - 1];
                matrix[n - 1 - i][n - j - 1] = matrix[j][n - 1 -i];
                matrix[j][n - 1 - i] = matrix[i][j];
                matrix[i][j] = new_value;
            }
        }
    }
}

fn main() {
    let mut vec: Vec<Vec<i32>> = vec![[5,1,9,11].to_vec(),[2,4,8,10].to_vec(),[13,3,6,7].to_vec(),[15,14,12,16].to_vec()];
    Solution::rotate(&mut vec);
    println!("{:?}",vec);
}