;; Namespace
(ns fp-hello-world-n-times)

;; Main function, not needed on HackerRank
(defn -main [& args]
  (
    ;; Solution
    (fn[n] (dotimes [n n]
      (println "Hello World")
    )) 
    
    ;; Number of times
    4
  ) 
)
