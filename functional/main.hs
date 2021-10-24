data Employee = Employee { name::String, age::Int }

findEmployee :: String -> Employee
setAge :: Employee -> Int -> Employee
updateEmployee :: String -> Int -> Employee
updateEmployee2 :: Employee -> Employee

findEmployee name = Employee { name = name, age = 30 }
updateEmployee2 e = e

setAge employee age = employee { age = age }

--updateEmployee name = setAge (findEmployee name)
updateEmployee name age = updateEmployee2(setAge (findEmployee name) age)

someFunc = putStrLn(name (findEmployee "yoshiken") ++ ":" ++ show (age (updateEmployee "yoshiken" 31)))
