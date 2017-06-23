$g28 = comdat any
$com1 = comdat exactmatch

; Mutable.
@g1 = global i32 0

; Immutable.
@g2 = constant i32 0

; External linkage.
@g3 = external global i32
@g4 = extern_weak global i32

; Linkage.
@g5 = appending global i32 0
@g6 = available_externally global i32 0
@g7 = common global i32 0
@g8 = internal global i32 0
@g9 = linkonce global i32 0
@g10 = linkonce_odr global i32 0
@g11 = private global i32 0
@g12 = weak global i32 0
@g13 = weak_odr global i32 0

; Visibility.
@g14 = default global i32 0
@g15 = hidden global i32 0
@g16 = protected global i32 0

; DLL storage class.
@g17 = dllimport global i32 0
@g18 = dllexport global i32 0

; Thread local storage.
@g19 = thread_local global i32 0
@g20 = thread_local(localdynamic) global i32 0
@g21 = thread_local(initialexec) global i32 0
@g22 = thread_local(localexec) global i32 0

; Unnamed address.
@g23 = local_unnamed_addr global i32 0
@g24 = unnamed_addr global i32 0

; Address space.
@g25 = addrspace(1) global i32 0

; Externally initialized.
@g26 = externally_initialized global i32 0

; Section.
@g27 = global i32 0, section "foo"

; Comdat.
@g28 = global i32 0, comdat
@g29 = global i32 0, comdat($com1)

; Align.
@g30 = global i32 0, align 8

; Metadata.
@g31 = global i32 0, !foo !{!"bar"}, !baz !{!"qux"}

; Full global definition.
@g32 = common default dllexport thread_local(localdynamic) unnamed_addr addrspace(1) externally_initialized global i32 0, section "foo", comdat($com1), align 8, !foo !{!"bar"}, !baz !{!"qux"}
