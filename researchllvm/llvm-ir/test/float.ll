; ModuleID = 'float.c'
source_filename = "float.c"
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@a = global float 0x41332C4FE0000000, align 4

!llvm.module.flags = !{!0}
!llvm.ident = !{!1}

!0 = !{i32 1, !"wchar_size", i32 4}
!1 = !{!"clang version 6.0.0 (http://llvm.org/git/clang.git 06b063e3bbf2ce05d25770a4c77cf1b748ce198c) (http://llvm.org/git/llvm.git 7e0324beccfab1b203d8a8e6d3249c8a118a85d0)"}
