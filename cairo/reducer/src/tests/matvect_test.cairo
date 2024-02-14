use algorithm::matvectmult::{matrixTrait, vecTrait, mapper, reducer, final_output};
use core::debug::PrintTrait;


#[test]
#[available_gas(2000000)]
fn testing_init_one() {
    //[[1,1,1],[1,1,1],[1,1,1]] @ [1,1,1] = [3,3,3]
    let mat = matrixTrait::init_one(3, 3); //(3*3 matrix with entry 1)
    let vec = vecTrait::init_one(3);
    let map_res = mapper(@mat, @vec);
    let (_, sum) = reducer(2, @map_res);
    assert(sum == 3, 'result should be 3');
}
#[test]
#[available_gas(2000000)]
fn testing_init_arr() {
    //[[1,2,3],[4,5,6],[7,8,9]] @ [1,1,1] = [6,15,24]

    let row1 = array![1, 2, 3];
    let row2 = array![4, 5, 6];
    let row3 = array![7, 8, 9];
    let matrix_array = array![row1, row2, row3];
    let mat = matrixTrait::init_array(3, 3, @matrix_array);
    assert(mat.get_size() == (3, 3), 'dimension mismatch');

    let vec_one = vecTrait::init_one(3);

    //testing multiplication [[1,2,3],[4,5,6],[7,8,9]] @ [1,1,1] = [6,15,24]
    let map_res = mapper(@mat, @vec_one);
    let (_, first_row_sum) = reducer(0, @map_res);
    assert(first_row_sum == 6, 'result should be 6');
    let (_, second_row_sum) = reducer(1, @map_res);
    assert(second_row_sum == 15, 'result should be 15');
    let (_, third_row_sum) = reducer(2, @map_res);
    assert(third_row_sum == 24, 'result should be 24');
}
#[test]
#[available_gas(2000000)]
fn testing_init_arr_two() {
    //testing multiplication with [[1,2,3],[4,5,6],[7,8,9]] @ [2,4,6] = [28,64,100]
    let row1 = array![1, 2, 3];
    let row2 = array![4, 5, 6];
    let row3 = array![7, 8, 9];
    let matrix_array = array![row1, row2, row3];
    let mat = matrixTrait::init_array(3, 3, @matrix_array);
    assert(mat.get_size() == (3, 3), 'dimension mismatch');
    let vec_test = array![2, 4, 6];
    let vec = vecTrait::init_array(3, @vec_test);
    let map_res = mapper(@mat, @vec);
    let (_, first_row_sum) = reducer(0, @map_res);
    assert(first_row_sum == 28, 'result should be 28');
    let (_, second_row_sum) = reducer(1, @map_res);
    assert(second_row_sum == 64, 'result should be 64');
    let (_, third_row_sum) = reducer(2, @map_res);
    assert(third_row_sum == 100, 'result should be 100');
}

#[test]
#[available_gas(2000000)]
fn testing_init_arr_three() {
    //testing multiplication with [[1,2,3],[4,5,6],[7,8,9]] @ [2,4,6] = [28,64,100]
    let row1 = array![1, 2, 3];
    let row2 = array![4, 5, 6];
    let row3 = array![7, 8, 9];
    let matrix_array = array![row1, row2, row3];
    let mat = matrixTrait::init_array(3, 3, @matrix_array);
    assert(mat.get_size() == (3, 3), 'dimension mismatch');
    let vec_test = array![2, 4, 6];
    let vec = vecTrait::init_array(3, @vec_test);
    let map_res = mapper(@mat, @vec);
    let res = final_output(3, @map_res);
    res.print();
}
