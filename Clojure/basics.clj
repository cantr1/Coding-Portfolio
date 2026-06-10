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


; square func
(defn square [n]
  (* n n))

; return the sum of squares with reduce
(reduce + (map square (range 1 6)))


; conditionals
(defn production-rate
  "Returns an assembly line's production rate per hour,
   taking into account its success rate"
  [speed]
  (cond (== speed 0) 0
    (and (> speed 0) (<= speed 4)) (* speed 221)
    (and (> speed 4) (<= speed 8)) (* speed 221 0.90)
    (== speed 9) (* speed 221 0.80)
    (== speed 10) (* speed 221 0.77)
  )
)

(defn working-items
  "Calculates how many working cars are produced per minute"
  [speed]
  (int (quot (production-rate speed) 60))
  )

(working-items 5)
