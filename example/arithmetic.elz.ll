%list = type opaque

@0 = global [28 x i8] c"1 + 2 = %d\0A3.14 + 4.56 = %f\0A", align 1

declare i64 @printf(i8* %format, ...)

declare i8* @elz_malloc(i64 %size)

declare %list* @new_list(i64 %size, i8** %elements)

declare i8* @list_index(%list* %list, i64 %index)

declare i64 @list_length(%list* %list)

define void @init() {
; <label>:0
	ret void
}

define i64 @main() {
; <label>:0
	call void @init()
	%1 = getelementptr [28 x i8], [28 x i8]* @0, i64 0, i64 0
	%2 = alloca i8*
	store i8* %1, i8** %2
	%3 = load i8*, i8** %2
	%4 = add i64 1, 2
	%5 = fadd double 0x40091EB851EB851F, 0x40123D70A3D70A3D
	%6 = call i64 (i8*, ...) @printf(i8* %3, i64 %4, double %5)
	ret i64 0
}
