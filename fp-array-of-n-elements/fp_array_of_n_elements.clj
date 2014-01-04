;; Namespace
(ns fp-array-of-n-elements)

;; Runner, similar to how HackerRank wraps the solution.
(defn run [solution] (
  let [input (line-seq (java.io.BufferedReader. *in*))] (
    println(count(solution (Integer. (first input))))
  )
))

;; Main function, not needed on HackerRank
(defn -main [] (run  
  ;; Solution
  (fn[n] (take n (iterate inc 1)))
))
