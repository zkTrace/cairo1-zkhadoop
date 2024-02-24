// === Header Start ===use array::ArrayTrait;
use cairomap::matvectmult::{Matrix, Vec, matrixTrait, vecTrait, mapper, reducer, final_output};// === Header End ===//return matrix, row, col
fn create_matrix() -> (Matrix,u32,u32){
    let row1 = array![4, 54, 61, 73, 1, 26, 59, 62, 35, 83];
    let row2 = array![4, 66, 62, 41, 9, 31, 95, 46, 5, 53];
    let row3 = array![77, 45, 48, 53, 36, 86, 33, 58, 22, 87];
    let row4 = array![84, 46, 17, 58, 98, 30, 56, 78, 48, 5];
    let row5 = array![0, 30, 17, 24, 38, 68, 46, 98, 30, 40];
    let row6 = array![70, 57, 55, 60, 8, 83, 74, 41, 64, 20];
    let row7 = array![52, 30, 4, 4, 63, 38, 77, 84, 9, 68];
    let row8 = array![19, 49, 72, 47, 76, 19, 14, 99, 98, 12];
    let row9 = array![21, 24, 44, 55, 53, 57, 31, 87, 35, 18];
    let row10 = array![66, 22, 15, 34, 58, 38, 21, 84, 83, 22];
    let matrix_array = array![row1, row2, row3, row4, row5, row6, row7, row8, row9, row10];
    let mat = matrixTrait::init_array(10, 10, @matrix_array);
    (mat,10,10)
}

//return row, vector_length
fn create_vector()->(Vec,u32){
    let vec_test = array![9, 2, 2, 4, 9, 10, 3, 1, 7, 9];
    let vec = vecTrait::init_array(10, @vec_test);
    (vec,10)
}
