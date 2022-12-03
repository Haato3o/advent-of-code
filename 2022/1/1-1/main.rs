use std::{fs::File, io::{BufReader, Read}};

fn main() {
    let input = File::open("input").unwrap();
    let mut buffer_reader = BufReader::new(input);
    let mut buffer = String::new();

    buffer_reader.read_to_string(&mut buffer).unwrap();

    let mut max: i64 = 0;

    for line in buffer.split("\n\n") {

        let calories: i64 = line.split("\n")
            .map(|a| a.parse::<i64>().unwrap())
            .sum();

        if calories > max {
            max = calories;
        }
    }

    println!("Result: {}", max);
}