
data Employee = Employee { name::String, age::Int }

findEmployee :: String -> Employee
setAge :: Employee -> Int -> Employee
updateEmployee :: String -> Int -> Employee

findEmployee name = Employee { name = name, age = 30 }
setAge employee age = employee { age = age }
updateEmployee name age = setAge (findEmployee name) age

str = "xyz"
someFunc3 = take 1 str

print = putStrLn(name (findEmployee "yoshiken") ++ ":" ++ show (age (updateEmployee "yoshiken" 31)))
