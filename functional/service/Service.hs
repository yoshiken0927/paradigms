module Service
    ( putEmployee,
    ) where

import Model
import Repository

putEmployee :: String -> Int -> Employee

infixl 1 |>
(|>) x f = f x

-- OK
-- putEmployee name = setAge (findEmployee name)
-- putEmployee name age = updateEmployee(setAge (findEmployee name) age)
-- putEmployee name age = setAge (findEmployee name) age |> updateEmployee
putEmployee name age = findEmployee name |> setAge age |> updateEmployee

-- NG
-- putEmployee name age = findEmployee name |> setAge(age) |>  updateEmployee
-- putEmployee name = setAge (findEmployee name) |> updateEmployee
