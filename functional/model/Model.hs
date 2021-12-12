module Model
    ( Employee(..),
      setAge,
    ) where

data Employee = Employee { name::String, age::Int, version::Int }

setAge :: Employee -> Int -> Employee

setAge e age = e { age = age }
