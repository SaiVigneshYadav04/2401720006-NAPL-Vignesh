;; (def nums [1 2 3 4 5 6])
;; 
;; (println "COLLECTION 1")
;; 
;; (println "Map:" (map #(* % %) nums))
;; 
;; (println "Filter:" (filter even? nums))
;; 
;; (println "Reduce:" (reduce + nums))
;; 
;; (println "Take:" (take 3 nums))
;; 
;; (println "Drop:" (drop 2 nums))
;; 
;; (println "Threading:"
;;          (->> nums
;;               (filter even?)
;;               (map #(* % %))))


; Collection 2
(def nums2 [10 15 20 25 30])

(println "COLLECTION 2")

(println "Map:" (map #(* % %) nums2))

(println "Filter:" (filter even? nums2))

(println "Reduce:" (reduce + nums2))

(println "Take:" (take 2 nums2))

(println "Drop:" (drop 2 nums2))

(println "Threading:"
         (->> nums2
              (filter even?)
              (map #(* % %))))