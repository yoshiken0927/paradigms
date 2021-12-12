module Repository
    ( findEmployee,
      updateEmployee,
    ) where

import Model

findEmployee :: String -> Employee
updateEmployee :: Employee -> Employee

findEmployee name = Employee { name = name, age = 30 }
-- updateEmployee e = e
updateEmployee e = Employee { name = name(e), age = 32 }
