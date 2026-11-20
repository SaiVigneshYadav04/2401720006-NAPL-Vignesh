
(println "1. Arithmetic Exception")

(try
  (/ 10 0)
  (catch ArithmeticException e
    (println "Error:" (.getMessage e)))
  (finally
    (println "Finally block executed")))

(println "\n2. Null Pointer Exception")

(try
  (.length ^String nil)
  (catch NullPointerException e
    (println "Error: Null value found"))
  (finally
    (println "Finally block executed")))

(println "\n3. Custom Exception")

(defn check-age [age]
  (if (< age 18)
    (throw (ex-info "Age must be 18 or above"
                    {:age age :minimum 18}))
    (println "Age is valid")))

(try
  (check-age 15)
  (catch clojure.lang.ExceptionInfo e
    (println "Error:" (ex-message e))
    (println "Exception Data:" (ex-data e))
    (println "Recovering from exception")) 
  (finally
    (println "Finally block executed")))
