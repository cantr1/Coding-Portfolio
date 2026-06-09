; basics of clojure programming
; clojure seems very REPL based
; to load functions here use (load-file "./basics.clj")
; then call functions with (basics/func-name argument)

; set name space
(ns basics)

(defn hello [name] (str "Hello, " name))
;hello "GitHub"

