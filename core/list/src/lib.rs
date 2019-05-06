use libc::c_void;

#[no_mangle]
pub extern "C" fn new_list(elems: &[*mut c_void]) -> List {
    List::new(elems)
}
#[no_mangle]
pub extern "C" fn list_len(l: &List) -> usize {
    l.len()
}

/// List is using to implement immutable list for Elz
pub struct List {
    elems: Vec<*mut c_void>,
}

impl List {
    fn new(elems: &[*mut c_void]) -> List {
        List {
            elems: Vec::from(elems),
        }
    }
    fn len(&self) -> usize {
        self.elems.len()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let x = 0 as *mut u8;
        let a: *mut c_void = x as *mut c_void;
        let test_array = [a];
        let l = new_list(&test_array);
        assert_eq!(list_len(&l), 1);
    }
}
