use std::io;

mod interpretmath;

use interpretmath::{
    ast,
    parser::{ParseError, Parser},
};

fn evaluate(expr: String) -> Result<f64, ParseError> {
    let expr = expr.split_whitespace().collect::<String>();
    let mut math_parser = Parser::new(&expr)?;
    let ast = math_parser.parse()?;
    println!("The generated AST is {:?}", ast);

    Ok(ast::eval(ast)?)
}

fn main() {
    println!("Hello! Welcome to Arithmetic expression evaluator.");
    println!("Example: 2*3+(4-5)+2^3/4");
    println!("Enter new expression: ");
    loop {
        let mut input = String::new();
        match io::stdin().read_line(&mut input) {
            Ok(_) => {
                match evaluate(input) {
                    Ok(val) => println!("The result is {}\n", val),
                    Err(_) => {
                        println!("Error Evaluating or Invalid expression\n");
                    }
                };
            }

            Err(error) => println!("error: {}", error),
        }
    }
}
