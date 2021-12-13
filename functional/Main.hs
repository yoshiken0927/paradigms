module Main where

import Model
import Service






main :: IO ()
main = putStrLn $ format $ putEmployee "yoshiken" 31


format :: Employee -> String
format e = name e ++ ":" ++ show(age e) ++ ":" ++ show(version e)

