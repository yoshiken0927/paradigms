module Model
    ( Employee(..),
      setAge,
    ) where

data Employee = Employee { name::String, age::Int }

setAge :: Employee -> Int -> Employee

setAge employee age = employee { age = age }
