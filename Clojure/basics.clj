; basics of clojure programming
; clojure seems very REPL based
; to load functions here use (load-file "./basics.clj")
; then call functions with (basics/func-name argument)

; set name space
(ns basics)

(defn hello [name] (str "Hello, " name))
;hello "GitHub"

(defn divide [x y] (quot x y))
;divide 25 5

; Vectors = [1 2 3], sequential data
; List = (1 2 3), can be quoted with = (quote (1 2 3)) or `(1 2 3)
; Set = {2 5 6 1}, no dupes
; Maps = {:language "Clojure" :favorite false}

; Work with Maps
(defrecord ProgrammingLanguage [language favorite])
(def new_language (->ProgrammingLanguage "Clojure" false))

(get new_language :language)
;more idomatic: 
(:favorite new_language)