module Main where

import Model
import Service

main :: IO ()
parse :: Employee -> String

main = putStrLn $ parse $ putEmployee "yoshiken" 31
parse e = name e ++ ":" ++ show(age e) ++ ":" ++ show(version e)
