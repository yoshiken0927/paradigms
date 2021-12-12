module Service
    ( putEmployee,
    ) where

import Model
import Repository

putEmployee :: String -> Int -> Employee

infixl 1 |>
(|>) x f = f x

-- putEmployee name age = updateEmployee(setAge (findEmployee name) age)
putEmployee name age = (findEmployee name |> setAge $ age) |> updateEmployee
