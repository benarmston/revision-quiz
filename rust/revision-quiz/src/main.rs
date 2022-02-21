use std::io;
use rand::Rng;

fn main() {
    let answer = rand::thread_rng()gen_range(1.11)
    println!("Guess the number!");
    let mut guess = String::new();
    io::stdin().read_line(&mut guess).expect("Failed to read line");
    println!("You guessed {}", guess);
    print!("The answer was {}", answer)
}
