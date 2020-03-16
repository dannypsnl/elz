use crate::ast::Tag;

pub(crate) trait SemanticTag {
    fn is_extern(&self) -> bool;
}

impl SemanticTag for Option<Tag> {
    fn is_extern(&self) -> bool {
        match self {
            Some(tag) => {
                tag.name.as_str() == "extern"
                    && tag.properties.len() == 1
                    && tag.properties.last() == Some(&"c".to_string())
            }
            None => false,
        }
    }
}
