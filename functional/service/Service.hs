module Service
    ( putEmployee,
    ) where

import Model
import Repository

-- TODO バリデーションとエラーハンドリング
putEmployee :: String -> Int -> Employee
putEmployee name age = do
  updateEmployee(setAge (findEmployee name) age)

-- infixl 1 |>
-- (|>) x f = f x
-- putEmployee name age = do
--   (findEmployee name |> setAge $ age) |> updateEmployee
