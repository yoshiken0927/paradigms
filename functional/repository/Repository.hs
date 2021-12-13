module Repository
    ( findEmployee,
      updateEmployee,
    ) where

import Model

findEmployee :: String -> Employee
findEmployee name = Employee {
  name = name,
  age = 30,
  version = 0
}


updateEmployee :: Employee -> Employee
updateEmployee e = Employee {
  name = name(e),
  age = age(e),
  version = version(e) + 1
}


