impl Solution {
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        let n = matrix[0].len();
        let inner_count = n / 2;
        let total_count = inner_count + (n % 2);
        for i in 0..total_count {
            for j in 0..inner_count {
                let newValue = matrix[n - 1 - j][i];
                matrix[n - 1 - j][i] = matrix[n - 1 - i][n - j - 1];
                matrix[n - 1 - i][n - j - 1] = matrix[j][n - 1 -i];
                matrix[j][n - 1 - i] = matrix[i][j];
                matrix[i][j] = newValue;
            }
        }
    }
}