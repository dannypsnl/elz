use crate::ast::Tag;

pub(crate) trait SemanticTag {
    fn is_builtin(&self) -> bool;
}

impl SemanticTag for Option<Tag> {
    fn is_builtin(&self) -> bool {
        match self {
            Some(tag) => tag.name == "builtin".to_string(),
            None => false,
        }
    }
}
