
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
// Implement the printing function of the vector
/////////////////////////////////////////////////
impl vecPrintImpl of PrintTrait<Vec> {
    fn print(self: Vec) {
        // self.size.print();
        let mut i = 0;
        loop {
            if (i >= self.size) {
                break;
            }
            let (_index, value) = self.data.at(i);
            let temp_val = *value;
            temp_val.print();
            i += 1;
        }
    }
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
        // if (*vec_value==0){
        //     i+=1;
        //     continue;
        // }
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
fn reducer(key: u32, mapper_result: @Array<(u32, felt252)>) -> (u32, felt252) {
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
    let vec_test = array![1, 2, 3, 4];
    let vec = vecTrait::init_array(4, @vec_test);
    (vec,4)
}


// #[cfg(test)]
// mod tests;
//This script will return the mapper result

fn main() {
  
    //Mapper Job
    let (mat,row,col) = create_matrix();
    let (vec,vec_length) = create_vector();
    assert(col==vec_length,'dimension mismatch in cairo map');
    let map_res = mapper(@mat, @vec);
    let total_length= row*col;
    let map_snapshot= @map_res;
    let mut i =0;
    let header: ByteArray = "{\n \"intermediary_values\": [ ";
    println!("{}",header);
    loop{
        if(i>=total_length){
            break;
        }
        let (k1,v1)= *map_snapshot.at(i);
        let kvpair: ByteArray=format!("[{},{}]", k1, v1);
       
        i+=1;
        print!("{}",kvpair);
        if(i<total_length){
            print!(",")
        }

    };
    let end: ByteArray = "]\n}";
    println!("{}",end);



}
