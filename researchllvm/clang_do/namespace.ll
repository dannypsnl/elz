; ModuleID = 'namespace.cc'
source_filename = "namespace.cc"
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

%"class.Test::Test" = type { i8 }

@_ZN4Test1tE = global %"class.Test::Test" undef, align 1

; Function Attrs: noinline nounwind optnone uwtable
define void @_ZN4Test3fooEv() #0 {
entry:
  ret void
}

attributes #0 = { noinline nounwind optnone uwtable "correctly-rounded-divide-sqrt-fp-math"="false" "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-jump-tables"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+fxsr,+mmx,+sse,+sse2,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.module.flags = !{!0}
!llvm.ident = !{!1}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{!"clang version 6.0.0 (http://llvm.org/git/clang.git 06b063e3bbf2ce05d25770a4c77cf1b748ce198c) (http://llvm.org/git/llvm.git 7e0324beccfab1b203d8a8e6d3249c8a118a85d0)"}
