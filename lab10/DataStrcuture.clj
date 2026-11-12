;; List
(def my-list '(10 20 30))

(println "Original List:" my-list)
(println "First Element:" (first my-list))
(println "Rest of List:" (rest my-list))

(def new-list (conj my-list 5))

(println "New List:" new-list)
(println "Original List:" my-list)

;; Vector
(def my-vector [10 20 30])

(println "Original Vector:" my-vector)
(println "Element at index 1:" (get my-vector 1))
(println "Count:" (count my-vector))

(def new-vector (conj my-vector 40))

(println "New Vector:" new-vector)
(println "Original Vector:" my-vector)

;; Map
(def my-map {:name "Sid" :age 20})

(println "Original Map:" my-map)
(println "Name:" (get my-map :name))

(def new-map (assoc my-map :city "Delhi"))

(println "New Map:" new-map)

(def updated-map (dissoc new-map :age))

(println "Updated Map:" updated-map)
(println "Original Map:" my-map)

;; Set 

(def my-set #{10 20 30})

(println "Original Set:" my-set)
(println "Contains 20:" (contains? my-set 20))

(def new-set (conj my-set 40))

(println "New Set:" new-set)

(def updated-set (disj new-set 10))

(println "Updated Set:" updated-set)
(println "Original Set:" my-set)