module Main where

import Model
import Service

main :: IO ()
main = putStrLn $ parse $ putEmployee "yoshiken" 31

parse :: Employee -> String
parse e = name e ++ ":" ++ show(age e) ++ ":" ++ show(version e)
