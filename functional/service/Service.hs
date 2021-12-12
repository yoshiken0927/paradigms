module Service
    ( putEmployee,
    ) where

import Model
import Repository

putEmployee :: String -> Int -> Employee

putEmployee name age = updateEmployee(setAge (findEmployee name) age)
