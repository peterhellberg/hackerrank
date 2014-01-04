;; Namespace
(ns fp-reverse-a-list)

;; Runner, similar to how HackerRank wraps the solution.
(defn run [solution] (
  let [input (line-seq (java.io.BufferedReader. *in*))] (
    println (apply str (map #(str % "\n") (#(solution input))))
  )
))

;; Main function, not needed on HackerRank
(defn -main [] (run  
  ;; Solution
  (fn[lst] (reduce conj '() lst))
))
