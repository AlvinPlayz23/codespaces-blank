/// Checks if two strings are equal in an ASCII case-insensitive manner.
pub fn ascii_eq_ignore_case(a: &str, b: &str) -> bool {
    a.eq_ignore_ascii_case(b)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_ascii_eq_ignore_case() {
        assert!(ascii_eq_ignore_case("HTML", "html"));
        assert!(ascii_eq_ignore_case("html", "HTML"));
        assert!(ascii_eq_ignore_case("hTmL", "HtMl"));
        assert!(!ascii_eq_ignore_case("html", "htm"));
        assert!(!ascii_eq_ignore_case("html", "css"));
    }
}
