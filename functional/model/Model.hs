module Model
    ( Employee(..),
      setAge,
    ) where

-- TODO カプセル化
data Employee = Employee {
  name::String,
  age::Int,
  version::Int
}

-- TODO バリデーションとエラーハンドリング
setAge :: Employee -> Int -> Employee
setAge e age = e { age = age }
