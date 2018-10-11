use super::ast::*;

pub fn visit_program(ast_tree: Vec<Top>) {
    for ast in ast_tree {
        match ast {
            Top::Import(chain, block) => {
                println!("chain: {:?}, block: {:?}", chain, block);
            }
            Top::GlobalBind(exported, name, e) => {
                println!("global bind: {} = {:?}, exported: {}", name, e, exported);
            }
            _ => println!("Not implement yet"),
        }
    }
}
