(ns collatz)

(defn even-number [n]
  (zero? (mod n 2)))

(defn collatz-next [n]
  (cond (even-number n) (quot n 2)
    :else (+ (* 3 n) 1)
  ))

(defn collatz
  "Returns the number of steps for num to reach 1
  according to the Collatz Conjecture."
  [num]
  (loop [n num
         steps 0]
    (if (= n 1) 
      steps
      (recur (collatz-next n) (inc steps))))
)

(collatz 100)