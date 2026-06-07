; Everything that comes after a semicolon (;) is a comment
; Assembler-time constants may be defined using 'equ'

section .text
; You should implement functions in the .text section

; the global directive makes a function visible to the test files
global expected_minutes_in_oven
expected_minutes_in_oven:
    ; write 40 to a register
    mov rax, 40
    ret

global remaining_minutes_in_oven
remaining_minutes_in_oven:
    ; call the above func, sub the register by rdi (the param)
    call expected_minutes_in_oven
    sub rax, rdi 
    ret

global preparation_time_in_minutes
preparation_time_in_minutes:
    ; takes an argument and doubles it
    ; store the paramenter in rax (always returned)
    mov rax, rdi
    imul rax, 2
    ret

global elapsed_time_in_minutes
elapsed_time_in_minutes:
    ; This function takes two numbers as arguments and must return a number
    call preparation_time_in_minutes
    add rax, rsi
    ret

%ifidn __OUTPUT_FORMAT__,elf64
section .note.GNU-stack noalloc noexec nowrite progbits
%endif
