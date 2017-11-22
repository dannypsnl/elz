; ModuleID = 'template.cc'
source_filename = "template.cc"
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

%class.Test = type { i32 }
%class.Test.0 = type { i8* }

@t = global %class.Test { i32 1 }, align 4
@t2 = global %class.Test.0 zeroinitializer, align 8

; Function Attrs: noinline nounwind optnone uwtable
define void @_Z3fooi(i32) #0 {
entry:
  %.addr = alloca i32, align 4
  store i32 %0, i32* %.addr, align 4
  ret void
}

attributes #0 = { noinline nounwind optnone uwtable "correctly-rounded-divide-sqrt-fp-math"="false" "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-jump-tables"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+fxsr,+mmx,+sse,+sse2,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.module.flags = !{!0}
!llvm.ident = !{!1}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{!"clang version 6.0.0 (http://llvm.org/git/clang.git 06b063e3bbf2ce05d25770a4c77cf1b748ce198c) (http://llvm.org/git/llvm.git 7e0324beccfab1b203d8a8e6d3249c8a118a85d0)"}
