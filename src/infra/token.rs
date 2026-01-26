#[derive(Debug, Clone, PartialEq)]
pub enum Token {
    Character(char),
    EndOfFile,
    // Future tokens will be added here or in module-specific token types.
}

pub trait Tokenizer {
    fn next_token(&mut self) -> crate::infra::error::Result<Token>;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_token_equality() {
        assert_eq!(Token::Character('a'), Token::Character('a'));
        assert_ne!(Token::Character('a'), Token::Character('b'));
        assert_eq!(Token::EndOfFile, Token::EndOfFile);
        assert_ne!(Token::Character('a'), Token::EndOfFile);
    }
}
