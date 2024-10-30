#[derive(Debug, Clone)]
pub enum Token {
    Add,
    Subtract,
    Multiply,
    Divide,
    Caret,
    LeftParen,
    RightParen,
    Num(f64),
    EOF,
}

impl Token {
    pub fn get_oper_prec(&self) -> OperPrec {
        use self::OperPrec::*;
        use self::Token::*;

        match *self {
            Add | Subtract => AddSub,
            Multiply | Divide => MulDiv,
            Caret => Power,

            _ => DefaultZero,
        }
    }
}

#[derive(Debug, PartialEq, PartialOrd)]
pub enum OperPrec {
    DefaultZero,
    AddSub,
    MulDiv,
    Power,
    Negative,
}

impl PartialEq for Token {
    fn eq(&self, other: &Self) -> bool {
        match (self, other) {
            (Token::Num(a), Token::Num(b)) => (a - b).abs() < f64::EPSILON, // Comparing floats with tolerance
            (Token::Add, Token::Add)
            | (Token::Substract, Token::Substract)
            | (Token::Multiply, Token::Multiply)
            | (Token::Divide, Token::Divide)
            | (Token::Caret, Token::Caret)
            | (Token::LeftParen, Token::LeftParen)
            | (Token::RightParen, Token::RightParen)
            | (Token::EOF, Token::EOF) => true,
            _ => false,
        }
    }
}
