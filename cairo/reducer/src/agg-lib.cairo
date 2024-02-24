
use array::ArrayTrait;
use core::dict::Felt252DictTrait;
use core::traits::TryInto;
use core::traits::Into;
use option::OptionTrait;
use core::debug::PrintTrait;
use core::fmt::Formatter;


/////////////////////////////////////////////////
//* A simple matrix structure 
//* Internal Representation: (row, col, value)
//* @param col_size: u32
//* @param row_size: u32
//* @param data: felt252
/////////////////////////////////////////////////
#[derive(Drop)]
struct Matrix {
    col_size: u32,
    row_size: u32,
    data: Array<(u32, u32, felt252)>
}

/////////////////////////////////////////////////
//* A simple vector structure 
//* Internal Representation: (i ,value)
//* @param col_size: u32
//* @param row_size: u32
//* @param data: felt252
/////////////////////////////////////////////////
#[derive(Drop)]
struct Vec {
    size: u32,
    data: Array<(u32, felt252)>
}

/////////////////////////////////////////////////
/// This is the interface for the matrix object
/////////////////////////////////////////////////
trait matrixTrait {
    
    
    /////////////////////////////////////////////////
    //* Initialze a one matrix with [row,id]
    //* @param col_size: u32
    //* @param row_size: u32
    //* @return Matrix object with (row,col,value=1)
    /////////////////////////////////////////////////
    fn init_one(row: u32, col: u32) -> Matrix;

/////////////////////////////////////////////////
    //* Initialze a matrix with array of array
    //* @param col_size: u32
    //* @param row_size: u32
    //* @param mat_arr: array of arrays
    //* @return Matrix object with (row,col,value)
    /////////////////////////////////////////////////
    fn init_array(row: u32, col: u32, mat_arr: @Array<Array<felt252>>) -> Matrix;
    
    /////////////////////////////////////////////////
    //* Return size of the matrix object
    //* @param snapshot of matrix
    //* @return Matrix object with (row,col)
    /////////////////////////////////////////////////
    fn get_size(self: @Matrix) -> (u32, u32);
}
trait vecTrait {
    /////////////////////////////////////////////////
    //* Initialze a vector with value 1
    //* @param col_size: u32
    //* @param row_size: u32
    //* @return Vector object with (index,value=1)
    /////////////////////////////////////////////////
    fn init_one(size: u32) -> Vec;

    /////////////////////////////////////////////////
    //* Initialze a vector with value an array
    //* @param col_size: u32
    //* @param row_size: u32
    //* @param vec_array: [v1,v2,...]
    //* @return Vector object with (index,value=1)
    /////////////////////////////////////////////////
    fn init_array(size: u32, vec_arr: @Array<felt252>) -> Vec;

    /////////////////////////////////////////////////
    //* Return the length of the array
    //* @param Vec: Snapshot of Vec object
    /////////////////////////////////////////////////
    fn get_size(self: @Vec) -> u32;
}


/////////////////////////////////////////////////
// Implementation fo MatrixTrait
/////////////////////////////////////////////////
impl matrixTraitImp of matrixTrait {
    fn init_one(row: u32, col: u32) -> Matrix {
        let mut matrix = Matrix { col_size: col, row_size: row, data: ArrayTrait::new() };
        let mut i: u32 = 0;
        loop {
            if (i >= row) {
                break;
            }
            let mut j: u32 = 0;
            loop {
                if (j >= col) { // let value: felt252 = 10;
                    break;
                }
                let value: felt252 = 1;
                matrix.data.append((i, j, value));
                j += 1;
            };
            i += 1;
        };
        matrix
    }
    fn init_array(row: u32, col: u32, mat_arr: @Array<Array<felt252>>) -> Matrix {
        let mut matrix = Matrix { col_size: col, row_size: row, data: ArrayTrait::new() };
        assert(row == mat_arr.len(), 'row mismatch');
        assert(row >= 1, 'empty matrix detected');
        assert(col == mat_arr.at(0).len(), 'col mismatch');
        assert(col >= 1, 'col mismatch');
        let mut i: u32 = 0;
        loop {
            if (i >= row) {
                break;
            }
            let mut j: u32 = 0;
            loop {
                if (j >= col) { // let value: felt252 = 10;
                    break;
                }
                let value = mat_arr.at(i).at(j);
                matrix.data.append((i, j, *value));
                j += 1;
            };
            i += 1;
        };
        matrix
    }
    fn get_size(self: @Matrix) -> (u32, u32) {
        (*self.row_size, *self.col_size)
    }
}

/////////////////////////////////////////////////
// Implementation fo VectorTrait
/////////////////////////////////////////////////
impl vecTraitImp of vecTrait {
    fn get_size(self: @Vec) -> u32 {
        *self.size
    }
    fn init_one(size: u32) -> Vec {
        let mut vec = Vec { size: size, data: ArrayTrait::new() };
        let mut i = 0;
        loop {
            if (i >= size) {
                break;
            }
            let value: felt252 = 1;
            vec.data.append((i, value));
            i += 1;
        };
        vec
    }

