use crate::ast::Tag;

pub(crate) trait CodegenTag {
    fn is_builtin(&self) -> bool;
}

impl CodegenTag for Option<Tag> {
    fn is_builtin(&self) -> bool {
        match self {
            Some(tag) => tag.name == "builtin".to_string(),
            None => false,
        }
    }
}
