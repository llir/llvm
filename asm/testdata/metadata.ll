declare !foo !0 !bar !{!"baz"} !qux !{!2} void @f1()

define void @f2() !foo !0 !bar !{!"baz"} !qux !{!2} {
	ret void
}

; --- [ Metadata definitions ] -------------------------------------------------

; Empty named metadata definition.
!foo = !{}

; Plain named metadata definition.
!bar = !{!0}

; Multiple metadata IDs.
!baz = !{!0, !1}

; Empty metadata definition.
!0 = !{}

; Plain metadata definition.
!1 = !{!0}

; Multiple metadata IDs.
!2 = !{!0, !1}

; Distinct.
!3 = distinct !{!2}

; Nested metadata.
!4 = !{!{!{!0}}}

; Metadata string.
!5 = !{!"foo"}

; Metadata constant.
!6 = !{i32 42}

; Metadata constant.
!7 = !{!{!"bar"}}
