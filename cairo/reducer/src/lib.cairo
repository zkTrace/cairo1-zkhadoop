// === Header Start ===
pub mod matvecdata_reducer;
pub mod matvectmult;
use cairoreduce::matvecdata_reducer::inter_result;
use cairoreduce::matvectmult::reducer;
// === Header End ===

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


   
  