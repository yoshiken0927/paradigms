module Model
    ( Employee(..),
      setAge,
    ) where

data Employee = Employee { name::String, age::Int }

setAge :: Int -> Employee -> Employee

setAge age employee = employee { age = age }