    fn init_array(size: u32, vec_arr: @Array<felt252>) -> Vec {
        assert(size == vec_arr.len(), 'size mismatch');
        assert(size >= 1, 'empty vector detected');
        let mut vec = Vec { size: size, data: ArrayTrait::new() };

        let mut i = 0;
        loop {
            if (i >= size) {
                break;
            }
            let value: felt252 = *vec_arr.at(i);
            vec.data.append((i, value));
            i += 1;
        };
        vec
    }
}

/////////////////////////////////////////////////
//* Mapper for Matrix Vector Multiplication
//* @param mat: Snapshot of matrix
//* @param vec: Snapshot of Vector
//* @return result: Array of (index,value) 
/////////////////////////////////////////////////
fn mapper(mat: @Matrix, vec: @Vec) -> Array<(u32, felt252)> {
    let (row_size, col_size) = mat.get_size();
    let vec_size = vec.get_size();
    assert(vec_size == col_size, 'Dimension mismatch');
    let total_length = row_size * col_size;
    assert(total_length == mat.data.len(), 'total len neq matrix len');
    let mut i = 0;
    let mut result = ArrayTrait::new();
    loop {
        if (i >= total_length) {
            break;
        };

        assert(i < total_length, 'index out of bound');
        let (row, col, mat_value) = mat.data.at(i);

        assert(*row < row_size, 'row mismatch');
        assert(*col < col_size && *col < vec_size, 'col mismatch');
        let (_vec_index, vec_value) = vec.data.at(*col);
        //don't need to record zero value
        if (*vec_value==0){
            continue;
        }
        let value: felt252 = *mat_value * *vec_value;
        let entry = (*row, value);
        result.append(entry);
        i += 1;
    };
    result
}

/////////////////////////////////////////////////
//* Reducer for Matrix Vector Multiplication
//* @param key: The specific index of the resulting vector
//* @param mapper_result: Result of the vector
//* @return result: Array of (index,value) 
/////////////////////////////////////////////////
pub fn reducer(key: u32, mapper_result: @Array<(u32, felt252)>) -> (u32, felt252) {
    let mut sum = 0;
    let mut i = 0;
    let total_length = mapper_result.len();
    loop {
        if (i >= total_length) {
            break;
        };
        let (row, value) = mapper_result.at(i);
        if (*row == key) {
            sum += *value;
        };
        i += 1;
    };
    (key, sum)
}

////////////////////////////////////////////////////////////////////////////////////
//* Created for debug purpose
//* Combine all the result from the mapper function, call reducer and return a result
//* @param size: The length of the resulting vector
//* @param mapper_result: Result of the vector
//* @return result: Array of (index,value) 
//////////////////////////////////////////////////////////////////////////////////////
fn final_output(size: u32, mapper_result: @Array<(u32, felt252)>) -> Vec {
    let mut temp_vec: Array<felt252> = Default::default();
    let mut i = 0;
    loop {
        if (i >= size) {
            break;
        };
        let (key, sum) = reducer(i, mapper_result);
        assert(key == i, 'order should match');

        temp_vec.append(sum);
        i += 1;
    };
    vecTrait::init_array(size, @temp_vec)
}

pub fn inter_result()->Array<(u32, felt252)>{
    let inter_res = array![(0,36), (0,108), (0,122), (0,292), (0,9), (0,260), (0,177), (0,62), (0,245), (0,747), (1,36), (1,132), (1,124), (1,164), (1,81), (1,310), (1,285), (1,46), (1,35), (1,477), (2,693), (2,90), (2,96), (2,212), (2,324), (2,860), (2,99), (2,58), (2,154), (2,783), (3,756), (3,92), (3,34), (3,232), (3,882), (3,300), (3,168), (3,78), (3,336), (3,45), (4,0), (4,60), (4,34), (4,96), (4,342), (4,680), (4,138), (4,98), (4,210), (4,360), (5,630), (5,114), (5,110), (5,240), (5,72), (5,830), (5,222), (5,41), (5,448), (5,180), (6,468), (6,60), (6,8), (6,16), (6,567), (6,380), (6,231), (6,84), (6,63), (6,612), (7,171), (7,98), (7,144), (7,188), (7,684), (7,190), (7,42), (7,99), (7,686), (7,108), (8,189), (8,48), (8,88), (8,220), (8,477), (8,570), (8,93), (8,87), (8,245), (8,162), (9,594), (9,44), (9,30), (9,136), (9,522), (9,380), (9,63), (9,84), (9,581), (9,198)];
    inter_res
}


//This script will return the mapper result

fn main(job_id: u32) {
  
    let map_res: Array<(u32, felt252)> = inter_result();

    // Reducer Job
    let (k2,v2): (u32, felt252) = reducer(job_id, @map_res);   

    //printing the output
    let header: ByteArray = "{\n \"Reducer_Result\": [ ";
    println!("{}",header);

    print!("[{},{}]",k2,v2);
    let end: ByteArray = "]\n}";
    println!("{}",end);

}


   
  