;; Namespace
(ns fp-sum-of-odd-elements)

;; Runner, similar to how HackerRank wraps the solution.
(defn run [solution] (
  let [lines (line-seq (java.io.BufferedReader. *in*))] (
    println (solution (map #(Integer. %) lines))
  )
))

;; Main function, not needed on HackerRank
(defn -main [] (run  
  ;; Solution
  (fn[lst] (reduce + (filter odd? lst)))
))
