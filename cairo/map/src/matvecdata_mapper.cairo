use array::ArrayTrait;
use algorithm::matvectmult::{Matrix, Vec, matrixTrait, vecTrait, mapper, reducer, final_output};//return matrix, row, col
fn create_matrix() -> (Matrix,u32,u32){
    let row1 = array![1, 2, 3, 4];
    let row2 = array![3, 4, 5, 5];
    let row3 = array![6, 7, 8, 9];
    let matrix_array = array![row1, row2, row3];
    let mat = matrixTrait::init_array(3, 4, @matrix_array);
    (mat,3,4)
}

//return row, vector_length
fn create_vector()->(Vec,u32){
    let vec_test = array![1, 1, 2, 3];
    let vec = vecTrait::init_array(4, @vec_test);
    (vec,4)
}
