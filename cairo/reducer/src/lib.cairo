
mod matvectmult;
mod matvecdata_mapper;
use algorithm::matvecdata_mapper::{create_matrix,create_vector};
use algorithm::matvectmult::{matrixTrait, vecTrait, mapper, reducer, final_output};
use core::debug::PrintTrait;
use core::fmt::Formatter;

#[cfg(test)]
mod tests;
//This script will return the mapper result

fn main() {
    //Mapper Job
    let (mat,row,col) = create_matrix();
    let (vec,vec_length) = create_vector();
    assert(col==vec_length,'dimension mismatch');
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
