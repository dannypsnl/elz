; ModuleID = 'structure.c'
source_filename = "structure.c"
target datalayout = "e-m:o-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-apple-macosx10.13.0"

%struct.Label = type { i32, i8* }

; Function Attrs: noinline nounwind optnone ssp uwtable
define { i32, i8* } @NewLabel(i32, i8*) #0 {
  %3 = alloca %struct.Label, align 8
  %4 = alloca i32, align 4
  %5 = alloca i8*, align 8
  %6 = alloca %struct.Label, align 8
  store i32 %0, i32* %4, align 4
  store i8* %1, i8** %5, align 8
  %7 = getelementptr inbounds %struct.Label, %struct.Label* %6, i32 0, i32 0
  %8 = load i32, i32* %4, align 4
  store i32 %8, i32* %7, align 8
  %9 = getelementptr inbounds %struct.Label, %struct.Label* %6, i32 0, i32 1
  %10 = load i8*, i8** %5, align 8
  store i8* %10, i8** %9, align 8
  %11 = load i32, i32* %4, align 4
  %12 = getelementptr inbounds %struct.Label, %struct.Label* %6, i32 0, i32 0
  store i32 %11, i32* %12, align 8
  %13 = load i8*, i8** %5, align 8
  %14 = getelementptr inbounds %struct.Label, %struct.Label* %6, i32 0, i32 1
  store i8* %13, i8** %14, align 8
  %15 = bitcast %struct.Label* %3 to i8*
  %16 = bitcast %struct.Label* %6 to i8*
  call void @llvm.memcpy.p0i8.p0i8.i64(i8* %15, i8* %16, i64 16, i32 8, i1 false)
  %17 = bitcast %struct.Label* %3 to { i32, i8* }*
  %18 = load { i32, i8* }, { i32, i8* }* %17, align 8
  ret { i32, i8* } %18
}

; Function Attrs: argmemonly nounwind
declare void @llvm.memcpy.p0i8.p0i8.i64(i8* nocapture writeonly, i8* nocapture readonly, i64, i32, i1) #1

attributes #0 = { noinline nounwind optnone ssp uwtable "correctly-rounded-divide-sqrt-fp-math"="false" "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-jump-tables"="false" "no-nans-fp-math"="false" "no-signed-zeros-fp-math"="false" "no-trapping-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="penryn" "target-features"="+cx16,+fxsr,+mmx,+sse,+sse2,+sse3,+sse4.1,+ssse3,+x87" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { argmemonly nounwind }

!llvm.module.flags = !{!0, !1}
!llvm.ident = !{!2}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{i32 7, !"PIC Level", i32 2}
!2 = !{!"clang version 6.0.0 (tags/RELEASE_600/final)"}
