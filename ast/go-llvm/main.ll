; ModuleID = 'main'
source_filename = "main"

@.str = private unnamed_addr constant [23 x i8] c"detect variant args: %d", align 1
@"main::float" = global float 0x4031B93CE0000000
@"main::string" = global [51 x i8] c"\5C\5Ca\E4\BD\A0\E5\A5\BD, llvm, $@#%^!&!)~!#*(@#+_)(*&GBJNLSfdlbc)"
@"main::A struct" = global { float } { float 0x40091EB860000000 }

declare [51 x i8] @"main::foo_string_string"([51 x i8])

define i32 @"main::add"(i32 %lv, i32 %rv) {
entry:
  %result = add i32 %lv, %rv
  ret i32 %result
}

declare i32 @printf(i8*, ...)

define i32 @main() {
entry:
  %tmp = call i32 @"main::add"(i32 2, i32 3)
  %0 = call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([23 x i8], [23 x i8]* @.str, i32 0, i32 0), i32 %tmp)
  ret i32 0
}
